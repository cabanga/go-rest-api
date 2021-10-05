package main

import "github.com/cabanga/go-rest-api/server"

/*
func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing!")
}
*/

func main() {
	server := server.NewServer()
	server.Run()
}
