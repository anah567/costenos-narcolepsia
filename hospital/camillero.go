package hospital

import (
	"errors"
	"fmt"
)

type Camillero struct {
	Persona
	episodios []RegistroEpisodio
}

func NuevoCamillero(
	id string,
	nombre string,
	edad int,
) (*Camillero, error) {

	persona, err := nuevaPersona(id, nombre, edad)
	if err != nil {
		return nil, fmt.Errorf(
			"camillero inválido: %w",
			err,
		)
	}

	return &Camillero{
		Persona: persona,
	}, nil
}

func (c *Camillero) Atender(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	if p == nil {
		return RegistroEpisodio{}, errors.New(
			"el paciente no puede ser nil",
		)
	}

	if p.Estado() == Despierto {
		return RegistroEpisodio{}, fmt.Errorf(
			"el paciente %s está despierto, no hay que moverlo",
			p.ID(),
		)
	}

	registro := nuevoRegistroEpisodio(
		p,
		c,
		ubicacion,
	)

	c.episodios = append(
		c.episodios,
		registro,
	)

	return registro, nil
}

func (c *Camillero) MisEpisodios() []RegistroEpisodio {
	copia := make(
		[]RegistroEpisodio,
		len(c.episodios),
	)

	copy(copia, c.episodios)

	return copia
}
