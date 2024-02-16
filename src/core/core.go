package core

import (
	"database/sql"
	"fmt"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	logDebugCannotLoadFile           = "cannot load definition file [%s]"
	logDebugCannotParseFile          = "cannot parse definition file [%s]"
	logDebugNoDefinitionsFoundInFile = "no definitions found in definition file [%s]"
)

func LoadSpecificationFromFile(filename string) (*definition.NodeClassSpecification, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf(logDebugCannotLoadFile, filename))
		return nil, err
	}

	spec := &definition.NodeClassSpecification{}
	err = yaml.Unmarshal(file, spec)
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf(logDebugCannotParseFile, filename))

		return nil, err
	}

	// if no definitions are found, return an error and a nil Specification
	if len(spec.Definitions) == 0 {
		log.Debug().Err(err).Msg(fmt.Sprintf(logDebugNoDefinitionsFoundInFile, filename))
		return nil, fmt.Errorf(logDebugNoDefinitionsFoundInFile, filename)
	}

	// TODO debug
	return spec, nil
}

func LoadNodeClassDefinitions(conn *sql.DB) {
	spec, _ := LoadSpecificationFromFile("./definition/_test/definitions.yaml")

	db.StoreNodeClassSpecification(conn, spec)
}
