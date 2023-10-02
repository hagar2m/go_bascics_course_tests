package main

import (
	"fmt"
	"html/template"
	"net/http"
    // "github.com/gin-gonic/gin"

	"test.com/museum/api"
	"test.com/museum/data"
)

func main()  {

		// r := gin.Default()
		// r.GET("/ping", func(c *gin.Context) {
		// 	c.JSON(200, gin.H {
		// 		"message": "pong",
		// 	})
		// })
		// r.Run()	
	

	server := http.NewServeMux()
	server.HandleFunc("/hello", helloHandler)
	server.HandleFunc("/template", templateHandler)
	server.HandleFunc("/api/exhibitions", api.GetApi)
	server.HandleFunc("/api/exhibitions/new", api.PostApi)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)


	err := http.ListenAndServe(":8080", server)
	if err == nil {
		fmt.Println("Errpr while opening the server")
	}
}	

func helloHandler(w http.ResponseWriter, r* http.Request) {
	w.Write([]byte("Hello"))
}

func templateHandler(w http.ResponseWriter, r* http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("INTERNAAL SERVER ERRORRR"))
		return
	}
	html.Execute(w, data.GetAll())
}
