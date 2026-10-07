package hospital

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Error específico que se usa cuando no se encuentra ninguna habitación disponible.
// Al tenerlo como una variable, después podemos reconocer este error con errors.Is.
var ErrSinHabitacion = errors.New("no hay habitación disponible")

// Hospital representa el sistema central del hospital.
// Guarda los médicos, todo el personal que puede atender, los pacientes, las habitaciones disponibles y el historial de episodios ocurridos.
type Hospital struct {
	nombre       string
	medicos      []*Medico
	personal     []Atendedor
	pacientes    []*Paciente
	habitaciones []*Habitacion
	historial    []RegistroEpisodio

	// Guarda qué miembro del personal debe atender el siguiente episodio.
	// Se utiliza para repartir las atenciones por turnos.
	siguientePersonal int
}

// NuevoHospital crea un hospital y valida que tenga un nombre válido.
func NuevoHospital(nombre string) (*Hospital, error) {

	// Elimina espacios innecesarios al inicio y al final del nombre.
	nombre = strings.TrimSpace(nombre)

	// No se permite crear un hospital sin nombre.
	if nombre == "" {
		return nil, errors.New(
			"el nombre del hospital no puede estar vacío",
		)
	}

	return &Hospital{
		nombre: nombre,
	}, nil
}

// Devuelve el nombre del hospital.
func (h *Hospital) Nombre() string {
	return h.nombre
}

// Devuelve una copia de la lista de médicos.
// Así no se modifica directamente la lista original guardada en el hospital.
func (h *Hospital) Medicos() []*Medico {
	copia := make([]*Medico, len(h.medicos))
	copy(copia, h.medicos)

	return copia
}

// Devuelve una copia de todos los pacientes admitidos en el hospital.
func (h *Hospital) Pacientes() []*Paciente {
	copia := make([]*Paciente, len(h.pacientes))
	copy(copia, h.pacientes)

	return copia
}

// Devuelve una copia de las habitaciones registradas en el hospital.
func (h *Hospital) Habitaciones() []*Habitacion {
	copia := make([]*Habitacion, len(h.habitaciones))
	copy(copia, h.habitaciones)

	return copia
}

// Devuelve una copia del historial completo de episodios.
// El historial original queda protegido dentro del hospital.
func (h *Hospital) Historial() []RegistroEpisodio {
	copia := make([]RegistroEpisodio, len(h.historial))
	copy(copia, h.historial)

	return copia
}

