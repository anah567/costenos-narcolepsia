import './style.css';

import {
    ObtenerEstadoHospital,
    RegistrarAtaque,
    DespertarPaciente,
    AsignarHabitacion
} from '../wailsjs/go/main/App';


let estado = null;
let seccionActual = 'dashboard';

// Iconos de la interfaz
function icono(nombre, clase = '') {

    const iconos = {

        dashboard: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M3 11.5 12 4l9 7.5"></path>
                <path d="M5.5 10.5V20h13v-9.5"></path>
                <path d="M9.5 20v-6h5v6"></path>
            </svg>
        `,

        pacientes: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <circle cx="12" cy="8" r="3.5"></circle>
                <path d="M5.5 20c.5-4 2.7-6 6.5-6s6 2 6.5 6"></path>
            </svg>
        `,

        grupo: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <circle cx="9" cy="8" r="3"></circle>
                <path d="M3.5 19c.4-3.5 2.2-5.3 5.5-5.3s5.1 1.8 5.5 5.3"></path>

                <circle cx="17" cy="9" r="2.4"></circle>
                <path d="M14.5 14.5c.8-.5 1.7-.8 2.7-.8 2.5 0 4 1.7 4.3 4.8"></path>
            </svg>
        `,

        personal: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M8 3v5a4 4 0 0 0 8 0V3"></path>
                <path d="M6 3h4"></path>
                <path d="M14 3h4"></path>
                <path d="M12 12v2a6 6 0 0 0 12 0v-1"></path>
                <circle cx="21" cy="10" r="2"></circle>
            </svg>
        `,

        cama: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M3 5v15"></path>
                <path d="M21 20v-8a3 3 0 0 0-3-3H9v7"></path>
                <path d="M3 16h18"></path>
                <path d="M7 9h2v7H3v-3a4 4 0 0 1 4-4Z"></path>
            </svg>
        `,

        luna: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M20.5 15.5A9 9 0 0 1 8.5 3.5 9 9 0 1 0 20.5 15.5Z"></path>
            </svg>
        `,

        reportes: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M5 20V10"></path>
                <path d="M10 20V4"></path>
                <path d="M15 20v-7"></path>
                <path d="M20 20V7"></path>
            </svg>
        `,

        actividad: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M3 12h4l2.2-6 4.1 12 2.3-6H21"></path>
            </svg>
        `,

        hospital: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <rect x="5" y="4" width="14" height="17" rx="2"></rect>
                <path d="M9 8h6"></path>
                <path d="M12 5v6"></path>
                <path d="M9 21v-5h6v5"></path>
            </svg>
        `,

        rayo: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M13 2 5 14h6l-1 8 9-13h-6Z"></path>
            </svg>
        `,

        reloj: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <circle cx="12" cy="12" r="9"></circle>
                <path d="M12 7v5l3 2"></path>
            </svg>
        `,

        sol: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <circle cx="12" cy="12" r="4"></circle>
                <path d="M12 2v2"></path>
                <path d="M12 20v2"></path>
                <path d="m4.9 4.9 1.4 1.4"></path>
                <path d="m17.7 17.7 1.4 1.4"></path>
                <path d="M2 12h2"></path>
                <path d="M20 12h2"></path>
                <path d="m4.9 19.1 1.4-1.4"></path>
                <path d="m17.7 6.3 1.4-1.4"></path>
            </svg>
        `,

        flecha: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M5 12h14"></path>
                <path d="m14 7 5 5-5 5"></path>
            </svg>
        `,

        mas: `
            <svg viewBox="0 0 24 24" class="${clase}" aria-hidden="true">
                <path d="M12 5v14"></path>
                <path d="M5 12h14"></path>
            </svg>
        `
    };

    return iconos[nombre] || '';
}
// INICIO

async function iniciarAplicacion() {

    try {

        estado = await ObtenerEstadoHospital();

        construirAplicacion();

        mostrarSeccion('dashboard');

    } catch (error) {

        console.error(error);

        document.querySelector('#app').innerHTML = `
            <div class="error-general">
                <h2>No se pudo iniciar el hospital</h2>
                <p>${error}</p>
            </div>
        `;
    }
}



// ACTUALIZAR DATOS

async function actualizarEstado() {

    estado = await ObtenerEstadoHospital();

    mostrarSeccion(seccionActual);
}


// ESTRUCTURA

function construirAplicacion() {

    document.querySelector('#app').innerHTML = `

        <div class="app-shell">

            <aside class="sidebar">

                <div class="brand">

                    <div class="brand-icon">
                        ${icono('hospital')}
                    </div>

                    <div>
                        <h2>Hospital Costeño</h2>
                        <p>Narcolepsia & Sueño</p>
                    </div>

                </div>


                <nav class="nav-menu">

                    ${crearNav('dashboard', 'Dashboard', 'dashboard')}

                    ${crearNav('pacientes', 'Pacientes', 'pacientes')}

                    ${crearNav('personal', 'Personal médico', 'personal')}

                    ${crearNav('cama', 'Habitaciones', 'habitaciones')}

                    ${crearNav('luna', 'Episodios', 'episodios')}

                    ${crearNav('reportes', 'Reportes', 'reportes')}

                </nav>


                <div class="sidebar-status">

                    <span class="status-dot"></span>

                    <div>
                        <strong>Sistema operativo</strong>
                        <p>Servicios activos</p>
                    </div>

                </div>

            </aside>


            <main class="main-content">

                <header class="topbar">

                    <div>
                        <strong>${estado.nombre}</strong>
                        <p>Sistema de gestión hospitalaria</p>
                    </div>


                    <div class="topbar-profile">

                        <div class="profile-icon">
                            HC
                        </div>

                        <div>
                            <strong>Hospital Costeño</strong>
                            <p>Sistema de atención</p>
                        </div>

                    </div>

                </header>


                <div id="contenido"></div>

            </main>

        </div>
    `;


    document
        .querySelectorAll('.nav-item')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => mostrarSeccion(
                    boton.dataset.section
                )
            );

        });
}


function crearNav(nombreIcono, nombre, seccion) {

    return `
        <button
            class="nav-item"
            data-section="${seccion}"
        >
            <span class="nav-icon">
                ${icono(nombreIcono)}
            </span>

            <span>
                ${nombre}
            </span>
        </button>
    `;
}

// NAVEGACIÓN

function mostrarSeccion(seccion) {

    seccionActual = seccion;

    document
        .querySelectorAll('.nav-item')
        .forEach(boton => {

            boton.classList.toggle(
                'active',
                boton.dataset.section === seccion
            );

        });


    switch (seccion) {

        case 'dashboard':
            renderDashboard();
            break;

        case 'pacientes':
            renderPacientes();
            break;

        case 'personal':
            renderPersonal();
            break;

        case 'habitaciones':
            renderHabitaciones();
            break;

        case 'episodios':
            renderEpisodios();
            break;

        case 'reportes':
            renderReportes();
            break;
    }
}

// DASHBOARD

function renderDashboard() {

    // Habitaciones que se muestran en el inicio
    const habitaciones = (estado.habitaciones || [])
        .map(h => `
            <div class="dashboard-room ${h.disponible ? '' : 'ocupada'}">

                <strong>${h.numero}</strong>

                <span class="dashboard-room-state">
                    <i></i>
                    ${h.disponible ? 'Libre' : 'Ocupada'}
                </span>

                <div class="dashboard-room-icon">
                    ${icono('cama')}
                </div>

                <p>
                    ${h.disponible ? 'Sin paciente' : h.ocupante}
                </p>

            </div>
        `)
        .join('');


    // Últimos episodios registrados
    const episodios = estado.episodios || [];

    const actividad = episodios.length === 0
        ? `
            <div class="dashboard-empty-activity">

                <div class="activity-empty-icon">
                    ${icono('luna')}
                </div>

                <div>
                    <strong>No hay actividad reciente</strong>
                    <p>Aún no se han registrado episodios de sueño.</p>
                </div>

            </div>
        `
        : [...episodios]
            .reverse()
            .slice(0, 3)
            .map(e => `
                <div class="dashboard-activity-item">

                    <div class="activity-icon">
                        ${icono('luna')}
                    </div>

                    <div>
                        <strong>${e.id}</strong>
                        <p>${e.resumen}</p>
                    </div>

                </div>
            `)
            .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page dashboard-page">

            <div class="hero dashboard-hero">

                <div class="hero-content">

                    <p class="eyebrow">
                        BIENVENIDA AL SISTEMA
                    </p>

                    <h1 class="dashboard-greeting">
                        Buenos días

                        <span class="sun">
                            ${icono('sol')}
                        </span>
                    </h1>

                    <p class="hero-description">
                        Administra pacientes con narcolepsia,
                        episodios de sueño y habitaciones.
                    </p>

                </div>


                <div class="hero-actions">

                    <button
                        class="btn-primary dashboard-main-button"
                        id="nuevoAtaque"
                    >
                        <span class="button-icon">
                            ${icono('mas')}
                        </span>

                        Registrar ataque de sueño
                    </button>


                    <button
                        class="btn-secondary dashboard-room-button"
                        id="verHabitaciones"
                    >
                        <span class="button-icon">
                            ${icono('cama')}
                        </span>

                        Ver habitaciones
                    </button>

                </div>

            </div>


            <div class="stats-grid dashboard-stats">

                ${crearStat(
        icono('pacientes'),
        'Pacientes totales',
        estado.pacientes?.length || 0,
        'pacientes registrados',
        'pink'
    )}

                ${crearStat(
        icono('luna'),
        'Pacientes dormidos',
        estado.dormidos || 0,
        'actualmente',
        'purple'
    )}

                ${crearStat(
        icono('cama'),
        'Habitaciones libres',
        estado.habitacionesLibres || 0,
        `de ${estado.habitaciones?.length || 0} habitaciones`,
        'green'
    )}

                ${crearStat(
        icono('actividad'),
        'Episodios de hoy',
        estado.totalEpisodios || 0,
        'episodios registrados',
        'yellow'
    )}

            </div>


            ${estado.pacientesEnPasillo > 0
            ? `
                        <div class="hallway-alert">

                            <strong>Atención</strong>

                            <span>
                                ${estado.pacientesEnPasillo}
                                paciente(s) dormido(s)
                                permanecen en el pasillo.
                            </span>

                        </div>
                    `
            : ''
        }


            <div class="dashboard-main-grid">

                <article class="dashboard-box hospital-status-box">

                    <div class="dashboard-box-title">

                        <div class="dashboard-title-icon pink-icon">
                            ${icono('hospital')}
                        </div>

                        <div>
                            <h2>Estado del hospital</h2>

                            <p>
                                Vista general de las habitaciones y pacientes
                            </p>
                        </div>

                    </div>


                    <div class="dashboard-rooms">
                        ${habitaciones}
                    </div>

                </article>


                <article class="dashboard-box quick-actions-box">

                    <div class="dashboard-box-title">

                        <div class="dashboard-title-icon pink-icon">
                            ${icono('rayo')}
                        </div>

                        <div>
                            <h2>Acciones rápidas</h2>

                            <p>
                                Gestiona las tareas principales del hospital
                            </p>
                        </div>

                    </div>


                    <div class="dashboard-quick-grid">

                        <button
                            class="quick-card quick-pink"
                            id="accionPacientes"
                        >
                            <span class="quick-card-icon">
                                ${icono('pacientes')}
                            </span>

                            <span>
                                <strong>Ver pacientes</strong>

                                <small>
                                    Consultar pacientes registrados
                                </small>
                            </span>

                            <b>
                                ${icono('flecha')}
                            </b>
                        </button>


                        <button
                            class="quick-card quick-purple"
                            id="accionAtaque"
                        >
                            <span class="quick-card-icon">
                                ${icono('luna')}
                            </span>

                            <span>
                                <strong>
                                    Registrar ataque de sueño
                                </strong>

                                <small>
                                    Crear un nuevo episodio
                                </small>
                            </span>

                            <b>
                                ${icono('flecha')}
                            </b>
                        </button>


                        <button
                            class="quick-card quick-green"
                            id="accionHabitacion"
                        >
                            <span class="quick-card-icon">
                                ${icono('cama')}
                            </span>

                            <span>
                                <strong>
                                    Asignar habitación
                                </strong>

                                <small>
                                    Buscar habitación disponible
                                </small>
                            </span>

                            <b>
                                ${icono('flecha')}
                            </b>
                        </button>


                        <button
                            class="quick-card quick-yellow"
                            id="accionDespertar"
                        >
                            <span class="quick-card-icon">
                                ${icono('sol')}
                            </span>

                            <span>
                                <strong>
                                    Despertar paciente
                                </strong>

                                <small>
                                    Cambiar estado del paciente
                                </small>
                            </span>

                            <b>
                                ${icono('flecha')}
                            </b>
                        </button>

                    </div>

                </article>

            </div>


            <article class="dashboard-box recent-activity">

                <div class="recent-activity-header">

                    <div class="dashboard-box-title">

                        <div class="dashboard-title-icon pink-icon">
                            ${icono('reloj')}
                        </div>

                        <div>
                            <h2>Actividad reciente</h2>

                            <p>
                                Últimos episodios y movimientos en el sistema
                            </p>
                        </div>

                    </div>


                    <button
                        class="activity-link"
                        id="verEpisodios"
                    >
                        Ver todos los episodios

                        <span class="activity-arrow">
                            ${icono('flecha')}
                        </span>
                    </button>

                </div>


                <div class="dashboard-activity">
                    ${actividad}
                </div>

            </article>

        </section>
    `;


    // Accesos del dashboard
    document
        .querySelector('#nuevoAtaque')
        .addEventListener(
            'click',
            () => mostrarSeccion('pacientes')
        );

    document
        .querySelector('#verHabitaciones')
        .addEventListener(
            'click',
            () => mostrarSeccion('habitaciones')
        );

    document
        .querySelector('#accionPacientes')
        .addEventListener(
            'click',
            () => mostrarSeccion('pacientes')
        );

    document
        .querySelector('#accionAtaque')
        .addEventListener(
            'click',
            () => mostrarSeccion('pacientes')
        );

    document
        .querySelector('#accionHabitacion')
        .addEventListener(
            'click',
            () => mostrarSeccion('habitaciones')
        );

    document
        .querySelector('#accionDespertar')
        .addEventListener(
            'click',
            () => mostrarSeccion('pacientes')
        );

    document
        .querySelector('#verEpisodios')
        .addEventListener(
            'click',
            () => mostrarSeccion('episodios')
        );
}


function crearStat(
    icono,
    titulo,
    valor,
    descripcion,
    clase
) {

    return `
        <article class="stat-card ${clase}">

            <div class="stat-icon">
                ${icono}
            </div>

            <div class="stat-content">

                <p>${titulo}</p>

                <strong>${valor}</strong>

                <span>
                    ${descripcion}
                </span>

            </div>

        </article>
    `;
}

// PACIENTES

function renderPacientes() {

    const pacientes = estado.pacientes || [];

    document.querySelector('#contenido').innerHTML = `

        <section class="page patients-page">

            <div class="patients-header">

                <div>
                    <p class="eyebrow">
                        GESTIÓN DE PACIENTES
                    </p>

                    <h1>
                        Pacientes
                    </h1>

                    <p class="page-description">
                        Consulta y administra el estado de los pacientes admitidos.
                    </p>
                </div>

            </div>


            <div class="patients-tools">

                <div class="patient-search">

                    <span class="search-icon">
                        <svg viewBox="0 0 24 24" aria-hidden="true">
                            <circle cx="11" cy="11" r="7"></circle>
                            <path d="m20 20-4-4"></path>
                        </svg>
                    </span>

                    <input
                        id="buscarPaciente"
                        type="text"
                        placeholder="Buscar paciente..."
                        autocomplete="off"
                    >

                </div>


                <select id="filtroEstado" class="patient-filter">
                    <option value="todos">
                        Todos los estados
                    </option>

                    <option value="Despierto">
                        Despiertos
                    </option>

                    <option value="DormidoEnCama">
                        Dormidos en habitación
                    </option>

                    <option value="DormidoEnPasillo">
                        Dormidos en pasillo
                    </option>
                </select>

            </div>


            <div class="patients-count">
                <span id="contadorPacientes">
                    ${pacientes.length} pacientes registrados
                </span>
            </div>


            <div
                class="patients-grid patients-reference-grid"
                id="listaPacientes"
            >
                ${crearTarjetasPacientes(pacientes)}
            </div>

        </section>
    `;


    // Buscar pacientes por nombre o identificación
    document
        .querySelector('#buscarPaciente')
        .addEventListener(
            'input',
            filtrarPacientes
        );


    // Filtrar pacientes por su estado
    document
        .querySelector('#filtroEstado')
        .addEventListener(
            'change',
            filtrarPacientes
        );


    activarBotonesPacientes();
}

function crearTarjetasPacientes(pacientes) {

    if (pacientes.length === 0) {

        return `
            <div class="patients-empty">
                No se encontraron pacientes.
            </div>
        `;
    }


    return pacientes
        .map(paciente => {

            const inicial =
                paciente.nombre
                    ? paciente.nombre.charAt(0).toUpperCase()
                    : '?';


            let claseNivel = 'level-moderate';

            if (paciente.nivel === 'Leve') {
                claseNivel = 'level-light';
            }

            if (paciente.nivel === 'Severo') {
                claseNivel = 'level-severe';
            }


            let claseEstado = 'patient-awake';

            if (paciente.estado === 'DormidoEnCama') {
                claseEstado = 'patient-sleeping';
            }

            if (paciente.estado === 'DormidoEnPasillo') {
                claseEstado = 'patient-hallway';
            }


            let nombreEstado = paciente.estado;

            if (paciente.estado === 'DormidoEnCama') {
                nombreEstado = 'Dormido en cama';
            }

            if (paciente.estado === 'DormidoEnPasillo') {
                nombreEstado = 'Dormido en pasillo';
            }


            let acciones = '';


            if (paciente.estado === 'Despierto') {

                acciones = `
                    <button
    class="patient-main-action btn-attack"
    data-id="${paciente.id}"
    data-nombre="${paciente.nombre}"
>
                        <span class="patient-action-icon">
                            ${icono('luna')}
                        </span>

                        Registrar ataque
                    </button>
                `;

            } else {

                acciones = `
                    <button
                        class="patient-main-action btn-wake"
                        data-id="${paciente.id}"
                        data-nombre="${paciente.nombre}"
                    >
                        <span class="patient-action-icon">
                            ${icono('sol')}
                        </span>

                        Despertar
                    </button>
                `;


                if (paciente.estado === 'DormidoEnPasillo') {

                    acciones += `
                        <button
                            class="patient-secondary-action btn-room"
                            data-id="${paciente.id}"
                        >
                            <span class="patient-action-icon">
                                ${icono('cama')}
                            </span>

                            Asignar habitación
                        </button>
                    `;
                }
            }


            return `

                <article
                    class="patient-card patient-reference-card"
                    data-name="${paciente.nombre.toLowerCase()}"
                    data-id="${paciente.id.toLowerCase()}"
                    data-state="${paciente.estado}"
                >

                    <div class="patient-card-header">

                        <div class="patient-avatar">
                            ${inicial}
                        </div>


                        <div class="patient-identity">

                            <strong>
                                ${paciente.nombre}
                            </strong>

                            <span>
                                ${paciente.id}
                            </span>

                        </div>


                        <span class="patient-level ${claseNivel}">
                            ${paciente.nivel}
                        </span>

                    </div>


                    <div class="patient-divider"></div>


                    <div class="patient-information">

                        <div class="patient-info-row">

                            <span>
                                Edad
                            </span>

                            <strong>
                                ${paciente.edad} años
                            </strong>

                        </div>


                        <div class="patient-info-row">

                            <span>
                                Estado
                            </span>

                            <strong class="patient-state ${claseEstado}">
                                ${nombreEstado}
                            </strong>

                        </div>


                        <div class="patient-info-row">

                            <span>
                                Ubicación
                            </span>

                            <strong>
                                ${paciente.ubicacion || '-'}
                            </strong>

                        </div>


                        <div class="patient-info-row">

                            <span>
                                Habitación
                            </span>

                            <strong>
                                ${paciente.habitacion || '-'}
                            </strong>

                        </div>

                    </div>


                    <div class="patient-card-actions">
                        ${acciones}
                    </div>

                </article>
            `;
        })
        .join('');
}


function filtrarPacientes() {

    const texto =
        document
            .querySelector('#buscarPaciente')
            .value
            .trim()
            .toLowerCase();


    const estadoSeleccionado =
        document
            .querySelector('#filtroEstado')
            .value;


    const pacientesFiltrados =
        (estado.pacientes || [])
            .filter(paciente => {

                const coincideBusqueda =
                    paciente.nombre
                        .toLowerCase()
                        .includes(texto)
                    ||
                    paciente.id
                        .toLowerCase()
                        .includes(texto);


                const coincideEstado =
                    estadoSeleccionado === 'todos'
                    ||
                    paciente.estado === estadoSeleccionado;


                return coincideBusqueda && coincideEstado;
            });


    document
        .querySelector('#listaPacientes')
        .innerHTML =
        crearTarjetasPacientes(
            pacientesFiltrados
        );


    document
        .querySelector('#contadorPacientes')
        .textContent =
        `${pacientesFiltrados.length} ${pacientesFiltrados.length === 1
            ? 'paciente'
            : 'pacientes'
        } registrados`;


    activarBotonesPacientes();
}


function activarBotonesPacientes() {

    // Registrar ataque
    document
        .querySelectorAll('.btn-attack')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => {

                    const pacienteID =
                        boton.dataset.id;

                    const pacienteNombre =
                        boton.dataset.nombre;

                    abrirModalAtaque(
                        pacienteID,
                        pacienteNombre
                    );
                }
            );

        });


    // Despertar paciente
    document
        .querySelectorAll('.btn-wake')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => {

                    despertar(
                        boton.dataset.id,
                        boton.dataset.nombre
                    );
                }
            );

        });


    // Asignar habitación
    document
        .querySelectorAll('.btn-room')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => {

                    asignarHabitacion(
                        boton.dataset.id
                    );
                }
            );

        });
}


