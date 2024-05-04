/*
 * Copyright (C) 2024 Paul Tatham <paul@nextmetaphor.io>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
	entityDefinition         = "definition"

	// HTTP header parameter constants
	paramID                            = "ID"
	paramNamespace                     = "namespace"
	paramNodeID                        = "nodeID"
	paramNodeClassID                   = "nodeClassID"
	paramNodeClassNamespace            = "nodeClassNamespace"
	paramNodeClassAttributeID          = "nodeClassAttributeID"
	paramSourceNodeID                  = "sourceNodeID"
	paramSourceNodeClassID             = "sourceNodeClassID"
	paramSourceNodeClassNamespace      = "sourceNodeClassNamespace"
	paramDestinationNodeID             = "destinationNodeID"
	paramDestinationNodeClassID        = "destinationNodeClassID"
	paramDestinationNodeClassNamespace = "destinationNodeClassNamespace"
	paramDestinationNodeNamespace      = "destinationNodeNamespace"
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

	pathDefinitionRoot = "/" + entityDefinition
)

var (
	dbConn *sql.DB
)

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")

	w.WriteHeader(http.StatusNoContent)
}

func writeHTTPResponse(returnCode int, data any, err error, w http.ResponseWriter, errorMessage string) {
	var b []byte
	if err == nil && data != nil {
		b, err = json.Marshal(data)
	}

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
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
	mux.HandleFunc(http.MethodOptions+" "+pathNamespaceRoot, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNamespaceRoot, selectNamespaceHandler)

	// nodeClass functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassRoot, selectNodeClassHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeClassRoot, createNodeClassHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassRootEntity, readNodeClassHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeClassRootEntity, updateNodeClassHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeClassRootEntity, deleteNodeClassHandler)

	// nodeClassEdge functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassEdgeRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassEdgeRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassEdgeRoot, selectNodeClassEdgeBySourceNodeClassHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeClassEdgeRoot, createNodeClassEdgeHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassEdgeRootEntity, readNodeClassEdgeHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeClassEdgeRootEntity, updateNodeClassEdgeHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeClassEdgeRootEntity, deleteNodeClassEdgeHandler)

	//nodeClassAttribute functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassAttributeRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassAttributeRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassAttributeRoot, selectNodeClassAttributeHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeClassAttributeRoot, createNodeClassAttributeHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassAttributeRootEntity, readNodeClassAttributeHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeClassAttributeRootEntity, updateNodeClassAttributeHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeClassAttributeRootEntity, deleteNodeClassAttributeHandler)

	// node functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeRoot, selectNodeHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeRoot, createNodeHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeRootEntity, readNodeHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeRootEntity, updateNodeHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeRootEntity, deleteNodeHandler)

	// nodeAttribute functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeAttributeRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeAttributeRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeAttributeRoot, selectNodeAttributeHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeAttributeRoot, createNodeAttributeHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeAttributeRootEntity, readNodeAttributeHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeAttributeRootEntity, updateNodeAttributeHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeAttributeRootEntity, deleteNodeAttributeHandler)

	// nodeEdge functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeEdgeRoot, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeEdgeRootEntity, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeEdgeRoot, selectNodeEdgeBySourceNodeHandler)
	mux.HandleFunc(http.MethodPost+" "+pathNodeEdgeRoot, createNodeEdgeHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeEdgeRootEntity, readNodeEdgeHandler)
	mux.HandleFunc(http.MethodPut+" "+pathNodeEdgeRootEntity, updateNodeEdgeHandler)
	mux.HandleFunc(http.MethodDelete+" "+pathNodeEdgeRootEntity, deleteNodeEdgeHandler)

	// graph functions
	mux.HandleFunc(http.MethodOptions+" "+pathNodeClassGraph, preflightHandler)
	mux.HandleFunc(http.MethodOptions+" "+pathNodeGraph, preflightHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeClassGraph, nodeClassGraphHandler)
	mux.HandleFunc(http.MethodGet+" "+pathNodeGraph, nodeGraphHandler)

	// definition functions
	mux.HandleFunc(http.MethodPost+" "+pathDefinitionRoot, loadDefinitionsHandler)
	mux.HandleFunc(http.MethodGet+" "+pathDefinitionRoot, saveDefinitionsHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}
