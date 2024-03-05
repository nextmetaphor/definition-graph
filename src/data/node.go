package data

type (
	Nodes struct {
		Nodes []Node `json:"nodes"`
	}

	Node struct {
		ID                 string          `json:"id"`
		NodeClassID        string          `json:"nodeClassID"`
		NodeClassNamespace string          `json:"nodeClassNamespace"`
		NodeAttributes     []NodeAttribute `json:"nodeAttributes"`
	}

	NodeAttribute struct {
		AttributeID string `json:"attributeID"`
		Value       string `json:"value"`
	}
)
