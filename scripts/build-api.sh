#!/usr/bin/env bash

set -e

buildPath=deploy/api/build

buildTags=()
if [[ "$TARGET" == "lambda" ]]; then
	buildTags+=(lambda.norpc)
	export GOOS=linux
	export GOARCH=arm64
	export CGO_ENABLED=0
fi
if [[ "$SWAGGER" == "true" ]]; then
	buildTags+=(swagger)
fi

# generate -tags
buildTags="${buildTags[@]}"
if [[ "$buildTags" != "" ]]; then
	buildTags="-tags ${buildTags// /,}"
fi

# DRY to build api services
gobuild() {
	svcPkg=$1
	svcName=$2
	rm -rf $buildPath/$svcName
	go build -trimpath -buildvcs=true $buildTags -ldflags "-s -w" -o $buildPath/$svcName/bootstrap $svcPkg

	# ! in case that we do not setup ssm parameter store cause we're poor :)
	# ! cp .env and stuff to build folder
	cp ./.env $buildPath/$svcName

	# # zip files for lambda
	# shopt -s dotglob # enable the globbing of hidden files for bash
	# shopt -u dotglob # disable the globbing of hidden files for bash

	setopt dotglob # enable the globbing of hidden files for zsh
	zip -j $buildPath/$svcName.zip $buildPath/$svcName/*
	setopt no_dotglob # disable the globbing of hidden files for zsh
}

# Start building
gobuild ./cmd/api main
