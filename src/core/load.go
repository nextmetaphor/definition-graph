/*
 * Copyright (C) 2024 Paul Tatham <paul@nextmetaphor.io>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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

func loadNodeSpecificationFromFile(filename string) (*definition.NodeSpecification, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf(logCannotLoadFile, filename))
		return nil, err
	}

	spec := &definition.NodeSpecification{}
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

func LoadNodeDefinitionsWithoutEdges(sourceDir []string, fileExtension string, conn *sql.DB) error {
	return loadDefinitions(sourceDir, fileExtension, func(filePath string, _ os.FileInfo) (err error) {
		log.Debug().Msg(fmt.Sprintf(logAboutToLoadFile, filePath))

		spec, err := loadNodeSpecificationFromFile(filePath)
		if (err == nil) && (spec != nil) {
			log.Debug().Msg(fmt.Sprintf(logSuccessfullyLoadedFile, filePath))
			err = db.StoreNodeSpecificationWithoutEdges(conn, spec)
		} else {
			log.Warn().Msgf(logSkippingFile, filePath, err)
		}

		return nil
	})
}

func LoadNodeDefinitionsOnlyEdges(sourceDir []string, fileExtension string, conn *sql.DB) error {
	return loadDefinitions(sourceDir, fileExtension, func(filePath string, _ os.FileInfo) (err error) {
		log.Debug().Msg(fmt.Sprintf(logAboutToLoadFile, filePath))

		spec, err := loadNodeSpecificationFromFile(filePath)
		if (err == nil) && (spec != nil) {

			log.Debug().Msg(fmt.Sprintf(logSuccessfullyLoadedFile, filePath))
			err = db.StoreNodeSpecificationOnlyEdges(conn, spec)
		} else {
			log.Warn().Msgf(logSkippingFile, filePath, err)
		}

		return nil
	})
}
