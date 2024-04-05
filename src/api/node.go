package api

import (
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotReadNode = "cannot read node"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeFunc = db.SelectNodes
	readNodeFunc   = db.ReadNode
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeFunc(dbConn, model.NodeClassKey{
		ID:        r.Header.Get(paramNodeClassID),
		Namespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeGraph)
}

func readNodeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readNodeFunc(dbConn, model.NodeKey{
		ID:                 r.Header.Get(paramID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotReadNode)
}
