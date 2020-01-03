package main

import (
	"strings"
	"testing"
)

func TestGoVersions(t *testing.T) {
	var versions []Version

	versions = SupportedPlatforms("go1.10")
	if !reflect.DeepEqual(versions, Platforms_1_0) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedPlatforms("go1.11")
	if !reflect.DeepEqual(versions, Platforms_1_1) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedPlatforms("go1.12")
	if !reflect.DeepEqual(versions, Platforms_1_2) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedPlatforms("go1.13")
	if !reflect.DeepEqual(versions, Platforms_1_3) {
		t.Fatalf("bad: %#v", versions)
	}
}
