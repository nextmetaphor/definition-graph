package api

import (
	"fmt"
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeClassNamespace := r.PathValue(entityNamespace)
	nodeClass := r.PathValue(entityNodeClass)

	b, err := core.SelectNodes(db, nodeClassNamespace, nodeClass)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,nodeClassNamespace,nodeClass")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func readNodeHandler(w http.ResponseWriter, r *http.Request) {
	namespace := r.PathValue(entityNamespace)
	nodeClassID := r.PathValue(entityNodeClass)
	nodeID := r.PathValue(entityNode)

	b, err := core.ReadNode(db, namespace, nodeClassID, nodeID)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,nodeClassNamespace,nodeClass")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}
