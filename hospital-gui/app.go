package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"costenos-narcolepsia/hospital"
)

// App representa la aplicación gráfica.
// La interfaz utiliza el mismo modelo del paquete hospital.
type App struct {
	ctx context.Context

	hospital *hospital.Hospital

	karen  *hospital.Medico
	julio  *hospital.Medico
	camilo *hospital.Camillero
}

// DTO
type PacienteDTO struct {
	ID         string `json:"id"`
	Nombre     string `json:"nombre"`
	Edad       int    `json:"edad"`
	Nivel      string `json:"nivel"`
	Estado     string `json:"estado"`
	Ubicacion  string `json:"ubicacion"`
	Habitacion string `json:"habitacion"`
}

type HabitacionDTO struct {
	Numero     int    `json:"numero"`
	Estado     string `json:"estado"`
	Disponible bool   `json:"disponible"`
	Ocupante   string `json:"ocupante"`
}

type PersonalDTO struct {
	ID           string `json:"id"`
	Nombre       string `json:"nombre"`
	Cargo        string `json:"cargo"`
	Especialidad string `json:"especialidad"`
	Episodios    int    `json:"episodios"`
}

type EpisodioDTO struct {
	ID             string `json:"id"`
	FechaHora      string `json:"fechaHora"`
	PacienteID     string `json:"pacienteId"`
	PacienteNombre string `json:"pacienteNombre"`
	Ubicacion      string `json:"ubicacion"`
	Habitacion     string `json:"habitacion"`
	AtendidoPor    string `json:"atendidoPor"`
	Nivel          string `json:"nivel"`
	Resumen        string `json:"resumen"`
}

type SeveroDTO struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Episodios int    `json:"episodios"`
}

type EstadoHospitalDTO struct {
	Nombre string `json:"nombre"`

	Pacientes    []PacienteDTO   `json:"pacientes"`
	Habitaciones []HabitacionDTO `json:"habitaciones"`
	Personal     []PersonalDTO   `json:"personal"`
	Episodios    []EpisodioDTO   `json:"episodios"`
	Severos      []SeveroDTO     `json:"severos"`

	PacientesEnPasillo int `json:"pacientesEnPasillo"`
	Dormidos           int `json:"dormidos"`
	HabitacionesLibres int `json:"habitacionesLibres"`
	TotalEpisodios     int `json:"totalEpisodios"`
}

// CREACIÓN DE LA APP
func NewApp() *App {

	h, err := hospital.NuevoHospital(
		"Hospital de los Costeños con Narcolepsia",
	)

	if err != nil {
		panic(err)
	}

	// PERSONAL
	karen, err := hospital.NuevoMedico(
		"D-001",
		"Dra. Karen Ospina",
		45,
		"Medicina del sueño",
	)

	if err != nil {
		panic(err)
	}

	julio, err := hospital.NuevoMedico(
		"D-002",
		"Dr. Julio Mendoza",
		52,
		"Neurología",
	)

	if err != nil {
		panic(err)
	}

	camilo, err := hospital.NuevoCamillero(
		"C-001",
		"Camilo Barros",
		30,
	)

	if err != nil {
		panic(err)
	}

	if err := h.ContratarMedico(karen); err != nil {
		panic(err)
	}

	if err := h.ContratarMedico(julio); err != nil {
		panic(err)
	}

	if err := h.ContratarPersonal(camilo); err != nil {
		panic(err)
	}

	// HABITACIONES
	for _, numero := range []int{101, 102, 103} {

		hab, err := hospital.NuevaHabitacion(
			numero,
			1,
		)

		if err != nil {
			panic(err)
		}

		if err := h.RegistrarHabitacion(hab); err != nil {
			panic(err)
		}
	}

	// PACIENTES
	p1 := crearPaciente(
		"P-001",
		"Rafael Escalona",
		34,
		hospital.Moderado,
	)

	p2 := crearPaciente(
		"P-002",
		"Yeimy Padilla",
		28,
		hospital.Severo,
	)

	p3 := crearPaciente(
		"P-003",
		"Kevin Pertuz",
		22,
		hospital.Leve,
	)

	p4 := crearPaciente(
		"P-004",
		"Wilfrido Berrio",
		61,
		hospital.Severo,
	)

	p5 := crearPaciente(
		"P-005",
		"Luz Marina Cuesta",
		40,
		hospital.Moderado,
	)

	for _, p := range []*hospital.Paciente{
		p1,
		p2,
		p3,
		p4,
		p5,
	} {

		if err := h.AdmitirPaciente(p); err != nil {
			panic(err)
		}
	}

	return &App{
		hospital: h,

		karen:  karen,
		julio:  julio,
		camilo: camilo,
	}
}

func crearPaciente(
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

	if err != nil {
		panic(err)
	}

	return p
}

// WAILS
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) NombreHospital() string {
	return a.hospital.Nombre()
}

// BUSCAR PACIENTE
func (a *App) buscarPaciente(id string) (*hospital.Paciente, error) {

	for _, p := range a.hospital.Pacientes() {

		if p.ID() == id {
			return p, nil
		}
	}

	return nil, fmt.Errorf(
		"no se encontró el paciente %s",
		id,
	)
}

