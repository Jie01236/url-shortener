package main

import (
	"github.com/axellelanca/urlshortener/cmd"
	_ "github.com/axellelanca/urlshortener/cmd/cli" // Importe le package 'cli' pour que ses init() soient exécutés
	_ "github.com/axellelanca/urlshortener/cmd/server" // Temporairement commenté pour la migration DB uniquement
)

func main() {
	// Exécute la commande racine Cobra
	cmd.Execute()
}
