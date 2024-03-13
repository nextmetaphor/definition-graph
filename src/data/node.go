package data

type (
	Nodes struct {
		Nodes []Node `json:"nodes"`
	}

	Node struct {
		ID                 string          `json:"id"`
		NodeClassID        string          `json:"node-class-id"`
		NodeClassNamespace string          `json:"node-class-namespace"`
		Attributes         []NodeAttribute `json:"attributes"`
	}

	NodeAttribute struct {
		NodeClassAttributeID string `json:"node-class-attribute-id"`
		Value                string `json:"value"`
	}

	NodeEdge struct {
		SourceNodeID             string `json:"source-node-id"`
		SourceNodeClassID        string `json:"source-node-class-id"`
		SourceNodeNamespace      string `json:"source-node-namespace"`
		DestinationNodeID        string `json:"destination-node-id"`
		DestinationNodeClassID   string `json:"destination-node-class-id"`
		DestinationNodeNamespace string `json:"destination-node-namespace"`
		Relationship             string `json:"relationship"`
	}
)