// OBTENER PACIENTES
func (a *App) ObtenerPacientes() []PacienteDTO {

	pacientes := a.hospital.Pacientes()

	resultado := make(
		[]PacienteDTO,
		0,
		len(pacientes),
	)

	for _, p := range pacientes {

		habitacion := "-"

		if p.Habitacion() != nil {
			habitacion = fmt.Sprintf(
				"%d",
				p.Habitacion().Numero(),
			)
		}

		ubicacion := p.Ubicacion()

		if ubicacion == "" {
			ubicacion = "-"
		}

		resultado = append(
			resultado,
			PacienteDTO{
				ID:         p.ID(),
				Nombre:     p.Nombre(),
				Edad:       p.Edad(),
				Nivel:      p.Nivel().String(),
				Estado:     p.Estado().String(),
				Ubicacion:  ubicacion,
				Habitacion: habitacion,
			},
		)
	}

	return resultado
}

// REGISTRAR ATAQUE
func (a *App) RegistrarAtaque(
	pacienteID string,
	ubicacion string,
) (string, error) {

	paciente, err := a.buscarPaciente(
		pacienteID,
	)

	if err != nil {
		return "", err
	}

	ubicacion = strings.TrimSpace(
		ubicacion,
	)

	if ubicacion == "" {
		return "",
			errors.New(
				"debe indicar dónde ocurrió el ataque",
			)
	}

	registro, err := a.hospital.AtenderAtaqueSueno(
		paciente,
		ubicacion,
	)

	if errors.Is(
		err,
		hospital.ErrSinHabitacion,
	) {

		return fmt.Sprintf(
			"%s | No hay habitación disponible: el paciente permanece en el pasillo.",
			registro.Resumen(),
		), nil
	}

	if err != nil {
		return "", err
	}

	return registro.Resumen(), nil
}

// DESPERTAR PACIENTE
func (a *App) DespertarPaciente(
	pacienteID string,
) (string, error) {

	paciente, err := a.buscarPaciente(
		pacienteID,
	)

	if err != nil {
		return "", err
	}

	if err := paciente.Despertar(); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s se despertó correctamente.",
		paciente.Nombre(),
	), nil
}

// ASIGNAR HABITACIÓN
func (a *App) AsignarHabitacion(
	pacienteID string,
) (string, error) {

	paciente, err := a.buscarPaciente(
		pacienteID,
	)

	if err != nil {
		return "", err
	}

	hab, err := a.hospital.AsignarHabitacion(
		paciente,
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Se asignó la habitación %d a %s.",
		hab.Numero(),
		paciente.Nombre(),
	), nil
}

// ESTADO COMPLETO DEL HOSPITAL
func (a *App) ObtenerEstadoHospital() EstadoHospitalDTO {

	estado := EstadoHospitalDTO{
		Nombre:    a.hospital.Nombre(),
		Pacientes: a.ObtenerPacientes(),
	}

	// HABITACIONES
	for _, hab := range a.hospital.Habitaciones() {

		ocupante := "-"

		ocupantes := hab.Ocupantes()

		if len(ocupantes) > 0 {
			ocupante = fmt.Sprintf(
				"%s · %s",
				ocupantes[0].ID(),
				ocupantes[0].Nombre(),
			)
		}

		disponible := hab.EstaDisponible()

		if disponible {
			estado.HabitacionesLibres++
		}

		estado.Habitaciones = append(
			estado.Habitaciones,
			HabitacionDTO{
				Numero:     hab.Numero(),
				Estado:     hab.Estado().String(),
				Disponible: disponible,
				Ocupante:   ocupante,
			},
		)
	}

	// PACIENTES DORMIDOS
	for _, p := range a.hospital.Pacientes() {

		if p.Estado() != hospital.Despierto {
			estado.Dormidos++
		}
	}

	estado.PacientesEnPasillo =
		len(a.hospital.PacientesEnPasillo())

		// EPISODIOS

	for _, episodio := range a.hospital.Historial() {

		paciente := episodio.Paciente()

		habitacion := "-"

		if episodio.Habitacion() != nil {
			habitacion = fmt.Sprintf(
				"%d",
				episodio.Habitacion().Numero(),
			)
		}

		estado.Episodios = append(
			estado.Episodios,
			EpisodioDTO{
				ID:             episodio.ID(),
				FechaHora:      episodio.FechaHora().Format("02/01/2006 15:04"),
				PacienteID:     paciente.ID(),
				PacienteNombre: paciente.Nombre(),
				Ubicacion:      episodio.Ubicacion(),
				Habitacion:     habitacion,
				AtendidoPor:    episodio.AtendidoPor().Nombre(),
				Nivel:          paciente.Nivel().String(),
				Resumen:        episodio.Resumen(),
			},
		)
	}

	estado.TotalEpisodios = len(estado.Episodios)

	// PERSONAL
	estado.Personal = []PersonalDTO{
		{
			ID:           a.karen.ID(),
			Nombre:       a.karen.Nombre(),
			Cargo:        "Médico",
			Especialidad: a.karen.Especialidad(),
			Episodios:    len(a.karen.MisEpisodios()),
		},
		{
			ID:           a.julio.ID(),
			Nombre:       a.julio.Nombre(),
			Cargo:        "Médico",
			Especialidad: a.julio.Especialidad(),
			Episodios:    len(a.julio.MisEpisodios()),
		},
		{
			ID:           a.camilo.ID(),
			Nombre:       a.camilo.Nombre(),
			Cargo:        "Camillero",
			Especialidad: "Atención y traslado",
			Episodios:    len(a.camilo.MisEpisodios()),
		},
	}

	// REPORTE SEVEROS
	reporte := a.hospital.ReporteSeveros()

	for _, p := range a.hospital.Pacientes() {

		cantidad, existe := reporte[p]

		if !existe {
			continue
		}

		estado.Severos = append(
			estado.Severos,
			SeveroDTO{
				ID:        p.ID(),
				Nombre:    p.Nombre(),
				Episodios: cantidad,
			},
		)
	}

	return estado
}
