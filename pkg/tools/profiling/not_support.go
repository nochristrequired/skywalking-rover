// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package profiling

import (
	"fmt"
)

var ErrNotSupport = fmt.Errorf("not support")

type NotSupport struct {
}

func NewNotSupport() *NotSupport {
	return &NotSupport{}
}

func (l *NotSupport) IsSupport(string) bool {
	// handle all not support file
	return true
}

func (l *NotSupport) AnalyzeSymbols(string) ([]*Symbol, error) {
	return nil, ErrNotSupport
}

func (l *NotSupport) ToModule(_ int32, _, _ string, _ []*ModuleRange) (*Module, error) {
	return nil, ErrNotSupport
}
