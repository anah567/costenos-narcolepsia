package hospital

// Atendedor define lo que debe poder hacer cualquier miembro del personal que pueda atender un ataque de sueño.
// Tanto Medico como Camillero cumplen con estos métodos.
type Atendedor interface {
	ID() string
	Nombre() string
	Atender(p *Paciente, ubicacion string) (RegistroEpisodio, error)
}

// Estas líneas comprueban que Medico y Camillero cumplen con la interfaz Atendedor.
// Si alguno deja de tener uno de los métodos que exige la interfaz,
// Go mostrará un error al compilar.
var (
	_ Atendedor = (*Medico)(nil)
	_ Atendedor = (*Camillero)(nil)
)
