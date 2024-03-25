package model

type (
	NodeClassAttributes []NodeClassAttribute

	NodeClassAttributeKey struct {
		ID                 string `json:"id"`
		NodeClassID        string `json:"node-class-id"`
		NodeClassNamespace string `json:"node-class-namespace"`
	}

	NodeClassAttribute struct {
		NodeClassAttributeKey
		Type        string  `json:"type"`
		IsRequired  int     `json:"is-required"`
		Description *string `json:"description,omitempty"`
	}
)
