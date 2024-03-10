package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

const (
	paramNamespace = "namespace"
	paramNodeClass = "nodeClass"
	paramNode      = "node"

	pathNodeClass  = "/nodeClass"
	pathSelectNode = "/node"

	pathReadNamespace = "/namespace/{" + paramNamespace + "}"
	pathReadNodeClass = pathReadNamespace + "/nodeClass/{" + paramNodeClass + "}"
	pathReadNode      = pathReadNodeClass + "/node/{" + paramNode + "}/"

	pathNodeClassGraph = "/nodeClassGraph"
	pathNodeGraph      = "/nodeGraph"
)

var (
	db *sql.DB
)

func Listen(conn *sql.DB) {
	db = conn

	mux := http.NewServeMux()

	// node functions
	fmt.Println(pathReadNode)
	mux.HandleFunc(pathReadNode, readNodeHandler)
	mux.HandleFunc(pathSelectNode, selectNodeHandler)

	// nodeClass functions
	mux.HandleFunc(pathNodeClass, nodeClassHandler)

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
