package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Largura)
}

func main() {
	teste := Circulo{10.0}
	fmt.Println(Circulo.Area(teste))
}
