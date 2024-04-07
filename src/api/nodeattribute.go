package api

import (
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotSelectNodeAttribute = "cannot select node attribute"
	logCannotCreateNodeAttribute = "cannot create node attribute"
	logCannotReadNodeAttribute   = "cannot read node attribute"
	logCannotUpdateNodeAttribute = "cannot update node attribute"
	logCannotDeleteNodeAttribute = "cannot delete node attribute"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeAttributeFunc = db.SelectNodeAttributeByNode
	createNodeAttributeFunc = db.CreateNodeAttribute
	readNodeAttributeFunc   = db.ReadNodeAttribute
	updateNodeAttributeFunc = db.UpdateNodeAttribute
	deleteNodeAttributeFunc = db.DeleteNodeAttribute
)

func selectNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeAttributeFunc(dbConn, model.NodeKey{
		ID:                 r.Header.Get(paramNodeID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeAttribute)
}

func createNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var na model.NodeAttribute
	err := json.NewDecoder(r.Body).Decode(&na)

	if err == nil {
		err = createNodeAttributeFunc(dbConn, na)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeAttribute)
}

func readNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	na, err := readNodeAttributeFunc(dbConn, model.NodeAttributeKey{
		NodeID:               r.Header.Get(paramNodeID),
		NodeClassID:          r.Header.Get(paramNodeClassID),
		NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
		NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
	})
	if na == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, na, err, w, logCannotReadNodeAttribute)
	}
}

func updateNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var na model.NodeAttribute
	err := json.NewDecoder(r.Body).Decode(&na)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeAttribute)
	} else {
		count, err := updateNodeAttributeFunc(dbConn, model.NodeAttributeKey{
			NodeID:               r.Header.Get(paramNodeID),
			NodeClassID:          r.Header.Get(paramNodeClassID),
			NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
			NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
		}, na)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeAttribute)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeAttribute)
		}
	}
}

func deleteNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeAttributeFunc(dbConn, model.NodeAttributeKey{
		NodeID:               r.Header.Get(paramNodeID),
		NodeClassID:          r.Header.Get(paramNodeClassID),
		NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
		NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeAttribute)
	}
}
