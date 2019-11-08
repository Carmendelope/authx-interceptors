/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package config

// Permission is a set of rules that uses the define primitive of the system.
type Permission struct {
	// Must is a list of primitives that the role MUST contains. If the role doesn't include
	// any primitive the role is not authorized.
	Must []string `json:"must,omitempty"`
	// Should is a list of primitives that the role SHOULD contains. The role must include at least one primitive.
	Should []string `json:"should,omitempty"`
	// MustNot is a list of primitive that the role MUST NOT include. The role must not include any primitive.
	MustNot []string `json:"must_not,omitempty"`
}

// Valid verifies if a list of primitives are valid for a set of rules.
func (p *Permission) Valid(primitives []string) bool {
	for _, must := range p.Must {
		check := false
		for _, pri := range primitives {
			if !check && pri == must {
				check = true
			}
		}
		if !check {
			return false
		}
	}
	counter := 0
	for _, should := range p.Should {
		for _, pri := range primitives {
			if pri == should {
				counter++
			}

		}
	}
	if len(p.Should) > 0 && counter == 0 {
		return false
	}
	for _, mustNo := range p.MustNot {
		for _, pri := range primitives {
			if pri == mustNo {
				return false
			}
		}
	}
	return true
}
