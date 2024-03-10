package api

import (
	"fmt"
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

func selectNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeClassNamespace := r.Header.Get("nodeClassNamespace")
	nodeClass := r.Header.Get("nodeClass")

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
	fmt.Println("I am here")
	namespace := r.PathValue(paramNamespace)
	nodeClassID := r.PathValue(paramNodeClass)
	nodeID := r.PathValue(paramNode)

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
