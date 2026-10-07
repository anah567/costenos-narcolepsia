package hospital

import (
	"errors"
	"fmt"
)

// Camillero representa al personal que también puede atender ataques de sueño.
// Incluye Persona para reutilizar id, nombre y edad.
// Guarda los episodios que ha atendido.
type Camillero struct {
	Persona
	episodios []RegistroEpisodio
}

// NuevoCamillero crea un camillero y valida sus datos usando nuevaPersona.
func NuevoCamillero(
	id string,
	nombre string,
	edad int,
) (*Camillero, error) {

	// Reutiliza nuevaPersona para validar el id, nombre y edad.
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

// Atender permite que el camillero atienda un ataque de sueño.
// Este método hace que Camillero también cumpla con la interfaz Atendedor.
func (c *Camillero) Atender(
	p *Paciente,
	ubicacion string,
) (RegistroEpisodio, error) {

	// Primero se valida que exista un paciente.
	if p == nil {
		return RegistroEpisodio{}, errors.New(
			"el paciente no puede ser nil",
		)
	}

	// Si el paciente está despierto, no hay un ataque de sueño que atender.
	if p.Estado() == Despierto {
		return RegistroEpisodio{}, fmt.Errorf(
			"el paciente %s está despierto, no hay que moverlo",
			p.ID(),
		)
	}

	// Crea un registro indicando que este camillero atendió el episodio.
	registro := nuevoRegistroEpisodio(
		p,
		c,
		ubicacion,
	)

	// Guarda el episodio en el historial personal del camillero.
	c.episodios = append(
		c.episodios,
		registro,
	)

	return registro, nil
}

// Devuelve una copia de todos los episodios atendidos por el camillero.
// Se devuelve una copia para no modificar directamente la lista original.
func (c *Camillero) MisEpisodios() []RegistroEpisodio {
	copia := make(
		[]RegistroEpisodio,
		len(c.episodios),
	)

	copy(copia, c.episodios)

	return copia
}
