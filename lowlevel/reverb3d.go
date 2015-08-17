package lowlevel

/*
#include <fmod.h>
*/
import "C"

type Reverb3D struct {
	cptr *C.FMOD_REVERB3D
}

/*
   'Reverb3D' API
*/

func (r *Reverb3D) Release() error {
	res := C.FMOD_Reverb3D_Release(r.cptr)
	return errs[res]
}

/*
   Reverb manipulation.
*/

func (r *Reverb3D) Set3DAttributes(position Vector, mindistance, maxdistance float64) error {
	cposition := position.toC()
	res := C.FMOD_Reverb3D_Set3DAttributes(r.cptr, &cposition, C.float(mindistance), C.float(maxdistance))
	return errs[res]
}

func (r *Reverb3D) Get3DAttributes() (Vector, float64, float64, error) {
	var cposition C.FMOD_VECTOR
	var mindistance, maxdistance C.float
	var position Vector
	res := C.FMOD_Reverb3D_Get3DAttributes(r.cptr, &cposition, &mindistance, &maxdistance)
	position.fromC(cposition)
	return position, float64(mindistance), float64(maxdistance), errs[res]
}

func (r *Reverb3D) SetProperties(properties ReverbProperties) error {
	cproperties := properties.toC()
	res := C.FMOD_Reverb3D_SetProperties(r.cptr, &cproperties)
	return errs[res]
}

func (r *Reverb3D) Properties() (ReverbProperties, error) {
	var cproperties C.FMOD_REVERB_PROPERTIES
	properties := NewReverbProperties()
	res := C.FMOD_Reverb3D_GetProperties(r.cptr, &cproperties)
	properties.fromC(cproperties)
	return properties, errs[res]
}

func (r *Reverb3D) SetActive(active bool) error {
	res := C.FMOD_Reverb3D_SetActive(r.cptr, getBool(active))
	return errs[res]
}

func (r *Reverb3D) IsActive() (bool, error) {
	var active C.FMOD_BOOL
	res := C.FMOD_Reverb3D_GetActive(r.cptr, &active)
	return setBool(active), errs[res]
}

/*
   Userdata set/get.
*/

func (r *Reverb3D) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_Reverb3D_SetUserData             (FMOD_REVERB3D *reverb3d, void *userdata);
	return ErrNoImpl
}

func (r *Reverb3D) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_Reverb3D_GetUserData             (FMOD_REVERB3D *reverb3d, void **userdata);
	return ErrNoImpl
}
