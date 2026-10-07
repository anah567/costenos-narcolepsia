package hospital

import (
	"errors"
	"fmt"
)

type Habitacion struct {
	numero    int
	capacidad int
	estado    EstadoHabitacion
	ocupantes []*Paciente
}

func NuevaHabitacion(numero, capacidad int) (*Habitacion, error) {
	if numero <= 0 {
		return nil, errors.New(
			"el número de habitación debe ser positivo",
		)
	}

	if capacidad <= 0 {
		return nil, fmt.Errorf(
			"habitación %d: la capacidad debe ser positiva",
			numero,
		)
	}

	return &Habitacion{
		numero:    numero,
		capacidad: capacidad,
		estado:    Disponible,
	}, nil
}

func (hab *Habitacion) Numero() int {
	return hab.numero
}

func (hab *Habitacion) Capacidad() int {
	return hab.capacidad
}

func (hab *Habitacion) Estado() EstadoHabitacion {
	return hab.estado
}

func (hab *Habitacion) Ocupantes() []*Paciente {
	copia := make([]*Paciente, len(hab.ocupantes))
	copy(copia, hab.ocupantes)

	return copia
}

func (hab *Habitacion) EstaDisponible() bool {
	return hab.estado == Disponible
}

func (hab *Habitacion) Ocupar(p *Paciente) error {
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	if !hab.EstaDisponible() {
		return fmt.Errorf(
			"la habitación %d está llena",
			hab.numero,
		)
	}

	if hab.indiceDe(p) != -1 {
		return fmt.Errorf(
			"el paciente %s ya está en la habitación %d",
			p.ID(),
			hab.numero,
		)
	}

	hab.ocupantes = append(hab.ocupantes, p)
	hab.actualizarEstado()

	return nil
}

func (hab *Habitacion) Liberar(p *Paciente) error {
	i := hab.indiceDe(p)

	if i == -1 {
		return fmt.Errorf(
			"el paciente no está en la habitación %d",
			hab.numero,
		)
	}

	hab.ocupantes = append(
		hab.ocupantes[:i],
		hab.ocupantes[i+1:]...,
	)

	hab.actualizarEstado()

	return nil
}

func (hab *Habitacion) indiceDe(p *Paciente) int {
	for i, ocupante := range hab.ocupantes {
		if ocupante == p {
			return i
		}
	}

	return -1
}

func (hab *Habitacion) actualizarEstado() {
	if len(hab.ocupantes) >= hab.capacidad {
		hab.estado = Ocupada
	} else {
		hab.estado = Disponible
	}
}
