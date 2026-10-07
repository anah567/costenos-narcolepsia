import './style.css';

import {
    ObtenerEstadoHospital,
    RegistrarAtaque,
    DespertarPaciente,
    AsignarHabitacion
} from '../wailsjs/go/main/App';


let estado = null;
let seccionActual = 'dashboard';


// ==========================================
// INICIO
// ==========================================

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


// ==========================================
// ACTUALIZAR DATOS
// ==========================================

async function actualizarEstado() {

    estado = await ObtenerEstadoHospital();

    mostrarSeccion(seccionActual);
}


// ==========================================
// ESTRUCTURA
// ==========================================

function construirAplicacion() {

    document.querySelector('#app').innerHTML = `

        <div class="app-shell">

            <aside class="sidebar">

                <div class="brand">

                    <div class="brand-icon">
                        +
                    </div>

                    <div>
                        <h2>Hospital Costeño</h2>
                        <p>Narcolepsia & Sueño</p>
                    </div>

                </div>


                <nav class="nav-menu">

                    ${crearNav('⌂', 'Dashboard', 'dashboard')}

                    ${crearNav('♙', 'Pacientes', 'pacientes')}

                    ${crearNav('✚', 'Personal médico', 'personal')}

                    ${crearNav('▣', 'Habitaciones', 'habitaciones')}

                    ${crearNav('☾', 'Episodios', 'episodios')}

                    ${crearNav('▥', 'Reportes', 'reportes')}

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


function crearNav(icono, nombre, seccion) {

    return `
        <button
            class="nav-item"
            data-section="${seccion}"
        >
            <span class="nav-icon">
                ${icono}
            </span>

            <span>
                ${nombre}
            </span>
        </button>
    `;
}


// ==========================================
// NAVEGACIÓN
// ==========================================

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


// ==========================================
// DASHBOARD
// ==========================================

function renderDashboard() {

    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            <div class="hero">

                <div class="hero-content">

                    <p class="eyebrow">
                        BIENVENIDA AL SISTEMA
                    </p>

                    <h1>
                        Buenos días
                    </h1>

                    <p class="hero-description">
                        Administra pacientes con narcolepsia,
                        episodios de sueño y habitaciones.
                    </p>

                    <div class="hero-actions">

                        <button
                            class="btn-primary"
                            id="nuevoAtaque"
                        >
                            + Registrar ataque de sueño
                        </button>

                        <button
                            class="btn-secondary"
                            id="verHabitaciones"
                        >
                            Ver habitaciones
                        </button>

                    </div>

                </div>

                <div class="hero-decoration">
                    ☾
                </div>

            </div>


            <div class="stats-grid">

                ${crearStat(
                    '♙',
                    'Pacientes',
                    estado.pacientes.length,
                    'pink'
                )}

                ${crearStat(
                    '☾',
                    'Dormidos',
                    estado.dormidos,
                    'purple'
                )}

                ${crearStat(
                    '▣',
                    'Habitaciones libres',
                    estado.habitacionesLibres,
                    'green'
                )}

                ${crearStat(
                    '⌁',
                    'Episodios',
                    estado.totalEpisodios,
                    'yellow'
                )}

            </div>


            <div class="section-title">

                <p class="eyebrow">
                    ESTADO ACTUAL
                </p>

                <h2>
                    Estado del hospital
                </h2>

            </div>


            <div class="dashboard-panels">

                <article class="status-card">

                    <div class="status-icon purple">
                        ☾
                    </div>

                    <div>

                        <span>
                            Pacientes dormidos
                        </span>

                        <strong>
                            ${
                                estado.dormidos === 0
                                    ? 'No hay pacientes dormidos'
                                    : `${estado.dormidos} paciente(s) dormido(s)`
                            }
                        </strong>

                    </div>

                </article>


                <article class="status-card">

                    <div class="status-icon green">
                        ▣
                    </div>

                    <div>

                        <span>
                            Disponibilidad
                        </span>

                        <strong>
                            ${estado.habitacionesLibres}
                            de
                            ${estado.habitaciones.length}
                            habitaciones disponibles
                        </strong>

                    </div>

                </article>

            </div>


            ${
                estado.pacientesEnPasillo > 0
                    ? `
                        <div class="hallway-alert">

                            <strong>
                                ⚠ Atención
                            </strong>

                            <span>
                                ${estado.pacientesEnPasillo}
                                paciente(s) dormido(s)
                                permanecen en el pasillo.
                            </span>

                        </div>
                    `
                    : ''
            }

        </section>
    `;


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
}


