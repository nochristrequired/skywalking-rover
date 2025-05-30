# Licensed to Apache Software Foundation (ASF) under one or more contributor
# license agreements. See the NOTICE file distributed with
# this work for additional information regarding copyright
# ownership. Apache Software Foundation (ASF) licenses this file to you under
# the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

GO_LINT = $(GO_PATH)/bin/golangci-lint

linter:
	$(GO_LINT) version || curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH)/bin v2.1.6

.PHONY: lint
lint: linter generate
	$(GO_LINT) run -v --timeout 5m ./...

.PHONY: safe-lint
safe-lint: linter generate
	git config --global --add safe.directory ${REPODIR}
	$(GO_LINT) run -v --timeout 5m ./...

.PHONY: container-lint
container-lint: COMMAND=lint
container-lint: container-command

.PHONY: container-safe-lint
container-safe-lint: COMMAND=safe-lint
container-safe-lint: container-command
