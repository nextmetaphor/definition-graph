package api

import (
	"database/sql"
	"fmt"
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
	pathNodeClass     = pathNodeClassRoot + "/{" + entityNodeClass + "}"

	pathNodeRoot = "/" + entityNode
	pathNode     = pathNodeRoot + "/{" + entityNode + "}"

	pathNodeClassGraph = "/" + entityGraph + "/" + entityNodeClass
	pathNodeGraph      = "/" + entityGraph + "/" + entityNode
)

var (
	db *sql.DB
)

func Listen(conn *sql.DB) {
	db = conn

	mux := http.NewServeMux()

	// node functions
	mux.HandleFunc(pathNode, readNodeHandler)
	mux.HandleFunc(pathNodeRoot, selectNodeHandler)

	// nodeClass functions
	mux.HandleFunc(pathNodeClassRoot, nodeClassHandler)
	mux.HandleFunc(pathNamespaceRoot, selectNamespaceHandler)

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