function crearStat(icono, titulo, valor, clase) {

    return `
        <article class="stat-card ${clase}">

            <div class="stat-icon">
                ${icono}
            </div>

            <div>
                <p>${titulo}</p>
                <strong>${valor}</strong>
            </div>

        </article>
    `;
}


// ==========================================
// PACIENTES
// ==========================================

function renderPacientes() {

    const tarjetas = estado.pacientes
        .map(p => {

            const despierto =
                p.estado === 'Despierto';

            const pasillo =
                p.estado === 'DormidoEnPasillo';

            return `

                <article class="patient-card">

                    <div class="patient-header">

                        <div class="patient-avatar">
                            ${p.nombre.charAt(0)}
                        </div>

                        <div class="patient-name">

                            <h3>
                                ${p.nombre}
                            </h3>

                            <span>
                                ${p.id}
                            </span>

                        </div>

                        <span
                            class="level-badge ${p.nivel.toLowerCase()}"
                        >
                            ${p.nivel}
                        </span>

                    </div>


                    <div class="patient-data">

                        <div>
                            <span>Edad</span>
                            <strong>
                                ${p.edad} años
                            </strong>
                        </div>

                        <div>
                            <span>Estado</span>
                            <strong>
                                ${p.estado}
                            </strong>
                        </div>

                        <div>
                            <span>Ubicación</span>
                            <strong>
                                ${p.ubicacion}
                            </strong>
                        </div>

                        <div>
                            <span>Habitación</span>
                            <strong>
                                ${p.habitacion}
                            </strong>
                        </div>

                    </div>


                    <div class="patient-actions">

                        ${
                            despierto
                                ? `
                                    <button
                                        class="patient-primary btn-ataque"
                                        data-id="${p.id}"
                                        data-nombre="${p.nombre}"
                                    >
                                        ☾ Registrar ataque
                                    </button>
                                `
                                : `
                                    <button
                                        class="patient-primary btn-despertar"
                                        data-id="${p.id}"
                                        data-nombre="${p.nombre}"
                                    >
                                        ☀ Despertar
                                    </button>
                                `
                        }


                        ${
                            pasillo
                                ? `
                                    <button
                                        class="btn-asignar"
                                        data-id="${p.id}"
                                        data-nombre="${p.nombre}"
                                    >
                                        ▣ Asignar habitación
                                    </button>
                                `
                                : ''
                        }

                    </div>

                </article>
            `;

        })
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            ${crearEncabezado(
                'GESTIÓN DE PACIENTES',
                'Pacientes',
                'Consulta y administra el estado de los pacientes admitidos.'
            )}

            <div class="page-counter">
                ${estado.pacientes.length}
                pacientes registrados
            </div>

            <div class="patients-grid">
                ${tarjetas}
            </div>

        </section>
    `;


    document
        .querySelectorAll('.btn-ataque')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => abrirModalAtaque(
                    boton.dataset.id,
                    boton.dataset.nombre
                )
            );

        });


    document
        .querySelectorAll('.btn-despertar')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => despertar(
                    boton.dataset.id,
                    boton.dataset.nombre
                )
            );

        });


    document
        .querySelectorAll('.btn-asignar')
        .forEach(boton => {

            boton.addEventListener(
                'click',
                () => asignarHabitacion(
                    boton.dataset.id
                )
            );

        });
}


// ==========================================
// PERSONAL
// ==========================================

function renderPersonal() {

    const tarjetas = estado.personal
        .map(persona => `

            <article class="staff-card">

                <div class="staff-avatar">
                    ${obtenerIniciales(persona.nombre)}
                </div>

                <span class="staff-role">
                    ${persona.cargo}
                </span>

                <h3>
                    ${persona.nombre}
                </h3>

                <p>
                    ${persona.id}
                </p>

                <div class="staff-specialty">
                    ${persona.especialidad}
                </div>

                <div class="staff-episodes">
                    ${persona.episodios}
                    episodio(s) atendido(s)
                </div>

            </article>

        `)
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            ${crearEncabezado(
                'EQUIPO HOSPITALARIO',
                'Personal médico',
                'Personal encargado de atender los episodios de sueño.'
            )}

            <div class="placeholder-grid">
                ${tarjetas}
            </div>

        </section>
    `;
}


// ==========================================
// HABITACIONES
// ==========================================

