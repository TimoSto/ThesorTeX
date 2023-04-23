package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int64
	Minor int64
	Patch int64
}

func (v Version) ToString() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}

func ParseString(version string) (Version, error) {
	// rm possible leading v
	version = strings.TrimLeft(version, "v")

	var ver Version

	dotParts := strings.SplitN(version, ".", 3)
	if len(dotParts) != 3 {
		return ver, fmt.Errorf("%s is not in dotted-tri format", version)
	}

	parsed := make([]int64, 3)

	for i, v := range dotParts[:3] {
		val, err := strconv.ParseInt(v, 10, 64)
		parsed[i] = val
		if err != nil {
			return ver, err
		}
	}

	ver.Major = parsed[0]
	ver.Minor = parsed[1]
	ver.Patch = parsed[2]

	return ver, nil
}

func (v Version) Slice() []int64 {
	return []int64{v.Major, v.Minor, v.Patch}
}

func Compare(versionA Version, versionB Version) int {
	return sliceCompare(versionA.Slice(), versionB.Slice())
}

func sliceCompare(v1 []int64, v2 []int64) int {
	if len(v1) == 0 {
		return 0
	}

	a := v1[0]
	b := v2[0]

	if a < b {
		return 1
	} else if a > b {
		return -1
	}

	return sliceCompare(v1[1:], v2[1:])
}
