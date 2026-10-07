package hospital

import (
	"errors"
	"fmt"
	"strings"
)

type Paciente struct {
	Persona
	nivel           NivelNarcolepsia
	estado          EstadoPaciente
	ubicacionActual string
	medicoAsignado  *Medico
	habitacion      *Habitacion
}

func NuevoPaciente(id, nombre string, edad int, nivel NivelNarcolepsia) (*Paciente, error) {
	persona, err := nuevaPersona(id, nombre, edad)
	if err != nil {
		return nil, fmt.Errorf("paciente inválido: %w", err)
	}
	if nivel < Leve || nivel > Severo {
		return nil, fmt.Errorf("paciente %s inválido: nivel de narcolepsia desconocido", id)
	}
	return &Paciente{Persona: persona, nivel: nivel, estado: Despierto}, nil
}

func (p *Paciente) Nivel() NivelNarcolepsia { return p.nivel }

func (p *Paciente) Estado() EstadoPaciente { return p.estado }

func (p *Paciente) Ubicacion() string { return p.ubicacionActual }

func (p *Paciente) MedicoAsignado() *Medico { return p.medicoAsignado }

func (p *Paciente) Habitacion() *Habitacion { return p.habitacion }

func (p *Paciente) SufrirAtaqueSueno(ubicacion string) error {
	ubicacion = strings.TrimSpace(ubicacion)
	if ubicacion == "" {
		return errors.New("la ubicación no puede estar vacía")
	}
	if p.estado != Despierto {
		return fmt.Errorf("el paciente %s ya está dormido", p.id)
	}
	p.estado = DormidoEnPasillo
	p.ubicacionActual = ubicacion
	return nil
}

func (p *Paciente) Despertar() error {
	if p.estado == Despierto {
		return fmt.Errorf("el paciente %s ya está despierto", p.id)
	}
	if p.habitacion != nil {
		if err := p.habitacion.Liberar(p); err != nil {
			return fmt.Errorf("no se pudo despertar al paciente %s: %w", p.id, err)
		}
		p.habitacion = nil
	}
	p.estado = Despierto
	return nil
}

func (p *Paciente) acostarEnCama(hab *Habitacion) {
	p.habitacion = hab
	p.estado = DormidoEnCama
	p.ubicacionActual = fmt.Sprintf("habitación %d", hab.Numero())
}

func (p *Paciente) asignarMedico(m *Medico) { p.medicoAsignado = m }
