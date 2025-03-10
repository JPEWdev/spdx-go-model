//
//
// This file was automatically generated by shacl2code. DO NOT MANUALLY MODIFY IT
//
// SPDX-License-Identifier: Apache-2.0

package spdx_v3_0_1

import (
    "strings"
)

// IRI Validation
func IsIRI(iri string) bool {
    if strings.HasPrefix(iri, "_:") {
        return false
    }
    if strings.Contains(iri, ":") {
        return true
    }
    return false
}

func IsBlankNode(iri string) bool {
    return strings.HasPrefix(iri, "_:")
}
