package main

import "fmt"

//Create interface
type Barril interface {
	Servir() float32
	GetNombre() string
}

type Antares struct {
	nombreCerveza     string
	porcentajeAlcohol float32
}

func (a *Antares) Servir() float32 {
	return a.porcentajeAlcohol
}

func (a *Antares) GetNombre() string {
	return a.nombreCerveza
}

// type Baum struct {
// 	nombreCerveza     string
// 	porcentajeAlcohol float32
// }

// func (b *Baum) Servir() float32 {
// 	return b.porcentajeAlcohol
// }

// func (b *Baum) GetNombre() string {
// 	return b.nombreCerveza
// }

func tipo(barril Barril) string {
	if barril.Servir() > 6.5 {
		return barril.GetNombre() + " es una bebida fuerte"
	} else {
		return barril.GetNombre() + " es una bebida suave"
	}
}

func main() {

	cerveza := &Antares{"Barley Wine", 8}

	fmt.Println(tipo(cerveza))

	// cervezas := []Barril{&Antares{"Barley Wine", 8}, &Baum{"Blonde", 4}}

	// for _, cerveza := range cervezas {
	// 	fmt.Println(tipo(cerveza))
	// }

}
