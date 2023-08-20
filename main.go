package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FortuneCookieMessage struct {
	Language string
	Message  string
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	fortunes := []FortuneCookieMessage{
		{Language: "Spanish", Message: "En el camino de las estrellas, tropezarás con un cangrejo danzante."},
		{Language: "English", Message: "Wandering through the galaxy, you'll meet a sneezing nebula."},
		{Language: "Spanish", Message: "La Luna ríe, pero las estrellas prefieren el beatbox."},
		{Language: "English", Message: "A cosmic snail will offer you philosophical advice."},
		{Language: "Spanish", Message: "Las nubes bailarán salsa mientras el viento toca la flauta."},
		{Language: "English", Message: "Beware the gravity-defying marshmallows on Saturn's rings."},
		{Language: "Spanish", Message: "Unicornios intergalácticos te traerán felicidad pixelada."},
		{Language: "English", Message: "In a parallel universe, dolphins write poetry."},
		{Language: "Spanish", Message: "El quásar susurra secretos cósmicos a los gatos espaciales."},
		{Language: "English", Message: "Embrace the chaos of quarks; they know the macarena."},
		{Language: "Spanish", Message: "Los agujeros de gusano son portales a dimensiones biscuit."},
		{Language: "English", Message: "Beware of rogue asteroids selling used starships."},
		{Language: "Spanish", Message: "Los cometas son cometas, excepto cuando son pasteles disfrazados."},
		{Language: "English", Message: "The moon's secret: it's a giant space cheese."},
		{Language: "Spanish", Message: "Los extraterrestres intercambian recetas de paella por secretos humanos."},
		{Language: "English", Message: "Binary stars whisper binary jokes to passing comets."},
		{Language: "Spanish", Message: "Los alienígenas intercambian tips de moda con los astronautas."},
		{Language: "English", Message: "Cosmic kangaroos bounce through asteroid fields."},
		{Language: "Spanish", Message: "El sol practica yoga para mantener su brillo radiante."},
		{Language: "English", Message: "Warp drive engaged: prepare for interstellar bumper cars."},
		{Language: "Spanish", Message: "Las constelaciones son el arte abstracto de los titanes espaciales."},
	}
	randomIndex := rand.Intn(len(fortunes))
	theFortune := fortunes[randomIndex]

	fmt.Println("Cookie says: ")
	fmt.Println((theFortune.Message))
}
