package definition

type (
	// Attributes TODO
	Attributes map[string]interface{}

	// NodeDefinition TODO
	NodeDefinition struct {
		// ClassID TODO
		ClassID string `yaml:"Class"`

		Attributes Attributes `yaml:"Attributes"`
	}

	// NodeSpecification TODO
	NodeSpecification struct {
		// Class allows the class for all the definitions within the document to be specified.
		ClassID string `yaml:"Class,omitempty"`

		// Definitions TODO
		Definitions map[string]NodeDefinition `yaml:"Definitions,omitempty"`
	}
)
