# Hospital de los Costeños con Narcolepsia

Proyecto desarrollado en Go para el ejercicio de Programación Orientada a Objetos.

El sistema representa un hospital encargado de atender pacientes con narcolepsia. Permite registrar pacientes, médicos, camilleros y habitaciones, además de manejar ataques de sueño, asignar habitaciones, liberar habitaciones y registrar los episodios atendidos.

También se desarrolló una interfaz gráfica utilizando Wails para visualizar y administrar la información del hospital.

## Funcionalidades principales

El sistema permite:

- Crear y admitir pacientes con diferentes niveles de narcolepsia.
- Registrar médicos y camilleros.
- Registrar habitaciones con una capacidad determinada.
- Registrar ataques de sueño.
- Asignar habitaciones disponibles a pacientes dormidos.
- Mantener al paciente en el pasillo cuando no existen habitaciones disponibles.
- Liberar una habitación cuando un paciente despierta.
- Reasignar habitaciones que quedan disponibles.
- Consultar los pacientes dormidos en el pasillo.
- Consultar los episodios atendidos por el personal.
- Consultar el estado y ocupación de las habitaciones.
- Generar un reporte de pacientes con narcolepsia severa y sus episodios del día.
- Visualizar la información mediante una interfaz gráfica.

## Estructura del proyecto

```text
costenos-narcolepsia/
├── go.mod
├── main.go
├── hospital/
│   ├── persona.go
│   ├── estados.go
│   ├── paciente.go
│   ├── medico.go
│   ├── camillero.go
│   ├── habitacion.go
│   ├── atendedor.go
│   ├── episodio.go
│   ├── hospital.go
│   └── hospital_test.go
├── hospital-gui/
│   ├── app.go
│   ├── main.go
│   ├── go.mod
│   ├── frontend/
│   │   └── src/
│   │       ├── main.js
│   │       └── style.css
│   └── build/
├── README.md
└── AI_USAGE.md
```

## Conceptos utilizados

### Encapsulamiento

Los atributos principales de las estructuras se mantienen privados utilizando nombres que comienzan con minúscula.

El acceso a estos datos se realiza mediante métodos, evitando modificar directamente el estado interno de las estructuras.

### Composición

`Paciente`, `Medico` y `Camillero` incluyen `Persona`.

Esto permite reutilizar información como ID, nombre y edad sin repetir estos atributos en cada estructura.

### Interfaces y polimorfismo

Se utiliza la interfaz `Atendedor`:

```go
type Atendedor interface {
	ID() string
	Nombre() string
	Atender(p *Paciente, ubicacion string) (RegistroEpisodio, error)
}
```

Esta interfaz es implementada por `Medico` y `Camillero`.

El hospital almacena el personal mediante:

```go
personal []Atendedor
```

Esto permite guardar diferentes tipos de personal en una misma colección y seleccionar cualquiera de ellos para atender un ataque de sueño.

De esta manera se utiliza polimorfismo, ya que el hospital trabaja con la interfaz `Atendedor` sin necesitar conocer específicamente si quien atiende es un médico o un camillero.

### Estados tipados

Se utilizan tipos propios y constantes para representar:

- Estado del paciente.
- Nivel de narcolepsia.
- Estado de las habitaciones.

Esto evita manejar estos valores como números o textos sueltos y permite representar mejor los posibles estados del sistema.

### Manejo de errores

Las operaciones que pueden fallar devuelven valores de tipo `error`.

Por ejemplo, cuando no existe una habitación disponible se utiliza:

```go
ErrSinHabitacion
```

Esto permite identificar específicamente este caso utilizando `errors.Is`.

Cuando ocurre este error, el paciente puede ser atendido y registrar su episodio, pero permanece dormido en el pasillo hasta que una habitación quede disponible.

## Escenario principal

El programa crea inicialmente:

- 2 médicos.
- 1 camillero.
- 3 habitaciones con capacidad para un paciente.
- 5 pacientes.
- Al menos 2 pacientes con narcolepsia severa.

Después se simulan cuatro ataques de sueño en diferentes ubicaciones.

Los primeros tres pacientes reciben una habitación disponible.

Cuando ocurre el cuarto ataque, las tres habitaciones se encuentran ocupadas, por lo que el paciente permanece dormido en el pasillo y se genera el error correspondiente por falta de habitaciones.

Posteriormente, uno de los pacientes se despierta y libera su habitación.

La habitación liberada puede ser asignada al paciente que estaba esperando en el pasillo.

Durante todo el proceso se registran los episodios de sueño y el personal que atendió cada uno.

## Consultas

El programa permite consultar:

1. Pacientes dormidos en el pasillo.
2. Episodios atendidos por cada médico.
3. Disponibilidad y ocupación de las habitaciones.
4. Pacientes con narcolepsia severa y cantidad de episodios del día.

Estas consultas retornan datos y la presentación de los resultados se realiza desde el programa principal o desde la interfaz gráfica.

## Pruebas

El proyecto contiene pruebas automáticas para comprobar el funcionamiento de la lógica principal.

Entre los casos probados se encuentran:

- Asignación de una habitación disponible.
- Comportamiento cuando no existen habitaciones disponibles.
- Consulta de pacientes en el pasillo.
- Liberación y reutilización de una habitación.
- Episodios atendidos por un médico.
- Reporte de pacientes con narcolepsia severa.

Para ejecutar las pruebas desde la carpeta principal:

```bash
go test ./...
```

También se puede verificar el formato y analizar posibles problemas con:

```bash
go fmt ./...
go vet ./...
go test ./...
```

## Ejecución por consola

Desde la carpeta principal del proyecto:

```bash
go run .
```

Esto ejecuta el escenario principal del hospital y muestra los resultados en la terminal.

## Interfaz gráfica

Como funcionalidad adicional se desarrolló una interfaz gráfica utilizando Wails.

La interfaz permite visualizar y administrar:

- Dashboard general del hospital.
- Pacientes.
- Personal médico.
- Habitaciones.
- Episodios.
- Reportes.

También permite realizar acciones como:

- Registrar ataques de sueño.
- Despertar pacientes.
- Asignar habitaciones.
- Buscar pacientes.
- Filtrar pacientes según su estado.
- Consultar los episodios registrados.

La interfaz se comunica con la lógica desarrollada en Go, por lo que las acciones realizadas desde la aplicación utilizan el modelo del hospital.

### Ejecutar la interfaz

Entrar a la carpeta:

```bash
cd hospital-gui
```

Ejecutar:

```bash
wails dev
```

Esto inicia la aplicación en modo de desarrollo.

## Verificaciones realizadas

Antes de finalizar el proyecto se ejecutaron las siguientes verificaciones sobre el código principal:

```bash
go fmt ./...
go vet ./...
go test ./...
```

Para el módulo de la interfaz también se verificó:

```bash
cd hospital-gui
go vet ./...
go build .
```

Además, se realizaron pruebas manuales de la interfaz para comprobar la navegación, registro de ataques, asignación y liberación de habitaciones, búsqueda, filtros, episodios y reportes.

## Tecnologías utilizadas

- Go
- Wails
- JavaScript
- HTML
- CSS
- Biblioteca estándar de Go
- Git
- GitHub