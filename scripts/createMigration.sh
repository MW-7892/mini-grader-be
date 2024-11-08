#!/bin/bash

PROJECT_PATH="$(git rev-parse --show-toplevel)"
goose -dir "${PROJECT_PATH}/database/migrations" create $1 sql
