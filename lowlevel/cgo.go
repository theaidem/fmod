package lowlevel

/*
#cgo pkg-config: fmod
*/
import "C"
import "runtime"

func init() {
	runtime.LockOSThread()
}
