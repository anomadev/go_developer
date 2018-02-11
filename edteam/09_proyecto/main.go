package main

import (
	"flag"
	"log"

	"github.com/go_developer/edteam/09_proyecto/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración...")
	}
}