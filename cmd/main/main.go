package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Ezinnem/workers/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.MembersRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7500", r))
	fmt.Println("Server is connected and running at port 7500")
}