package lowlevel

/*
#include <fmod_common.h>
*/
import "C"
import "unsafe"

var null = unsafe.Pointer(uintptr(0))

func getBool(b bool) C.FMOD_BOOL {
	if b {
		return 1
	}
	return 0
}

func setBool(b C.FMOD_BOOL) bool {
	if b == 1 {
		return true
	}
	return false
}
