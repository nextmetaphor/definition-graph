package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	entityNamespace = "namespace"
	entityNodeClass = "nodeclass"
	entityNode      = "node"
	entityGraph     = "graph"

	pathNamespaceRoot = "/" + entityNamespace
	pathNamespace     = pathNamespaceRoot + "/{" + entityNamespace + "}"

	pathNodeClassRoot = "/" + entityNodeClass
	pathNodeClass     = pathNodeClassRoot + "/{" + entityNamespace + "}" + "/{" + entityNodeClass + "}"

	pathNodeRoot = "/" + entityNode
	pathNode     = pathNodeRoot + "/{" + entityNode + "}"

	pathNodeClassGraph = "/" + entityGraph + "/" + entityNodeClass
	pathNodeGraph      = "/" + entityGraph + "/" + entityNode
)

var (
	db *sql.DB
)

func writeHTTPResponse(returnCode int, data any, err error, w http.ResponseWriter, errorMessage string) {
	var b []byte
	if err == nil && data != nil {
		b, err = json.Marshal(data)
	}

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err != nil {
		log.Error().Err(err).Msg(errorMessage)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(returnCode)
		if _, err = fmt.Fprintf(w, string(b)); err != nil {
			log.Error().Err(err).Msg(errorMessage)
		}
	}
}

func Listen(conn *sql.DB) {
	db = conn

	mux := http.NewServeMux()

	// node functions
	mux.HandleFunc(pathNode, readNodeHandler)
	mux.HandleFunc(pathNodeRoot, selectNodeHandler)

	// namespace functions
	mux.HandleFunc("GET "+pathNamespaceRoot, selectNamespaceHandler)

	// nodeClass functions
	mux.HandleFunc("GET "+pathNodeClassRoot, selectNodeClassHandler)
	mux.HandleFunc("POST "+pathNodeClassRoot, createNodeClassHandler)
	mux.HandleFunc("GET "+pathNodeClass, readNodeClassHandler)
	mux.HandleFunc("PUT "+pathNodeClass, updateNodeClassHandler)
	mux.HandleFunc("DELETE "+pathNodeClass, deleteNodeClassHandler)

	// graph functions
	mux.HandleFunc(pathNodeClassGraph, nodeClassGraphHandler)
	mux.HandleFunc(pathNodeGraph, nodeGraphHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}
