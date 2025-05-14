package main

import router "main/internal/routes"

func main() {
	r := router.SetupRoutes()
	r.Run()
}
