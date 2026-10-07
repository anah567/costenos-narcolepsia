package hospital

import (
	"fmt"
	"sync/atomic"
	"time"
)

var secuenciaEpisodios atomic.Int64

type RegistroEpisodio struct {
	id          string
	fechaHora   time.Time
	paciente    *Paciente
	atendidoPor Atendedor
	ubicacion   string
	habitacion  *Habitacion
}

func nuevoRegistroEpisodio(
	p *Paciente,
	a Atendedor,
	ubicacion string,
) RegistroEpisodio {

	return RegistroEpisodio{
		id:          fmt.Sprintf("E-%03d", secuenciaEpisodios.Add(1)),
		fechaHora:   time.Now(),
		paciente:    p,
		atendidoPor: a,
		ubicacion:   ubicacion,
		habitacion:  p.Habitacion(),
	}
}

func (e RegistroEpisodio) ID() string {
	return e.id
}

func (e RegistroEpisodio) FechaHora() time.Time {
	return e.fechaHora
}

func (e RegistroEpisodio) Paciente() *Paciente {
	return e.paciente
}

func (e RegistroEpisodio) AtendidoPor() Atendedor {
	return e.atendidoPor
}

func (e RegistroEpisodio) Ubicacion() string {
	return e.ubicacion
}

func (e RegistroEpisodio) Habitacion() *Habitacion {
	return e.habitacion
}

func (e RegistroEpisodio) Resumen() string {
	destino := "sin habitación (sigue en el pasillo)"

	if e.habitacion != nil {
		destino = fmt.Sprintf(
			"habitación %d",
			e.habitacion.Numero(),
		)
	}

	return fmt.Sprintf(
		"%s  %s  %-6s %-24s -> %s  (atendió %s)",
		e.fechaHora.Format("2006-01-02 15:04"),
		e.id,
		e.paciente.ID(),
		e.ubicacion,
		destino,
		e.atendidoPor.Nombre(),
	)
}
