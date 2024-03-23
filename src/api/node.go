package api

import (
	db2 "github.com/nextmetaphor/definition-graph/db"
	"net/http"
)

const (
	logCannotReadNode = "cannot read node"
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeClassNamespace := r.Header.Get(entityNamespace)
	nodeClass := r.Header.Get(entityNodeClass)

	data, err := db2.SelectNodes(db, nodeClass, nodeClassNamespace)
	writeHTTPResponse(data, err, w, logCannotSelectNodeGraph)
}

func readNodeHandler(w http.ResponseWriter, r *http.Request) {
	namespace := r.Header.Get(entityNamespace)
	nodeClassID := r.Header.Get(entityNodeClass)
	nodeID := r.PathValue(entityNode)

	data, err := db2.ReadNodeByID(db, namespace, nodeClassID, nodeID)
	writeHTTPResponse(data, err, w, logCannotReadNode)
}
