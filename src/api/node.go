package api

import (
	"github.com/nextmetaphor/definition-graph/db"
	"net/http"
)

const (
	logCannotReadNode = "cannot read node"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeFunc = db.SelectNodes
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeClassNamespace := r.Header.Get(entityNamespace)
	nodeClass := r.Header.Get(entityNodeClass)

	data, err := selectNodeFunc(dbConn, nodeClass, nodeClassNamespace)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeGraph)
}

func readNodeHandler(w http.ResponseWriter, r *http.Request) {
	namespace := r.Header.Get(entityNamespace)
	nodeClassID := r.Header.Get(entityNodeClass)
	nodeID := r.PathValue(entityNode)

	data, err := db.ReadNodeByID(dbConn, namespace, nodeClassID, nodeID)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotReadNode)
}
