package hospital

type Atendedor interface {
	ID() string
	Nombre() string
	Atender(p *Paciente, ubicacion string) (RegistroEpisodio, error)
}

var (
	_ Atendedor = (*Medico)(nil)
	_ Atendedor = (*Camillero)(nil)
)
