package api

import (
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotSelectNamespaces = "cannot select namespaces"
	logCannotSelectNodeClass  = "cannot select nodeClass"
	logCannotCreateNodeClass  = "cannot create nodeClass"
	logCannotReadNodeClass    = "cannot read nodeClass"
	logCannotDeleteNodeClass  = "cannot delete nodeClass"
)

// function indirection to allow unit test stubs to be created
var (
	selectNamespacesFunc = db.SelectNamespaces
	selectNodeClassFunc  = db.SelectNodeClass
	createNodeClassFunc  = db.CreateNodeClass
	readNodeClassFunc    = db.ReadNodeClass
	updateNodeClassFunc  = db.UpdateNodeClass
	deleteNodeClassFunc  = db.DeleteNodeClass
)

func selectNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNamespacesFunc(dbConn)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNamespaces)
}

func selectNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeClassFunc(dbConn)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClass)
}

func createNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err == nil {
		err = createNodeClassFunc(dbConn, nc)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeClass)
}

func readNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := readNodeClassFunc(dbConn, model.NodeClassKey{
		ID:        r.Header.Get(paramID),
		Namespace: r.Header.Get(paramNamespace),
	})
	if nc == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClass)
	} else {
		writeHTTPResponse(http.StatusOK, nc, err, w, logCannotReadNodeClass)
	}
}

func updateNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotReadNodeClass)
	} else {
		count, err := updateNodeClassFunc(dbConn, model.NodeClassKey{
			ID:        r.Header.Get(paramID),
			Namespace: r.Header.Get(paramNamespace),
		}, nc)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClass)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotReadNodeClass)
		}
	}
}

func deleteNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeClassFunc(dbConn, model.NodeClassKey{
		ID:        r.Header.Get(paramID),
		Namespace: r.Header.Get(paramNamespace),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeClass)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeClass)
	}
}
