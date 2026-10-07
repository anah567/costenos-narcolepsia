package hospital

import (
	"errors"
	"testing"
)

func preparar(t *testing.T, habitaciones int) *Hospital {
	t.Helper()

	h, err := NuevoHospital("Hospital de prueba")
	if err != nil {
		t.Fatal(err)
	}

	m, err := NuevoMedico("D-1", "Dr. Prueba", 40, "Sueño")
	if err != nil {
		t.Fatal(err)
	}

	c, err := NuevoCamillero("C-1", "Camillero Prueba", 30)
	if err != nil {
		t.Fatal(err)
	}

	if err := h.ContratarMedico(m); err != nil {
		t.Fatal(err)
	}

	if err := h.ContratarPersonal(c); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= habitaciones; i++ {
		hab, err := NuevaHabitacion(100+i, 1)
		if err != nil {
			t.Fatal(err)
		}

		if err := h.RegistrarHabitacion(hab); err != nil {
			t.Fatal(err)
		}
	}

	return h
}

func admitir(
	t *testing.T,
	h *Hospital,
	id string,
	nivel NivelNarcolepsia,
) *Paciente {

	t.Helper()

	p, err := NuevoPaciente(
		id,
		"Paciente "+id,
		30,
		nivel,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := h.AdmitirPaciente(p); err != nil {
		t.Fatal(err)
	}

	return p
}

func TestAsignarHabitacionConHabitacionLibre(t *testing.T) {
	h := preparar(t, 1)
	p := admitir(t, h, "P-1", Leve)

	if err := p.SufrirAtaqueSueno("cafetería"); err != nil {
		t.Fatal(err)
	}

	hab, err := h.AsignarHabitacion(p)

	if err != nil {
		t.Fatalf(
			"se esperaba una habitación, llegó error: %v",
			err,
		)
	}

	if hab.Numero() != 101 || hab.Estado() != Ocupada {
		t.Errorf(
			"se esperaba habitación 101 Ocupada, llegó %d %s",
			hab.Numero(),
			hab.Estado(),
		)
	}

	if p.Estado() != DormidoEnCama {
		t.Errorf(
			"se esperaba paciente DormidoEnCama, llegó %s",
			p.Estado(),
		)
	}
}

func TestAsignarHabitacionSinHabitacionLibre(t *testing.T) {
	h := preparar(t, 1)

	primero := admitir(t, h, "P-1", Leve)
	segundo := admitir(t, h, "P-2", Severo)

	if _, err := h.AtenderAtaqueSueno(primero, "cafetería"); err != nil {
		t.Fatal(err)
	}

	if err := segundo.SufrirAtaqueSueno("pasillo 2"); err != nil {
		t.Fatal(err)
	}

	hab, err := h.AsignarHabitacion(segundo)

	if !errors.Is(err, ErrSinHabitacion) {
		t.Fatalf(
			"se esperaba ErrSinHabitacion, llegó %v",
			err,
		)
	}

	if hab != nil {
		t.Errorf(
			"no se esperaba habitación, llegó %d",
			hab.Numero(),
		)
	}

	if segundo.Estado() != DormidoEnPasillo {
		t.Errorf(
			"el paciente debía seguir DormidoEnPasillo, llegó %s",
			segundo.Estado(),
		)
	}
}

func TestPacientesEnPasillo(t *testing.T) {
	h := preparar(t, 1)

	enCama := admitir(t, h, "P-1", Leve)
	enPasillo := admitir(t, h, "P-2", Severo)
	admitir(t, h, "P-3", Moderado)

	if _, err := h.AtenderAtaqueSueno(enCama, "cafetería"); err != nil {
		t.Fatal(err)
	}

	_, err := h.AtenderAtaqueSueno(enPasillo, "pasillo 2")

	if err != nil && !errors.Is(err, ErrSinHabitacion) {
		t.Fatal(err)
	}

	resultado := h.PacientesEnPasillo()

	if len(resultado) != 1 || resultado[0] != enPasillo {
		t.Fatalf(
			"se esperaba solo a P-2 en el pasillo, llegaron %d pacientes",
			len(resultado),
		)
	}
}

func TestDespertarLiberaHabitacionParaPacienteDelPasillo(t *testing.T) {
	h := preparar(t, 1)

	primero := admitir(t, h, "P-1", Leve)
	segundo := admitir(t, h, "P-2", Severo)

	if _, err := h.AtenderAtaqueSueno(primero, "cafetería"); err != nil {
		t.Fatal(err)
	}

	_, err := h.AtenderAtaqueSueno(segundo, "pasillo 2")

	if err != nil && !errors.Is(err, ErrSinHabitacion) {
		t.Fatal(err)
	}

	if err := primero.Despertar(); err != nil {
		t.Fatal(err)
	}

	hab, err := h.AsignarHabitacion(segundo)

	if err != nil {
		t.Fatal(err)
	}

	if hab.Numero() != 101 {
		t.Fatalf(
			"se esperaba a P-2 en la habitación 101, llegó habitación %d",
			hab.Numero(),
		)
	}
}

func TestEpisodiosDelMedicoYReporteSeveros(t *testing.T) {
	h := preparar(t, 3)

	severo := admitir(t, h, "P-1", Severo)
	admitir(t, h, "P-2", Severo)
	leve := admitir(t, h, "P-3", Leve)

	if _, err := h.AtenderAtaqueSueno(severo, "cafetería"); err != nil {
		t.Fatal(err)
	}

	if _, err := h.AtenderAtaqueSueno(leve, "recepción"); err != nil {
		t.Fatal(err)
	}

	medico := h.Medicos()[0]
	episodios := medico.MisEpisodios()

	if len(episodios) != 1 || episodios[0].Paciente() != severo {
		t.Errorf(
			"se esperaba 1 episodio del médico con P-1, llegaron %d",
			len(episodios),
		)
	}

	reporte := h.ReporteSeveros()

	if len(reporte) != 2 || reporte[severo] != 1 {
		t.Errorf(
			"se esperaban 2 severos y 1 episodio para P-1, llegó %v",
			reporte,
		)
	}
}
