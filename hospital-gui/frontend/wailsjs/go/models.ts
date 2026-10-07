export namespace main {
	
	export class EpisodioDTO {
	    id: string;
	    resumen: string;
	
	    static createFrom(source: any = {}) {
	        return new EpisodioDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.resumen = source["resumen"];
	    }
	}
	export class SeveroDTO {
	    id: string;
	    nombre: string;
	    episodios: number;
	
	    static createFrom(source: any = {}) {
	        return new SeveroDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nombre = source["nombre"];
	        this.episodios = source["episodios"];
	    }
	}
	export class PersonalDTO {
	    id: string;
	    nombre: string;
	    cargo: string;
	    especialidad: string;
	    episodios: number;
	
	    static createFrom(source: any = {}) {
	        return new PersonalDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nombre = source["nombre"];
	        this.cargo = source["cargo"];
	        this.especialidad = source["especialidad"];
	        this.episodios = source["episodios"];
	    }
	}
	export class HabitacionDTO {
	    numero: number;
	    estado: string;
	    disponible: boolean;
	    ocupante: string;
	
	    static createFrom(source: any = {}) {
	        return new HabitacionDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.numero = source["numero"];
	        this.estado = source["estado"];
	        this.disponible = source["disponible"];
	        this.ocupante = source["ocupante"];
	    }
	}
	export class PacienteDTO {
	    id: string;
	    nombre: string;
	    edad: number;
	    nivel: string;
	    estado: string;
	    ubicacion: string;
	    habitacion: string;
	
	    static createFrom(source: any = {}) {
	        return new PacienteDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nombre = source["nombre"];
	        this.edad = source["edad"];
	        this.nivel = source["nivel"];
	        this.estado = source["estado"];
	        this.ubicacion = source["ubicacion"];
	        this.habitacion = source["habitacion"];
	    }
	}
	export class EstadoHospitalDTO {
	    nombre: string;
	    pacientes: PacienteDTO[];
	    habitaciones: HabitacionDTO[];
	    personal: PersonalDTO[];
	    episodios: EpisodioDTO[];
	    severos: SeveroDTO[];
	    pacientesEnPasillo: number;
	    dormidos: number;
	    habitacionesLibres: number;
	    totalEpisodios: number;
	
	    static createFrom(source: any = {}) {
	        return new EstadoHospitalDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nombre = source["nombre"];
	        this.pacientes = this.convertValues(source["pacientes"], PacienteDTO);
	        this.habitaciones = this.convertValues(source["habitaciones"], HabitacionDTO);
	        this.personal = this.convertValues(source["personal"], PersonalDTO);
	        this.episodios = this.convertValues(source["episodios"], EpisodioDTO);
	        this.severos = this.convertValues(source["severos"], SeveroDTO);
	        this.pacientesEnPasillo = source["pacientesEnPasillo"];
	        this.dormidos = source["dormidos"];
	        this.habitacionesLibres = source["habitacionesLibres"];
	        this.totalEpisodios = source["totalEpisodios"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	

}

