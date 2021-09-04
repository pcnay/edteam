package main

import (
	"fmt"
)

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// Para convertilo a método de la -Struct-
// Se declara por fuera de la -Struct-
// No se utilizan las palabras reservas "this", "self"
// El argumento receptor debe ser claro, como c Course
// Se movera a otros archivo, para manejarlo como paquetes.

func (cursos Course) ImprimeClases() {
	text := "Las clases son :"
	for _, clases := range cursos.Classes {
		text += clases + ", "
	}
	fmt.Printf("Imprimiendo las clases : %v \n", text[:len(text)-2])
}

/* Es una funcion normal.
func ImprimeClases(c Course) {
	text := "Las clases son : "
	for _, class := range c.Classes {
		text += class + ", "
	}
	// text[:len(text)-2] = Para borrar la , y el espacio del último
	fmt.Printf("Imprimiendo las Clases : %v \n", text[:len(text)-2])

}
*/

func main() {
	// Struct = Coleccion de datos
	// Instanciando la "struct"
	Go := Course{

		// Existe otra manera de hacerlo sin utilizar los campos, pero se deben de agregar  en el mismo orden que se definio.
		"Go Desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",

			/*
				Name:    "Go Desde Cero",
				Price:   12.34,
				IsFree:  false,
				UserIDs: []uint{12, 56, 89},
				Classes: map[uint]string{
					1: "Introduccion",
					2: "Estructuras",
					3: "Maps",
			*/
		},
	}

	// Definiendo otra instancia.
	css := Course{
		Name:   "CSS Desde Cero",
		IsFree: true,
	}

	// Otra forma de instanciar la "Struct"
	js := Course{}
	js.Name = "JS desde Cero"
	js.UserIDs = []uint{12, 67}

	//ImprimeClases(Go)
	//ImprimeClases(css)

	// Para llamar al método recien creado.
	Course.ImprimeClases(Go)
	Course.ImprimeClases(css)

}
