package main

import (
	"nastha-test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// pool, err := utils.DBConnect()
	// if err != nil {
	// 	log.Fatalf("Gagal koneksi ke database: %v", err)
	// }
	// defer pool.Close()

	// services.RunSeed(pool)
	router.CombineRouter(r)

	r.Run(":8080")

}
