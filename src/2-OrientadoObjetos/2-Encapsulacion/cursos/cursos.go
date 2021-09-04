package cursos

import (
	"fmt"
)

// Se movera al archivo "cursos.go" para que este organizado y como buena practica, se debe tener en un solo sitio la "struct" y sus metodos.

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// No cambie el importe, ya que los métodos tienen el mismo comportamiento que las funciones
// Se envia una copia del valor, se tiene que enviar como puntero
// Con esto Go cuando se llamen los métodos se encarga de desreferenciar el puntero es decir *, & con solo llamar el nombre de la instancia de la clase
// Es mas legible trabajar}
// Se coloco el puntero si se van a cambiar valores en la struct.
func (cursos *Course) CambiarPrecio(precio float64) {
	cursos.Price = precio
}

// Para convertilo a método de la -Struct-
// Se declara por fuera de la -Struct-
// No se utilizan las palabras reservas "this", "self"
// El argumento receptor debe ser claro, como c Course
// NOTA: Del expositor.
// Cuando se este manejando "Interfaces", si un método utiliza puntero, en todos los demas métodos se debe utilizar punteros, por lo que de una vez se cambiara.
func (cursos *Course) ImprimeClases() {
	text := "Las clases son :"
	for _, clases := range cursos.Classes {
		text += clases + ", "
	}
	fmt.Printf("Imprimiendo las clases : %v \n", text[:len(text)-2])
}
