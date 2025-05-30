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

package module

import (
	"context"
	"reflect"
	"testing"
)

const testModuleName = "test"

func TestNewManager(t *testing.T) {
	mod := &testModule{conf: &testModuleConfig{}}
	closeCount := 0

	// new manager
	mgr := NewManager([]Module{mod}, func(error) {
		closeCount++
	})

	// find module
	findModule := mgr.FindModule(testModuleName)
	if !reflect.DeepEqual(findModule, mod) {
		t.Fatal("find module not correct")
	}

	// close modules
	mgr.ShutdownModules(nil)
	if closeCount == 0 {
		t.Fatalf("the shutdown method is not called")
	}
}

type testModuleConfig struct {
	Config
}

type testModule struct {
	conf *testModuleConfig
}

func (t *testModule) Name() string {
	return testModuleName
}

func (t *testModule) RequiredModules() []string {
	return nil
}

func (t *testModule) Config() ConfigInterface {
	return t.conf
}

func (t *testModule) Start(context.Context, *Manager) error {
	return nil
}

func (t *testModule) NotifyStartSuccess() {
}

func (t *testModule) Shutdown(context.Context, *Manager) error {
	return nil
}