// PERSONAL

function renderPersonal() {

    const personal = estado.personal || [];

    const tarjetas = personal
        .map(persona => {

            const partesNombre =
                persona.nombre
                    .replace('Dra. ', '')
                    .replace('Dr. ', '')
                    .trim()
                    .split(' ');

            const iniciales =
                partesNombre
                    .slice(0, 2)
                    .map(parte => parte.charAt(0))
                    .join('')
                    .toUpperCase();


            const esMedico =
                persona.cargo === 'Médico';


            const claseCargo =
                esMedico
                    ? 'staff-doctor'
                    : 'staff-orderly';


            const claseAvatar =
                esMedico
                    ? 'staff-avatar-doctor'
                    : 'staff-avatar-orderly';


            const detalleTitulo =
                esMedico
                    ? 'Especialidad'
                    : 'Rol';


            const detalle =
                persona.especialidad || persona.cargo;


            return `

                <article class="staff-card">

                    <div class="staff-card-top">

                        <div class="staff-avatar ${claseAvatar}">
                            ${iniciales}
                        </div>

                        <span class="staff-role ${claseCargo}">
                            ${persona.cargo}
                        </span>

                    </div>


                    <div class="staff-identity">

                        <h2>
                            ${persona.nombre}
                        </h2>

                        <span>
                            ${persona.id}
                        </span>

                    </div>


                    <div class="staff-divider"></div>


                    <div class="staff-details">

                        <div class="staff-detail-row">

                            <div class="staff-detail-icon">
                                ${esMedico
                    ? icono('personal')
                    : icono('cama')
                }
                            </div>

                            <div>
                                <span>
                                    ${detalleTitulo}
                                </span>

                                <strong>
                                    ${detalle}
                                </strong>
                            </div>

                        </div>


                        <div class="staff-detail-row">

                            <div class="staff-detail-icon">
                                ${icono('actividad')}
                            </div>

                            <div>
                                <span>
                                    Episodios atendidos
                                </span>

                                <strong>
                                    ${persona.episodios || 0}
                                </strong>
                            </div>

                        </div>

                    </div>


                    <button
                        class="staff-episodes-button"
                        data-personal="${persona.id}"
                    >
                        <span>
                            ${icono('reportes')}
                        </span>

                        Ver episodios
                    </button>

                </article>
            `;
        })
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page staff-page">

            <div class="staff-header">

                <p class="eyebrow">
                    EQUIPO HOSPITALARIO
                </p>

                <h1>
                    Personal médico
                </h1>

                <p class="page-description">
                    Personal encargado de atender los episodios de sueño.
                </p>

            </div>


            <div class="staff-grid">

                ${tarjetas}

            </div>

        </section>
    `;


    // Ir a los episodios atendidos
    document
        .querySelectorAll('.staff-episodes-button')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => mostrarSeccion('episodios')
            );

        });
}


// HABITACIONES

function renderHabitaciones() {

    const habitaciones =
        estado.habitaciones || [];


    const tarjetas =
        habitaciones
            .map(habitacion => {

                const disponible =
                    habitacion.disponible;


                const claseEstado =
                    disponible
                        ? 'room-available'
                        : 'room-occupied';


                const textoEstado =
                    disponible
                        ? 'Disponible'
                        : 'Ocupada';


                const paciente =
                    habitacion.ocupante &&
                        habitacion.ocupante !== '-'
                        ? habitacion.ocupante
                        : 'Sin paciente';


                const capacidad =
                    disponible
                        ? '0 / 1'
                        : '1 / 1';


                return `

                    <article class="room-card">

                        <div class="room-visual ${claseEstado}">

                            <div class="room-icon">
                                ${icono('cama')}
                            </div>


                            <span class="room-status">

                                <span class="room-status-dot"></span>

                                ${textoEstado}

                            </span>

                        </div>


                        <div class="room-number">

                            <span>
                                HABITACIÓN
                            </span>

                            <strong>
                                ${habitacion.numero}
                            </strong>

                        </div>


                        <div class="room-divider"></div>


                        <div class="room-information">

                            <div class="room-info-row">

                                <div class="room-info-icon">
                                    ${icono('paciente')}
                                </div>


                                <div class="room-info-text">

                                    <span>
                                        Paciente asignado
                                    </span>

                                    <strong>
                                        ${paciente}
                                    </strong>

                                </div>

                            </div>


                            <div class="room-info-row">

                                <div class="room-info-icon">
                                    ${icono('personal')}
                                </div>


                                <div class="room-info-text">

                                    <span>
                                        Capacidad
                                    </span>

                                    <strong>
                                        ${capacidad}
                                    </strong>

                                </div>

                            </div>

                        </div>

                    </article>
                `;
            })
            .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page rooms-page">

            <div class="rooms-header">

                <p class="eyebrow">
                    GESTIÓN DE RECURSOS
                </p>

                <h1>
                    Habitaciones
                </h1>

                <p class="page-description">
                    Consulta la disponibilidad y ocupación de las habitaciones.
                </p>

            </div>


            <div class="rooms-grid">

                ${tarjetas}

            </div>

        </section>
    `;
}

