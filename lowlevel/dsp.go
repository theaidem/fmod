package lowlevel

/*
#include <fmod.h>
*/
import "C"
import "unsafe"

type DSP struct {
	cptr *C.FMOD_DSP
}

/*
   'DSP' API
*/

// Frees a DSP object.
//
// This will free the DSP object.
// NOTE: If DSP is not removed from the Channel, ChannelGroup or System object with "Channel.RemoveDSP" or "ChannelGroup.RemoveDSP",
// after being added with "Channel.AddDSP" or "ChannelGroup.AddDSP", it will not release and will instead return FMOD_ERR_DSP_INUSE.
func (d *DSP) Release() error {
	res := C.FMOD_DSP_Release(d.cptr)
	return errs[res]
}

// Retrieves the parent System object that was used to create this object.
func (d *DSP) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_DSP_GetSystemObject(d.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   Connection / disconnection / input and output enumeration.
*/

// Adds the specified DSP unit as an input of the DSP object.
//
// input: The DSP unit to add as an input of the current unit.
//
// connection: The connection between the 2 units. Optional. Specify 0 or NULL to ignore.
//
// typ: The type of connection between the 2 units. See "DSPConnectionType".
//
// If you want to add a unit as an output of another unit, then add 'this' unit as an input of that unit instead.
// Inputs are automatically mixed together, then the mixed data is sent to the unit's output(s).
// To find the number of inputs or outputs a unit has use "DSP.NumInputs" or "DSP.NumOutputs".
// Note: The connection pointer retrieved here will become invalid if you disconnect the 2 dsp units that use it.
func (d *DSP) AddInput(input DSP, typ DSPConnectionType) (DspConnection, error) {
	var dspConn DspConnection
	res := C.FMOD_DSP_AddInput(d.cptr, input.cptr, &dspConn.cptr, C.FMOD_DSPCONNECTION_TYPE(typ))
	return dspConn, errs[res]
}

// Disconnect the DSP unit from the specified target.
//
// target: The unit that this unit is to be removed from. Specify 0 or NULL to disconnect the unit from all outputs and inputs.
//
// connection: If there is more than one connection between 2 dsp units, this can be used to define which of the connections should be disconnected.
//
// Note that when you disconnect a unit, it is up to you to reconnect the network so that data flow can continue.
// Important note: If you have a handle to the connection pointer that binds these 2 DSP units, then it will become invalid.
// The connection is then sent back to a freelist to be re-used again by a later addInput command.
func (d *DSP) DisconnectFrom(target DSP, connection DspConnection) error {
	res := C.FMOD_DSP_DisconnectFrom(d.cptr, target.cptr, connection.cptr)
	return errs[res]
}

// Helper function to disconnect either all inputs or all outputs of a dsp unit.
//
// inputs: true = disconnect all inputs to this DSP unit. false = leave input connections alone.
//
// outputs: true = disconnect all outputs to this DSP unit. false = leave output connections alone.
//
// This function is optimized to be faster than disconnecting inputs and outputs manually one by one.
// Important note: If you have a handle to DSPConnection pointers that bind any of the inputs or outputs to this DSP unit, then they will become invalid.
// The connections are sent back to a freelist to be re-used again by a later addInput command.
func (d *DSP) DisconnectAll(inputs, outputs bool) error {
	res := C.FMOD_DSP_DisconnectAll(d.cptr, getBool(inputs), getBool(outputs))
	return errs[res]
}

// Retrieves the number of inputs connected to the DSP unit.
//
// Inputs are units that feed data to this unit. When there are multiple inputs, they are mixed together.
//
// Performance warning! Because this function needs to flush the dsp queue before it can determine how many units are available,
// this function may block significantly while the background mixer thread operates.
func (d *DSP) NumInputs() (int, error) {
	var numinputs C.int
	res := C.FMOD_DSP_GetNumInputs(d.cptr, &numinputs)
	return int(numinputs), errs[res]
}

// Retrieves the number of outputs connected to the DSP unit.
//
// Outputs are units that this unit feeds data to. When there are multiple outputs, the data is split and sent to each unit individually.
//
// Performance warning! Because this function needs to flush the dsp queue before it can determine how many units are available,
// this function may block significantly while the background mixer thread operates.
func (d *DSP) NumOutputs() (int, error) {
	var numoutputs C.int
	res := C.FMOD_DSP_GetNumOutputs(d.cptr, &numoutputs)
	return int(numoutputs), errs[res]
}

// Retrieves a pointer to a DSP unit which is acting as an input to this unit.
//
// index: Index of the input unit to retrieve.
//
// An input is a unit which feeds audio data to this unit.
// If there are more than 1 input to this unit, the inputs will be mixed, and the current unit processes the mixed result.
// Find out the number of input units to this unit by calling "DSP.NumInputs".
//
// Performance warning! Because this function needs to flush the dsp queue before it can determine if the specified numerical input is available or not,
// this function may block significantly while the background mixer thread operates.
//
// Note: The connection pointer retrieved here will become invalid if you disconnect the 2 dsp units that use it.
func (d *DSP) Input(index int) (DSP, DspConnection, error) {
	var input DSP
	var inputconnection DspConnection
	res := C.FMOD_DSP_GetInput(d.cptr, C.int(index), &input.cptr, &inputconnection.cptr)
	return input, inputconnection, errs[res]
}

// Retrieves a pointer to a DSP unit which is acting as an output to this unit.
//
// index: Index of the output unit to retrieve.
//
// An output is a unit which this unit will feed data too once it has processed its data.
// Find out the number of output units to this unit by calling "DSP.NumOutputs".
//
// Performance warning! Because this function needs to flush the dsp queue before it can determine if the specified numerical output is available or not,
// this function may block significantly while the background mixer thread operates.
//
// Note: The connection pointer retrieved here will become invalid if you disconnect the 2 dsp units that use it.
func (d *DSP) Output(index int) (DSP, DspConnection, error) {
	var output DSP
	var outputconnection DspConnection
	res := C.FMOD_DSP_GetOutput(d.cptr, C.int(index), &output.cptr, &outputconnection.cptr)
	return output, outputconnection, errs[res]
}

/*
   DSP unit control.
*/

// Enables or disables a unit for being processed.
//
// active: true = unit is activated, false = unit is deactivated.
//
// This does not connect or disconnect a unit in any way, it just disables it so that it is not processed.
// If a unit is disabled, and has inputs, they will also cease to be processed.
// To disable a unit but allow the inputs of the unit to continue being processed, use "DSP.SetBypass" instead.
func (d *DSP) SetActive(active bool) error {
	res := C.FMOD_DSP_SetActive(d.cptr, getBool(active))
	return errs[res]
}

// Retrieves the active state of a DSP unit.
func (d *DSP) IsActive() (bool, error) {
	var active C.FMOD_BOOL
	res := C.FMOD_DSP_GetActive(d.cptr, &active)
	return setBool(active), errs[res]
}

// Enables or disables the read callback of a DSP unit so that it does or doesn't process the data coming into it.
// A DSP unit that is disabled still processes its inputs, it will just be 'dry'.
//
// bypass: Boolean to cause the read callback of the DSP unit to be bypassed or not. Default = false.
//
// If a unit is bypassed, it will still process its inputs.
// To disable the unit and all of its inputs, use "DSP.SetActive" instead.
func (d *DSP) SetBypass(bypass bool) error {
	res := C.FMOD_DSP_SetBypass(d.cptr, getBool(bypass))
	return errs[res]
}

// Retrieves the bypass state of the DSP unit.
// If a unit is bypassed, it will still process its inputs, unlike "DSP.SetActive" (when set to false) which causes inputs to stop processing as well.
func (d *DSP) Bypass() (bool, error) {
	var bypass C.FMOD_BOOL
	res := C.FMOD_DSP_GetBypass(d.cptr, &bypass)
	return setBool(bypass), errs[res]
}

// Allows the user to scale the affect of a DSP effect, through control of the 'wet' mix, which is the post-processed signal and the 'dry' which is the pre-processed signal.
//
// prewet: Floating point value from 0 to 1, describing a linear scale of the 'wet' (pre-processed signal) mix of the effect. Default = 1.0. Scale can be lower than 0 (negating) and higher than 1 (amplifying).
//
// postwet: Floating point value from 0 to 1, describing a linear scale of the 'wet' (post-processed signal) mix of the effect. Default = 1.0. Scale can be lower than 0 (negating) and higher than 1 (amplifying).
//
// dry: Floating point value from 0 to 1, describing a linear scale of the 'dry' (pre-processed signal) mix of the effect. Default = 0.0. Scale can be lower than 0 and higher than 1 (amplifying).
//
// The dry signal path is silent by default, because dsp effects transform the input and pass the newly processed result to the output. It does not add to the input.
func (d *DSP) SetWetDryMix(prewet, postwet, dry float64) error {
	res := C.FMOD_DSP_SetWetDryMix(d.cptr, C.float(prewet), C.float(postwet), C.float(dry))
	return errs[res]
}

// Retrieves the wet/dry scale of a DSP effect, through the 'wet' mix, which is the post-processed signal and the 'dry' mix which is the pre-processed signal.
// The dry signal path is silent by default, because dsp effects transform the input and pass the newly processed result to the output. It does not add to the input.
func (d *DSP) WetDryMix() (float64, float64, float64, error) {
	var prewet, postwet, dry C.float
	res := C.FMOD_DSP_GetWetDryMix(d.cptr, &prewet, &postwet, &dry)
	return float64(prewet), float64(postwet), float64(dry), errs[res]
}

// Sets the signal format of a dsp unit so that the signal is processed on the speakers specified.
// Also defines the number of channels in the unit that a read callback will process, and the output signal of the unit.
//
// channelmask: A series of bits specified by "ChannelMask" to determine which speakers are represented by the channels in the signal.
//
// numchannels: The number of channels to be processed on this unit and sent to the outputs connected to it. Maximum of FMOD_MAX_CHANNEL_WIDTH.
//
// source_speakermode: The source speaker mode where the signal came from.
//
// Setting the number of channels on a unit will force a down or up mix to that channel count before processing the DSP read callback.
// This channelcount is then sent to the outputs of the unit.
// source_speakermode is informational, when channelmask describes what bits are active, and numchannels describes how many channels are in a buffer,
// source_speakermode describes where the channels originated from.
// For example if numchannels = 2 then this could describe for the DSP if the original signal started from a stereo signal or a 5.1 signal.
// It could also describe the signal as all monaural, for example if numchannels was 16 and the speakermode was FMOD_SPEAKERMODE_MONO.
func (d *DSP) SetChannelFormat(channelmask ChannelMask, numchannels int, source_speakermode SpeakerMode) error {
	res := C.FMOD_DSP_SetChannelFormat(d.cptr, C.FMOD_CHANNELMASK(channelmask), C.int(numchannels), C.FMOD_SPEAKERMODE(source_speakermode))
	return errs[res]
}

// Gets the input signal format for a dsp units read/process callback, to determine which speakers the signal will be processed on and how many channels will be processed.
//
// source_speakermode is informational, when channelmask describes what bits are active, and numchannels describes how many channels are in a buffer,
// source_speakermode describes where the channels originated from.
// For example if numchannels = 2 then this could describe for the DSP if the original signal started from a stereo signal or a 5.1 signal.
// In the 5.1 signal the channels described might only represent 2 surround speakers for example.
func (d *DSP) ChannelFormat() (ChannelMask, int, SpeakerMode, error) {
	var channelmask C.FMOD_CHANNELMASK
	var numchannels C.int
	var source_speakermode C.FMOD_SPEAKERMODE
	res := C.FMOD_DSP_GetChannelFormat(d.cptr, &channelmask, &numchannels, &source_speakermode)
	return ChannelMask(channelmask), int(numchannels), SpeakerMode(source_speakermode), errs[res]
}

// Call the DSP process function to retrieve the output signal format for a DSP based on input values.
//
// inmask: Channel bitmask representing the speakers enabled for the incoming signal.
// For example a 5.1 signal could have inchannels 2 that represent CHANNELMASK_SURROUND_LEFT and CHANNELMASK_SURROUND_RIGHT.
//
// inchannels: Number of channels for the incoming signal.
//
// inspeakermode: Speaker mode for the incoming signal.
//
// A DSP unit may be an up mixer or down mixer for example. In this case if you specified 6 in for a downmixer, it may provide you with 2 out for example.
// Generally the input values will be reproduced for the output values, but some DSP units will want to alter the output format.
func (d *DSP) OutputChannelFormat(inmask ChannelMask, inchannels int, inspeakermode SpeakerMode) (ChannelMask, int, SpeakerMode, error) {
	var outmask C.FMOD_CHANNELMASK
	var outchannels C.int
	var outspeakermode C.FMOD_SPEAKERMODE
	res := C.FMOD_DSP_GetOutputChannelFormat(d.cptr, C.FMOD_CHANNELMASK(inmask), C.int(inchannels), C.FMOD_SPEAKERMODE(inspeakermode), &outmask, &outchannels, &outspeakermode)
	return ChannelMask(outmask), int(outchannels), SpeakerMode(outspeakermode), errs[res]
}

// Calls the DSP unit's reset function, which will clear internal buffers and reset the unit back to an initial state.
//
// Calling this function is useful if the DSP unit relies on a history to process itself (ie an echo filter).
//
// If you disconnected the unit and reconnected it to a different part of the network with a different sound, you would want to call this to reset the units state (ie clear and reset the echo filter)
// so that you dont get left over artifacts from the place it used to be connected.
func (d *DSP) Reset() error {
	res := C.FMOD_DSP_Reset(d.cptr)
	return errs[res]
}

/*
   DSP parameter control.
*/

// Sets a DSP unit's floating point parameter by index. To find out the parameter names and range, see the see also field.
//
// index: Parameter index for this unit. Find the number of parameters with "DSP.NumParameters".
//
// value: Floating point parameter value to be passed to the DSP unit.
//
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) SetParameterFloat(index int, value float64) error {
	res := C.FMOD_DSP_SetParameterFloat(d.cptr, C.int(index), C.float(value))
	return errs[res]
}

