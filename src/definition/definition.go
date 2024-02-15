package definition

type (
	// NodeClassAttribute TODO
	NodeClassAttribute struct {
		Description string `yaml:"Description,omitempty"`
		Type        string `yaml:"Type"`
		IsRequired  int    `yaml:"IsRequired"`
	}

	// NodeClassAttributes TODO
	NodeClassAttributes map[string]NodeClassAttribute

	// NodeClassDefinition TODO
	NodeClassDefinition struct {
		// Description TODO
		Description string `yaml:"Description,omitempty"`

		// Attributes TODO
		Attributes NodeClassAttributes `yaml:"Attributes"`
	}

	// NodeClassSpecification TODO
	NodeClassSpecification struct {
		// Definitions TODO
		Definitions map[string]NodeClassDefinition `yaml:"Definitions,omitempty"`
	}
)
