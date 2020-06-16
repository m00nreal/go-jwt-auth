package api

import (
	"fmt"
	"github.com/m00nreal/go-jwt-auth/api/auto"
	"github.com/m00nreal/go-jwt-auth/api/config"
	"github.com/m00nreal/go-jwt-auth/api/router"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}