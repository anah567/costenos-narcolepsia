package hospital

// EstadoPaciente define las diferentes situaciones en las que puede estar un paciente.
// Se usa un tipo propio para evitar manejar los estados como números sueltos.
type EstadoPaciente int

// iota asigna valores consecutivos automáticamente:
// Despierto = 0, DormidoEnPasillo = 1 y DormidoEnCama = 2.
const (
	Despierto EstadoPaciente = iota
	DormidoEnPasillo
	DormidoEnCama
)

// String permite mostrar el estado con un texto entendible en vez de mostrar su número.
func (e EstadoPaciente) String() string {
	switch e {
	case Despierto:
		return "Despierto"
	case DormidoEnPasillo:
		return "DormidoEnPasillo"
	case DormidoEnCama:
		return "DormidoEnCama"
	default:
		return "EstadoDesconocido"
	}
}

// NivelNarcolepsia representa qué tan fuerte es la narcolepsia de un paciente.
type NivelNarcolepsia int

// Los niveles también usan iota para asignar sus valores:
// Leve = 0, Moderado = 1 y Severo = 2.
const (
	Leve NivelNarcolepsia = iota
	Moderado
	Severo
)

// String convierte el nivel guardado en un texto fácil de mostrar.
func (n NivelNarcolepsia) String() string {
	switch n {
	case Leve:
		return "Leve"
	case Moderado:
		return "Moderado"
	case Severo:
		return "Severo"
	default:
		return "NivelDesconocido"
	}
}

// EstadoHabitacion indica si una habitación todavía tiene espacio o está ocupada.
type EstadoHabitacion int

// Disponible toma el valor 0 y Ocupada el valor 1 gracias a iota.
const (
	Disponible EstadoHabitacion = iota
	Ocupada
)

// String convierte el estado de la habitación en texto para poder mostrarlo fácilmente.
func (e EstadoHabitacion) String() string {
	switch e {
	case Disponible:
		return "Disponible"
	case Ocupada:
		return "Ocupada"
	default:
		return "EstadoDesconocido"
	}
}
