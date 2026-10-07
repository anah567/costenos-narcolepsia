package hospital

import (
	"errors"
	"strings"
)

// Persona guarda los datos básicos que tienen los pacientes y médicos.
type Persona struct {
	id     string
	nombre string
	edad   int
}

// Crea una persona y valida que sus datos sean correctos.
func nuevaPersona(id, nombre string, edad int) (Persona, error) {
	// Quita espacios innecesarios al inicio y al final.
	id, nombre = strings.TrimSpace(id), strings.TrimSpace(nombre)

	// El id y el nombre son obligatorios.
	if id == "" || nombre == "" {
		return Persona{}, errors.New("id y nombre no pueden estar vacíos")
	}

	// No se permite una edad negativa.
	if edad < 0 {
		return Persona{}, errors.New("la edad no puede ser negativa")
	}

	return Persona{
		id:     id,
		nombre: nombre,
		edad:   edad,
	}, nil
}

// Devuelve el id de la persona.
func (p Persona) ID() string {
	return p.id
}

// Devuelve el nombre de la persona.
func (p Persona) Nombre() string {
	return p.nombre
}

// Devuelve la edad de la persona.
func (p Persona) Edad() int {
	return p.edad
}
