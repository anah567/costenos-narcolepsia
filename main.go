package main

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"costenos-narcolepsia/hospital"
)

// main ejecuta el escenario principal del hospital.
// Aquí se crean los datos de prueba y se usan las funciones del paquete hospital.
// La lógica del negocio está dentro del paquete hospital, no directamente en main.
func main() {
	// Crea el hospital donde se realizará todo el escenario.
	h, err := hospital.NuevoHospital("Hospital de los Costeños con Narcolepsia")
	verificar(err)

	// Crea dos médicos y un camillero.
	// Los tres podrán participar en la atención de los ataques de sueño.
	karen, err := hospital.NuevoMedico(
		"D-001",
		"Dra. Karen Ospina",
		45,
		"Medicina del sueño",
	)
	verificar(err)

	julio, err := hospital.NuevoMedico(
		"D-002",
		"Dr. Julio Mendoza",
		52,
		"Neurología",
	)
	verificar(err)

	camilo, err := hospital.NuevoCamillero(
		"C-001",
		"Camilo Barros",
		30,
	)
	verificar(err)

	// Los médicos se contratan con ContratarMedico porque además
	// de ser personal general deben quedar registrados como médicos.
	verificar(h.ContratarMedico(karen))
	verificar(h.ContratarMedico(julio))

	// El camillero se registra directamente como personal Atendedor.
	verificar(h.ContratarPersonal(camilo))

	// Crea las habitaciones 101, 102 y 103.
	// Cada una tiene capacidad para un solo paciente.
	for _, numero := range []int{101, 102, 103} {
		hab, err := hospital.NuevaHabitacion(numero, 1)
		verificar(err)

		verificar(h.RegistrarHabitacion(hab))
	}

	// Crea cinco pacientes con diferentes niveles de narcolepsia.
	// P-002 y P-004 son Severos para poder generar el reporte de severidad.
	p1 := nuevoPaciente("P-001", "Rafael Escalona", 34, hospital.Moderado)
	p2 := nuevoPaciente("P-002", "Yeimy Padilla", 28, hospital.Severo)
	p3 := nuevoPaciente("P-003", "Kevin Pertuz", 22, hospital.Leve)
	p4 := nuevoPaciente("P-004", "Wilfrido Berrio", 61, hospital.Severo)
	p5 := nuevoPaciente("P-005", "Luz Marina Cuesta", 40, hospital.Moderado)

	// Admite los cinco pacientes dentro del hospital.
	for _, p := range []*hospital.Paciente{p1, p2, p3, p4, p5} {
		verificar(h.AdmitirPaciente(p))
	}

	fmt.Println("== Ataques de sueño ==")

	// Esta lista guarda qué paciente sufrirá cada ataque
	// y la ubicación donde ocurre.
	ataques := []struct {
		paciente  *hospital.Paciente
		ubicacion string
	}{
		{p1, "cafetería"},
		{p2, "clase de champeta"},
		{p3, "fila de radiología"},
		{p4, "pasillo 2, segundo piso"},
	}

	// Ejecuta los cuatro ataques de sueño uno por uno.
	for _, a := range ataques {
		episodio, err := h.AtenderAtaqueSueno(
			a.paciente,
			a.ubicacion,
		)

		switch {
		// Si no hay habitación, el episodio sí ocurrió y se muestra,
		// pero también se informa que el paciente quedó en el pasillo.
		case errors.Is(err, hospital.ErrSinHabitacion):
			fmt.Printf(
				"  %s\n    !! %v\n",
				episodio.Resumen(),
				err,
			)

		// Si ocurrió cualquier otro error, se considera un problema
		// inesperado y verificar termina el programa.
		case err != nil:
			verificar(err)

		// Si no ocurrió ningún error, simplemente se muestra el episodio.
		default:
			fmt.Printf(
				"  %s\n",
				episodio.Resumen(),
			)
		}
	}

	// Después de cuatro ataques y solamente tres habitaciones,
	// P-004 debe aparecer dormido en el pasillo.
	fmt.Println("\n== Pacientes dormidos en el pasillo (después de los ataques) ==")
	imprimirPasillo(h.PacientesEnPasillo())

	// P-001 se despierta y al hacerlo libera su habitación.
	fmt.Printf("\n== %s se despierta ==\n", p1.ID())
	verificar(p1.Despertar())

	fmt.Printf(
		"  %s ahora está %s\n",
		p1.Nombre(),
		p1.Estado(),
	)

	// La habitación que liberó P-001 ahora puede asignarse a P-004,
	// que era el paciente que había quedado en el pasillo.
	fmt.Printf("\n== Asignando habitación a %s ==\n", p4.ID())

	hab, err := h.AsignarHabitacion(p4)

	if err != nil {
		fmt.Printf("  %v\n", err)
	} else {
		fmt.Printf(
			"  %s -> habitación %d (%s)\n",
			p4.Nombre(),
			hab.Numero(),
			p4.Estado(),
		)
	}

	// Consulta 5.1:
	// muestra los pacientes que siguen dormidos en el pasillo.
	fmt.Println("\n== 5.1 Pacientes dormidos en el pasillo ==")
	imprimirPasillo(h.PacientesEnPasillo())

	// Consulta 5.2:
	// muestra los episodios atendidos por cada médico.
	for _, m := range h.Medicos() {
		fmt.Printf(
			"\n== 5.2 Episodios atendidos por %s ==\n",
			m.Nombre(),
		)

		episodios := m.MisEpisodios()

		if len(episodios) == 0 {
			fmt.Println("  (ninguno)")
		}

		for _, e := range episodios {
			fmt.Printf(
				"  %s\n",
				e.Resumen(),
			)
		}
	}

	// Consulta 5.3:
	// muestra el estado de cada habitación y quién la está ocupando.
	fmt.Println("\n== 5.3 Disponibilidad de camas ==")

	for _, hab := range h.Habitaciones() {
		// Si no hay ocupantes, se muestra "-" como valor inicial.
		ocupante := "-"

		// Si existe un ocupante, se muestra su ID y nombre.
		if o := hab.Ocupantes(); len(o) > 0 {
			ocupante = o[0].ID() + " " + o[0].Nombre()
		}

		fmt.Printf(
			"  habitación %d  %-10s %s\n",
			hab.Numero(),
			hab.Estado(),
			ocupante,
		)
	}

	// Consulta 5.4:
	// muestra los pacientes Severos y cuántos episodios han tenido hoy.
	fmt.Println("\n== 5.4 Reporte de severidad (pacientes Severos) ==")

	reporte := h.ReporteSeveros()

	// El reporte es un map y los maps no garantizan un orden al recorrerlos.
	// Por eso primero guardamos los pacientes en una lista.
	severos := make(
		[]*hospital.Paciente,
		0,
		len(reporte),
	)

	for p := range reporte {
		severos = append(
			severos,
			p,
		)
	}

	// Ordena los pacientes por ID para que el reporte siempre
	// se muestre de forma organizada.
	sort.Slice(
		severos,
		func(i, j int) bool {
			return severos[i].ID() < severos[j].ID()
		},
	)

	// Finalmente muestra cada paciente severo y su cantidad de episodios.
	for _, p := range severos {
		fmt.Printf(
			"  %s  %-18s %d episodio(s) hoy\n",
			p.ID(),
			p.Nombre(),
			reporte[p],
		)
	}
}

// nuevoPaciente es una función auxiliar de main.
// Crea un paciente y evita repetir el mismo manejo de errores
// cada vez que necesitamos uno para el escenario.
func nuevoPaciente(
	id string,
	nombre string,
	edad int,
	nivel hospital.NivelNarcolepsia,
) *hospital.Paciente {

	p, err := hospital.NuevoPaciente(
		id,
		nombre,
		edad,
		nivel,
	)

	verificar(err)

	return p
}

// imprimirPasillo recibe la lista de pacientes que están en el pasillo
// y muestra sus datos de una forma fácil de leer.
func imprimirPasillo(pacientes []*hospital.Paciente) {
	// Si la lista está vacía, significa que no queda nadie en el pasillo.
	if len(pacientes) == 0 {
		fmt.Println(
			"  (nadie, los pasillos están despejados)",
		)
	}

	for _, p := range pacientes {
		fmt.Printf(
			"  %s  %-18s (%s)\n",
			p.ID(),
			p.Nombre(),
			p.Ubicacion(),
		)
	}
}

// verificar centraliza el manejo de errores inesperados en main.
// Si recibe un error, log.Fatal lo muestra y termina la ejecución.
// Si err es nil, simplemente continúa normalmente.
func verificar(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