// ContratarPersonal agrega una persona que pueda atender ataques de sueño.
// Recibe un Atendedor, por lo que puede recibir distintos tipos, como un Medico o un Camillero.
func (h *Hospital) ContratarPersonal(a Atendedor) error {
	if a == nil {
		return errors.New(
			"el miembro del personal no puede ser nil",
		)
	}

	// Revisa los IDs para evitar contratar dos veces al mismo miembro del personal.
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

// ContratarMedico registra un médico en el hospital.
// Primero lo agrega al personal general porque también es un Atendedor, y después lo guarda en la lista específica de médicos.
func (h *Hospital) ContratarMedico(m *Medico) error {
	if m == nil {
		return errors.New("el médico no puede ser nil")
	}

	// Como Medico cumple la interfaz Atendedor, puede enviarse a ContratarPersonal.
	if err := h.ContratarPersonal(m); err != nil {
		return err
	}

	h.medicos = append(h.medicos, m)

	return nil
}

// RegistrarHabitacion agrega una habitación al hospital.
func (h *Hospital) RegistrarHabitacion(hab *Habitacion) error {
	if hab == nil {
		return errors.New(
			"la habitación no puede ser nil",
		)
	}

	// Comprueba que no exista otra habitación con el mismo número.
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

// AdmitirPaciente agrega un paciente al hospital.
// Antes verifica que exista y que no haya sido admitido anteriormente.
func (h *Hospital) AdmitirPaciente(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	// Busca el ID para evitar registrar dos veces al mismo paciente.
	if h.estaAdmitido(p.ID()) {
		return fmt.Errorf(
			"el paciente %s ya está admitido",
			p.ID(),
		)
	}

	h.pacientes = append(h.pacientes, p)

	return nil
}

// AsignarHabitacion busca una habitación disponible para un paciente que se encuentra dormido en el pasillo.
func (h *Hospital) AsignarHabitacion(
	p *Paciente,
) (*Habitacion, error) {

	if p == nil {
		return nil, errors.New(
			"el paciente no puede ser nil",
		)
	}

	// Solo se pueden asignar habitaciones a pacientes admitidos en el hospital.
	if !h.estaAdmitido(p.ID()) {
		return nil, fmt.Errorf(
			"el paciente %s no está admitido",
			p.ID(),
		)
	}

	// La habitación se asigna únicamente si el paciente está dormido en el pasillo.
	if p.Estado() != DormidoEnPasillo {
		return nil, fmt.Errorf(
			"el paciente %s no está dormido en el pasillo (estado: %s)",
			p.ID(),
			p.Estado(),
		)
	}

	// Recorre las habitaciones hasta encontrar la primera que tenga espacio.
	for _, hab := range h.habitaciones {

		// Si está ocupada, continue salta directamente a la siguiente habitación.
		if !hab.EstaDisponible() {
			continue
		}

		// Agrega al paciente a la lista de ocupantes de la habitación.
		if err := hab.Ocupar(p); err != nil {
			return nil, err
		}

		// También actualiza al paciente para indicar que ahora está en una cama.
		p.acostarEnCama(hab)

		return hab, nil
	}

	// Si se recorrieron todas las habitaciones y ninguna estaba disponible, se devuelve ErrSinHabitacion y el paciente permanece en el pasillo.
	return nil, fmt.Errorf(
		"%w: el paciente %s sigue en el pasillo",
		ErrSinHabitacion,
		p.ID(),
	)
}

// RegistrarEpisodio agrega un episodio al historial general del hospital.
func (h *Hospital) RegistrarEpisodio(
	e RegistroEpisodio,
) error {

	// Todo episodio debe estar relacionado con un paciente.
	if e.Paciente() == nil {
		return errors.New(
			"el episodio no tiene paciente",
		)
	}

	// Solo se registran episodios de pacientes admitidos en este hospital.
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

// AtenderAtaqueSueno coordina todo el proceso cuando un paciente sufre un ataque de sueño: cambia su estado, intenta conseguirle
// una habitación, busca quién lo atienda y registra el episodio.
func (h *Hospital) AtenderAtaqueSueno(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	// Primero se valida que exista el paciente.
	if p == nil {
		return RegistroEpisodio{}, errors.New(
			"el paciente no puede ser nil",
		)
	}

	// El paciente debe estar admitido para poder ser atendido.
	if !h.estaAdmitido(p.ID()) {
		return RegistroEpisodio{}, fmt.Errorf(
			"el paciente %s no está admitido",
			p.ID(),
		)
	}

	// Debe existir por lo menos un Atendedor disponible en el hospital.
	if len(h.personal) == 0 {
		return RegistroEpisodio{}, errors.New(
			"no hay personal para atender al paciente",
		)
	}

	// Cambia al paciente de Despierto a DormidoEnPasillo y guarda el lugar donde ocurrió el ataque.
	if err := p.SufrirAtaqueSueno(ubicacion); err != nil {
		return RegistroEpisodio{}, err
	}

	// Intenta conseguir una habitación para el paciente.
	_, errHabitacion := h.AsignarHabitacion(p)

	// ErrSinHabitacion es un caso esperado: el episodio puede continuar aunque el paciente tenga que quedarse dormido en el pasillo.
	// Cualquier otro error sí detiene el proceso.
	if errHabitacion != nil &&
		!errors.Is(errHabitacion, ErrSinHabitacion) {

		return RegistroEpisodio{}, errHabitacion
	}

	// despachar selecciona al siguiente Atendedor.
	// Como devuelve la interfaz Atendedor, aquí puede atender un Medico o un Camillero.
	registro, err := h.despachar().Atender(
		p,
		ubicacion,
	)

	if err != nil {
		return RegistroEpisodio{}, err
	}

	// Guarda el episodio atendido en el historial general del hospital.
	if err := h.RegistrarEpisodio(registro); err != nil {
		return RegistroEpisodio{}, err
	}

	// Devuelve el registro y también el posible ErrSinHabitacion.
	// Así sabemos que el episodio sí fue atendido aunque no hubiera cama.
	return registro, errHabitacion
}

// PacientesEnPasillo busca los pacientes que actualmente están dormidos en el pasillo y los devuelve en una lista.
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

// ReporteSeveros crea un reporte de los pacientes con narcolepsia Severa.
// El mapa relaciona cada paciente severo con la cantidad de episodios que ha tenido durante el día actual.
func (h *Hospital) ReporteSeveros() map[*Paciente]int {
	reporte := make(map[*Paciente]int)

	// Primero agrega todos los pacientes severos con contador en cero.
	// Así también aparecen en el reporte aunque hoy no hayan tenido episodios.
	for _, p := range h.pacientes {
		if p.Nivel() == Severo {
			reporte[p] = 0
		}
	}

	// Guarda la fecha y hora actual para comparar los episodios con el día de hoy.
	hoy := time.Now()

	// Recorre todo el historial para contar los episodios de pacientes severos.
	for _, e := range h.historial {

		// Primero comprueba si el paciente está dentro del mapa de severos y después verifica que el episodio haya ocurrido hoy.
		if _, esSevero := reporte[e.Paciente()]; esSevero &&
			mismoDia(e.FechaHora(), hoy) {

			reporte[e.Paciente()]++
		}
	}

	return reporte
}

// despachar selecciona quién atenderá el siguiente episodio.
// Reparte las atenciones por turnos entre todos los elementos de personal.
func (h *Hospital) despachar() Atendedor {

	// El operador % hace que al llegar al final de la lista se vuelva al inicio.
	// Por ejemplo, con 3 personas los índices serán: 0, 1, 2, 0, 1, 2...
	a := h.personal[h.siguientePersonal%len(h.personal)]

	// Deja preparado el turno para la siguiente atención.
	h.siguientePersonal++

	return a
}

// estaAdmitido busca un paciente por su ID.
// Devuelve true si está admitido y false si no lo encuentra.
func (h *Hospital) estaAdmitido(id string) bool {
	for _, p := range h.pacientes {
		if p.ID() == id {
			return true
		}
	}

	return false
}

// mismoDia compara dos fechas.
// Devuelve true solamente si tienen el mismo año, mes y día.
func mismoDia(a, b time.Time) bool {
	a1, m1, d1 := a.Date()
	a2, m2, d2 := b.Date()

	return a1 == a2 &&
		m1 == m2 &&
		d1 == d2
}
