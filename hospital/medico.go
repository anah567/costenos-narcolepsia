package hospital

import (
	"errors"
	"fmt"
	"strings"
)

// Medico guarda los datos propios de un médico.
// Incluye Persona para reutilizar id, nombre y edad.
// También guarda su especialidad, los pacientes que tiene a cargo y los episodios de sueño que ha atendido.
type Medico struct {
	Persona
	especialidad string
	pacientes    []*Paciente
	episodios    []RegistroEpisodio
}

// NuevoMedico crea un médico y valida que sus datos sean correctos.
func NuevoMedico(
	id string,
	nombre string,
	edad int,
	especialidad string,
) (*Medico, error) {

	// Reutiliza nuevaPersona para validar el id, nombre y edad.
	persona, err := nuevaPersona(id, nombre, edad)
	if err != nil {
		return nil, fmt.Errorf("médico inválido: %w", err)
	}

	// Quita espacios innecesarios de la especialidad.
	especialidad = strings.TrimSpace(especialidad)

	// Un médico debe tener una especialidad.
	if especialidad == "" {
		return nil, fmt.Errorf(
			"médico %s inválido: la especialidad no puede estar vacía",
			id,
		)
	}

	return &Medico{
		Persona:      persona,
		especialidad: especialidad,
	}, nil
}

// Devuelve la especialidad del médico.
func (m *Medico) Especialidad() string {
	return m.especialidad
}

// Devuelve una copia de los pacientes que están a cargo del médico.
// La copia evita que la lista original se modifique directamente desde afuera.
func (m *Medico) Pacientes() []*Paciente {
	copia := make([]*Paciente, len(m.pacientes))
	copy(copia, m.pacientes)

	return copia
}

// DiagnosticarPaciente pone a un paciente bajo el cuidado de este médico.
func (m *Medico) DiagnosticarPaciente(p *Paciente) error {

	// Primero se valida que exista un paciente.
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	// Recorre los pacientes del médico para evitar agregar el mismo dos veces.
	for _, existente := range m.pacientes {
		if existente == p {
			return fmt.Errorf(
				"el paciente %s ya está a cargo de %s",
				p.ID(),
				m.Nombre(),
			)
		}
	}

	// Agrega al paciente a la lista del médico.
	m.pacientes = append(m.pacientes, p)

	// También guarda este médico dentro del paciente.
	// Así la relación queda registrada en ambos lados.
	p.asignarMedico(m)

	return nil
}

// AtenderEmergenciaSueno verifica que realmente exista una emergencia  y asigna este médico al paciente si todavía no tiene uno.
func (m *Medico) AtenderEmergenciaSueno(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	// Si está despierto, no existe un ataque de sueño que atender.
	if p.Estado() == Despierto {
		return fmt.Errorf(
			"el paciente %s está despierto, no hay emergencia",
			p.ID(),
		)
	}

	// Si todavía no tiene médico, este médico queda asignado al paciente.
	if p.MedicoAsignado() == nil {
		return m.DiagnosticarPaciente(p)
	}

	return nil
}

// Atender permite que el médico atienda un ataque de sueño.
// Este método también hace que Medico cumpla con la interfaz Atendedor.
func (m *Medico) Atender(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	// Antes de registrar el episodio, verifica que la emergencia pueda atenderse.
	if err := m.AtenderEmergenciaSueno(p); err != nil {
		return RegistroEpisodio{}, err
	}

	// Crea el registro con el paciente, el médico y el lugar del ataque.
	registro := nuevoRegistroEpisodio(
		p,
		m,
		ubicacion,
	)

	// Guarda el episodio dentro del historial personal del médico.
	m.episodios = append(
		m.episodios,
		registro,
	)

	return registro, nil
}

// Devuelve una copia de todos los episodios atendidos por este médico.
// Esta lista se usa después para la consulta de episodios por médico.
func (m *Medico) MisEpisodios() []RegistroEpisodio {
	copia := make(
		[]RegistroEpisodio,
		len(m.episodios),
	)

	copy(copia, m.episodios)

	return copia
}
