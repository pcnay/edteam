package main

import (
	"fmt"
)

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
	fmt.Println(css)

	// Otra forma de instanciar la "Struct"
	js := Course{}
	js.Name = "JS desde Cero"
	js.UserIDs = []uint{12, 67}

	// Para llamar al método recien creado.
	//Course.ImprimeClases(Go)
	//Course.ImprimeClases(css)

	// Utilizando el nuevo método.
	// No cambie el importe, ya que los métodos tienen el mismo comportamiento que las funciones
	// Se envia una copia del valor, se tiene que enviar como puntero
	// Se igual se debe modificar la instancia, debe ser como puntero, pasando la dirección del puntero
	Go.CambiarPrecio(70.5)
	fmt.Printf("Imprimiendo el nuevo precio %v \n ", Go.Price)

}
