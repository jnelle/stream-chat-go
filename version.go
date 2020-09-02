package stream_chat //nolint: golint

import (
	"fmt"
)

const (
	versionMajor = 2
	versionMinor = 4
	versionPatch = 2
)

// Version returns the version of the library.
func Version() string {
	return "v" + fmtVersion()
}

func versionHeader() string {
	return "stream-go-client-" + fmtVersion()
}

func fmtVersion() string {
	return fmt.Sprintf("%d.%d.%d",
		versionMajor,
		versionMinor,
		versionPatch)
}
