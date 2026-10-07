package hospital

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrSinHabitacion = errors.New("no hay habitación disponible")

type Hospital struct {
	nombre            string
	medicos           []*Medico
	personal          []Atendedor
	pacientes         []*Paciente
	habitaciones      []*Habitacion
	historial         []RegistroEpisodio
	siguientePersonal int
}

func NuevoHospital(nombre string) (*Hospital, error) {
	nombre = strings.TrimSpace(nombre)

	if nombre == "" {
		return nil, errors.New(
			"el nombre del hospital no puede estar vacío",
		)
	}

	return &Hospital{
		nombre: nombre,
	}, nil
}

func (h *Hospital) Nombre() string {
	return h.nombre
}

func (h *Hospital) Medicos() []*Medico {
	copia := make([]*Medico, len(h.medicos))
	copy(copia, h.medicos)

	return copia
}

func (h *Hospital) Pacientes() []*Paciente {
	copia := make([]*Paciente, len(h.pacientes))
	copy(copia, h.pacientes)

	return copia
}

func (h *Hospital) Habitaciones() []*Habitacion {
	copia := make([]*Habitacion, len(h.habitaciones))
	copy(copia, h.habitaciones)

	return copia
}

func (h *Hospital) Historial() []RegistroEpisodio {
	copia := make([]RegistroEpisodio, len(h.historial))
	copy(copia, h.historial)

	return copia
}

func (h *Hospital) ContratarPersonal(a Atendedor) error {
	if a == nil {
		return errors.New(
			"el miembro del personal no puede ser nil",
		)
	}

	for _, existente := range h.personal {
		if existente.ID() == a.ID() {
			return fmt.Errorf(
				"%s ya está contratado",
				a.ID(),
			)
		}
	}

	h.personal = append(h.personal, a)

	return nil
}

func (h *Hospital) ContratarMedico(m *Medico) error {
	if m == nil {
		return errors.New("el médico no puede ser nil")
	}

	if err := h.ContratarPersonal(m); err != nil {
		return err
	}

	h.medicos = append(h.medicos, m)

	return nil
}

func (h *Hospital) RegistrarHabitacion(hab *Habitacion) error {
	if hab == nil {
		return errors.New(
			"la habitación no puede ser nil",
		)
	}

	for _, existente := range h.habitaciones {
		if existente.Numero() == hab.Numero() {
			return fmt.Errorf(
				"la habitación %d ya está registrada",
				hab.Numero(),
			)
		}
	}

	h.habitaciones = append(
		h.habitaciones,
		hab,
	)

	return nil
}

func (h *Hospital) AdmitirPaciente(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	if h.estaAdmitido(p.ID()) {
		return fmt.Errorf(
			"el paciente %s ya está admitido",
			p.ID(),
		)
	}

	h.pacientes = append(h.pacientes, p)

	return nil
}

func (h *Hospital) AsignarHabitacion(
	p *Paciente,
) (*Habitacion, error) {

	if p == nil {
		return nil, errors.New(
			"el paciente no puede ser nil",
		)
	}

	if !h.estaAdmitido(p.ID()) {
		return nil, fmt.Errorf(
			"el paciente %s no está admitido",
			p.ID(),
		)
	}

	if p.Estado() != DormidoEnPasillo {
		return nil, fmt.Errorf(
			"el paciente %s no está dormido en el pasillo (estado: %s)",
			p.ID(),
			p.Estado(),
		)
	}

	for _, hab := range h.habitaciones {
		if !hab.EstaDisponible() {
			continue
		}

		if err := hab.Ocupar(p); err != nil {
			return nil, err
		}

		p.acostarEnCama(hab)

		return hab, nil
	}

	return nil, fmt.Errorf(
		"%w: el paciente %s sigue en el pasillo",
		ErrSinHabitacion,
		p.ID(),
	)
}

func (h *Hospital) RegistrarEpisodio(
	e RegistroEpisodio,
) error {

	if e.Paciente() == nil {
		return errors.New(
			"el episodio no tiene paciente",
		)
	}

	if !h.estaAdmitido(e.Paciente().ID()) {
		return fmt.Errorf(
			"episodio %s: el paciente %s no está admitido",
			e.ID(),
			e.Paciente().ID(),
		)
	}

	h.historial = append(
		h.historial,
		e,
	)

	return nil
}

func (h *Hospital) AtenderAtaqueSueno(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	if p == nil {
		return RegistroEpisodio{}, errors.New(
			"el paciente no puede ser nil",
		)
	}

	if !h.estaAdmitido(p.ID()) {
		return RegistroEpisodio{}, fmt.Errorf(
			"el paciente %s no está admitido",
			p.ID(),
		)
	}

	if len(h.personal) == 0 {
		return RegistroEpisodio{}, errors.New(
			"no hay personal para atender al paciente",
		)
	}

	if err := p.SufrirAtaqueSueno(ubicacion); err != nil {
		return RegistroEpisodio{}, err
	}

	_, errHabitacion := h.AsignarHabitacion(p)

	if errHabitacion != nil &&
		!errors.Is(errHabitacion, ErrSinHabitacion) {

		return RegistroEpisodio{}, errHabitacion
	}

	registro, err := h.despachar().Atender(
		p,
		ubicacion,
	)

	if err != nil {
		return RegistroEpisodio{}, err
	}

	if err := h.RegistrarEpisodio(registro); err != nil {
		return RegistroEpisodio{}, err
	}

	return registro, errHabitacion
}

func (h *Hospital) PacientesEnPasillo() []*Paciente {
	var resultado []*Paciente

	for _, p := range h.pacientes {
		if p.Estado() == DormidoEnPasillo {
			resultado = append(
				resultado,
				p,
			)
		}
	}

	return resultado
}

func (h *Hospital) ReporteSeveros() map[*Paciente]int {
	reporte := make(map[*Paciente]int)

	for _, p := range h.pacientes {
		if p.Nivel() == Severo {
			reporte[p] = 0
		}
	}

	hoy := time.Now()

	for _, e := range h.historial {
		if _, esSevero := reporte[e.Paciente()]; esSevero &&
			mismoDia(e.FechaHora(), hoy) {

			reporte[e.Paciente()]++
		}
	}

	return reporte
}
func (h *Hospital) despachar() Atendedor {
	a := h.personal[h.siguientePersonal%len(h.personal)]
	h.siguientePersonal++
	return a
}

func (h *Hospital) estaAdmitido(id string) bool {
	for _, p := range h.pacientes {
		if p.ID() == id {
			return true
		}
	}

	return false
}

func mismoDia(a, b time.Time) bool {
	a1, m1, d1 := a.Date()
	a2, m2, d2 := b.Date()

	return a1 == a2 &&
		m1 == m2 &&
		d1 == d2
}
