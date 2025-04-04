// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/DePINNetwork/depin-sdk/data/transactions/logic"
)

// TEAL syntax highlighter as TM grammar.
// See the following resources for more info:
//
// 1. tmLanguage grammar
// https://raw.githubusercontent.com/martinring/tmlanguage/master/tmlanguage.json
// 2. tmLanguage description
// https://macromates.com/manual/en/language_grammars
// 3. Oniguruma regular expressions
// https://macromates.com/manual/en/regular_expressions

type tmLanguage struct {
	Schema     string             `json:"$schema,omitempty"`
	Name       string             `json:"name"`
	Patterns   []pattern          `json:"patterns,omitempty"`
	Repository map[string]pattern `json:"repository,omitempty"`
	ScopeName  string             `json:"scopeName"`
}

type pattern struct {
	Include  string             `json:"include,omitempty"`
	Name     string             `json:"name,omitempty"`
	Begin    string             `json:"begin,omitempty"`
	End      string             `json:"end,omitempty"`
	Match    string             `json:"match,omitempty"`
	Captures map[string]pattern `json:"captures,omitempty"`
	Patterns []pattern          `json:"patterns,omitempty"`
}

func buildSyntaxHighlight(version uint64) *tmLanguage {
	tm := tmLanguage{
		Schema:    "https://raw.githubusercontent.com/martinring/tmlanguage/master/tmlanguage.json",
		Name:      "Algorand TEAL",
		ScopeName: "source.teal",
		Patterns: []pattern{
			{Include: "#invalid"},
			{Include: "#comments"},
			{Include: "#strings"},
			{Include: "#literals"},
			{Include: "#labels"},
			{Include: "#keywords"},
			{Include: "#pragmas"},
		},
		Repository: map[string]pattern{},
	}
	tm.Repository["invalid"] = pattern{
		Patterns: []pattern{{
			Name:  "invalid.illegal.teal",
			Match: "^\\s+.*$",
		}},
	}
	tm.Repository["comments"] = pattern{
		Name:  "comment.line.double-slash.teal",
		Begin: "//",
		End:   "$",
	}
	tm.Repository["strings"] = pattern{
		Name:  "string.quoted.double.teal",
		Begin: "\"",
		End:   "\"",
		Patterns: []pattern{{
			Name:  "constant.character.escape.teal",
			Match: "\\\\(x[0-9A-Fa-f]{2}|.|$)",
		}},
	}
	tm.Repository["pragmas"] = pattern{
		Name:  "support.function.teal",
		Match: "^#pragma\\b.*$",
	}
	tm.Repository["labels"] = pattern{
		Patterns: []pattern{
			{
				Name:  "support.variable.teal",
				Match: "^\\w+:.*$",
			},
			{
				Match: "\\b(?<=b|bz|bnz)\\s+(\\w+)\\b",
				Captures: map[string]pattern{
					"1": {Name: "support.variable.teal"},
				},
			},
		},
	}
	literals := pattern{
		Patterns: []pattern{
			{
				Name:  "constant.numeric.teal",
				Match: "\\b([0-9]+)\\b",
			},
			{
				Name:  "constant.numeric.teal",
				Match: "\\b(?<=int\\s+)(0x[0-9]+)\\b",
			},
			{
				Name:  "string.quoted.double.teal",
				Match: "\\b(?<=byte\\s+)(0x[0-9]+)\\b",
			},
		},
	}
	var allNamedFields []string
	allNamedFields = append(allNamedFields, logic.TxnTypeNames[:]...)
	allNamedFields = append(allNamedFields, logic.OnCompletionNames[:]...)
	accumulated := make(map[string]bool)
	opSpecs := logic.OpcodesByVersion(version)
	for _, spec := range opSpecs {
		for _, imm := range spec.OpDetails.Immediates {
			if imm.Group != nil && !accumulated[imm.Group.Name] {
				for _, name := range imm.Group.Names {
					spec, ok := imm.Group.SpecByName(name)
					if !ok || spec.Version() > version {
						continue
					}
					allNamedFields = append(allNamedFields, name)
				}
				accumulated[imm.Group.Name] = true
			}
		}
	}

	var seen = make(map[string]bool, len(allNamedFields))
	var dedupe = make([]string, 0, len(allNamedFields))
	for _, name := range allNamedFields {
		if name != "" && !seen[name] {
			dedupe = append(dedupe, name)
		}
		seen[name] = true
	}

	literals.Patterns = append(literals.Patterns, pattern{
		Name:  "variable.parameter.teal",
		Match: fmt.Sprintf("\\b(%s)\\b", strings.Join(dedupe, "|")),
	})
	tm.Repository["literals"] = literals

	keywords := pattern{
		Patterns: []pattern{
			{
				Match: "\\b(base64|b64|base32|b32)(?:\\(|\\s+)([a-zA-Z0-9\\+\\/\\=]+)(?:\\)|\\s?|$)",
				Captures: map[string]pattern{
					"1": {Name: "support.class.teal"},
					"2": {Name: "string.quoted.triple.teal"},
				},
			},
			{
				Match: "^(addr)\\s+([A-Z2-7\\=]+)",
				Captures: map[string]pattern{
					"1": {Name: "keyword.other.teal"},
					"2": {Name: "string.unquoted.teal"},
				},
			},
		},
	}
	var allAccess []string
	var allArithmetics []string

	var keys []string
	for key := range logic.OpGroups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, grp := range keys {
		names := logic.OpGroups[grp]
		sort.Strings(names)
		switch grp {
		case "Flow Control":
			keywords.Patterns = append(keywords.Patterns, pattern{
				Name:  "keyword.control.teal",
				Match: fmt.Sprintf("^(%s)\\b", strings.Join(names, "|")),
			})
		case "Loading Values":
			loading := []string{"int", "byte", "addr"}
			loading = append(loading, names...)
			keywords.Patterns = append(keywords.Patterns, pattern{
				Name:  "keyword.other.teal",
				Match: fmt.Sprintf("^(%s)\\b", strings.Join(loading, "|")),
			})
		case "State Access", "Box Access":
			allAccess = append(allAccess, names...)
		// For these, accumulate into allArithmetics,
		// and only add to keyword.Patterns later, when all
		// have been collected.
		case "Arithmetic", "Byte Array Manipulation", "Byte Array Arithmetic",
			"Byte Array Logic", "Cryptography", "Inner Transactions":
			escape := map[rune]bool{
				'*': true,
				'+': true,
				'=': true,
				'-': true,
				'!': true,
				'~': true,
				'^': true,
				'$': true,
				'?': true,
				'|': true,
				'<': true,
				'>': true,
			}
			for _, op := range names {
				escaped := make([]byte, 0, len(op)*2)
				for _, ch := range op {
					if _, ok := escape[ch]; ok {
						escaped = append(escaped, '\\')
					}
					escaped = append(escaped, byte(ch))
				}
				allArithmetics = append(allArithmetics, string(escaped))
			}
		default:
			panic(fmt.Sprintf("Unknown ops group: %s", grp))
		}
	}
	keywords.Patterns = append(keywords.Patterns, pattern{
		Name:  "keyword.other.unit.teal",
		Match: fmt.Sprintf("^(%s)\\b", strings.Join(allAccess, "|")),
	})
	keywords.Patterns = append(keywords.Patterns, pattern{
		Name:  "keyword.operator.teal",
		Match: fmt.Sprintf("^(%s)\\b", strings.Join(allArithmetics, "|")),
	})
	tm.Repository["keywords"] = keywords

	return &tm
}
