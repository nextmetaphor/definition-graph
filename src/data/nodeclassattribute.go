package data

type (
	NodeClassAttributes []NodeClassAttribute

	NodeClassAttribute struct {
		ID                 string  `json:"id"`
		NodeClassID        string  `json:"node-class-id"`
		NodeClassNamespace string  `json:"node-class-namespace"`
		Type               string  `json:"type"`
		IsRequired         int     `json:"is-required"`
		Description        *string `json:"description,omitempty"`
	}
)
