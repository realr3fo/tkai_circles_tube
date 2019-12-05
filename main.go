package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"tkai_circles_tube/app"
	"tkai_circles_tube/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/tube/volume", controllers.GetTubeVolume).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
