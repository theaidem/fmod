package lowlevel

/*
#include <fmod.h>
*/
import "C"

type DspConnection struct {
	cptr *C.FMOD_DSPCONNECTION
}

/*
   'DSPConnection' API
*/

func (d *DspConnection) Input() (DSP, error) {
	var input DSP
	res := C.FMOD_DSPConnection_GetInput(d.cptr, &input.cptr)
	return input, errs[res]
}

func (d *DspConnection) Output() (DSP, error) {
	var output DSP
	res := C.FMOD_DSPConnection_GetOutput(d.cptr, &output.cptr)
	return output, errs[res]
}

func (d *DspConnection) SetMix(volume float64) error {
	res := C.FMOD_DSPConnection_SetMix(d.cptr, C.float(volume))
	return errs[res]
}

func (d *DspConnection) Mix() (float64, error) {
	var volume C.float
	res := C.FMOD_DSPConnection_GetMix(d.cptr, &volume)
	return float64(volume), errs[res]
}

// NOTE: Not implement yet
func (d *DspConnection) SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_SetMixMatrix       (FMOD_DSPCONNECTION *dspconnection, float *matrix, int outchannels, int inchannels, int inchannel_hop);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (d *DspConnection) MixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop *C.int) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_GetMixMatrix       (FMOD_DSPCONNECTION *dspconnection, float *matrix, int *outchannels, int *inchannels, int inchannel_hop);
	return ErrNoImpl
}

func (d *DspConnection) Type() (DSPConnectionType, error) {
	var typ C.FMOD_DSPCONNECTION_TYPE
	res := C.FMOD_DSPConnection_GetType(d.cptr, &typ)
	return DSPConnectionType(typ), errs[res]
}

/*
   Userdata set/get.
*/

// NOTE: Not implement yet
func (d *DspConnection) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_SetUserData        (FMOD_DSPCONNECTION *dspconnection, void *userdata);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (d *DspConnection) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_DSPConnection_GetUserData        (FMOD_DSPCONNECTION *dspconnection, void **userdata);
	return ErrNoImpl
}