function renderHabitaciones() {

    const tarjetas = estado.habitaciones
        .map(h => `

            <article
                class="room-card ${
                    h.disponible
                        ? ''
                        : 'room-occupied'
                }"
            >

                <div class="room-top">

                    <div class="room-icon">
                        ▣
                    </div>

                    <span
                        class="${
                            h.disponible
                                ? 'available-badge'
                                : 'occupied-badge'
                        }"
                    >
                        ${
                            h.disponible
                                ? 'Disponible'
                                : 'Ocupada'
                        }
                    </span>

                </div>

                <p>
                    HABITACIÓN
                </p>

                <h2>
                    ${h.numero}
                </h2>

                <div class="room-footer">

                    <span>
                        ${
                            h.disponible
                                ? 'Sin ocupante'
                                : 'Paciente'
                        }
                    </span>

                    <strong>
                        ${h.ocupante}
                    </strong>

                </div>

            </article>

        `)
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            ${crearEncabezado(
                'GESTIÓN DE RECURSOS',
                'Habitaciones',
                'Consulta la disponibilidad y ocupación de las habitaciones.'
            )}

            <div class="rooms-grid">
                ${tarjetas}
            </div>

        </section>
    `;
}


// ==========================================
// EPISODIOS
// ==========================================

function renderEpisodios() {

    if (estado.episodios.length === 0) {

        document.querySelector('#contenido').innerHTML = `

            <section class="page">

                ${crearEncabezado(
                    'SEGUIMIENTO CLÍNICO',
                    'Historial de episodios',
                    'Ataques de sueño registrados y personal que realizó la atención.'
                )}

                <div class="empty-state">

                    <div>☾</div>

                    <h3>
                        No hay episodios registrados
                    </h3>

                    <p>
                        Cuando un paciente sufra un ataque,
                        el episodio aparecerá aquí.
                    </p>

                    <button
                        class="btn-primary"
                        id="irPacientes"
                    >
                        Registrar primer ataque
                    </button>

                </div>

            </section>
        `;


        document
            .querySelector('#irPacientes')
            .addEventListener(
                'click',
                () => mostrarSeccion('pacientes')
            );

        return;
    }


    const filas = [...estado.episodios]
        .reverse()
        .map(e => `

            <article class="episode-card">

                <div class="episode-icon">
                    ☾
                </div>

                <div class="episode-info">

                    <strong>
                        ${e.id}
                    </strong>

                    <p>
                        ${e.resumen}
                    </p>

                </div>

            </article>

        `)
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            ${crearEncabezado(
                'SEGUIMIENTO CLÍNICO',
                'Historial de episodios',
                'Ataques de sueño registrados y personal que realizó la atención.'
            )}

            <div class="episodes-list">
                ${filas}
            </div>

        </section>
    `;
}


// ==========================================
// REPORTES
// ==========================================

function renderReportes() {

    const severos = estado.severos
        .map(p => `

            <div class="severe-row">

                <span>
                    ${p.id} · ${p.nombre}
                </span>

                <strong>
                    ${p.episodios}
                    episodio(s) hoy
                </strong>

            </div>

        `)
        .join('');


    const medicos = estado.personal
        .filter(p => p.cargo === 'Médico')
        .map(m => `

            <div class="severe-row">

                <span>
                    ${m.nombre}
                </span>

                <strong>
                    ${m.episodios}
                    episodio(s)
                </strong>

            </div>

        `)
        .join('');


    document.querySelector('#contenido').innerHTML = `

        <section class="page">

            ${crearEncabezado(
                'CONSULTAS DEL HOSPITAL',
                'Reportes',
                'Consultas generadas directamente desde el modelo del hospital.'
            )}


            <div class="reports-grid">

                <article class="report-card pink">

                    <div class="report-icon">
                        ⚠
                    </div>

                    <p>
                        Pacientes en pasillo
                    </p>

                    <strong>
                        ${estado.pacientesEnPasillo}
                        paciente(s)
                    </strong>

                </article>


                <article class="report-card green">

                    <div class="report-icon">
                        ▣
                    </div>

                    <p>
                        Camas disponibles
                    </p>

                    <strong>
                        ${estado.habitacionesLibres}
                        de
                        ${estado.habitaciones.length}
                    </strong>

                </article>


                <article class="report-detail purple">

                    <p class="eyebrow">
                        EPISODIOS POR MÉDICO
                    </p>

                    <h3>
                        Historial de atención
                    </h3>

                    ${medicos}

                </article>


                <article class="report-detail yellow">

                    <p class="eyebrow">
                        NARCOLEPSIA SEVERA
                    </p>

                    <h3>
                        Episodios registrados hoy
                    </h3>

                    ${severos}

                </article>

            </div>

        </section>
    `;
}


// ==========================================
// MODAL ATAQUE
// ==========================================

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


// ==========================================
// DESPERTAR
// ==========================================

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


// ==========================================
// ASIGNAR HABITACIÓN
// ==========================================

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


// ==========================================
// TOAST
// ==========================================

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


// ==========================================
// UTILIDADES
// ==========================================

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


// ==========================================
// INICIAR
// ==========================================

iniciarAplicacion();