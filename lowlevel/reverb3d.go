package lowlevel

/*
#include <fmod.h>
*/
import "C"
import "unsafe"

type Reverb3D struct {
	cptr *C.FMOD_REVERB3D
}

/*
   'Reverb3D' API
*/

// Releases the memory for a reverb object and makes it inactive.
// If no reverb objects are created, the ambient reverb will be the only audible reverb. By default this ambient reverb setting is set to OFF.
func (r *Reverb3D) Release() error {
	res := C.FMOD_Reverb3D_Release(r.cptr)
	return errs[res]
}

/*
   Reverb manipulation.
*/

// Sets the 3d properties of a 'virtual' reverb object.
//
// position: Pointer to a vector containing the 3d position of the center of the reverb in 3d space. Default = { 0,0,0 }.
//
// mindistance: The distance from the centerpoint that the reverb will have full effect at. Default = 0.0.
//
// maxdistance: The distance from the centerpoint that the reverb will not have any effect. Default = 0.0.
//
//The 3D reverb object is a sphere having 3D attributes (position, minimum distance, maximum distance) and reverb properties.
//
// The properties and 3D attributes of all reverb objects collectively determine, along with the listener's position, the settings of and input gains into a single 3D reverb DSP.
//
// Please note that this only applies to software channels. When the listener is within the sphere of effect of one or more 3d reverbs, the listener's 3D reverb properties are
// a weighted combination of such 3d reverbs. When the listener is outside all of the reverbs, the 3D reverb setting is set to the default ambient reverb setting.
func (r *Reverb3D) Set3DAttributes(position Vector, mindistance, maxdistance float64) error {
	cposition := position.toC()
	res := C.FMOD_Reverb3D_Set3DAttributes(r.cptr, &cposition, C.float(mindistance), C.float(maxdistance))
	return errs[res]
}

// Retrieves the 3d attributes of a Reverb object.
//
// The 3D reverb object is a sphere having 3D attributes (position, minimum distance, maximum distance) and reverb properties.
//
// The properties and 3D attributes of all reverb objects collectively determine, along with the listener's position, the settings of and input gains into a single 3D reverb DSP.
//
// Please note that this only applies to software channels. When the listener is within the sphere of effect of one or more 3d reverbs, the listener's 3D reverb properties are
// a weighted combination of such 3d reverbs. When the listener is outside all of the reverbs, the 3D reverb setting is set to the default ambient reverb setting.
func (r *Reverb3D) Get3DAttributes() (Vector, float64, float64, error) {
	var cposition C.FMOD_VECTOR
	var mindistance, maxdistance C.float
	var position Vector
	res := C.FMOD_Reverb3D_Get3DAttributes(r.cptr, &cposition, &mindistance, &maxdistance)
	position.fromC(cposition)
	return position, float64(mindistance), float64(maxdistance), errs[res]
}

// Sets reverb parameters for the current reverb object.
// Reverb parameters can be set manually, or automatically using the pre-defined presets given in the fmod.h header.
//
// properties:  "ReverbProperties" structure which defines the attributes for the reverb
func (r *Reverb3D) SetProperties(properties ReverbProperties) error {
	cproperties := properties.toC()
	res := C.FMOD_Reverb3D_SetProperties(r.cptr, &cproperties)
	return errs[res]
}

// Retrieves the current reverb environment.
func (r *Reverb3D) Properties() (*ReverbProperties, error) {
	var cproperties C.FMOD_REVERB_PROPERTIES
	properties := NewReverbProperties()
	res := C.FMOD_Reverb3D_GetProperties(r.cptr, &cproperties)
	properties.fromC(cproperties)
	return properties, errs[res]
}

// Disables or enables a reverb object so that it does or does not contribute to the 3d scene.
func (r *Reverb3D) SetActive(active bool) error {
	res := C.FMOD_Reverb3D_SetActive(r.cptr, getBool(active))
	return errs[res]
}

// Retrieves the active state of the reverb object.
func (r *Reverb3D) IsActive() (bool, error) {
	var active C.FMOD_BOOL
	res := C.FMOD_Reverb3D_GetActive(r.cptr, &active)
	return setBool(active), errs[res]
}

/*
   Userdata set/get.
*/

// Sets a user value that the Reverb object will store internally. Can be retrieved with "Reverb.UserData".
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
//
// It can be useful if an FMOD callback passes an object of this type as a parameter,
// and the user does not know which object it is (if many of these types of objects exist). Using Reverb::getUserData would help in the identification of the object.
func (r *Reverb3D) SetUserData(userdata interface{}) error {
	data := *(*[]*C.char)(unsafe.Pointer(&userdata))
	res := C.FMOD_Reverb3D_SetUserData(r.cptr, unsafe.Pointer(&data))
	return errs[res]
}

// Retrieves the user value that that was set by calling the "Reverb.SetUserData" function.
func (r *Reverb3D) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_Reverb3D_GetUserData(r.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}
