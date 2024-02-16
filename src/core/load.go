package core

import (
	"database/sql"
	"fmt"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	definitionFormat = ".%s"

	logCannotProcessFiles       = "cannot process files in directory [%s]"
	logProcessingFile           = "processing file [%s] in directory [%s]"
	logIgnoringFile             = "ignoring file [%s] in directory [%s]"
	logAboutToLoadFile          = "about to load file [%s]"
	logSuccessfullyLoadedFile   = "successfully loaded file [%s]"
	logSkippingFile             = "skipping file [%s] due to error [%s]"
	logCannotLoadFile           = "cannot load definition file [%s]"
	logCannotParseFile          = "cannot parse definition file [%s]"
	logNoDefinitionsFoundInFile = "no definitions found in definition file [%s]"
)

type (
	processFileFuncType = func(filePath string, fileInfo os.FileInfo) (err error)
)

func loadDefinitions(sourceDir []string, fileExtension string, processFileFunc processFileFuncType) error {
	for _, dir := range sourceDir {
		err := filepath.Walk(dir, func(filePath string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotProcessFiles, filePath)
				return err
			}
			if !fileInfo.IsDir() {
				if strings.HasSuffix(fileInfo.Name(), fmt.Sprintf(definitionFormat, fileExtension)) {
					log.Debug().Msg(fmt.Sprintf(logProcessingFile, fileInfo.Name(), filePath))
					return processFileFunc(filePath, fileInfo)
				}

				log.Debug().Msg(fmt.Sprintf(logIgnoringFile, fileInfo.Name(), filePath))
			}
			return nil
		})

		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf(logCannotProcessFiles, dir))
			return err
		}
	}
	return nil
}

func loadNodeClassSpecificationFromFile(filename string) (*definition.NodeClassSpecification, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf(logCannotLoadFile, filename))
		return nil, err
	}

	spec := &definition.NodeClassSpecification{}
	err = yaml.Unmarshal(file, spec)
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf(logCannotParseFile, filename))

		return nil, err
	}

	// if no definitions are found, return an error and a nil Specification
	if len(spec.Definitions) == 0 {
		log.Warn().Err(err).Msg(fmt.Sprintf(logNoDefinitionsFoundInFile, filename))
		return nil, fmt.Errorf(logNoDefinitionsFoundInFile, filename)
	}

	// TODO debug
	return spec, nil
}

func LoadNodeClassDefinitions(sourceDir []string, fileExtension string, conn *sql.DB) error {
	return loadDefinitions(sourceDir, fileExtension, func(filePath string, _ os.FileInfo) (err error) {
		log.Debug().Msg(fmt.Sprintf(logAboutToLoadFile, filePath))

		spec, err := loadNodeClassSpecificationFromFile(filePath)
		if (err == nil) && (spec != nil) {
			log.Debug().Msg(fmt.Sprintf(logSuccessfullyLoadedFile, filePath))
			err = db.StoreNodeClassSpecification(conn, spec)
		} else {
			log.Warn().Msgf(logSkippingFile, filePath, err)
		}

		return nil
	})
}
