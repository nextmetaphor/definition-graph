package model

type (
	NodeAttributes []NodeAttribute

	NodeAttributeKey struct {
		NodeID               string `json:"node-id"`
		NodeClassID          string `json:"node-class-id"`
		NodeClassNamespace   string `json:"node-class-namespace"`
		NodeClassAttributeID string `json:"node-class-attribute-id"`
	}

	NodeAttribute struct {
		NodeAttributeKey
		Value string `json:"value"`
	}
)
