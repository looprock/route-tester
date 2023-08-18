#!/bin/sh
set -x
gcloud auth print-access-token | docker login -u oauth2accesstoken \
    --password-stdin https://gcr.io
PLATFORMS="linux/amd64"
docker buildx build \
	-t "gcr.io/trackhearing-dev/route-tester:latest" \
	--platform "${PLATFORMS}" \
	--push \
	.