// EPISODIOS

function renderEpisodios() {

    const episodios =
        estado.episodios || [];


    const filas =
        episodios.length > 0

            ? [...episodios]
                .reverse()
                .map(episodio => {

                    const nivelClase =
                        episodio.nivel === 'Severo'
                            ? 'episode-level-severe'
                            : episodio.nivel === 'Moderado'
                                ? 'episode-level-moderate'
                                : 'episode-level-mild';


                    const destino =
                        episodio.habitacion !== '-'
                            ? `${episodio.ubicacion} → ${episodio.habitacion}`
                            : `${episodio.ubicacion} → Sin habitación`;


                    return `

                        <article class="episode-detail-row">

                            <div class="episode-id-section">

                                <div class="episode-icon">
                                    ${icono('luna')}
                                </div>

                                <div>
                                    <strong class="episode-code">
                                        ${episodio.id}
                                    </strong>

                                    <span class="episode-date">
                                        ${episodio.fechaHora}
                                    </span>
                                </div>

                            </div>


                            <div class="episode-column">

                                <div class="episode-column-icon">
                                    ${icono('paciente')}
                                </div>

                                <div>
                                    <span>
                                        Paciente
                                    </span>

                                    <strong>
                                        ${episodio.pacienteId} ·
                                        ${episodio.pacienteNombre}
                                    </strong>
                                </div>

                            </div>


                            <div class="episode-column">

                                <div class="episode-column-icon">
                                    ${icono('cama')}
                                </div>

                                <div>
                                    <span>
                                        Ubicación / habitación
                                    </span>

                                    <strong>
                                        ${destino}
                                    </strong>
                                </div>

                            </div>


                            <div class="episode-column">

                                <div class="episode-column-icon">
                                    ${icono('personal')}
                                </div>

                                <div>
                                    <span>
                                        Atendido por
                                    </span>

                                    <strong>
                                        ${episodio.atendidoPor}
                                    </strong>
                                </div>

                            </div>


                            <span class="
                                episode-level
                                ${nivelClase}
                            ">
                                ${episodio.nivel}
                            </span>

                        </article>
                    `;
                })
                .join('')

            : `

                <div class="episodes-empty">

                    <div class="episodes-empty-icon">
                        ${icono('luna')}
                    </div>

                    <strong>
                        No hay episodios registrados
                    </strong>

                    <span>
                        Los ataques de sueño aparecerán aquí
                        cuando sean registrados.
                    </span>

                </div>
            `;


    document.querySelector('#contenido').innerHTML = `

        <section class="page episodes-page">

            <div class="episodes-header">

                <p class="eyebrow">
                    SEGUIMIENTO CLÍNICO
                </p>

                <h1>
                    Historial de episodios
                </h1>

                <p class="page-description">
                    Ataques de sueño registrados y personal que realizó la atención.
                </p>

            </div>


            <div class="episodes-detail-list">

                ${filas}

            </div>

        </section>
    `;
}

