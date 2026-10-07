package hospital

import (
	"errors"
	"fmt"
	"strings"
)

// Paciente guarda la información propia de un paciente.
// Persona está incluida para reutilizar id, nombre y edad sin repetirlos.
type Paciente struct {
	Persona
	nivel           NivelNarcolepsia
	estado          EstadoPaciente
	ubicacionActual string
	medicoAsignado  *Medico
	habitacion      *Habitacion
}

// NuevoPaciente crea un paciente con sus datos básicos y su nivel de narcolepsia.
// Al crearse, el paciente siempre empieza en estado Despierto.
func NuevoPaciente(id, nombre string, edad int, nivel NivelNarcolepsia) (*Paciente, error) {

	// Reutiliza la creación de Persona para validar id, nombre y edad.
	persona, err := nuevaPersona(id, nombre, edad)
	if err != nil {
		return nil, fmt.Errorf("paciente inválido: %w", err)
	}

	// Verifica que el nivel recibido sea Leve, Moderado o Severo.
	if nivel < Leve || nivel > Severo {
		return nil, fmt.Errorf("paciente %s inválido: nivel de narcolepsia desconocido", id)
	}

	return &Paciente{
		Persona: persona,
		nivel:   nivel,
		estado:  Despierto,
	}, nil
}

// Devuelve el nivel de narcolepsia del paciente.
func (p *Paciente) Nivel() NivelNarcolepsia {
	return p.nivel
}

// Devuelve el estado actual del paciente.
func (p *Paciente) Estado() EstadoPaciente {
	return p.estado
}

// Devuelve la ubicación actual del paciente.
func (p *Paciente) Ubicacion() string {
	return p.ubicacionActual
}

// Devuelve el médico que tiene asignado el paciente.
// Si todavía no tiene médico, devuelve nil.
func (p *Paciente) MedicoAsignado() *Medico {
	return p.medicoAsignado
}

// Devuelve la habitación donde está el paciente.
// Si está en el pasillo o no tiene habitación, devuelve nil.
func (p *Paciente) Habitacion() *Habitacion {
	return p.habitacion
}

// SufrirAtaqueSueno registra que el paciente sufrió un ataque de sueño.
// Cuando ocurre, queda dormido en el pasillo hasta que el hospital pueda asignarle una habitación.
func (p *Paciente) SufrirAtaqueSueno(ubicacion string) error {

	// Limpia espacios innecesarios de la ubicación.
	ubicacion = strings.TrimSpace(ubicacion)

	// No se puede registrar un ataque sin saber dónde ocurrió.
	if ubicacion == "" {
		return errors.New("la ubicación no puede estar vacía")
	}

	// Solo un paciente despierto puede sufrir un nuevo ataque de sueño.
	if p.estado != Despierto {
		return fmt.Errorf("el paciente %s ya está dormido", p.id)
	}

	// El paciente primero queda en el pasillo y se guarda dónde ocurrió el ataque.
	p.estado = DormidoEnPasillo
	p.ubicacionActual = ubicacion

	return nil
}

// Despertar cambia al paciente nuevamente a Despierto.
// Si estaba en una habitación, primero libera esa habitación para que otro paciente pueda usarla.
func (p *Paciente) Despertar() error {
	// Evita intentar despertar a alguien que ya está despierto.
	if p.estado == Despierto {
		return fmt.Errorf("el paciente %s ya está despierto", p.id)
	}

	// Si tenía una habitación, se libera antes de cambiar su estado.
	if p.habitacion != nil {
		if err := p.habitacion.Liberar(p); err != nil {
			return fmt.Errorf("no se pudo despertar al paciente %s: %w", p.id, err)
		}

		// El paciente deja de tener una habitación asignada.
		p.habitacion = nil
	}

	// El paciente vuelve al estado despierto.
	p.estado = Despierto

	// Al despertar, ya no permanece en la ubicación donde estaba dormido.
	p.ubicacionActual = ""

	return nil
}

// acostarEnCama guarda la habitación asignada y cambia el estado del paciente.
// Este método se usa cuando el hospital logra encontrarle una habitación disponible.
func (p *Paciente) acostarEnCama(hab *Habitacion) {
	p.habitacion = hab
	p.estado = DormidoEnCama
	p.ubicacionActual = fmt.Sprintf("habitación %d", hab.Numero())
}

// Guarda el médico que quedó asignado al paciente.
func (p *Paciente) asignarMedico(m *Medico) {
	p.medicoAsignado = m
}
