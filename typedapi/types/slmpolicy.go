// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e16324dcde9297dd1149c1ef3d6d58afe272e646

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SLMPolicy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e16324dcde9297dd1149c1ef3d6d58afe272e646/specification/slm/_types/SnapshotLifecycle.ts#L76-L82
type SLMPolicy struct {
	Config     *Configuration `json:"config,omitempty"`
	Name       string         `json:"name"`
	Repository string         `json:"repository"`
	Retention  *Retention     `json:"retention,omitempty"`
	Schedule   string         `json:"schedule"`
}

func (s *SLMPolicy) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "config":
			if err := dec.Decode(&s.Config); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = o

		case "retention":
			if err := dec.Decode(&s.Retention); err != nil {
				return err
			}

		case "schedule":
			if err := dec.Decode(&s.Schedule); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSLMPolicy returns a SLMPolicy.
func NewSLMPolicy() *SLMPolicy {
	r := &SLMPolicy{}

	return r
}
