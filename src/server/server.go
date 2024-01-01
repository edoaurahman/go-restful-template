package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	// Inisialisasi server Gin
	r := gin.Default()

	// Mengambil URL MongoDB dari variabel lingkungan
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	// Menjalankan server pada port tertentu
	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
