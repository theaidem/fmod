package lowlevel

/*
#include <fmod.h>
*/
import "C"

type DSP struct {
	cptr *C.FMOD_DSP
}

/*
   'DSP' API
*/

func (d *DSP) Release() error {
	res := C.FMOD_DSP_Release(d.cptr)
	return errs[res]
}

func (d *DSP) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_DSP_GetSystemObject(d.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   Connection / disconnection / input and output enumeration.
*/

func (d *DSP) AddInput(input DSP, typ DSPConnectionType) (DspConnection, error) {
	var dspConn DspConnection
	res := C.FMOD_DSP_AddInput(d.cptr, input.cptr, &dspConn.cptr, C.FMOD_DSPCONNECTION_TYPE(typ))
	return dspConn, errs[res]
}

func (d *DSP) DisconnectFrom(target DSP, connection DspConnection) error {
	res := C.FMOD_DSP_DisconnectFrom(d.cptr, target.cptr, connection.cptr)
	return errs[res]
}

func (d *DSP) DisconnectAll(inputs, outputs bool) error {
	res := C.FMOD_DSP_DisconnectAll(d.cptr, getBool(inputs), getBool(outputs))
	return errs[res]
}

func (d *DSP) NumInputs() (int, error) {
	var numinputs C.int
	res := C.FMOD_DSP_GetNumInputs(d.cptr, &numinputs)
	return int(numinputs), errs[res]
}

func (d *DSP) NumOutputs() (int, error) {
	var numoutputs C.int
	res := C.FMOD_DSP_GetNumOutputs(d.cptr, &numoutputs)
	return int(numoutputs), errs[res]
}

func (d *DSP) Input(index int) (DSP, DspConnection, error) {
	var input DSP
	var inputconnection DspConnection
	res := C.FMOD_DSP_GetInput(d.cptr, C.int(index), &input.cptr, &inputconnection.cptr)
	return input, inputconnection, errs[res]
}

func (d *DSP) Output(index int) (DSP, DspConnection, error) {
	var output DSP
	var outputconnection DspConnection
	res := C.FMOD_DSP_GetOutput(d.cptr, C.int(index), &output.cptr, &outputconnection.cptr)
	return output, outputconnection, errs[res]
}

/*
   DSP unit control.
*/

func (d *DSP) SetActive(active bool) error {
	res := C.FMOD_DSP_SetActive(d.cptr, getBool(active))
	return errs[res]
}

func (d *DSP) IsActive() (bool, error) {
	var active C.FMOD_BOOL
	res := C.FMOD_DSP_GetActive(d.cptr, &active)
	return setBool(active), errs[res]
}

func (d *DSP) SetBypass(bypass bool) error {
	res := C.FMOD_DSP_SetBypass(d.cptr, getBool(bypass))
	return errs[res]
}

func (d *DSP) Bypass() (bool, error) {
	var bypass C.FMOD_BOOL
	res := C.FMOD_DSP_GetBypass(d.cptr, &bypass)
	return setBool(bypass), errs[res]
}

func (d *DSP) SetWetDryMix(prewet, postwet, dry float64) error {
	res := C.FMOD_DSP_SetWetDryMix(d.cptr, C.float(prewet), C.float(postwet), C.float(dry))
	return errs[res]
}

func (d *DSP) WetDryMix() (float64, float64, float64, error) {
	var prewet, postwet, dry C.float
	res := C.FMOD_DSP_GetWetDryMix(d.cptr, &prewet, &postwet, &dry)
	return float64(prewet), float64(postwet), float64(dry), errs[res]
}

func (d *DSP) SetChannelFormat(channelmask ChannelMask, numchannels int, source_speakermode SpeakerMode) error {
	res := C.FMOD_DSP_SetChannelFormat(d.cptr, C.FMOD_CHANNELMASK(channelmask), C.int(numchannels), C.FMOD_SPEAKERMODE(source_speakermode))
	return errs[res]
}

func (d *DSP) ChannelFormat() (ChannelMask, int, SpeakerMode, error) {
	var channelmask C.FMOD_CHANNELMASK
	var numchannels C.int
	var source_speakermode C.FMOD_SPEAKERMODE
	res := C.FMOD_DSP_GetChannelFormat(d.cptr, &channelmask, &numchannels, &source_speakermode)
	return ChannelMask(channelmask), int(numchannels), SpeakerMode(source_speakermode), errs[res]
}

func (d *DSP) OutputChannelFormat(inmask ChannelMask, inchannels int, inspeakermode SpeakerMode) (ChannelMask, int, SpeakerMode, error) {
	var outmask C.FMOD_CHANNELMASK
	var outchannels C.int
	var outspeakermode C.FMOD_SPEAKERMODE
	res := C.FMOD_DSP_GetOutputChannelFormat(d.cptr, C.FMOD_CHANNELMASK(inmask), C.int(inchannels), C.FMOD_SPEAKERMODE(inspeakermode), &outmask, &outchannels, &outspeakermode)
	return ChannelMask(outmask), int(outchannels), SpeakerMode(outspeakermode), errs[res]
}

func (d *DSP) Reset() error {
	res := C.FMOD_DSP_Reset(d.cptr)
	return errs[res]
}

/*
   DSP parameter control.
*/

func (d *DSP) SetParameterFloat(index int, value float64) error {
	res := C.FMOD_DSP_SetParameterFloat(d.cptr, C.int(index), C.float(value))
	return errs[res]
}

func (d *DSP) SetParameterInt(index, value int) error {
	res := C.FMOD_DSP_SetParameterInt(d.cptr, C.int(index), C.int(value))
	return errs[res]
}

func (d *DSP) SetParameterBool(index int, value bool) error {
	res := C.FMOD_DSP_SetParameterBool(d.cptr, C.int(index), getBool(value))
	return errs[res]
}

func (d *DSP) SetParameterData(index C.int, data *interface{}, length C.uint) error {
	//FMOD_RESULT F_API FMOD_DSP_SetParameterData(FMOD_DSP *dsp, int index, void *data, unsigned int length);
	return ErrNoImpl
}

func (d *DSP) ParameterFloat(index C.int, value *C.float, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterFloat(FMOD_DSP *dsp, int index, float *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

func (d *DSP) ParameterInt(index C.int, value *C.int, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterInt(FMOD_DSP *dsp, int index, int *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

func (d *DSP) ParameterBool(index C.int, value *C.FMOD_BOOL, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterBool(FMOD_DSP *dsp, int index, FMOD_BOOL *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

func (d *DSP) ParameterData(index C.int, data **interface{}, length *C.uint, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterData(FMOD_DSP *dsp, int index, void **data, unsigned int *length, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

func (d *DSP) NumParameters() (int, error) {
	var numparams C.int
	res := C.FMOD_DSP_GetNumParameters(d.cptr, &numparams)
	return int(numparams), errs[res]
}

func (d *DSP) ParameterInfo(index C.int, desc **C.FMOD_DSP_PARAMETER_DESC) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterInfo             (FMOD_DSP *dsp, int index, FMOD_DSP_PARAMETER_DESC **desc);
	return ErrNoImpl
}

func (d *DSP) DataParameterIndex(datatype C.int, index *C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetDataParameterIndex        (FMOD_DSP *dsp, int datatype, int *index);
	return ErrNoImpl
}

func (d *DSP) ShowConfigDialog(hwnd *interface{}, show C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_DSP_ShowConfigDialog             (FMOD_DSP *dsp, void *hwnd, FMOD_BOOL show);
	return ErrNoImpl
}

/*
   DSP attributes.
*/

func (d *DSP) Info(name *C.char, version *C.uint, channels, configwidth, configheight *C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetInfo(FMOD_DSP *dsp, char *name, unsigned int *version, int *channels, int *configwidth, int *configheight);
	return ErrNoImpl
}

func (d *DSP) Type() (DSPType, error) {
	var typ C.FMOD_DSP_TYPE
	res := C.FMOD_DSP_GetType(d.cptr, &typ)
	return DSPType(typ), errs[res]
}

func (d *DSP) Idle() (bool, error) {
	var idle C.FMOD_BOOL
	res := C.FMOD_DSP_GetIdle(d.cptr, &idle)
	return setBool(idle), errs[res]
}

/*
   Userdata set/get.
*/

func (d *DSP) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_DSP_SetUserData                  (FMOD_DSP *dsp, void *userdata);
	return ErrNoImpl
}

func (d *DSP) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_DSP_GetUserData                  (FMOD_DSP *dsp, void **userdata);
	return ErrNoImpl
}

/*
   Metering.
*/

func (d *DSP) SetMeteringEnabled(inputEnabled, outputEnabled bool) error {
	res := C.FMOD_DSP_SetMeteringEnabled(d.cptr, getBool(inputEnabled), getBool(outputEnabled))
	return errs[res]
}

func (d *DSP) MeteringEnabled() (bool, bool, error) {
	var inputEnabled, outputEnabled C.FMOD_BOOL
	res := C.FMOD_DSP_GetMeteringEnabled(d.cptr, &inputEnabled, &outputEnabled)
	return setBool(inputEnabled), setBool(outputEnabled), errs[res]
}

func (d *DSP) MeteringInfo() (DSPMeteringInfo, DSPMeteringInfo, error) {
	var cinputInfo, coutputInfo C.FMOD_DSP_METERING_INFO
	var inputInfo, outputInfo DSPMeteringInfo
	res := C.FMOD_DSP_GetMeteringInfo(d.cptr, &cinputInfo, &coutputInfo)
	inputInfo.fromC(cinputInfo)
	outputInfo.fromC(coutputInfo)
	return inputInfo, outputInfo, errs[res]
}
