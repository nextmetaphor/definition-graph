package model

type (
	Namespaces struct {
		Namespace []Namespace `json:"namespaces"`
	}

	Namespace struct {
		Namespace string `json:"namespace"`
	}
)
