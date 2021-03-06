package nanomsg

// #include <nanomsg/nn.h>
import "C"

// Version holds the nanomsg version which is used.
var Version struct {
	Major int
	Minor int
	Patch int
}

func init() {
	Version.Major = int(C.NN_VERSION_MAJOR)
	Version.Minor = int(C.NN_VERSION_MINOR)
	Version.Patch = int(C.NN_VERSION_PATCH)
}