// REPORTES

function renderReportes() {

    const personal =
        estado.personal || [];

    const severos =
        estado.severos || [];

    const habitaciones =
        estado.habitaciones || [];

    const totalHabitaciones =
        habitaciones.length;

    const libres =
        estado.habitacionesLibres || 0;

    const pasillo =
        estado.pacientesEnPasillo || 0;


    // Arma la lista del personal con la cantidad de episodios atendidos.
    const historialPersonal =
        personal
            .map(persona => {

                const iniciales =
                    persona.nombre
                        .replace('Dra. ', '')
                        .replace('Dr. ', '')
                        .split(' ')
                        .map(parte => parte.charAt(0))
                        .slice(0, 2)
                        .join('')
                        .toUpperCase();


                const claseAvatar =
                    persona.cargo === 'Camillero'
                        ? 'report-avatar-orderly'
                        : 'report-avatar-doctor';


                return `

                    <div class="report-person-row">

                        <div class="
                            report-person-avatar
                            ${claseAvatar}
                        ">
                            ${iniciales}
                        </div>


                        <div class="report-person-info">

                            <strong>
                                ${persona.nombre}
                            </strong>

                            <span>
                                ${persona.id}
                            </span>

                        </div>


                        <strong class="report-count">
                            ${persona.episodios}
                            episodio(s)
                        </strong>

                    </div>
                `;
            })
            .join('');


    // Arma el reporte de los pacientes con narcolepsia severa.
    const reporteSeveros =
        severos.length > 0

            ? severos
                .map(paciente => {

                    const inicial =
                        paciente.nombre
                            .charAt(0)
                            .toUpperCase();


                    return `

                        <div class="report-severe-row">

                            <div class="report-severe-avatar">
                                ${inicial}
                            </div>


                            <div class="report-severe-info">

                                <strong>
                                    ${paciente.id} ·
                                    ${paciente.nombre}
                                </strong>

                                <span>
                                    Narcolepsia severa
                                </span>

                            </div>


                            <strong class="report-count">
                                ${paciente.episodios}
                                episodio(s)
                            </strong>

                        </div>
                    `;
                })
                .join('')

            : `

                <div class="report-empty">
                    No hay pacientes con narcolepsia severa.
                </div>
            `;


    document.querySelector('#contenido').innerHTML = `

        <section class="page reports-page">

            <div class="reports-header">

                <p class="eyebrow">
                    CONSULTAS DEL HOSPITAL
                </p>

                <h1>
                    Reportes
                </h1>

                <p class="page-description">
                    Consultas generadas directamente desde el modelo del hospital.
                </p>

            </div>


            <div class="reports-summary-grid">


                <!-- Pacientes en pasillo -->

                <article class="
                    report-summary-card
                    report-summary-hallway
                ">

                   <div class="report-background-icon">
    ${icono('grupo')}
</div>

<div class="report-summary-icon">
    ${icono('grupo')}
</div>


                    <div class="report-summary-content">

                        <span>
                            Pacientes en pasillo
                        </span>

                        <strong>
                            ${pasillo}
                        </strong>

                        <p>
                            paciente(s) actualmente
                        </p>

                    </div>

                </article>


                <!-- Camas disponibles -->

                <article class="
                    report-summary-card
                    report-summary-beds
                ">

                    <div class="report-background-icon">
                        ${icono('cama')}
                    </div>


                    <div class="report-summary-icon">
                        ${icono('cama')}
                    </div>


                    <div class="report-summary-content">

                        <span>
                            Camas disponibles
                        </span>

                        <strong>
                            ${libres} de ${totalHabitaciones}
                        </strong>

                        <p>
                            habitaciones disponibles
                        </p>

                    </div>

                </article>

            </div>


            <div class="reports-detail-grid">


                <!-- Episodios atendidos por el personal -->

                <article class="
                    report-detail-card
                    report-history-card
                ">

                    <div class="report-card-heading">

                        <div class="report-heading-icon">
                            ${icono('luna')}
                        </div>


                        <div>

                            <span class="report-label">
                                EPISODIOS POR PERSONAL
                            </span>

                            <h2>
                                Historial de atención
                            </h2>

                        </div>

                    </div>


                    <p class="report-description">
                        Cantidad de episodios atendidos por cada
                        miembro del personal.
                    </p>


                    <div class="report-rows">

                        ${historialPersonal}

                    </div>

                </article>


                <!-- Pacientes con narcolepsia severa -->

                <article class="
                    report-detail-card
                    report-severe-card
                ">

                    <div class="report-card-heading">

                        <div class="report-heading-icon">
                            ${icono('actividad')}
                        </div>


                        <div>

                            <span class="report-label">
                                NARCOLEPSIA SEVERA
                            </span>

                            <h2>
                                Pacientes severos
                            </h2>

                        </div>

                    </div>


                    <p class="report-description">
                        Pacientes con nivel severo y cantidad
                        de episodios registrados.
                    </p>


                    <div class="report-rows">

                        ${reporteSeveros}

                    </div>

                </article>

            </div>

        </section>
    `;
}