// Sets a DSP unit's integer parameter by index. To find out the parameter names and range, see the see also field.
//
// index: Parameter index for this unit. Find the number of parameters with "DSP.NumParameters".
//
// value: Integer parameter value to be passed to the DSP unit.
//
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) SetParameterInt(index, value int) error {
	res := C.FMOD_DSP_SetParameterInt(d.cptr, C.int(index), C.int(value))
	return errs[res]
}

// Sets a DSP unit's boolean parameter by index. To find out the parameter names and range, see the see also field.
//
// index: Parameter index for this unit. Find the number of parameters with "DSP.NumParameters".
//
// value: Boolean parameter value to be passed to the DSP unit. Should be TRUE or FALSE.
//
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) SetParameterBool(index int, value bool) error {
	res := C.FMOD_DSP_SetParameterBool(d.cptr, C.int(index), getBool(value))
	return errs[res]
}

// NOTE: Not implement yet
// Sets a DSP unit's binary data parameter by index. To find out the parameter names and range, see the see also field.
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) SetParameterData(index C.int, data *interface{}, length C.uint) error {
	//FMOD_RESULT F_API FMOD_DSP_SetParameterData(FMOD_DSP *dsp, int index, void *data, unsigned int length);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a DSP unit's floating point parameter by index. To find out the parameter names and range, see the see also field.
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) ParameterFloat(index C.int, value *C.float, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterFloat(FMOD_DSP *dsp, int index, float *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a DSP unit's integer parameter by index. To find out the parameter names and range, see the see also field.
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) ParameterInt(index C.int, value *C.int, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterInt(FMOD_DSP *dsp, int index, int *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a DSP unit's boolean parameter by index. To find out the parameter names and range, see the see also field.
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) ParameterBool(index C.int, value *C.FMOD_BOOL, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterBool(FMOD_DSP *dsp, int index, FMOD_BOOL *value, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a DSP unit's data block parameter by index. To find out the parameter names and range, see the see also field.
// The parameter properties (such as min/max values) can be retrieved with "DSP.ParameterInfo".
func (d *DSP) ParameterData(index C.int, data **interface{}, length *C.uint, valuestr *C.char, valuestrlen C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterData(FMOD_DSP *dsp, int index, void **data, unsigned int *length, char *valuestr, int valuestrlen);
	return ErrNoImpl
}

// Retrieves the number of parameters a DSP unit has to control its behaviour.
// Use this to enumerate all parameters of a DSP unit with "DSP.ParameterInfo".
func (d *DSP) NumParameters() (int, error) {
	var numparams C.int
	res := C.FMOD_DSP_GetNumParameters(d.cptr, &numparams)
	return int(numparams), errs[res]
}

// NOTE: Not implement yet
// Retrieve information about a specified parameter within the DSP unit.
// Use "DSP.NumParameters" to find out the number of parameters for this DSP unit.
func (d *DSP) ParameterInfo(index C.int, desc **C.FMOD_DSP_PARAMETER_DESC) error {
	//FMOD_RESULT F_API FMOD_DSP_GetParameterInfo             (FMOD_DSP *dsp, int index, FMOD_DSP_PARAMETER_DESC **desc);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieve the index of the first data parameter of a particular data type.
// The return code can therefore be used to check whether the DSP supports specific functionality through data parameters of certain types without the need to pass in 'index'.
func (d *DSP) DataParameterIndex(datatype C.int, index *C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetDataParameterIndex        (FMOD_DSP *dsp, int datatype, int *index);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Display or hide a DSP unit configuration dialog box inside the target window.
//
// Dialog boxes are used by DSP plugins that prefer to use a graphical user interface to modify their parameters rather than using the other method of enumerating the parameters and
// using "DSP.SetParameterFloat" / "DSP.SetParameterInt" / "DSP.SetParameterBool" / "DSP.SetParameterData".
//
// These are usually VST plugins. FMOD Studio plugins do not have configuration dialog boxes.
// To find out what size window to create to store the configuration screen, use "DSP.Info" where you can get the width and height.
func (d *DSP) ShowConfigDialog(hwnd *interface{}, show C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_DSP_ShowConfigDialog             (FMOD_DSP *dsp, void *hwnd, FMOD_BOOL show);
	return ErrNoImpl
}

/*
   DSP attributes.
*/

// NOTE: Not implement yet
// TODO: add more docs
// Retrieves information about the current DSP unit, including name, version, default channels and width and height of configuration dialog box if it exists.
func (d *DSP) Info(name *C.char, version *C.uint, channels, configwidth, configheight *C.int) error {
	//FMOD_RESULT F_API FMOD_DSP_GetInfo(FMOD_DSP *dsp, char *name, unsigned int *version, int *channels, int *configwidth, int *configheight);
	return ErrNoImpl
}

// Retrieves the pre-defined type of a FMOD registered DSP unit.
// This is only valid for built in FMOD effects. Any user plugins will simply return "DSP_TYPE_UNKNOWN".
func (d *DSP) Type() (DSPType, error) {
	var typ C.FMOD_DSP_TYPE
	res := C.FMOD_DSP_GetType(d.cptr, &typ)
	return DSPType(typ), errs[res]
}

// Retrieves the idle state of a DSP. A DSP is idle when no signal is coming into it.
// This can be a useful method of determining if a DSP sub branch is finished processing, so it can be disconnected for example.
//
// The idle state takes into account things like tails of echo filters, even if a wavetable or dsp has finished generating sound.
// When all nodes in a graph have finished processing, only then will it set the top level DSP state to idle.
func (d *DSP) Idle() (bool, error) {
	var idle C.FMOD_BOOL
	res := C.FMOD_DSP_GetIdle(d.cptr, &idle)
	return setBool(idle), errs[res]
}

/*
   Userdata set/get.
*/

// Sets a user value that the DSP object will store internally. Can be retrieved with "DSP.UserData".
//
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
//
// It can be useful if an FMOD callback passes an object of this type as a parameter, and the user does not know which object it is (if many of these types of objects exist).
// Using "DSP.UserData" would help in the identification of the object.
func (d *DSP) SetUserData(userdata interface{}) error {
	data := *(*[]*C.char)(unsafe.Pointer(&userdata))
	res := C.FMOD_DSP_SetUserData(d.cptr, unsafe.Pointer(&data))
	return errs[res]
}

// Retrieves the user value that that was set by calling the "DSP.SetUserData" function.
func (d *DSP) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_DSP_GetUserData(d.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}

/*
   Metering.
*/

// Enable metering for a DSP unit so that "DSP.MeteringInfo" will return metering information, and so that FMOD Studio profiler tool can visualize the levels.
//
// inputEnabled: Enable metering for the input signal (pre-processing). Specify true to turn on input level metering, false to turn it off.
//
// outputEnabled: Enable metering for the output signal (post-processing). Specify true to turn on output level metering, false to turn it off.
//
// "INIT_PROFILE_METER_ALL" with "System.Init" will automatically turn on metering for all DSP units inside the FMOD mixer graph.
func (d *DSP) SetMeteringEnabled(inputEnabled, outputEnabled bool) error {
	res := C.FMOD_DSP_SetMeteringEnabled(d.cptr, getBool(inputEnabled), getBool(outputEnabled))
	return errs[res]
}

// Retrieve the information about metering for a particular DSP to see if it is enabled or not.
//
// "INIT_PROFILE_METER_ALL" with "System.Init" will automatically turn on metering for all DSP units inside the FMOD mixer graph.
func (d *DSP) MeteringEnabled() (bool, bool, error) {
	var inputEnabled, outputEnabled C.FMOD_BOOL
	res := C.FMOD_DSP_GetMeteringEnabled(d.cptr, &inputEnabled, &outputEnabled)
	return setBool(inputEnabled), setBool(outputEnabled), errs[res]
}

// Retrieve the metering information for a particular DSP.
//
// "INIT_PROFILE_METER_ALL" with "System.Init" will automatically turn on metering for all DSP units inside the FMOD mixer graph.
func (d *DSP) MeteringInfo() (DSPMeteringInfo, DSPMeteringInfo, error) {
	var cinputInfo, coutputInfo C.FMOD_DSP_METERING_INFO
	var inputInfo, outputInfo DSPMeteringInfo
	res := C.FMOD_DSP_GetMeteringInfo(d.cptr, &cinputInfo, &coutputInfo)
	inputInfo.fromC(cinputInfo)
	outputInfo.fromC(coutputInfo)
	return inputInfo, outputInfo, errs[res]
}
