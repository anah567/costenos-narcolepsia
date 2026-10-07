package hospital

import (
	"errors"
	"fmt"
	"strings"
)

type Medico struct {
	Persona
	especialidad string
	pacientes    []*Paciente
	episodios    []RegistroEpisodio
}

func NuevoMedico(
	id string,
	nombre string,
	edad int,
	especialidad string,
) (*Medico, error) {

	persona, err := nuevaPersona(id, nombre, edad)
	if err != nil {
		return nil, fmt.Errorf("médico inválido: %w", err)
	}

	especialidad = strings.TrimSpace(especialidad)

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

func (m *Medico) Especialidad() string {
	return m.especialidad
}

func (m *Medico) Pacientes() []*Paciente {
	copia := make([]*Paciente, len(m.pacientes))
	copy(copia, m.pacientes)

	return copia
}

func (m *Medico) DiagnosticarPaciente(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	for _, existente := range m.pacientes {
		if existente == p {
			return fmt.Errorf(
				"el paciente %s ya está a cargo de %s",
				p.ID(),
				m.Nombre(),
			)
		}
	}

	m.pacientes = append(m.pacientes, p)
	p.asignarMedico(m)

	return nil
}

func (m *Medico) AtenderEmergenciaSueno(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	if p.Estado() == Despierto {
		return fmt.Errorf(
			"el paciente %s está despierto, no hay emergencia",
			p.ID(),
		)
	}

	if p.MedicoAsignado() == nil {
		return m.DiagnosticarPaciente(p)
	}

	return nil
}

func (m *Medico) Atender(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	if err := m.AtenderEmergenciaSueno(p); err != nil {
		return RegistroEpisodio{}, err
	}

	registro := nuevoRegistroEpisodio(
		p,
		m,
		ubicacion,
	)

	m.episodios = append(
		m.episodios,
		registro,
	)

	return registro, nil
}

func (m *Medico) MisEpisodios() []RegistroEpisodio {
	copia := make(
		[]RegistroEpisodio,
		len(m.episodios),
	)

	copy(copia, m.episodios)

	return copia
}
