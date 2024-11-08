#!/bin/bash

PROJECT_PATH="$(git rev-parse --show-toplevel)"
CURRENT_PATH="$(pwd)"

cd "${PROJECT_PATH}"
go run "${PROJECT_PATH}/scripts/scriptMigrateUp/main.go"
cd "${CURRENT_PATH}"
