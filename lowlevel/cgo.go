package lowlevel

/*
#cgo pkg-config: --define-variable=prefix=.. fmod
*/
import "C"
import "runtime"

func init() {
	runtime.LockOSThread()
}
