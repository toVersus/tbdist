package main

import (
	"log"
	"net/http"

	"github.com/toversus/tbdist/router"
	"github.com/toversus/tbdist/server"
)

func main() {
	r := &router.Router{}
	r.HandleFunc("/tasks/pending", http.MethodGet, server.GetPendingTasks)
	r.HandleFunc("/tasks/doing", http.MethodGet, server.GetDoingTasks)
	r.HandleFunc("/tasks/done", http.MethodGet, server.GetDoneTasks)
	r.HandleFunc("/tasks", http.MethodPost, server.AddTask)
	r.HandleFunc(`/tasks/\d`, http.MethodPut, server.UpdateTask)

	log.Fatal(http.ListenAndServe(":8080", r))
}
