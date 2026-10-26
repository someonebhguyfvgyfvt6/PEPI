// Punto de entrada de la API: carga configuración, conecta la base de datos,
// levanta el servidor HTTP y lo apaga limpiamente al recibir Ctrl+C / SIGTERM.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/someonebhguyfvgyfvt6/pepi-backend/config"
	deliveryhttp "github.com/someonebhguyfvgyfvt6/pepi-backend/internal/delivery/http"
	"github.com/someonebhguyfvgyfvt6/pepi-backend/pkg/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	db, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("conectado a PostgreSQL")

	app := deliveryhttp.NewRouter()

	// El servidor corre en una goroutine para que main pueda esperar señales.
	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("error del servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("apagando servidor...")
	// Deja terminar las peticiones en curso antes de cerrar.
	if err := app.ShutdownWithTimeout(cfg.ShutdownTimeout); err != nil {
		log.Printf("error en el apagado: %v", err)
	}
	log.Println("servidor detenido")
}
