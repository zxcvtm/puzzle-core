package main

import (
    "app/workspace/routes"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "github.com/urfave/negroni"
    "fmt"
    "net/http"
)
func main() {
    r := mux.NewRouter()

    n := negroni.Classic() // Includes some default middlewares

    n.Use(cors.New(cors.Options{
        AllowedOrigins   : []string{"*"},
        AllowedMethods   : []string{"HEAD","GET","POST","PUT","DELETE","PATCH","OPTIONS"},
        AllowedHeaders   : []string{"Origin","Authorization","X-Requested-With","Content-Type","Accept","Signature"},
        ExposedHeaders   : []string{"Content-Length"},
        AllowCredentials : true,
    }))

    n.UseHandler(r)
    //Routes
    apiRoutes(r)
    //Run server
    fmt.Println("Server is about to run...")
    _ = http.ListenAndServe(":3000", n)
}

func apiRoutes(r *mux.Router) {
    routes.SocketApi(r)
}