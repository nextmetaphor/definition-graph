package model

type (
	Nodes []Node

	NodeKey struct {
		ID                 string `json:"ID"`
		NodeClassID        string `json:"nodeClassID"`
		NodeClassNamespace string `json:"nodeClassNamespace"`
	}

	Node struct {
		NodeKey
	}
)
