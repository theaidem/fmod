package lowlevel

/*
#include <fmod.h>
*/
import "C"
import "unsafe"

type DspConnection struct {
	cptr *C.FMOD_DSPCONNECTION
}

/*
   'DSPConnection' API
*/

// Retrieves the DSP unit that is the input of this connection.
// A DSPConnection joins 2 DSP units together (think of it as the line between 2 circles).
// Each DSPConnection has 1 input and 1 output.
//
// Note! If a "DSP.AddInput" just occurred, the connection might not be ready because the DSP system is still queued to connect in the background.
// If so the function will return FMOD_ERR_NOTREADY and the input will be null. Poll until it is ready.
func (d *DspConnection) Input() (DSP, error) {
	var input DSP
	res := C.FMOD_DSPConnection_GetInput(d.cptr, &input.cptr)
	return input, errs[res]
}

// Retrieves the DSP unit that is the output of this connection.
// A DSPConnection joins 2 DSP units together (think of it as the line between 2 circles).
// Each DSPConnection has 1 input and 1 output.
//
// Note! If a "DSP.AddInput" just occurred, the connection might not be ready because the DSP system is still queued to connect in the background.
// If so the function will return FMOD_ERR_NOTREADY and the input will be null. Poll until it is ready.
func (d *DspConnection) Output() (DSP, error) {
	var output DSP
	res := C.FMOD_DSPConnection_GetOutput(d.cptr, &output.cptr)
	return output, errs[res]
}

// Sets the volume of the connection so that the input is scaled by this value before being passed to the output.
//
// volume: Volume or mix level of the connection. 0.0 = silent, 1.0 = full volume.
func (d *DspConnection) SetMix(volume float64) error {
	res := C.FMOD_DSPConnection_SetMix(d.cptr, C.float(volume))
	return errs[res]
}

// Retrieves the volume of the connection - the scale level of the input before being passed to the output.
func (d *DspConnection) Mix() (float64, error) {
	var volume C.float
	res := C.FMOD_DSPConnection_GetMix(d.cptr, &volume)
	return float64(volume), errs[res]
}

// NOTE: Not implement yet
// Sets a NxN panning matrix on a DSP connection.
// Skipping/hop is supported, so memory for the matrix can be wider than the width of the inchannels parameter.
//
// matrix: Pointer to an array of floating point matrix data, where rows represent output speakers, and columns represent input channels.
//
// outchannels: Number of output channels in the matrix being specified.
//
// inchannels: Number of input channels in the matrix being specified.
//
// inchannel_hop: Number of floating point values stored in memory for a row, so that the memory can be skipped through correctly to read the right values,
// if the intended matrix memory to be read from is wider than the matrix stored in the DSPConnection.
func (d *DspConnection) SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_SetMixMatrix       (FMOD_DSPCONNECTION *dspconnection, float *matrix, int outchannels, int inchannels, int inchannel_hop);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Returns the panning matrix set by the user, for a connection.
func (d *DspConnection) MixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop *C.int) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_GetMixMatrix       (FMOD_DSPCONNECTION *dspconnection, float *matrix, int *outchannels, int *inchannels, int inchannel_hop);
	return ErrNoImpl
}

// Returns the type of the connection between 2 DSP units.
// This can be "DSPCONNECTION_TYPE_STANDARD", "DSPCONNECTION_TYPE_SIDECHAIN", "DSPCONNECTION_TYPE_SEND" or "DSPCONNECTION_TYPE_SEND_SIDECHAIN".
func (d *DspConnection) Type() (DSPConnectionType, error) {
	var typ C.FMOD_DSPCONNECTION_TYPE
	res := C.FMOD_DSPConnection_GetType(d.cptr, &typ)
	return DSPConnectionType(typ), errs[res]
}

/*
   Userdata set/get.
*/

// Sets a user value that the DSPConnection object will store internally. Can be retrieved with "DSPConnection.UserData".
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
func (d *DspConnection) SetUserData(userdata interface{}) error {
	res := C.FMOD_DSPConnection_SetUserData(d.cptr, unsafe.Pointer(&userdata))
	return errs[res]
}

//Retrieves the user value that that was set by calling the  DSPConnection.SetUserData.
func (d *DspConnection) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_DSPConnection_GetUserData(d.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}
