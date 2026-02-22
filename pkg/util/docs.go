package util

import (
	"fmt"

	"github.com/circa10a/shoutrrr/internal/meta"
)

// DocsURL returns a full documentation URL for the current version of Shoutrrr with the path appended.
// If the path contains a leading slash, it is stripped.
func DocsURL(path string) string {
	// strip leading slash if present
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// Use commit for dev builds, version for releases
	version := meta.GetVersion()
	if version == "unknown" || version == "dev" {
		version = meta.GetCommit()
	}

	return fmt.Sprintf("https://caleblemoine.dev/shoutrrr/%s/%s", version, path)
}
