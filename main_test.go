package main

import (
	"strings"
	"testing"
)

//Version combines the OS and architecture
type Version struct {
	OS   string
	Arch string

	Default bool // indicates whether Version is a build target
}

func (vs *Version) String() string {
	return fmt.Sprintf("%s/%s", vs.OS, vs.Arch)
}

var (
	Versions_1_0 = []Version{
		{"darwin", "amd64", true},
		{"linux", "amd64", true},
		{"windows", "amd64", true},
	}

	Versions_1_1 = append(Versions_1_0)
	Versions_1_2 = append(Versions_1_1)
	Versions_1_3 = append(Versions_1_2)
)

func TestGoVersions(t *testing.T) {
	var versions []Version

	versions = SupportedVersions("go1.10")
	if !reflect.DeepEqual(versions, Versions_1_0) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedVersions("go1.11")
	if !reflect.DeepEqual(versions, Versions_1_1) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedVersions("go1.12")
	if !reflect.DeepEqual(versions, Versions_1_2) {
		t.Fatalf("bad: %#v", versions)
	}

	versions = SupportedVersions("go1.13")
	if !reflect.DeepEqual(versions, Versions_1_3) {
		t.Fatalf("bad: %#v", versions)
	}
}
