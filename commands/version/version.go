package version

import (
	"fmt"
	"os"
)

const (
	versionAPINewBreaking  = 0
	versionAPIAddedFeature = 0
	versionAPIBugFix       = 0
)

// V returns the version.
func V() (version string) {
	version = fmt.Sprintf("%s v%d.%d.%d", os.Args[0], versionAPINewBreaking, versionAPIAddedFeature, versionAPIBugFix)
	return
}
