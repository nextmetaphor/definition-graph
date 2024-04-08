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
	paramNodeID                        = "node-id"
	paramNodeClassID                   = "node-class-id"
	paramNodeClassNamespace            = "node-class-namespace"
	paramNodeClassAttributeID          = "node-class-attribute-id"
	paramSourceNodeID                  = "source-node-" + paramID
	paramSourceNodeClassID             = "source-" + paramNodeClassID
	paramSourceNodeClassNamespace      = "source-" + paramNodeClassNamespace
	paramDestinationNodeID             = "destination-node-" + paramID
	paramDestinationNodeClassID        = "destination-" + paramNodeClassID
	paramDestinationNodeClassNamespace = "destination-" + paramNodeClassNamespace
	paramDestinationNodeNamespace      = "destination-node-" + paramNamespace
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
	pathNodeRoot       = "/" + entityNode
	pathNodeRootEntity = pathNodeRoot + "/"

	// NodeAttribute URL paths
	pathNodeAttributeRoot       = "/" + entityNodeAttribute
	pathNodeAttributeRootEntity = pathNodeAttributeRoot + "/"

	// NodeEdge URL paths
	pathNodeEdgeRoot       = "/" + entityNodeEdge
	pathNodeEdgeRootEntity = pathNodeEdgeRoot + "/"

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
	mux.HandleFunc("POST "+pathNodeRoot, createNodeHandler)
	mux.HandleFunc("GET "+pathNodeRootEntity, readNodeHandler)
	mux.HandleFunc("PUT "+pathNodeRootEntity, updateNodeHandler)
	mux.HandleFunc("DELETE "+pathNodeRootEntity, deleteNodeHandler)

	// nodeAttribute functions
	mux.HandleFunc(pathNodeAttributeRoot, selectNodeAttributeHandler)
	mux.HandleFunc("POST "+pathNodeAttributeRoot, createNodeAttributeHandler)
	mux.HandleFunc("GET "+pathNodeAttributeRootEntity, readNodeAttributeHandler)
	mux.HandleFunc("PUT "+pathNodeAttributeRootEntity, updateNodeAttributeHandler)
	mux.HandleFunc("DELETE "+pathNodeAttributeRootEntity, deleteNodeAttributeHandler)

	// nodeEdge functions
	mux.HandleFunc(pathNodeEdgeRoot, selectNodeEdgeBySourceNodeHandler)
	mux.HandleFunc("POST "+pathNodeEdgeRoot, createNodeEdgeHandler)
	mux.HandleFunc("GET "+pathNodeEdgeRootEntity, readNodeEdgeHandler)
	mux.HandleFunc("PUT "+pathNodeEdgeRootEntity, updateNodeEdgeHandler)
	mux.HandleFunc("DELETE "+pathNodeEdgeRootEntity, deleteNodeEdgeHandler)

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
