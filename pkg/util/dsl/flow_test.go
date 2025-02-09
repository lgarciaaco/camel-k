/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dsl

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWriteYaml(t *testing.T) {
	// yaml in conventional form as marshalled by the go runtime
	yaml := `- from:
    parameters:
      period: 3600001
    steps:
    - to: log:info
    uri: timer:tick
`

	yamlReader := bytes.NewReader([]byte(yaml))
	flows, err := FromYamlDSL(yamlReader)
	assert.NoError(t, err)
	assert.NotNil(t, flows)
	assert.Len(t, flows, 1)

	flow := map[string]interface{}{}
	err = json.Unmarshal(flows[0].RawMessage, &flow)
	assert.NoError(t, err)

	assert.NotNil(t, flow["from"])
	assert.Nil(t, flow["xx"])

	data, err := ToYamlDSL(flows)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, yaml, string(data))
}
