package api

import (
	"database/sql"
	"fmt"
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

const (
	pathNodeClassGraph = "/nodeClassGraph"
	pathNodeClass      = "/nodeGraph"
)

var (
	db *sql.DB
)

func nodeClassGraphHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNodeClassGraph(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func nodeGraphHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNodeGraph(db)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}
func Listen(conn *sql.DB) {
	db = conn
	http.HandleFunc(pathNodeClassGraph, nodeClassGraphHandler)
	http.HandleFunc(pathNodeClass, nodeGraphHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
