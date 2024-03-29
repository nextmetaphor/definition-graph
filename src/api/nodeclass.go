package api

import (
	"encoding/json"
	db2 "github.com/nextmetaphor/definition-graph/db"
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
	selectNamespacesFunc = db2.SelectNamespaces
	selectNodeClassFunc  = db2.SelectNodeClass
)

func selectNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNamespacesFunc(db)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNamespaces)
}

func selectNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeClassFunc(db)
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClass)
}

func createNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err == nil {
		err = db2.CreateNodeClass(db, nc)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeClass)
}

func readNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	nc, err := db2.ReadNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	})
	if nc == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClass)
	} else {
		writeHTTPResponse(http.StatusOK, nc, err, w, logCannotReadNodeClass)
	}
}

func updateNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err == nil {
		err = db2.CreateNodeClass(db, nc)
	}

	count, err := db2.UpdateNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	}, nc)
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClass)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotReadNodeClass)
	}
}

func deleteNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	count, err := db2.DeleteNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeClass)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeClass)
	}
}
