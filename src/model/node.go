package model

type (
	Nodes struct {
		Nodes []Node `json:"nodes"`
	}

	NodeKey struct {
		ID                 string `json:"id"`
		NodeClassID        string `json:"node-class-id"`
		NodeClassNamespace string `json:"node-class-namespace"`
	}

	Node struct {
		NodeKey
	}
)
