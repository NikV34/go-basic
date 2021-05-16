package main

import (
	"chat/api"
	"chat/user"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

var rdb *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()
	rdb.SAdd(user.ChannelsKey, "general", "random")
}

func main() {
	// var dir string

	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))

	r.Path("/").Methods("GET").Handler(fileServer)
	r.Path("/chat").Methods("GET").HandlerFunc(api.H(rdb, api.ChatWebSocketHandler))
	r.Path("/user/{user}/channels").Methods("GET").HandlerFunc(api.H(rdb, api.UserChannelsHandler))
	r.Path("/users").Methods("GET").HandlerFunc(api.H(rdb, api.UsersHandler))
	
	r.Path("/").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "static/index.html")})
	r.Path("/styles.css").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "static/styles.css")})
	r.Path("/main.js").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "static/main.js")})
	
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	fmt.Println("chat service started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}