package hospital

type EstadoPaciente int

const (
	Despierto EstadoPaciente = iota
	DormidoEnPasillo
	DormidoEnCama
)

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

type NivelNarcolepsia int

const (
	Leve NivelNarcolepsia = iota
	Moderado
	Severo
)

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

type EstadoHabitacion int

const (
	Disponible EstadoHabitacion = iota
	Ocupada
)

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
