##!/usr/bin/env sh
set -x on

app_name='go-1CNativeAddin'
namespace='khorevaa'

docker run --rm --privileged  -v $PWD:/go/src/khorevaa/1CNativeAddin -v /var/run/docker.sock:/var/run/docker.sock -w /go/src/khorevaa/1CNativeAddin  goreleaser/goreleaser:latest-cgo goreleaser --snapshot --skip-publish --rm-dist
#  goreleaser/goreleaser:latest release --snapshot --skip-publish --rm-dist