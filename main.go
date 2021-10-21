package main

import (
	"me/framework"
	"me/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// core.Use(middleware.Timeout(1 * timeSecond))

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
