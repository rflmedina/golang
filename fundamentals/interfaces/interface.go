package main

import "fmt"

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func escreverArea(f forma) {
	fmt.Print("A área da forma é %0.2f", f.area())
}

func main() {
	r := retangulo{10, 25}
	escreverArea(r)
}