// MODAL ATAQUE

function abrirModalAtaque(
    pacienteID,
    pacienteNombre
) {

    const modal =
        document.createElement('div');

    modal.className =
        'modal-overlay';


    modal.innerHTML = `

        <div class="modal-card">

            <button
                class="modal-close"
                id="cerrarModal"
            >
                ×
            </button>


            <div class="modal-icon">
                ☾
            </div>


            <p class="eyebrow">
                NUEVO EPISODIO
            </p>


            <h2>
                Registrar ataque de sueño
            </h2>


            <p class="modal-description">
                Indica dónde ocurrió el ataque de
                <strong>
                    ${pacienteNombre}
                </strong>.
            </p>


            <div class="modal-patient">

                <div class="modal-avatar">
                    ${pacienteNombre.charAt(0)}
                </div>

                <div>

                    <strong>
                        ${pacienteNombre}
                    </strong>

                    <span>
                        ${pacienteID}
                    </span>

                </div>

            </div>


            <label class="modal-label">
                Ubicación del ataque
            </label>


            <input
                id="ubicacionAtaque"
                class="modal-input"
                placeholder="Ej. cafetería"
                autocomplete="off"
            >


            <p
                id="modalError"
                class="modal-error"
            ></p>


            <div class="modal-actions">

                <button
                    class="btn-secondary"
                    id="cancelarAtaque"
                >
                    Cancelar
                </button>

                <button
                    class="btn-primary"
                    id="confirmarAtaque"
                >
                    Registrar ataque
                </button>

            </div>

        </div>
    `;


    document.body.appendChild(
        modal
    );


    const input =
        modal.querySelector(
            '#ubicacionAtaque'
        );

    const confirmar =
        modal.querySelector(
            '#confirmarAtaque'
        );

    const error =
        modal.querySelector(
            '#modalError'
        );


    input.focus();


    const cerrar = () => {
        modal.remove();
    };


    modal
        .querySelector('#cerrarModal')
        .addEventListener(
            'click',
            cerrar
        );


    modal
        .querySelector('#cancelarAtaque')
        .addEventListener(
            'click',
            cerrar
        );


    confirmar.addEventListener(
        'click',
        async () => {

            const ubicacion =
                input.value.trim();


            if (!ubicacion) {

                error.textContent =
                    'Debes indicar la ubicación.';

                return;
            }


            try {

                confirmar.disabled = true;

                confirmar.textContent =
                    'Registrando...';


                const mensaje =
                    await RegistrarAtaque(
                        pacienteID,
                        ubicacion
                    );


                cerrar();

                await actualizarEstado();

                mostrarToast(
                    mensaje,
                    'success'
                );


            } catch (err) {

                error.textContent =
                    String(err);

                confirmar.disabled = false;

                confirmar.textContent =
                    'Registrar ataque';
            }

        }
    );


    input.addEventListener(
        'keydown',
        event => {

            if (event.key === 'Enter') {
                confirmar.click();
            }

        }
    );
}


