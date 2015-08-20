package lowlevel

/*
#include <fmod_common.h>
*/
import "C"

// not gophered yet!
type DSPDesc C.FMOD_DSP_DESCRIPTION

// When creating a DSP unit, declare one of these and provide the relevant callbacks and name for FMOD to use when it creates and uses a DSP unit of this type.
//
// Members marked with [r] mean the variable is modified by FMOD and is for reading purposes only. Do not change this value.
// Members marked with [w] mean the variable can be written to. The user can set the value.
//
// There are 2 different ways to change a parameter in this architecture.
// One is to use "DSP.SetParameterFloat" / "DSP.SetParameterInt" / "DSP.SetParameterBool" / "DSP.SetParameterData".
// This is platform independant and is dynamic, so new unknown plugins can have their parameters enumerated and used.
// The other is to use "DSP.ShowConfigDialog". This is platform specific and requires a GUI, and will display a dialog box to configure the plugin.
/*
type DSPDesc struct {

	// [w] The plugin SDK version this plugin is built for.  set to this to "PLUGIN_SDK_VERSION" defined above.
	Pluginsdkversion uint

	// [w] The identifier of the DSP. This will also be used as the name of DSP and shouldn't change between versions.
	Name string //char  name[32];

	// [w] Plugin writer's version number.
	Version uint

	// [w] Number of input buffers to process.  Use 0 for DSPs that only generate sound and 1 for effects that process incoming sound.
	Numinputbuffers int

	// [w] Number of audio output buffers.  Only one output buffer is currently supported.
	Numoutputbuffers int

	// [w] Create callback.  This is called when DSP unit is created.  Can be null.
	Create C.FMOD_DSP_CREATE_CALLBACK

	// [w] Release callback.  This is called just before the unit is freed so the user can do any cleanup needed for the unit.  Can be null.
	Release C.FMOD_DSP_RELEASE_CALLBACK

	// [w] Reset callback.  This is called by the user to reset any history buffers that may need resetting for a filter, when it is to be used or re-used for the first time to its initial clean state.
	// Use to avoid clicks or artifacts.
	Reset C.FMOD_DSP_RESET_CALLBACK

	// [w] Read callback.  Processing is done here.  Can be null.
	Read C.FMOD_DSP_READ_CALLBACK

	// [w] Process callback.  Can be specified instead of the read callback if any channel format changes occur between input and output.
	// This also replaces shouldiprocess and should return an error if the effect is to be bypassed.  Can be null.
	Process C.FMOD_DSP_PROCESS_CALLBACK

	// [w] Set position callback.  This is called if the unit wants to update its position info but not process data, or
	// reset a cursor position internally if it is reading data from a certain source.  Can be null.
	Setposition C.FMOD_DSP_SETPOSITION_CALLBACK

	// [w] Number of parameters used in this filter.  The user finds this with "DSP.NumParameters"
	Numparameters int

	// [w] Variable number of parameter structures.
	Paramdesc C.FMOD_DSP_PARAMETER_DESC

	// [w] This is called when the user calls "DSP.SetParameterFloat". Can be null.
	Setparameterfloat C.FMOD_DSP_SETPARAM_FLOAT_CALLBACK

	// [w] This is called when the user calls "DSP.SetParameterInt".   Can be null.
	Setparameterint C.FMOD_DSP_SETPARAM_INT_CALLBACK

	// [w] This is called when the user calls "DSP.SetParameterBool".  Can be null.
	Setparameterbool C.FMOD_DSP_SETPARAM_BOOL_CALLBACK

	// [w] This is called when the user calls "DSP.SetParameterData".  Can be null.
	Setparameterdata C.FMOD_DSP_SETPARAM_DATA_CALLBACK

	// [w] This is called when the user calls "DSP.ParameterFloat". Can be null.
	Getparameterfloat C.FMOD_DSP_GETPARAM_FLOAT_CALLBACK

	// [w] This is called when the user calls "DSP.ParameterInt".   Can be null.
	Getparameterint C.FMOD_DSP_GETPARAM_INT_CALLBACK

	// [w] This is called when the user calls "DSP.ParameterBool".  Can be null.
	Getparameterbool C.FMOD_DSP_GETPARAM_BOOL_CALLBACK

	// [w] This is called when the user calls "DSP.ParameterData".  Can be null.
	Getparameterdata C.FMOD_DSP_GETPARAM_DATA_CALLBACK

	// [w] This is called before processing.  You can detect if inputs are idle and return FMOD_OK to process, or any other error code to avoid processing the effect.
	// Use a count down timer to allow effect tails to process before idling!
	Shouldiprocess C.FMOD_DSP_SHOULDIPROCESS_CALLBACK

	// [w] Optional. Specify 0 to ignore. This is user data to be attached to the DSP unit during creation.  Access via "DSP.UserData".
	Userdata *interface{} //void *userdata

	// [w] Register callback.  This is called when DSP unit is loaded/registered.  Useful for 'global'/per system object init for plugin.  Can be null.
	Sys_register C.FMOD_DSP_SYSTEM_REGISTER_CALLBACK

	// [w] Deregister callback.  This is called when DSP unit is unloaded/deregistered.  Useful as 'global'/per system object shutdown for plugin.  Can be null.
	Sys_deregister C.FMOD_DSP_SYSTEM_DEREGISTER_CALLBACK

	// [w] System mix stage callback.  This is called when the mixer starts to execute or is just finishing executing.  Useful for 'global'/per system object once a mix update calls for a plugin.  Can be null.
	Sys_mix C.FMOD_DSP_SYSTEM_MIX_CALLBACK
}
*/
