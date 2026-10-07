package hospital

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Lleva la cuenta de los episodios para generar un ID diferente para cada uno.
// atomic permite aumentar el contador de forma segura.
var secuenciaEpisodios atomic.Int64

// RegistroEpisodio guarda toda la información de un ataque de sueño:
// cuándo ocurrió, qué paciente fue atendido, quién lo atendió,
// dónde ocurrió y si terminó en una habitación.
type RegistroEpisodio struct {
	id          string
	fechaHora   time.Time
	paciente    *Paciente
	atendidoPor Atendedor
	ubicacion   string
	habitacion  *Habitacion
}

// Crea el registro de un nuevo episodio con la información del paciente y del atendedor.
func nuevoRegistroEpisodio(
	p *Paciente,
	a Atendedor,
	ubicacion string,
) RegistroEpisodio {

	return RegistroEpisodio{

		// Aumenta el contador y crea IDs como E-001, E-002, E-003...
		id: fmt.Sprintf("E-%03d", secuenciaEpisodios.Add(1)),

		// Guarda la fecha y hora exactas en las que se crea el episodio.
		fechaHora: time.Now(),

		paciente: p,

		// Se guarda como Atendedor porque puede ser un Medico o un Camillero.
		atendidoPor: a,

		// Guarda el lugar donde ocurrió el ataque de sueño.
		ubicacion: ubicacion,

		// Guarda la habitación que recibió el paciente.
		// Puede ser nil si no había habitación disponible.
		habitacion: p.Habitacion(),
	}
}

// Devuelve el identificador del episodio.
func (e RegistroEpisodio) ID() string {
	return e.id
}

// Devuelve la fecha y hora en que ocurrió el episodio.
func (e RegistroEpisodio) FechaHora() time.Time {
	return e.fechaHora
}

// Devuelve el paciente relacionado con el episodio.
func (e RegistroEpisodio) Paciente() *Paciente {
	return e.paciente
}

// Devuelve quién atendió el episodio.
// Puede ser un Medico o un Camillero porque ambos son Atendedor.
func (e RegistroEpisodio) AtendidoPor() Atendedor {
	return e.atendidoPor
}

// Devuelve el lugar donde ocurrió el ataque de sueño.
func (e RegistroEpisodio) Ubicacion() string {
	return e.ubicacion
}

// Devuelve la habitación asignada durante el episodio.
// Si no se consiguió habitación, devuelve nil.
func (e RegistroEpisodio) Habitacion() *Habitacion {
	return e.habitacion
}

// Crea un texto con los datos principales del episodio para poder mostrarlo en pantalla.
func (e RegistroEpisodio) Resumen() string {

	// Por defecto se supone que el paciente quedó en el pasillo.
	destino := "sin habitación (sigue en el pasillo)"

	// Si sí tiene habitación, se muestra su número como destino.
	if e.habitacion != nil {
		destino = fmt.Sprintf(
			"habitación %d",
			e.habitacion.Numero(),
		)
	}

	// Une la fecha, el ID, el paciente, la ubicación, el destino y la persona que atendió en un solo texto.
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
