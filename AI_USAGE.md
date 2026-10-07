# Uso de Inteligencia Artificial

## Herramienta utilizada

Durante el desarrollo del proyecto se utilizó ChatGPT como herramienta de apoyo.

La inteligencia artificial se utilizó principalmente para:

- Analizar y comprender algunos requisitos del ejercicio.
- Resolver dudas sobre conceptos de Go.
- Comprender el uso de composición, interfaces y polimorfismo.
- Revisar partes del código durante el desarrollo.
- Identificar posibles errores.
- Apoyar la revisión de las pruebas.
- Mejorar la organización y legibilidad de algunas partes del código.
- Comprender errores encontrados durante la integración de la interfaz gráfica.
- Revisar la comunicación entre la interfaz y la lógica implementada en Go.

Las sugerencias realizadas por la inteligencia artificial fueron revisadas, ejecutadas y probadas antes de incorporarlas al proyecto.

## Prompts importantes

Algunas de las solicitudes que tuvieron mayor importancia durante el desarrollo fueron:

### Prompt 1

> Explícame cómo puedo aplicar composición, interfaces y polimorfismo en Go para este proyecto del hospital.

Esta solicitud se utilizó como apoyo para comprender cómo representar los conceptos de programación orientada a objetos solicitados en el ejercicio utilizando las características propias de Go.

### Prompt 2

> Revisa esta parte del código y explícame qué hace cada estructura y método de una manera sencilla para poder entenderlo y sustentarlo.

Esta solicitud se utilizó durante la revisión del proyecto para comprender mejor la responsabilidad de las diferentes estructuras y métodos y poder explicar las decisiones tomadas durante el desarrollo.

### Prompt 3

> Ayúdame a revisar por qué esta acción de la interfaz no está funcionando correctamente y qué parte debería verificar.

Esta solicitud se utilizó durante las pruebas de la interfaz para identificar problemas de integración entre JavaScript, Wails y la lógica desarrollada en Go.

Las respuestas obtenidas se utilizaron como orientación y apoyo durante el desarrollo. Las soluciones fueron revisadas y probadas antes de incorporarlas al proyecto.

## Casos en los que fue necesario corregir sugerencias de IA

### Caso 1: referencia a `Greet` después de modificar la aplicación de Wails

Durante la construcción de la interfaz se modificó el código inicial generado por Wails.

En una de las versiones, el frontend todavía intentaba importar y utilizar el método `Greet`, que pertenecía al ejemplo inicial de Wails, aunque ese método ya había sido eliminado del código Go.

Esto produjo un error al compilar el frontend porque se estaba intentando utilizar una función que ya no existía.

Se revisó la integración entre Go y el frontend y se eliminaron las referencias anteriores. Las llamadas fueron reemplazadas por los métodos utilizados realmente por el hospital, entre ellos:

```javascript
ObtenerEstadoHospital()
RegistrarAtaque()
DespertarPaciente()
AsignarHabitacion()
```

Con esta corrección se logró nuevamente la comunicación entre la interfaz y los métodos expuestos desde Go.

Este problema permitió comprender la importancia de mantener sincronizados los métodos disponibles en Go con las funciones utilizadas desde el frontend de Wails.

### Caso 2: uso de `window.confirm()` para despertar un paciente

Durante la implementación de la acción para despertar pacientes se utilizó inicialmente:

```javascript
window.confirm()
```

para solicitar una confirmación antes de realizar la operación.

Durante las pruebas de la interfaz se observó que al presionar el botón para despertar un paciente no se ejecutaba correctamente la acción esperada.

Primero se revisó el método `Despertar()` en Go para comprobar que la lógica encargada de cambiar el estado del paciente y liberar la habitación estuviera funcionando correctamente.

Después de comprobar que la lógica de Go era correcta, se identificó el problema en el comportamiento de la confirmación utilizada desde JavaScript dentro de la aplicación.

La solución fue eliminar esa confirmación y realizar directamente la llamada al método conectado con Go:

```javascript
const mensaje = await DespertarPaciente(pacienteID);
```

Después del cambio se realizaron nuevamente las pruebas y se comprobó que:

- El paciente cambiaba correctamente al estado despierto.
- La habitación quedaba disponible.
- La ubicación y habitación del paciente se actualizaban.
- El Dashboard mostraba los nuevos valores.
- La habitación liberada podía volver a ser asignada.

Este caso mostró la importancia de probar el funcionamiento completo de una solución y no asumir que funciona únicamente porque el código compila.

## Qué aprendí durante el desarrollo

Durante el proyecto aprendí y reforcé diferentes conceptos de Go y programación orientada a objetos.

Entre ellos se encuentran:

- Crear estructuras utilizando `struct`.
- Utilizar métodos con receivers.
- Trabajar con punteros.
- Mantener atributos privados utilizando nombres que comienzan con minúscula.
- Reutilizar información mediante composición y embedding.
- Crear e implementar interfaces.
- Utilizar polimorfismo mediante una interfaz.
- Trabajar con slices.
- Crear constantes tipadas utilizando `iota`.
- Manejar errores utilizando `error`.
- Utilizar `errors.Is` para reconocer errores específicos.
- Crear pruebas utilizando el paquete `testing`.
- Separar la lógica del sistema del escenario ejecutado en `main`.
- Organizar un proyecto Go utilizando diferentes archivos dentro de un mismo paquete.
- Conectar una interfaz gráfica con métodos implementados en Go.
- Comprobar que los cambios realizados en el modelo se reflejen correctamente en la interfaz.

Uno de los conceptos más importantes del proyecto fue comprender cómo una interfaz permite trabajar con diferentes tipos utilizando un comportamiento común.

En este proyecto, tanto `Medico` como `Camillero` cumplen la interfaz `Atendedor`.

Por esta razón, el hospital puede almacenarlos utilizando:

```go
personal []Atendedor
```

y seleccionar diferentes tipos de personal para atender los episodios sin necesitar trabajar directamente con un único tipo concreto.

También fue importante comprender la separación entre la lógica principal del hospital y la interfaz gráfica.

La lógica del sistema permanece dentro del paquete `hospital`, mientras que la interfaz utiliza los métodos disponibles para consultar y modificar el estado del hospital.

De esta manera, la interfaz funciona como una forma de interacción con el sistema y no reemplaza la lógica implementada en Go.

## Pruebas y revisión

Durante el desarrollo, las sugerencias realizadas por la inteligencia artificial fueron revisadas y probadas antes de incorporarlas definitivamente.

Para comprobar el código principal se utilizaron:

```bash
go fmt ./...
go vet ./...
go test ./...
go run .
```

Para comprobar el módulo correspondiente a la interfaz gráfica también se utilizaron:

```bash
go vet ./...
go build .
wails dev
```

Además, se realizaron pruebas manuales de la interfaz para comprobar:

- Navegación entre las diferentes secciones.
- Registro de ataques de sueño.
- Asignación de habitaciones.
- Manejo de un paciente cuando no existen habitaciones disponibles.
- Liberación de una habitación cuando un paciente despierta.
- Reasignación de una habitación disponible.
- Registro y visualización de episodios.
- Actualización de los reportes.
- Actualización de los datos del Dashboard.
- Búsqueda de pacientes.
- Filtros según el estado de los pacientes.
- Funcionamiento de las acciones rápidas.

La inteligencia artificial fue utilizada como una herramienta de apoyo para comprender conceptos, revisar problemas y analizar posibles soluciones. Las modificaciones incorporadas al proyecto fueron ejecutadas y verificadas para comprender su funcionamiento.