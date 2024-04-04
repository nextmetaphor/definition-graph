package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	entityNamespace          = "namespace"
	entityNodeClass          = "nodeclass"
	entityNodeClassEdge      = "nodeclassedge"
	entityNodeClassAttribute = "nodeclassattribute"
	entityNode               = "node"
	entityNodeAttribute      = "nodeattribute"
	entityNodeEdge           = "nodeedge"
	entityGraph              = "graph"

	// HTTP header parameter constants
	paramID                            = "id"
	paramNamespace                     = "namespace"
	paramNodeClassID                   = "node-class-id"
	paramNodeClassNamespace            = "node-class-namespace"
	paramSourceNodeClassID             = "source-" + paramNodeClassID
	paramSourceNodeClassNamespace      = "source-" + paramNodeClassNamespace
	paramDestinationNodeClassID        = "destination-" + paramNodeClassID
	paramDestinationNodeClassNamespace = "destination-" + paramNodeClassNamespace
	paramRelationship                  = "relationship"

	// Namespace URL paths
	pathNamespaceRoot = "/" + entityNamespace

	// NodeClass URL paths
	pathNodeClassRoot       = "/" + entityNodeClass
	pathNodeClassRootEntity = pathNodeClassRoot + "/"

	// NodeClassEdge URL paths
	pathNodeClassEdgeRoot       = "/" + entityNodeClassEdge
	pathNodeClassEdgeRootEntity = pathNodeClassEdgeRoot + "/"

	// NodeClassAttribute URL paths
	pathNodeClassAttributeRoot       = "/" + entityNodeClassAttribute
	pathNodeClassAttributeRootEntity = pathNodeClassAttributeRoot + "/"

	// Node URL paths
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

	// namespace functions
	mux.HandleFunc(pathNamespaceRoot, selectNamespaceHandler)

	// nodeClass functions
	mux.HandleFunc(pathNodeClassRoot, selectNodeClassHandler)
	mux.HandleFunc("POST "+pathNodeClassRoot, createNodeClassHandler)
	mux.HandleFunc("GET "+pathNodeClassRootEntity, readNodeClassHandler)
	mux.HandleFunc("PUT "+pathNodeClassRootEntity, updateNodeClassHandler)
	mux.HandleFunc("DELETE "+pathNodeClassRootEntity, deleteNodeClassHandler)

	// nodeClassEdge functions
	mux.HandleFunc(pathNodeClassEdgeRoot, selectNodeClassEdgeBySourceNodeClassHandler)
	mux.HandleFunc("POST "+pathNodeClassEdgeRoot, createNodeClassEdgeHandler)
	mux.HandleFunc("GET "+pathNodeClassEdgeRootEntity, readNodeClassEdgeHandler)
	mux.HandleFunc("PUT "+pathNodeClassEdgeRootEntity, updateNodeClassEdgeHandler)
	mux.HandleFunc("DELETE "+pathNodeClassEdgeRootEntity, deleteNodeClassEdgeHandler)

	//nodeClassAttribute functions
	mux.HandleFunc(pathNodeClassAttributeRoot, selectNodeClassAttributeHandler)
	mux.HandleFunc("POST "+pathNodeClassAttributeRoot, createNodeClassAttributeHandler)
	mux.HandleFunc("GET "+pathNodeClassAttributeRootEntity, readNodeClassAttributeHandler)
	mux.HandleFunc("PUT "+pathNodeClassAttributeRootEntity, updateNodeClassAttributeHandler)
	mux.HandleFunc("DELETE "+pathNodeClassAttributeRootEntity, deleteNodeClassAttributeHandler)

	// node functions
	mux.HandleFunc(pathNodeRoot, selectNodeHandler)
	mux.HandleFunc(pathNode, readNodeHandler)

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
