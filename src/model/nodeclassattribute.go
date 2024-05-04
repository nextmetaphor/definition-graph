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

package model

type (
	NodeClassAttributes []NodeClassAttribute

	NodeClassAttributeKey struct {
		ID                 string `json:"ID"`
		NodeClassID        string `json:"nodeClassID"`
		NodeClassNamespace string `json:"nodeClassNamespace"`
	}

	NodeClassAttribute struct {
		NodeClassAttributeKey
		Type        string  `json:"type"`
		IsRequired  int     `json:"isRequired"`
		Description *string `json:"description,omitempty"`
	}
)
