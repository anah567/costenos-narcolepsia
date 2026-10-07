package hospital

import (
	"errors"
	"strings"
)

type Persona struct {
	id     string
	nombre string
	edad   int
}

func nuevaPersona(id, nombre string, edad int) (Persona, error) {
	id, nombre = strings.TrimSpace(id), strings.TrimSpace(nombre)

	if id == "" || nombre == "" {
		return Persona{}, errors.New("id y nombre no pueden estar vacíos")
	}

	if edad < 0 {
		return Persona{}, errors.New("la edad no puede ser negativa")
	}

	return Persona{
		id:     id,
		nombre: nombre,
		edad:   edad,
	}, nil
}

func (p Persona) ID() string {
	return p.id
}

func (p Persona) Nombre() string {
	return p.nombre
}

func (p Persona) Edad() int {
	return p.edad
}
