package utilities

import (
	"strings"
	"time"
)

//
func DateTimeNow() time.Time {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().UTC().Local()

}

func DateTimeNowAddHoursUnix(hours int64) int64 {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().Add(time.Hour * time.Duration(hours)).Unix()

}
func DateTimeNowAddHours(hours int64) time.Time {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().Add(time.Hour * time.Duration(hours))

}

const (
	// Velocidade média de leitura em palavras por minuto
	VelocidadeLeitura = 200
)

func CalcularTempoLeitura(texto string) int {
	// Dividir o texto em palavras
	palavras := strings.Fields(texto)

	// Calcular o tempo de leitura em minutos
	tempoLeitura := len(palavras) / VelocidadeLeitura

	return tempoLeitura
}
