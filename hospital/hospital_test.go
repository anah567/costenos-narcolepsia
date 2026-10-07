package hospital

import (
	"errors"
	"testing"
)

// preparar crea un hospital listo para usar en las pruebas.
// Agrega un médico, un camillero y la cantidad de habitaciones
// que necesite cada test. Así evitamos repetir esta preparación en todos.
func preparar(t *testing.T, habitaciones int) *Hospital {
	// Marca esta función como una función auxiliar de los tests.
	// Así, si ocurre un error, Go indica la línea del test que la llamó
	// y no solamente una línea dentro de esta función.
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

	// El médico se registra como médico y también queda dentro
	// del personal general que puede atender episodios.
	if err := h.ContratarMedico(m); err != nil {
		t.Fatal(err)
	}

	// El camillero se agrega directamente como personal Atendedor.
	if err := h.ContratarPersonal(c); err != nil {
		t.Fatal(err)
	}

	// Crea las habitaciones necesarias para cada prueba.
	// Todas tienen capacidad para un solo paciente.
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

// admitir crea un paciente de prueba y lo admite en el hospital.
// Recibe el ID y el nivel para poder crear distintos pacientes según cada test.
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

// Comprueba que un paciente dormido en el pasillo pueda recibir
// una habitación cuando existe una disponible.
func TestAsignarHabitacionConHabitacionLibre(t *testing.T) {
	h := preparar(t, 1)
	p := admitir(t, h, "P-1", Leve)

	// El paciente primero debe quedar dormido en el pasillo
	// para poder solicitar una habitación.
	if err := p.SufrirAtaqueSueno("cafetería"); err != nil {
		t.Fatal(err)
	}

	hab, err := h.AsignarHabitacion(p)

	// En este caso no debería ocurrir ningún error porque hay una habitación libre.
	if err != nil {
		t.Fatalf(
			"se esperaba una habitación, llegó error: %v",
			err,
		)
	}

	// Comprueba que se haya asignado la habitación 101
	// y que ahora aparezca como Ocupada.
	if hab.Numero() != 101 || hab.Estado() != Ocupada {
		t.Errorf(
			"se esperaba habitación 101 Ocupada, llegó %d %s",
			hab.Numero(),
			hab.Estado(),
		)
	}

	// También comprueba que el estado del paciente cambió correctamente.
	if p.Estado() != DormidoEnCama {
		t.Errorf(
			"se esperaba paciente DormidoEnCama, llegó %s",
			p.Estado(),
		)
	}
}

// Comprueba qué ocurre cuando todas las habitaciones están ocupadas.
// El sistema debe devolver ErrSinHabitacion y dejar al paciente en el pasillo.
func TestAsignarHabitacionSinHabitacionLibre(t *testing.T) {
	h := preparar(t, 1)

	primero := admitir(t, h, "P-1", Leve)
	segundo := admitir(t, h, "P-2", Severo)

	// El primer paciente ocupa la única habitación disponible.
	if _, err := h.AtenderAtaqueSueno(primero, "cafetería"); err != nil {
		t.Fatal(err)
	}

	// El segundo paciente sufre un ataque y queda inicialmente en el pasillo.
	if err := segundo.SufrirAtaqueSueno("pasillo 2"); err != nil {
		t.Fatal(err)
	}

	hab, err := h.AsignarHabitacion(segundo)

	// errors.Is permite comprobar específicamente que el error
	// recibido sea el error ErrSinHabitacion.
	if !errors.Is(err, ErrSinHabitacion) {
		t.Fatalf(
			"se esperaba ErrSinHabitacion, llegó %v",
			err,
		)
	}

	// Como no había espacio, no debe haberse asignado ninguna habitación.
	if hab != nil {
		t.Errorf(
			"no se esperaba habitación, llegó %d",
			hab.Numero(),
		)
	}

	// El paciente debe permanecer dormido en el pasillo.
	if segundo.Estado() != DormidoEnPasillo {
		t.Errorf(
			"el paciente debía seguir DormidoEnPasillo, llegó %s",
			segundo.Estado(),
		)
	}
}

// Comprueba que PacientesEnPasillo devuelva solamente
// los pacientes cuyo estado actual sea DormidoEnPasillo.
func TestPacientesEnPasillo(t *testing.T) {
	h := preparar(t, 1)

	enCama := admitir(t, h, "P-1", Leve)
	enPasillo := admitir(t, h, "P-2", Severo)

	// Este tercer paciente está admitido, pero permanece despierto.
	admitir(t, h, "P-3", Moderado)

	// El primer paciente ocupa la única habitación.
	if _, err := h.AtenderAtaqueSueno(enCama, "cafetería"); err != nil {
		t.Fatal(err)
	}

	// El segundo también sufre un ataque, pero ya no hay habitación.
	_, err := h.AtenderAtaqueSueno(enPasillo, "pasillo 2")

	// ErrSinHabitacion es esperado en esta prueba.
	// Cualquier error diferente sí significa que algo salió mal.
	if err != nil && !errors.Is(err, ErrSinHabitacion) {
		t.Fatal(err)
	}

	resultado := h.PacientesEnPasillo()

	// Solo P-2 debe aparecer en el resultado:
	// P-1 está en cama y P-3 está despierto.
	if len(resultado) != 1 || resultado[0] != enPasillo {
		t.Fatalf(
			"se esperaba solo a P-2 en el pasillo, llegaron %d pacientes",
			len(resultado),
		)
	}
}

// Comprueba que cuando un paciente se despierta,
// su habitación queda disponible para otro paciente que estaba en el pasillo.
func TestDespertarLiberaHabitacionParaPacienteDelPasillo(t *testing.T) {
	h := preparar(t, 1)

	primero := admitir(t, h, "P-1", Leve)
	segundo := admitir(t, h, "P-2", Severo)

	// P-1 ocupa la única habitación.
	if _, err := h.AtenderAtaqueSueno(primero, "cafetería"); err != nil {
		t.Fatal(err)
	}

	// P-2 sufre un ataque, pero debe quedarse en el pasillo.
	_, err := h.AtenderAtaqueSueno(segundo, "pasillo 2")

	if err != nil && !errors.Is(err, ErrSinHabitacion) {
		t.Fatal(err)
	}

	// Al despertar P-1, la habitación 101 debe quedar libre.
	if err := primero.Despertar(); err != nil {
		t.Fatal(err)
	}

	// Ahora la habitación liberada se intenta asignar a P-2.
	hab, err := h.AsignarHabitacion(segundo)

	if err != nil {
		t.Fatal(err)
	}

	// Comprueba que P-2 recibió exactamente la habitación que se liberó.
	if hab.Numero() != 101 {
		t.Fatalf(
			"se esperaba a P-2 en la habitación 101, llegó habitación %d",
			hab.Numero(),
		)
	}
}

// Comprueba dos consultas:
// los episodios atendidos por el médico y el reporte de pacientes severos.
func TestEpisodiosDelMedicoYReporteSeveros(t *testing.T) {
	h := preparar(t, 3)

	severo := admitir(t, h, "P-1", Severo)

	// Este paciente severo no tendrá ningún episodio.
	// Sirve para comprobar que igualmente aparezca en ReporteSeveros con valor 0.
	admitir(t, h, "P-2", Severo)

	leve := admitir(t, h, "P-3", Leve)

	// El primer ataque lo atiende el médico porque es el primero
	// en el orden de personal del hospital.
	if _, err := h.AtenderAtaqueSueno(severo, "cafetería"); err != nil {
		t.Fatal(err)
	}

	// El segundo ataque lo atiende el siguiente miembro del personal,
	// que en esta preparación es el camillero.
	if _, err := h.AtenderAtaqueSueno(leve, "recepción"); err != nil {
		t.Fatal(err)
	}

	medico := h.Medicos()[0]
	episodios := medico.MisEpisodios()

	// El médico solamente debe tener el episodio del primer paciente.
	if len(episodios) != 1 || episodios[0].Paciente() != severo {
		t.Errorf(
			"se esperaba 1 episodio del médico con P-1, llegaron %d",
			len(episodios),
		)
	}

	reporte := h.ReporteSeveros()

	// Deben aparecer los dos pacientes severos.
	// P-1 debe tener un episodio y P-2 debe aparecer con cero.
	if len(reporte) != 2 || reporte[severo] != 1 {
		t.Errorf(
			"se esperaban 2 severos y 1 episodio para P-1, llegó %v",
			reporte,
		)
	}
}
