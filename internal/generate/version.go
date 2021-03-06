package generate

import (
	"fmt"
	"sort"
	"strings"

	"github.com/blang/semver"
)

type Versioned interface {
	version() string
}

type VersionLister interface {
	versionList() Versions
}

type Version struct {
	value   string
	version semver.Version
}

func VersionNew(value string) Version {
	value = strings.TrimSuffix(value, ".")

	version, err := semver.ParseTolerant(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse version '%s' : %s", value, err))
	}

	return Version{
		value:   value,
		version: version,
	}
}

func (v Version) GTE(other Version) bool {
	return v.version.GTE(other.version)
}

func (v Version) LTE(other Version) bool {
	return v.version.LTE(other.version)
}

func supportedByVersion(versioned Versioned, targetVersion *Version) bool {
	if versioned.version() == "" {
		return true
	}

	if VersionNew(versioned.version()).LTE(*targetVersion) {
		return true
	}

	return false
}

type Versions []Version

func (vv Versions) Len() int {
	return len(vv)
}

func (vv Versions) Less(i, j int) bool {
	return vv[i].version.LT(vv[j].version)
}

func (vv Versions) Swap(i, j int) {
	vv[i], vv[j] = vv[j], vv[i]
}

func (vv Versions) sort() {
	sort.Sort(vv)
}

func (vv Versions) dedupe() Versions {
	valuesMap := make(map[string]Version)
	for _, v := range vv {
		valuesMap[v.value] = v
	}

	var newVersions []Version
	for _, v := range valuesMap {
		newVersions = append(newVersions, v)
	}

	return newVersions
}

func (vv Versions) versionStringsGreaterThanOrEqual(version Version) []string {
	var eligibleVersions []string

	for _, possibleVersion := range vv {
		if possibleVersion.version.GTE(version.version) {
			eligibleVersions = append(eligibleVersions, possibleVersion.value)
		}
	}

	return eligibleVersions
}

func (vv Versions) String() string {
	str := ""
	for i, v := range vv {
		str += v.value
		if i < len(vv)-1 {
			str += " "
		}
	}

	return str
}