// DESPERTAR

async function despertar(
    pacienteID,
    nombre
) {

    try {

        const mensaje =
            await DespertarPaciente(
                pacienteID
            );

        await actualizarEstado();

        mostrarToast(
            mensaje,
            'success'
        );

    } catch (error) {

        console.error(
            'Error al despertar:',
            error
        );

        mostrarToast(
            `No se pudo despertar a ${nombre}: ${String(error)}`,
            'error'
        );
    }
}


// ASIGNAR HABITACIÓN

async function asignarHabitacion(
    pacienteID
) {

    try {

        const mensaje =
            await AsignarHabitacion(
                pacienteID
            );


        await actualizarEstado();

        mostrarToast(
            mensaje,
            'success'
        );


    } catch (error) {

        mostrarToast(
            String(error),
            'error'
        );
    }
}


// TOAST

function mostrarToast(
    mensaje,
    tipo = 'success'
) {

    const anterior =
        document.querySelector(
            '.toast'
        );


    if (anterior) {
        anterior.remove();
    }


    const toast =
        document.createElement('div');


    toast.className =
        `toast ${tipo}`;


    toast.textContent =
        mensaje;


    document.body.appendChild(
        toast
    );


    setTimeout(
        () => toast.classList.add('visible'),
        20
    );


    setTimeout(
        () => {

            toast.classList.remove(
                'visible'
            );

            setTimeout(
                () => toast.remove(),
                250
            );

        },
        3500
    );
}


// UTILIDADES

function crearEncabezado(
    etiqueta,
    titulo,
    descripcion
) {

    return `

        <div class="page-heading">

            <p class="eyebrow">
                ${etiqueta}
            </p>

            <h1>
                ${titulo}
            </h1>

            <p>
                ${descripcion}
            </p>

        </div>
    `;
}


function obtenerIniciales(nombre) {

    return nombre
        .replace('Dra.', '')
        .replace('Dr.', '')
        .trim()
        .split(' ')
        .slice(0, 2)
        .map(parte => parte.charAt(0))
        .join('')
        .toUpperCase();
}


// INICIAR

iniciarAplicacion();