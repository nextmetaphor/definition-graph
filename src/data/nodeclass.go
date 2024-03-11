package data

type (
	Namespaces struct {
		Namespace []Namespace `json:"namespaces"`
	}

	Namespace struct {
		Namespace string `json:"namespace"`
	}

	NodeClasses struct {
		NodeClasses []NodeClass `json:"nodeClasses"`
	}

	NodeClass struct {
		ID          string `json:"id"`
		Namespace   string `json:"namespace"`
		Description string `json:"description"`
	}
)
