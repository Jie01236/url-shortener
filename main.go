package main

import (
<<<<<<< HEAD
	"github.com/axellelanca/urlshortener/cmd"
	_ "github.com/axellelanca/urlshortener/cmd/cli" // Importe le package 'cli' pour que ses init() soient exécutés
	_ "github.com/axellelanca/urlshortener/cmd/server" // Temporairement commenté pour la migration DB uniquement
=======
	_ "github.com/Quanghng/url-shortener/cmd/cli"    // Importe le package 'cli' pour que ses init() soient exécutés
	_ "github.com/Quanghng/url-shortener/cmd/server" // Importe le package 'server' pour que ses init() soient exécutés
>>>>>>> 45cfe642b207426e144b1e381838a64ed43d2f37
)

func main() {
	// Exécute la commande racine Cobra
	cmd.Execute()
}
