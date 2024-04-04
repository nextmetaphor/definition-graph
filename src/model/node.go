package model

type (
	Nodes []Node

	NodeKey struct {
		ID                 string `json:"id"`
		NodeClassID        string `json:"node-class-id"`
		NodeClassNamespace string `json:"node-class-namespace"`
	}

	Node struct {
		NodeKey
	}
)
