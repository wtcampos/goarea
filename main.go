package goarea

import "math"

// Pi e uma proporcao numerica definida entre o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ eh responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {

	return math.Pow(raio, 2) * Pi
}

// Rect eh responsavel por retornar a area de um retangulo
func Rect(base, altura float64) float64 {

	return base * altura
}

// Nao eh visivel, pois comeca com _ (Privado)
func _TrianguloEQ(base, altura float64) float64 {

	return (base * altura) / 2
}
