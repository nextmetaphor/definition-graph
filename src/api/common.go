package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	entityNamespace     = "namespace"
	entityNodeClass     = "nodeclass"
	entityNodeClassEdge = "nodeclassedge"
	entityNode          = "node"
	entityGraph         = "graph"

	pathNamespaceRoot = "/" + entityNamespace
	pathNamespace     = pathNamespaceRoot + "/{" + entityNamespace + "}"

	pathNodeClassRoot = "/" + entityNodeClass
	pathNodeClass     = pathNodeClassRoot + "/{" + entityNamespace + "}" + "/{" + entityNodeClass + "}"

	pathNodeClassEdgeRoot       = "/" + entityNodeClassEdge
	pathNodeClassEdgeRootEntity = pathNodeClassEdgeRoot + "/"

	pathNodeRoot = "/" + entityNode
	pathNode     = pathNodeRoot + "/{" + entityNode + "}"

	pathNodeClassGraph = "/" + entityGraph + "/" + entityNodeClass
	pathNodeGraph      = "/" + entityGraph + "/" + entityNode
)

var (
	dbConn *sql.DB
)

func writeHTTPResponse(returnCode int, data any, err error, w http.ResponseWriter, errorMessage string) {
	var b []byte
	if err == nil && data != nil {
		b, err = json.Marshal(data)
	}

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

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
	dbConn = conn

	mux := http.NewServeMux()

	// node functions
	mux.HandleFunc(pathNode, readNodeHandler)
	mux.HandleFunc(pathNodeRoot, selectNodeHandler)

	// namespace functions
	mux.HandleFunc(pathNamespaceRoot, selectNamespaceHandler)

	// nodeClass functions
	mux.HandleFunc(pathNodeClassRoot, selectNodeClassHandler)
	mux.HandleFunc("POST "+pathNodeClassRoot, createNodeClassHandler)
	mux.HandleFunc("GET "+pathNodeClass, readNodeClassHandler)
	mux.HandleFunc("PUT "+pathNodeClass, updateNodeClassHandler)
	mux.HandleFunc("DELETE "+pathNodeClass, deleteNodeClassHandler)

	// nodeClassEdge functions
	mux.HandleFunc(pathNodeClassEdgeRoot, selectNodeClassEdgeBySourceNodeClassHandler)
	mux.HandleFunc("POST "+pathNodeClassEdgeRoot, createNodeClassEdgeHandler)
	mux.HandleFunc("GET "+pathNodeClassEdgeRootEntity, readNodeClassEdgeHandler)
	mux.HandleFunc("PUT "+pathNodeClassEdgeRootEntity, updateNodeClassEdgeHandler)
	mux.HandleFunc("DELETE "+pathNodeClassEdgeRootEntity, deleteNodeClassEdgeHandler)

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
