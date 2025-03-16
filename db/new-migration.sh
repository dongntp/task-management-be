#!/bin/bash

atlas migrate diff "$1" --dir "file://db/migrations" --to "file://db/schema.sql" --dev-url "docker://postgres/15/dev"
