package goarea

import "math"

//Pi é uma propoção numérica definida pela relação entre
//o prímetro de uma circusferência e seu diâmetro

const Pi = 3.1416

//Circ é responsável por calcular a área da circumferâcia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
