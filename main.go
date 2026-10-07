package main

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"costenos-narcolepsia/hospital"
)

func main() {
	h, err := hospital.NuevoHospital("Hospital de los Costeños con Narcolepsia")
	verificar(err)

	karen, err := hospital.NuevoMedico("D-001", "Dra. Karen Ospina", 45, "Medicina del sueño")
	verificar(err)
	julio, err := hospital.NuevoMedico("D-002", "Dr. Julio Mendoza", 52, "Neurología")
	verificar(err)
	camilo, err := hospital.NuevoCamillero("C-001", "Camilo Barros", 30)
	verificar(err)
	verificar(h.ContratarMedico(karen))
	verificar(h.ContratarMedico(julio))
	verificar(h.ContratarPersonal(camilo))

	for _, numero := range []int{101, 102, 103} {
		hab, err := hospital.NuevaHabitacion(numero, 1)
		verificar(err)
		verificar(h.RegistrarHabitacion(hab))
	}

	p1 := nuevoPaciente("P-001", "Rafael Escalona", 34, hospital.Moderado)
	p2 := nuevoPaciente("P-002", "Yeimy Padilla", 28, hospital.Severo)
	p3 := nuevoPaciente("P-003", "Kevin Pertuz", 22, hospital.Leve)
	p4 := nuevoPaciente("P-004", "Wilfrido Berrio", 61, hospital.Severo)
	p5 := nuevoPaciente("P-005", "Luz Marina Cuesta", 40, hospital.Moderado)
	for _, p := range []*hospital.Paciente{p1, p2, p3, p4, p5} {
		verificar(h.AdmitirPaciente(p))
	}

	fmt.Println("== Ataques de sueño ==")
	ataques := []struct {
		paciente  *hospital.Paciente
		ubicacion string
	}{
		{p1, "cafetería"},
		{p2, "clase de champeta"},
		{p3, "fila de radiología"},
		{p4, "pasillo 2, segundo piso"},
	}
	for _, a := range ataques {
		episodio, err := h.AtenderAtaqueSueno(a.paciente, a.ubicacion)
		switch {
		case errors.Is(err, hospital.ErrSinHabitacion):
			fmt.Printf("  %s\n    !! %v\n", episodio.Resumen(), err)
		case err != nil:
			verificar(err)
		default:
			fmt.Printf("  %s\n", episodio.Resumen())
		}
	}

	fmt.Println("\n== Pacientes dormidos en el pasillo (después de los ataques) ==")
	imprimirPasillo(h.PacientesEnPasillo())

	fmt.Printf("\n== %s se despierta ==\n", p1.ID())
	verificar(p1.Despertar())
	fmt.Printf("  %s ahora está %s\n", p1.Nombre(), p1.Estado())

	fmt.Printf("\n== Asignando habitación a %s ==\n", p4.ID())
	hab, err := h.AsignarHabitacion(p4)
	if err != nil {
		fmt.Printf("  %v\n", err)
	} else {
		fmt.Printf("  %s -> habitación %d (%s)\n", p4.Nombre(), hab.Numero(), p4.Estado())
	}

	fmt.Println("\n== 5.1 Pacientes dormidos en el pasillo ==")
	imprimirPasillo(h.PacientesEnPasillo())

	for _, m := range h.Medicos() {
		fmt.Printf("\n== 5.2 Episodios atendidos por %s ==\n", m.Nombre())
		episodios := m.MisEpisodios()
		if len(episodios) == 0 {
			fmt.Println("  (ninguno)")
		}
		for _, e := range episodios {
			fmt.Printf("  %s\n", e.Resumen())
		}
	}

	fmt.Println("\n== 5.3 Disponibilidad de camas ==")
	for _, hab := range h.Habitaciones() {
		ocupante := "-"
		if o := hab.Ocupantes(); len(o) > 0 {
			ocupante = o[0].ID() + " " + o[0].Nombre()
		}
		fmt.Printf("  habitación %d  %-10s %s\n", hab.Numero(), hab.Estado(), ocupante)
	}

	fmt.Println("\n== 5.4 Reporte de severidad (pacientes Severos) ==")
	reporte := h.ReporteSeveros()
	severos := make([]*hospital.Paciente, 0, len(reporte))
	for p := range reporte {
		severos = append(severos, p)
	}
	sort.Slice(severos, func(i, j int) bool { return severos[i].ID() < severos[j].ID() })
	for _, p := range severos {
		fmt.Printf("  %s  %-18s %d episodio(s) hoy\n", p.ID(), p.Nombre(), reporte[p])
	}
}

func nuevoPaciente(id, nombre string, edad int, nivel hospital.NivelNarcolepsia) *hospital.Paciente {
	p, err := hospital.NuevoPaciente(id, nombre, edad, nivel)
	verificar(err)
	return p
}

func imprimirPasillo(pacientes []*hospital.Paciente) {
	if len(pacientes) == 0 {
		fmt.Println("  (nadie, los pasillos están despejados)")
	}
	for _, p := range pacientes {
		fmt.Printf("  %s  %-18s (%s)\n", p.ID(), p.Nombre(), p.Ubicacion())
	}
}

func verificar(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
