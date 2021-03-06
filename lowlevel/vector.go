package lowlevel

/*
#include <fmod_common.h>
*/
import "C"
import "unsafe"

type Vector struct {
	X, Y, Z float32
}

func NewVector() Vector {
	return Vector{}
}

func (v *Vector) fromC(cv C.FMOD_VECTOR) {
	v.X = float32(cv.x)
	v.Y = float32(cv.y)
	v.Z = float32(cv.z)
}

func (v *Vector) toC() C.FMOD_VECTOR {
	var cv C.FMOD_VECTOR
	cv.x = C.float(v.X)
	cv.y = C.float(v.Y)
	cv.z = C.float(v.Z)
	return cv
}

func (v *Vector) toCp() **C.FMOD_VECTOR {
	var cv C.FMOD_VECTOR
	cv.x = C.float(v.X)
	cv.y = C.float(v.Y)
	cv.z = C.float(v.Z)
	data := *(**C.FMOD_VECTOR)(unsafe.Pointer(&cv))
	return &data
}
