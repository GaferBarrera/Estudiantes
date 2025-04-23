/* Crear un programa para gestionar estudiantes en una universidad. El sistema debe permitir:
Registrar estudiantes con su información básica. (Nombre, Edad, Carrera, Calificaciones) Tip: Las calificaciones que sea un map en dónde el key es la materia y el value es la calificación.
Almacenar sus calificaciones por materia.
Calcular el promedio de calificaciones de cada estudiante.
Determinar si aprueban o no usando condicionales.
Agrupar a los estudiantes por carrera.
Permitir modificar sus calificaciones usando punteros.
Calcular el promedio general de todos en el sistema. */

package main

import "fmt"

type Estudiante struct {
	nombre  string
	edad    uint
	carrera string

	materia map[string]uint
}

func main() {

	estudiantes := []Estudiante{
		{"Angelica", 29, "Logistica", map[string]uint{
			"Matematicas": 45,
			"Fisica":      3,
			"Lenguaje":    98},
		},
		{"Juan", 19, "Admon Empresas", map[string]uint{
			"Matematicas": 56,
			"Fisica":      1,
			"Lenguaje":    86,
			"Sociales":    78},
		},
		{"Pedro", 25, "Ing. en sistemas", map[string]uint{
			"Matematicas": 82,
			"Fisica":      94,
			"Lenguaje":    76},
		},
	}

	fmt.Println("----------------------------------------")
	promedioGeneral(estudiantes)

	fmt.Println("----------------------------------------")
	promedioEstudiante(estudiantes)

	fmt.Println("----------------------------------------")
	ActualizaNota(&estudiantes[0], "Matimaticas", 100)

	fmt.Println("----------------------------------------")
	Aprobado(estudiantes)

}

// Metodo - para calculcar promedio de alumno
func (e *Estudiante) CalcularPromedio() uint {
	var promedioNota uint
	for _, nota := range e.materia {
		promedioNota += nota
	}
	return promedioNota / uint(len(e.materia))
}

// Metodo para calculcar si apruebas o no
func (e *Estudiante) EsAprobado() bool {
	if e.CalcularPromedio() > 60 {
		return true
	} else {
		return false
	}
}

// Funcion para si todos los estudiantes aprobaron
func Aprobado(estudiantes []Estudiante) {
	for _, nota := range estudiantes {
		if nota.EsAprobado() {
			fmt.Printf("El Alumno: %v, ha sido aprobado\n", nota.nombre)
		} else {
			fmt.Printf("El Alumno: %v, ha sido reprobado\n", nota.nombre)
		}
	}
}

// Funcion para calcular el promedio general
func promedioGeneral(estudiantes []Estudiante) {
	var promedioGeneral uint
	for _, nota := range estudiantes {
		promedioGeneral += nota.CalcularPromedio() / uint(len(nota.materia))
	}
	fmt.Printf("El promedio total de todos los estudiantes es de %v\n", promedioGeneral)
}

//Funcion para calular promedio de alumno

func promedioEstudiante(estudiantes []Estudiante) {
	for _, estudiantes := range estudiantes {
		fmt.Printf("El Alumno %v, su promedio es de %v\n", estudiantes.nombre, estudiantes.CalcularPromedio())
	}
}

// No actualizamos la base
func ActualizaNota(estudiante *Estudiante, materia string, nota uint) {
	estudiante.materia[materia] = nota
}

// Cambiar las notas

/*

Forma de consola

func (e *Estudiante) Actualizar(nuevaNota uint) {
	for materia, nota := range e.materia {
		var cambio bool
		fmt.Printf("Desea cambiar nota %v de la materia %v\n", nota, materia)
		fmt.Scanln(&cambio)
		if cambio {
			fmt.Printf("Coloque la nueva nota de la materia: %v\n", materia)
			fmt.Scanln(&nuevaNota)
			nota = nuevaNota
		} else {
			fmt.Printf("NO se realizo cambio la nota de: %v\n", materia)
		}
	}
}

func CambiarNotas(estudiantes *[]Estudiante) {
	for i := range *estudiantes {
		fmt.Printf("El Alumno: %s, desea cambiar sus notas\n", (*estudiantes)[i].nombre)
		var cambio bool
		fmt.Scanln(&cambio)
		if cambio {
			(*estudiantes)[i].Actualizar()

		}
	}
}
*/
