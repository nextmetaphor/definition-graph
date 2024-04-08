package api

import (
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotCreateNode = "cannot create node"
	logCannotReadNode   = "cannot read node"
	logCannotUpdateNode = "cannot update node"
	logCannotDeleteNode = "cannot delete node"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeFunc = db.SelectNodes
	createNodeFunc = db.CreateNode
	readNodeFunc   = db.ReadNode
	updateNodeFunc = db.UpdateNode
	deleteNodeFunc = db.DeleteNode
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeFunc(dbConn, model.NodeClassKey{
		ID:        r.Header.Get(paramNodeClassID),
		Namespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeGraph)
}

func createNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node model.Node
	err := json.NewDecoder(r.Body).Decode(&node)

	if err == nil {
		err = createNodeFunc(dbConn, node)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNode)
}

func readNodeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readNodeFunc(dbConn, model.NodeKey{
		ID:                 r.Header.Get(paramID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	if data == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNode)
	} else {
		writeHTTPResponse(http.StatusOK, data, err, w, logCannotReadNode)
	}
}

func updateNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node model.Node
	err := json.NewDecoder(r.Body).Decode(&node)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNode)
	} else {
		count, err := updateNodeFunc(dbConn, model.NodeKey{
			ID:                 r.Header.Get(paramID),
			NodeClassID:        r.Header.Get(paramNodeClassID),
			NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
		}, node)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNode)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNode)
		}
	}
}

func deleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeFunc(dbConn, model.NodeKey{
		ID:                 r.Header.Get(paramID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNode)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNode)
	}
}
