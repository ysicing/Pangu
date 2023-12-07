#!/usr/bin/env bash

addlicense -f hack/licenses/licenses.tpl -ignore web/** -ignore "**/*.md" -ignore vendor/** -ignore "**/*.yml" -ignore "**/*.yaml" -ignore "**/*.rb" -ignore "**/*.sh" ./**
