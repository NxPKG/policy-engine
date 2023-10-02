// © 2022-2023 Khulnasoft Limited All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package embed

import (
	_ "embed"
)

//go:embed khulnasoft.rego
var khulnasoftRego []byte

//go:embed khulnasoft/internal/relations_cache.rego
var khulnasoftRelationsCache []byte

//go:embed khulnasoft/terraform.rego
var khulnasoftTerraformRego []byte

//go:embed khulnasoft/relations.rego
var khulnasoftRelationsRego []byte

//go:embed khulnasoft/internal/relations.rego
var khulnasoftInternalRelationsRego []byte

var KhulnasoftBuiltins map[string][]byte = map[string][]byte{
	"khulnasoft.rego":                          khulnasoftRego,
	"khulnasoft/internal/relations_cache.rego": khulnasoftRelationsCache,
}

var KhulnasoftLib map[string][]byte = map[string][]byte{
	"khulnasoft/terraform.rego":          khulnasoftTerraformRego,
	"khulnasoft/relations.rego":          khulnasoftRelationsRego,
	"khulnasoft/internal/relations.rego": khulnasoftInternalRelationsRego,
}
