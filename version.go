package main

// version is set from the release tag by GoReleaser in the release pipeline.
var version = "dev"

func VersionString() string {
	return version
}
