package data

type (
	NodeClassesOuter struct {
		NodeClasses []NodeClass `json:"nodeClasses"`
	}

	NodeClass struct {
		ID          string               `json:"id"`
		Namespace   string               `json:"namespace"`
		Description string               `json:"description"`
		Attributes  []NodeClassAttribute `json:"attributes,omitempty"`
		Edges       []NodeClassEdge      `json:"edges,omitempty"`
	}

	NodeClassAttribute struct {
		ID          string `json:"id"`
		Type        string `json:"type,omitempty"`
		IsRequired  int    `json:"is-required,omitempty"`
		Description string `json:"description,omitempty"`
	}

	NodeClassEdge struct {
		SourceNodeClassID             string `json:"source-node-class-id"`
		SourceNodeClassNamespace      string `json:"source-node-class-namespace"`
		DestinationNodeClassID        string `json:"destination-node-class-id"`
		DestinationNodeClassNamespace string `json:"destination-node-class-namespace"`
	}
)
