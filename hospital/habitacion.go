package hospital

import (
	"errors"
	"fmt"
)

// Habitacion guarda la información de una habitación del hospital.
// También lleva el control de su capacidad, estado y pacientes que la ocupan.
type Habitacion struct {
	numero    int
	capacidad int
	estado    EstadoHabitacion
	ocupantes []*Paciente
}

// NuevaHabitacion crea una habitación y valida que el número y la capacidad sean válidos.
// Toda habitación nueva comienza en estado Disponible.
func NuevaHabitacion(numero, capacidad int) (*Habitacion, error) {
	// El número de la habitación debe ser mayor que cero.
	if numero <= 0 {
		return nil, errors.New(
			"el número de habitación debe ser positivo",
		)
	}

	// La habitación debe permitir por lo menos un paciente.
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

// Devuelve el número de la habitación.
func (hab *Habitacion) Numero() int {
	return hab.numero
}

// Devuelve la cantidad máxima de pacientes que caben en la habitación.
func (hab *Habitacion) Capacidad() int {
	return hab.capacidad
}

// Devuelve el estado actual de la habitación: Disponible u Ocupada.
func (hab *Habitacion) Estado() EstadoHabitacion {
	return hab.estado
}

// Devuelve una copia de la lista de pacientes que están en la habitación.
// Se hace una copia para no modificar directamente la lista original desde afuera.
func (hab *Habitacion) Ocupantes() []*Paciente {
	copia := make([]*Paciente, len(hab.ocupantes))
	copy(copia, hab.ocupantes)

	return copia
}

// Indica si todavía hay espacio disponible en la habitación.
func (hab *Habitacion) EstaDisponible() bool {
	return hab.estado == Disponible
}

// Ocupar agrega un paciente a la habitación si todavía tiene espacio.
func (hab *Habitacion) Ocupar(p *Paciente) error {
	// Primero se valida que exista un paciente.
	if p == nil {
		return errors.New("el paciente no puede ser nil")
	}

	// Si la habitación ya está llena, no se puede agregar otro paciente.
	if !hab.EstaDisponible() {
		return fmt.Errorf(
			"la habitación %d está llena",
			hab.numero,
		)
	}

	// indiceDe devuelve -1 cuando el paciente no está en la habitación.
	// Si devuelve otro valor, significa que ya está dentro y no se debe repetir.
	if hab.indiceDe(p) != -1 {
		return fmt.Errorf(
			"el paciente %s ya está en la habitación %d",
			p.ID(),
			hab.numero,
		)
	}

	// Agrega el paciente a la lista de ocupantes.
	hab.ocupantes = append(hab.ocupantes, p)

	// Revisa si con el nuevo paciente la habitación quedó llena.
	hab.actualizarEstado()

	return nil
}

// Liberar saca a un paciente de la habitación y vuelve a revisar su disponibilidad.
func (hab *Habitacion) Liberar(p *Paciente) error {
	// Busca la posición del paciente dentro de la lista de ocupantes.
	i := hab.indiceDe(p)

	// Si devuelve -1, significa que el paciente no estaba en esta habitación.
	if i == -1 {
		return fmt.Errorf(
			"el paciente no está en la habitación %d",
			hab.numero,
		)
	}

	// Elimina al paciente de la lista usando todo lo que está antes
	// y después de su posición.
	hab.ocupantes = append(
		hab.ocupantes[:i],
		hab.ocupantes[i+1:]...,
	)

	// Al salir un paciente, la habitación puede volver a estar disponible.
	hab.actualizarEstado()

	return nil
}

// Busca al paciente en la lista de ocupantes.
// Devuelve su posición si lo encuentra y -1 si no está.
func (hab *Habitacion) indiceDe(p *Paciente) int {
	for i, ocupante := range hab.ocupantes {
		if ocupante == p {
			return i
		}
	}

	return -1
}

// Actualiza el estado según la cantidad de pacientes y la capacidad.
// Si llegó a su capacidad máxima queda Ocupada; si todavía hay espacio queda Disponible.
func (hab *Habitacion) actualizarEstado() {
	if len(hab.ocupantes) >= hab.capacidad {
		hab.estado = Ocupada
	} else {
		hab.estado = Disponible
	}
}
