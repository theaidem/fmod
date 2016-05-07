package lowlevel

/*
#include <stdlib.h>
#include <fmod.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// The main object for the FMOD Low Level System.
// When using FMOD Studio, this system object will be automatically instantiated as part of `StudioSystem.Initialize()`.
type System struct {
	cptr *C.FMOD_SYSTEM
}

/*
   'System' API
*/

// FMOD System creation function.
// This must be called to create an FMOD System object before you can do anything else.
// Use this function to create 1, or multiple instances of FMOD System objects.
func SystemCreate() (*System, error) {
	var s System
	res := C.FMOD_System_Create(&s.cptr)
	runtime.SetFinalizer(&s, (*System).Release)
	return &s, errs[res]
}

// Closes and frees a system object and its resources.
// This function also calls "System.Close()", so calling close before this function is not necessary.
func (s *System) Release() error {
	res := C.FMOD_System_Release(s.cptr)
	return errs[res]
}

/*
   Setup functions.
*/

// This function selects the output mode for the platform.
// This is for selecting different OS specific APIs which might have different features.
// See "OutputType" for different output types you can select.
func (s *System) SetOutput(output OutputType) error {
	res := C.FMOD_System_SetOutput(s.cptr, C.FMOD_OUTPUTTYPE(output))
	return errs[res]
}

// Retrieves the current output system FMOD is using to address the hardware.
func (s *System) Output() (OutputType, error) {
	var output C.FMOD_OUTPUTTYPE
	res := C.FMOD_System_GetOutput(s.cptr, &output)
	return OutputType(output), errs[res]
}

// Retrieves the number of soundcard devices on the machine, specific to the output mode set with "System.SetOutput".
// If "System.SetOutput" is not called it will return the number of drivers available for the default output type. Use this for enumerating sound devices.
// Use "System.DriverInfo()" to get the device's name.
func (s *System) NumDrivers() (int, error) {
	var numdrivers C.int
	res := C.FMOD_System_GetNumDrivers(s.cptr, &numdrivers)
	return int(numdrivers), errs[res]
}

// Retrieves identification information about a sound device specified by its index, and specific to the output mode set with "System.SetOutput".
func (s *System) DriverInfo(id int, name string) (Guid, int, SpeakerMode, int, error) {
	var guid C.FMOD_GUID
	var systemrate C.int
	var speakermode C.FMOD_SPEAKERMODE
	var speakermodechannels C.int
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	namelen := len(name)
	res := C.FMOD_System_GetDriverInfo(s.cptr, C.int(id), cname, C.int(namelen), &guid, &systemrate, &speakermode, &speakermodechannels)
	return Guid(guid), int(systemrate), SpeakerMode(speakermode), int(speakermodechannels), errs[res]
}

// Selects a soundcard driver.
// This function is used when an output mode has enumerated more than one output device, and you need to select between them.
// If this function is called after FMOD is already initialized with "System.Init", the current driver will be shutdown and the newly selected driver will be initialized / started.
// The driver that you wish to change to must support the current output format, sample rate, and number of channels.
// If it does not, FMOD_ERR_OUTPUT_INIT is returned and driver state is cleared.
// You should now call "System.SetDriver" with your original driver index to restore driver state (providing that driver is still available / connected) or make another selection.
// driver: Driver number to select. 0 = primary or main sound device as selected by the operating system settings.
// Use "System.NumDrivers" and "System.DriverInfo" to determine available devices.
func (s *System) SetDriver(driver int) error {
	res := C.FMOD_System_SetDriver(s.cptr, C.int(driver))
	return errs[res]
}

// Returns the currently selected driver number.
// Drivers are enumerated when selecting a driver with "System.SetDriver" or other driver related functions such as "System.NumDrivers" or "System.DriverInfo".
// ID. 0 = primary or main sound device as selected by the operating system settings.
func (s *System) Driver() (int, error) {
	var driver C.int
	res := C.FMOD_System_GetDriver(s.cptr, &driver)
	return int(driver), errs[res]
}

// Sets the maximum number of software mixed channels possible.
// numsoftwarechannels: The maximum number of mixable voices to be allocated by FMOD, default = 64.
// This function cannot be called after FMOD is already activated, it must be called before "System.Init", or after "System.Close".
func (s *System) SetSoftwareChannels(numsoftwarechannels int) error {
	res := C.FMOD_System_SetSoftwareChannels(s.cptr, C.int(numsoftwarechannels))
	return errs[res]
}

// Retrieves the maximum number of software mixed channels possible.
func (s *System) SoftwareChannels() (int, error) {
	var numsoftwarechannels C.int
	res := C.FMOD_System_GetSoftwareChannels(s.cptr, &numsoftwarechannels)
	return int(numsoftwarechannels), errs[res]
}

// Sets the output format for the software mixer.
// If loading Studio banks, this must be called with speakermode corresponding to the project's output format if there is a possibility of the output audio device not matching the project's format.
// Any differences between the project format and the system's speakermode will cause the mix to sound wrong.
// If not loading Studio banks, do not call this unless you explicity want to change a setting from the default.
// FMOD will default to the speaker mode and sample rate that the OS / output prefers.
//
// samplerate: Sample rate in Hz, that the software mixer will run at. Specify values between 8000 and 192000.
//
// speakermode: Speaker setup for the software mixer.
//
// numrawspeakers: Number of output channels / speakers to initialize the sound card to in FMOD_SPEAKERMODE_RAW mode. Optional. Specify 0 to ignore. Maximum of FMOD_MAX_CHANNEL_WIDTH.
//
// This function cannot be called after FMOD is already activated, it must be called before "System.Init", or after "System.Close".
func (s *System) SetSoftwareFormat(samplerate int, speakermode SpeakerMode, numrawspeakers int) error {
	res := C.FMOD_System_SetSoftwareFormat(s.cptr, C.int(samplerate), C.FMOD_SPEAKERMODE(speakermode), C.int(numrawspeakers))
	return errs[res]
}

// Retrieves the output format for the software mixer.
// Note that the settings returned here may differ from the settings provided by the user with "System.SetSoftwareFormat".
// This is because the driver may require certain settings to initialize.
func (s *System) SoftwareFormat() (int, SpeakerMode, int, error) {
	var samplerate C.int
	var speakermode C.FMOD_SPEAKERMODE
	var numrawspeakers C.int
	res := C.FMOD_System_GetSoftwareFormat(s.cptr, &samplerate, &speakermode, &numrawspeakers)
	return int(samplerate), SpeakerMode(speakermode), int(numrawspeakers), errs[res]
}

// Sets the FMOD internal mixing buffer size. This function is used if you need to control mixer latency or granularity.
// Smaller buffersizes lead to smaller latency, but can lead to stuttering/skipping/unstable sound on slower machines or soundcards with bad drivers.
//
// bufferlength: The mixer engine block size in samples. Use this to adjust mixer update granularity. Default = 1024.
// (milliseconds = 1024 at 48khz = 1024 / 48000 * 1000 = 21.33ms). This means the mixer updates every 21.33ms.
//
// numbuffers: The mixer engine number of buffers used. Use this to adjust mixer latency. Default = 4.
// To get the total buffersize multiply the bufferlength by the numbuffers value. By default this would be 4*1024.
//
// The FMOD software mixer mixes to a ringbuffer.
// The size of this ringbuffer is determined here.
// It mixes a block of sound data every 'bufferlength' number of samples, and there are 'numbuffers' number of these blocks that make up the entire ringbuffer.
// Adjusting these values can lead to extremely low latency performance (smaller values), or greater stability in sound output (larger values).
//
// Warning! The 'buffersize' is generally best left alone. Making the granularity smaller will just increase CPU usage (cache misses and DSP network overhead).
// Making it larger affects how often you hear commands update such as volume/pitch/pan changes.
// Anything above 20ms will be noticable and sound parameter changes will be obvious instead of smooth.
//
// FMOD chooses the most optimal size by default for best stability, depending on the output type, and if the drivers are emulated or not (for example DirectSound is emulated using waveOut on NT).
// It is not recommended changing this value unless you really need to. You may get worse performance than the default settings chosen by FMOD.
// To convert from milliseconds to 'samples', simply multiply the value in milliseconds by the sample rate of the output (ie 48000 if that is what it is set to), then divide by 1000.
//
//// This function cannot be called after FMOD is already activated, it must be called before "System.Init", or after "System.Close".
func (s *System) SetDSPBufferSize(bufferlength uint32, numbuffers int) error {
	res := C.FMOD_System_SetDSPBufferSize(s.cptr, C.uint(bufferlength), C.int(numbuffers))
	return errs[res]
}

// Retrieves the buffer size settings for the FMOD software mixing engine.
//See documentation on "System.SetDSPBufferSize" for more information about these values.
func (s *System) DSPBufferSize() (uint32, int, error) {
	var bufferlength C.uint
	var numbuffers C.int
	res := C.FMOD_System_GetDSPBufferSize(s.cptr, &bufferlength, &numbuffers)
	return uint32(bufferlength), int(numbuffers), errs[res]
}

// TODO: add more docs
// NOTE: Not implement yet
// Specify user callbacks for FMOD's internal file manipulation functions. This function is useful for replacing FMOD's file system with a game system's own file reading API.
func (s *System) SetFileSystem(useropen C.FMOD_FILE_OPEN_CALLBACK, userclose C.FMOD_FILE_CLOSE_CALLBACK, userread C.FMOD_FILE_READ_CALLBACK, userseek C.FMOD_FILE_SEEK_CALLBACK, userasyncread C.FMOD_FILE_ASYNCREAD_CALLBACK, userasynccancel C.FMOD_FILE_ASYNCCANCEL_CALLBACK, blockalign C.int) error {
	//FMOD_RESULT F_API FMOD_System_SetFileSystem             (FMOD_SYSTEM *system, FMOD_FILE_OPEN_CALLBACK useropen, FMOD_FILE_CLOSE_CALLBACK userclose, FMOD_FILE_READ_CALLBACK userread, FMOD_FILE_SEEK_CALLBACK userseek, FMOD_FILE_ASYNCREAD_CALLBACK userasyncread, FMOD_FILE_ASYNCCANCEL_CALLBACK userasynccancel, int blockalign);
	return ErrNoImpl
}

// TODO: add more docs
// NOTE: Not implement yet
// Function to allow a user to 'piggyback' on FMOD's file reading routines.
// This allows users to capture data as FMOD reads it, which may be useful for ripping the raw data that FMOD reads for hard to support sources (for example internet streams).
//
// NOTE! Do not use this to 'override' FMOD's file system! That is what setFileSystem is for.
// This function is purely for 'snooping' and letting FMOD do its own file access, but if you want to capture what FMOD is reading you can do it with this function.
func (s *System) AttachFileSystem(useropen C.FMOD_FILE_OPEN_CALLBACK, userclose C.FMOD_FILE_CLOSE_CALLBACK, userread C.FMOD_FILE_READ_CALLBACK, userseek C.FMOD_FILE_SEEK_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_System_AttachFileSystem          (FMOD_SYSTEM *system, FMOD_FILE_OPEN_CALLBACK useropen, FMOD_FILE_CLOSE_CALLBACK userclose, FMOD_FILE_READ_CALLBACK userread, FMOD_FILE_SEEK_CALLBACK userseek);
	return ErrNoImpl
}

// Sets advanced features like configuring memory and cpu usage for FMOD_CREATECOMPRESSEDSAMPLE usage.
func (s *System) SetAdvancedSettings(settings *AdvancedSettings) error {
	var csettings = settings.toC()
	res := C.FMOD_System_SetAdvancedSettings(s.cptr, &csettings)
	return errs[res]
}

// Retrieves the advanced settings value set for the system object.
func (s *System) AdvancedSettings() (*AdvancedSettings, error) {
	var settings C.FMOD_ADVANCEDSETTINGS
	settings.cbSize = C.int(unsafe.Sizeof(settings))
	res := C.FMOD_System_GetAdvancedSettings(s.cptr, &settings)
	/*ॐ*/
	as := new(AdvancedSettings)
	as.fromC(settings)
	return as, errs[res]
}

// TODO: add more docs
// NOTE: Not implement yet
// Sets a system callback to catch various fatal or informational events.
func (s *System) SetCallback(callback C.FMOD_SYSTEM_CALLBACK, callbackmask C.FMOD_SYSTEM_CALLBACK_TYPE) error {
	//FMOD_RESULT F_API FMOD_System_SetCallback               (FMOD_SYSTEM *system, FMOD_SYSTEM_CALLBACK callback, FMOD_SYSTEM_CALLBACK_TYPE callbackmask);
	return ErrNoImpl
}

/*
   Plug-in support.
*/

// NOTE: Not implement yet
// Specify a base search path for plugins so they can be placed somewhere else than the directory of the main executable.
func (s *System) SetPluginPath(path *C.char) error {
	//FMOD_RESULT F_API FMOD_System_SetPluginPath             (FMOD_SYSTEM *system, const char *path);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Loads an FMOD plugin. This could be a DSP, file format or output plugin.
func (s *System) LoadPlugin(filename *C.char, handle *C.uint, priority C.uint) error {
	//FMOD_RESULT F_API FMOD_System_LoadPlugin                (FMOD_SYSTEM *system, const char *filename, unsigned int *handle, unsigned int priority);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Unloads a plugin from memory.
func (s *System) UnloadPlugin(handle C.uint) error {
	//FMOD_RESULT F_API FMOD_System_UnloadPlugin              (FMOD_SYSTEM *system, unsigned int handle);
	return ErrNoImpl
}

// Retrieves the number of available plugins loaded into FMOD at the current time.
// plugintype: Plugin type such as PLUGINTYPE_OUTPUT, PLUGINTYPE_CODEC or PLUGINTYPE_DSP.
func (s *System) NumPlugins(plugintype PluginType) (int, error) {
	var numplugins C.int
	res := C.FMOD_System_GetNumPlugins(s.cptr, C.FMOD_PLUGINTYPE(plugintype), &numplugins)
	return int(numplugins), errs[res]
}

// NOTE: Not implement yet
//Retrieves the handle of a plugin based on its type and relative index. Use "System.NumPlugins" to enumerate plugins.
func (s *System) PluginHandle(plugintype C.FMOD_PLUGINTYPE, index C.int, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetPluginHandle           (FMOD_SYSTEM *system, FMOD_PLUGINTYPE plugintype, int index, unsigned int *handle);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves information to display for the selected plugin.
func (s *System) PluginInfo(handle C.uint, plugintype *C.FMOD_PLUGINTYPE, name *C.char, version *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetPluginInfo             (FMOD_SYSTEM *system, unsigned int handle, FMOD_PLUGINTYPE *plugintype, char *name, int namelen, unsigned int *version);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Selects an output type based on the enumerated list of outputs including FMOD and 3rd party output plugins.
func (s *System) SetOutputByPlugin(handle C.uint) error {
	//FMOD_RESULT F_API FMOD_System_SetOutputByPlugin         (FMOD_SYSTEM *system, unsigned int handle);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Returns the currently selected output as an id in the list of output plugins.
// This function can be called after FMOD is already activated. You can use it to change the output mode at runtime.
// If SYSTEM_CALLBACK_DEVICELISTCHANGED is specified use the setOutput call to change to "OUTPUTTYPE_NOSOUND" if no more sound card drivers exist.
func (s *System) OutputByPlugin(handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetOutputByPlugin         (FMOD_SYSTEM *system, unsigned int *handle);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Creates a DSP unit object which is either built in or loaded as a plugin, to be inserted into a DSP network, for the purposes of sound filtering or sound generation.
// This function creates a DSP unit that can be enumerated by using "System.NumPlugins" and "System.PluginInfo".
//
// A DSP unit can generate or filter incoming data.
// To be active, a unit must be inserted into the FMOD DSP network to be heard.
// Use functions such as "ChannelGroup.AddDSP", "Channel.AddDSP" or "DSP.AddInput" to do this.
func (s *System) CreateDSPByPlugin(handle C.uint, dsp **C.FMOD_DSP) error {
	//FMOD_RESULT F_API FMOD_System_CreateDSPByPlugin         (FMOD_SYSTEM *system, unsigned int handle, FMOD_DSP **dsp);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieve the description structure for a pre-existing DSP plugin.
func (s *System) DSPInfoByPlugin(handle C.uint, description **C.FMOD_DSP_DESCRIPTION) error {
	//FMOD_RESULT F_API FMOD_System_GetDSPInfoByPlugin        (FMOD_SYSTEM *system, unsigned int handle, const FMOD_DSP_DESCRIPTION **description);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Creates a file format codec to be used by FMOD for opening custom file types.
func (s *System) RegisterCodec(description *C.FMOD_CODEC_DESCRIPTION, handle *C.uint, priority C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterCodec             (FMOD_SYSTEM *system, FMOD_CODEC_DESCRIPTION *description, unsigned int *handle, unsigned int priority);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Register a user-defined DSP effect for use with the System.
// This function allows you to register statically-linked DSP effects.
// Once registered, you can create instances of the DSP effect by using System::createDSPByPlugin.
func (s *System) RegisterDSP(description *C.FMOD_DSP_DESCRIPTION, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterDSP               (FMOD_SYSTEM *system, const FMOD_DSP_DESCRIPTION *description, unsigned int *handle);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Register a user-defined output mode for use with the System.
// This function allows you to register statically-linked output modes.
// Once registered, you can use the output mode with "System.SetOutputByPlugin".
func (s *System) RegisterOutput(description *C.FMOD_OUTPUT_DESCRIPTION, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterOutput            (FMOD_SYSTEM *system, const FMOD_OUTPUT_DESCRIPTION *description, unsigned int *handle);
	return ErrNoImpl
}

/*
   Init/Close.
*/

// Initializes the system object, and the sound device. This has to be called at the start of the user's program.
// You must create a system object with "SystemCreate".
//
// maxchannels: The maximum number of channels to be used in FMOD.
// They are also called 'virtual channels' as you can play as many of these as you want, even if you only have a small number of software voices. See remarks for more.
//
// flags: See "InitFlags". This can be a selection of flags bitwise OR'ed together to change the behaviour of FMOD at initialization time.
//
// extradriverdata: Driver specific data that can be passed to the output plugin.
// For example the filename for the wav writer plugin. See FMOD_OUTPUTTYPE for what each output mode might take here. Optional. Specify 0 or NULL to ignore.
//
// Virtual channels.
// These types of voices are the ones you work with using the Channel API.
// The advantage of virtual channels are, unlike older versions of FMOD, you can now play as many sounds as you like without fear of ever running out of voices, or playsound failing.
// You can also avoid 'channel stealing' if you specify enough virtual voices.
//
// As an example, you can play 1000 sounds at once, even on a 32 channel soundcard.
//
// FMOD will only play the most important/closest/loudest (determined by volume/distance/geometry and priority settings) voices, and the other 968 voices will be virtualized without expense to the CPU.
// The voice's cursor positions are updated.
//
// When the priority of sounds change or emulated sounds get louder than audible ones, they will swap the actual voice resource over and play the voice from its correct position in time as it should be heard.
//
// What this means is you can play all 1000 sounds, if they are scattered around the game world, and as you move around the world you will hear the closest or most important 32,
// and they will automatically swap in and out as you move.
//
// Currently the maximum channel limit is 4093.
func (s *System) Init(maxchannels int, flags InitFlags, extradriverdata interface{}) error {
	res := C.FMOD_System_Init(s.cptr, C.int(maxchannels), C.FMOD_INITFLAGS(flags), unsafe.Pointer(uintptr(extradriverdata.(int))))
	return errs[res]
}

// Closes the system object without freeing the object's memory, so the system handle will still be valid.
// Closing the output renders objects created with this system object invalid.
// Make sure any sounds, channelgroups, geometry and dsp objects are released before closing the system object.
func (s *System) Close() error {
	res := C.FMOD_System_Close(s.cptr)
	return errs[res]
}

/*
   General post-init system functions.
*/

// TODO: add more docs
// Updates the FMOD system. This should be called once per 'game' tick, or once per frame in your application.
func (s *System) Update() error {
	res := C.FMOD_System_Update(s.cptr)
	return errs[res]
}

// TODO: add more docs
// This function allows the user to specify the position of their actual physical speaker to account for non standard setups.
// It also allows the user to disable speakers from 3D consideration in a game.
// The funtion is for describing the 'real world' speaker placement to provide a more natural panning solution for 3d sound.
// Graphical configuration screens in an application could draw icons for speaker placement that the user could position at their will.
func (s *System) SetSpeakerPosition(speaker Speaker, x, y float32, active bool) error {
	res := C.FMOD_System_SetSpeakerPosition(s.cptr, C.FMOD_SPEAKER(speaker), C.float(x), C.float(y), getBool(active))
	return errs[res]
}

// Retrieves the current speaker position information for the selected speaker.
func (s *System) SpeakerPosition(speaker Speaker) (float32, float32, bool, error) {
	var x, y C.float
	var active C.FMOD_BOOL
	res := C.FMOD_System_GetSpeakerPosition(s.cptr, C.FMOD_SPEAKER(speaker), &x, &y, &active)
	return float32(x), float32(y), setBool(active), errs[res]
}

// Sets the internal buffersize for streams opened after this call.
// Larger values will consume more memory, whereas smaller values may cause buffer under-run/starvation/stuttering caused by large delays in disk access (ie netstream),
// or cpu usage in slow machines, or by trying to play too many streams at once.
//
// filebuffersize: Size of stream file buffer. Default is 16384 (TIMEUNIT_RAWBYTES).
//
// filebuffersizetype: Type of unit for stream file buffer size.
// Must be TIMEUNIT_MS, TIMEUNIT_PCM, TIMEUNIT_PCMBYTES or TIMEUNIT_RAWBYTES. Default is TIMEUNIT_RAWBYTES.
//
// Note this function does not affect streams created with OPENUSER, as the buffer size is specified in "System.CreateSound".
// This function does not affect latency of playback. All streams are pre-buffered (unless opened with OPENONLY), so they will always start immediately.
// Seek and Play operations can sometimes cause a reflush of this buffer.
//
// If TIMEUNIT_RAWBYTES is used, the memory allocated is 2 * the size passed in, because fmod allocates a double buffer.
// If TIMEUNIT_MS, TIMEUNIT_PCM or TIMEUNIT_PCMBYTES is used, and the stream is infinite (such as a shoutcast netstream), or VBR,
// then FMOD cannot calculate an accurate compression ratio to work with when the file is opened.
// This means it will then base the buffersize on TIMEUNIT_PCMBYTES, or in other words the number of PCM bytes, but this will be incorrect for some compressed formats.
// Use TIMEUNIT_RAWBYTES for these type (infinite / undetermined length) of streams for more accurate read sizes.
//
// Note to determine the actual memory usage of a stream, including sound buffer and other overhead, use "MemoryGetStats" before and after creating a sound.
// Note that the stream may still stutter if the codec uses a large amount of cpu time, which impacts the smaller, internal 'decode' buffer.
// The decode buffer size is changeable via CREATESOUNDEXINFO.
func (s *System) SetStreamBufferSize(filebuffersize uint32, filebuffersizetype TimeUnit) error {
	res := C.FMOD_System_SetStreamBufferSize(s.cptr, C.uint(filebuffersize), C.FMOD_TIMEUNIT(filebuffersizetype))
	return errs[res]
}

// Returns the current internal buffersize settings for streamable sounds.
func (s *System) StreamBufferSize() (uint32, TimeUnit, error) {
	var filebuffersize C.uint
	var filebuffersizetype C.FMOD_TIMEUNIT
	res := C.FMOD_System_GetStreamBufferSize(s.cptr, &filebuffersize, &filebuffersizetype)
	return uint32(filebuffersize), TimeUnit(filebuffersizetype), errs[res]
}

// Sets the global doppler scale, distance factor and log rolloff scale for all 3D sound in FMOD.
// dopplerscale: Scaling factor for doppler shift. Default = 1.0.
//
// distancefactor: Relative distance factor to FMOD's units. Default = 1.0. (1.0 = 1 metre).
//
// rolloffscale: Scaling factor for 3D sound rolloff or attenuation for FMOD_3D_INVERSEROLLOFF based sounds only (which is the default type). Default = 1.0.
//
// The doppler scale is a general scaling factor for how much the pitch varies due to doppler shifting in 3D sound.
// Doppler is the pitch bending effect when a sound comes towards the listener or moves away from it, much like the effect you hear when a train goes past you with its horn sounding.
// With "dopplerscale" you can exaggerate or diminish the effect. FMOD's effective speed of sound at a doppler factor of 1.0 is 340 m/s.
//
// The distance factor is the FMOD 3D engine relative distance factor, compared to 1.0 meters.
// Another way to put it is that it equates to "how many units per meter does your engine have".
// For example, if you are using feet then "scale" would equal 3.28.
//
// Note! This only affects doppler!
// If you keep your min/max distance, custom rolloff curves and positions in scale relative to each other the volume rolloff will not change.
// If you set this, the mindistance of a sound will automatically set itself to this value when it is created in case the user forgets to set the mindistance to match the new distancefactor.
//
// The rolloff scale sets the global attenuation rolloff factor for FMOD_3D_INVERSEROLLOFF based sounds only (which is the default).
// Volume for a sound set to FMOD_3D_INVERSEROLLOFF will scale at mindistance / distance.
// This gives an inverse attenuation of volume as the source gets further away (or closer).
// Setting this value makes the sound drop off faster or slower. The higher the value, the faster volume will attenuate, and conversely the lower the value, the slower it will attenuate.
// For example a rolloff factor of 1 will simulate the real world, where as a value of 2 will make sounds attenuate 2 times quicker.
//
// Note! "rolloffscale" has no effect when using FMOD_3D_LINEARROLLOFF, FMOD_3D_LINEARSQUAREROLLOFF or FMOD_3D_CUSTOMROLLOFF.
func (s *System) Set3DSettings(dopplerscale, distancefactor, rolloffscale float32) error {
	res := C.FMOD_System_Set3DSettings(s.cptr, C.float(dopplerscale), C.float(distancefactor), C.float(rolloffscale))
	return errs[res]
}

// Retrieves the global doppler scale, distance factor and rolloff scale for all 3D sound in FMOD.
func (s *System) Get3DSettings() (float32, float32, float32, error) {
	var dopplerscale, distancefactor, rolloffscale C.float
	res := C.FMOD_System_Get3DSettings(s.cptr, &dopplerscale, &distancefactor, &rolloffscale)
	return float32(dopplerscale), float32(distancefactor), float32(rolloffscale), errs[res]
}

// Sets the number of 3D 'listeners' in the 3D sound scene. This function is useful mainly for split-screen game purposes.
//
// numlisteners: Number of listeners in the scene. Valid values are from 1 to MAX_LISTENERS inclusive. Default = 1.
//
// If the number of listeners is set to more than 1, then panning and doppler are turned off.
// All sound effects will be mono. FMOD uses a 'closest sound to the listener' method to determine what should be heard in this case.
func (s *System) Set3DNumListeners(numlisteners int) error {
	res := C.FMOD_System_Set3DNumListeners(s.cptr, C.int(numlisteners))
	return errs[res]
}

// Retrieves the number of 3D listeners.
func (s *System) Get3DNumListeners() (int, error) {
	var numlisteners C.int
	res := C.FMOD_System_Get3DNumListeners(s.cptr, &numlisteners)
	return int(numlisteners), errs[res]
}

// This updates the position, velocity and orientation of the specified 3D sound listener.
//
// listener: Listener ID in a multi-listener environment. Specify 0 if there is only 1 listener.
//
// pos: The position of the listener in world space, measured in distance units.
// You can specify 0 or NULL to not update the position.
//
// vel: The velocity of the listener measured in distance units per second.
// You can specify 0 or NULL to not update the velocity of the listener.
//
// forward: The forwards orientation of the listener. This vector must be of unit length and perpendicular to the up vector.
// You can specify 0 or NULL to not update the forwards orientation of the listener.
//
// up: The upwards orientation of the listener. This vector must be of unit length and perpendicular to the forwards vector.
// You can specify 0 or NULL to not update the upwards orientation of the listener.
//
// By default, FMOD uses a left-handed co-ordinate system. This means +X is right, +Y is up, and +Z is forwards.
// To change this to a right-handed coordinate system, use FMOD_INIT_3D_RIGHTHANDED. This means +X is right, +Y is up, and +Z is backwards or towards you.
//
// To map to another coordinate system, flip/negate and exchange these values.
//
// Orientation vectors are expected to be of UNIT length. This means the magnitude of the vector should be 1.0.
//
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
//
// Always remember to use units per second, not units per frame as this is a common mistake and will make the doppler effect sound wrong.
//
// For example, Do not just use (pos - lastpos) from the last frame's data for velocity, as this is not correct.
// You need to time compensate it so it is given in units per second.
func (s *System) Set3DListenerAttributes(listener int, pos, vel, forward, up Vector) error {
	var cpos C.FMOD_VECTOR = pos.toC()
	var cvel C.FMOD_VECTOR = vel.toC()
	var cforward C.FMOD_VECTOR = forward.toC()
	var cup C.FMOD_VECTOR = up.toC()
	res := C.FMOD_System_Set3DListenerAttributes(s.cptr, C.int(listener), &cpos, &cvel, &cforward, &cup)
	return errs[res]
}

// This retrieves the position, velocity and orientation of the specified 3D sound listener.
func (s *System) Get3DListenerAttributes(listener int) (pos, vel, forward, up Vector, err error) {
	var cpos, cvel, cforward, cup C.FMOD_VECTOR
	res := C.FMOD_System_Get3DListenerAttributes(s.cptr, C.int(listener), &cpos, &cvel, &cforward, &cup)
	err = errs[res]
	pos.fromC(cpos)
	vel.fromC(cvel)
	forward.fromC(cforward)
	up.fromC(cup)
	return
}

// NOTE: Not implement yet
// When FMOD wants to calculate 3d volume for a channel, this callback can be used to override the internal volume calculation based on distance.
//
// callback: Pointer to a C function of type FMOD_3D_ROLLOFF_CALLBACK, that is used to override the FMOD volume calculation.
// Default is 0 or NULL. Setting the callback to null will return 3d calculation back to FMOD.
func (s *System) Set3DRolloffCallback(callback C.FMOD_3D_ROLLOFF_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_System_Set3DRolloffCallback      (FMOD_SYSTEM *system, FMOD_3D_ROLLOFF_CALLBACK callback);
	return ErrNoImpl
}

// Suspend mixer thread and relinquish usage of audio hardware while maintaining internal state.
// Used on mobile platforms when entering a backgrounded state to reduce CPU to 0%.
// All internal state will be maintained, i.e. created sound and channels will stay available in memory.
func (s *System) MixerSuspend() error {
	res := C.FMOD_System_MixerSuspend(s.cptr)
	return errs[res]
}

// Resume mixer thread and reacquire access to audio hardware.
// Used on mobile platforms when entering the foreground after being suspended.
// All internal state will resume, i.e. created sound and channels are still valid and playback will continue.
func (s *System) MixerResume() error {
	res := C.FMOD_System_MixerResume(s.cptr)
	return errs[res]
}

// NOTE: Not implement yet
// Gets the default matrix used to convert from one speaker mode to another.
// The gain for source channel 's' to target channel 't' is matrix[t * matrixhop + s].
// If 'sourcespeakermode' or 'targetspeakermode' is SPEAKERMODE_RAW, this function will return error.
func (s *System) DefaultMixMatrix(sourcespeakermode, targetspeakermode C.FMOD_SPEAKERMODE, matrix *C.float, matrixhop C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetDefaultMixMatrix       (FMOD_SYSTEM *system, FMOD_SPEAKERMODE sourcespeakermode, FMOD_SPEAKERMODE targetspeakermode, float *matrix, int matrixhop);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Gets the a speaker mode's channel count.
func (s *System) SpeakerModeChannels(mode C.FMOD_SPEAKERMODE, channels *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetSpeakerModeChannels    (FMOD_SYSTEM *system, FMOD_SPEAKERMODE mode, int *channels);
	return ErrNoImpl
}

/*
   System information functions.
*/

// Returns the current version of FMOD Studio being used.
func (s *System) Version() (uint32, error) {
	var version C.uint
	res := C.FMOD_System_GetVersion(s.cptr, &version)
	return uint32(version), errs[res]
}

// NOTE: Not implement yet
// Retrieves a pointer to the system level output device module.
// This means a pointer to a DirectX "LPDIRECTSOUND", or a WINMM handle, or with something like with OUTPUTTYPE_NOSOUND output, the handle will be null or 0.
// Must be called after "System.Init".
func (s *System) OutputHandle(handle **interface{}) error {
	//FMOD_RESULT F_API FMOD_System_GetOutputHandle           (FMOD_SYSTEM *system, void **handle);
	return ErrNoImpl
}

// Retrieves the number of currently playing channels.
func (s *System) ChannelsPlaying() (int, error) {
	var channels C.int
	res := C.FMOD_System_GetChannelsPlaying(s.cptr, &channels)
	return int(channels), errs[res]
}

// TODO: Redo to native output instead map
// Retrieves in percent of CPU time - the amount of cpu usage that FMOD is taking for streaming/mixing and "System.Update" combined.
// This value is slightly smoothed to provide more stable readout (and to round off spikes that occur due to multitasking/operating system issues).
//
// NOTE! On ps3 and xbox360, the dsp and stream figures are NOT main cpu/main thread usage.
// On PS3 this is the percentage of SPU being used. On Xbox 360 it is the percentage of a hardware thread being used which is on a totally different CPU than the main one.
//
// Do not be alarmed if the usage for these platforms reaches over 50%, this is normal and should be ignored if you are playing a lot of compressed sounds and are using effects.
// The only value on the main cpu / main thread to take note of here that will impact your framerate is the update value, and this is typically very low (ie less than 1%).
func (s *System) CPUUsage() (map[string]float32, error) {
	var dsp, stream, geometry, update, total C.float
	res := C.FMOD_System_GetCPUUsage(s.cptr, &dsp, &stream, &geometry, &update, &total)
	cpu := map[string]float32{
		"dsp":      float32(dsp),
		"stream":   float32(stream),
		"geometry": float32(geometry),
		"update":   float32(update),
		"total":    float32(total),
	}
	return cpu, errs[res]
}

// TODO: Redo to native output instead map
// Retrieves the amount of dedicated sound ram available if the platform supports it.
// Most platforms use main ram to store audio data, so this function usually isn't necessary.
func (s *System) SoundRAM() (map[string]int, error) {
	var currentalloced, maxalloced, total C.int
	res := C.FMOD_System_GetSoundRAM(s.cptr, &currentalloced, &maxalloced, &total)
	ram := map[string]int{
		"currentalloced": int(currentalloced),
		"maxalloced":     int(maxalloced),
		"total":          int(total),
	}
	return ram, errs[res]
}

/*
   Sound/DSP/Channel/FX creation and retrieval.
*/

// TODO: add more docs
// Loads a sound into memory, or opens it for streaming.
//
// name_or_data: Name of the file or URL to open encoded in a UTF-8 string, or a pointer to a preloaded sound memory block
// if FMOD_OPENMEMORY/FMOD_OPENMEMORY_POINT is used. For CD playback the name should be a drive letter with a colon, example "D:" (windows only).
//
// mode: Behaviour modifier for opening the sound. See FMOD_MODE.
//
// exinfo: Pointer to a FMOD_CREATESOUNDEXINFO which lets the user provide extended information while playing the sound. Optional. Specify 0 or NULL to ignore.
func (s *System) CreateSound(name_or_data string, mode Mode, exinfo *CreatesSoundExInfo) (*Sound, error) {
	var sound Sound
	defer runtime.SetFinalizer(&sound, (*Sound).Release)
	cname_or_data := C.CString(name_or_data)
	defer C.free(unsafe.Pointer(cname_or_data))
	// FIX me
	res := C.FMOD_System_CreateSound(s.cptr, cname_or_data, C.FMOD_MODE(mode), (*C.FMOD_CREATESOUNDEXINFO)(null), &sound.cptr)
	return &sound, errs[res]
}

// TODO: add more docs
// Opens a sound for streaming.
// This function is a helper function that is the same as "System.CreateSound" but has the CREATESTREAM flag added internally.
func (s *System) CreateStream(name_or_data string, mode Mode, exinfo *CreatesSoundExInfo) (*Sound, error) {
	var sound Sound
	defer runtime.SetFinalizer(&sound, (*Sound).Release)
	cname_or_data := C.CString(name_or_data)
	defer C.free(unsafe.Pointer(cname_or_data))
	// FIX me
	res := C.FMOD_System_CreateStream(s.cptr, cname_or_data, C.FMOD_MODE(mode), (*C.FMOD_CREATESOUNDEXINFO)(null), &sound.cptr)
	return &sound, errs[res]
}

// Creates a user defined DSP unit object to be inserted into a DSP network, for the purposes of sound filtering or sound generation.
//
// description: Pointer of an DSP_DESCRIPTION structure containing information about the unit to be created.
// Some members of DSP_DESCRIPTION are referenced directly inside FMOD so the structure should be allocated statically or at least remain in memory for the lifetime of the system.
//
// A DSP unit can generate or filter incoming data.
// The data is created or filtered through use of the read callback that is defined by the user.
// See the definition for the DSP_DESCRIPTION structure to find out what each member means.
// To be active, a unit must be inserted into the FMOD DSP network to be heard.
// Use functions such as "ChannelGroup.AddDSP", "Channel.AddDSP" or "DSP.AddInput" to do this.
func (s *System) CreateDSP(description *DSPDesc) (*DSP, error) {
	var dsp DSP
	defer runtime.SetFinalizer(&dsp, (*DSP).Release)
	res := C.FMOD_System_CreateDSP(s.cptr, (*C.FMOD_DSP_DESCRIPTION)(description), &dsp.cptr)
	return &dsp, errs[res]
}

// Creates an FMOD defined built in DSP unit object to be inserted into a DSP network, for the purposes of sound filtering or sound generation.
// This function is used to create special effects that come built into FMOD.
//
// typ: A pre-defined DSP effect or sound generator described by a DSP_TYPE.
//
// Note! Winamp DSP and VST plugins will only return the first plugin of this type that was loaded!
// To access all VST or Winamp DSP plugins the "System.CreateDSPByPlugin" function!
// Use the index returned by "System.LoadPlugin" if you don't want to enumerate them all.
func (s *System) CreateDSPByType(typ DSPType) (*DSP, error) {
	// TODO Finalizer
	var dsp DSP
	res := C.FMOD_System_CreateDSPByType(s.cptr, C.FMOD_DSP_TYPE(typ), &dsp.cptr)
	return &dsp, errs[res]
}

// Creates a channel group object. These objects can be used to assign channels to for group channel settings, such as volume.
// Channel groups are also used for sub-mixing. Any channels that are assigned to a channel group get submixed into that channel group's DSP.
//
// name: Label to give to the channel group for identification purposes. Optional (can be null).
//
// See the channel group class definition for the types of operations that can be performed on 'groups' of channels.
// The channel group can for example be used to have 2 seperate groups of master volume, instead of one global master volume.
// A channel group can be used for sub-mixing, ie so that a set of channels can be mixed into a channel group, then can have effects applied to it without affecting other channels.
func (s *System) CreateChannelGroup(name string) (*ChannelGroup, error) {
	var channelgroup ChannelGroup
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.FMOD_System_CreateChannelGroup(s.cptr, cname, &channelgroup.cptr)
	defer runtime.SetFinalizer(&channelgroup, (*ChannelGroup).Release)
	return &channelgroup, errs[res]
}

// Creates a sound group, which can store handles to multiple Sound pointers.
//
// name: Name of sound group.
//
// Once a SoundGroup is created, "Sound.SetSoundGroup" is used to put a sound in a SoundGroup.
func (s *System) CreateSoundGroup(name string) (*SoundGroup, error) {
	var soundgroup = SoundGroup{name: name}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.FMOD_System_CreateSoundGroup(s.cptr, cname, &soundgroup.cptr)
	defer runtime.SetFinalizer(&soundgroup, (*SoundGroup).Release)
	return &soundgroup, errs[res]
}

// Creates a 'virtual reverb' object. This object reacts to 3d location and morphs the reverb environment based on how close it is to the reverb object's center.
// Multiple reverb objects can be created to achieve a multi-reverb environment.
// 1 Physical reverb object is used for all 3d reverb objects (slot 0 by default).
//
// The 3D reverb object is a sphere having 3D attributes (position, minimum distance, maximum distance) and reverb properties.
// The properties and 3D attributes of all reverb objects collectively determine, along with the listener's position, the settings of and input gains into a single 3D reverb DSP.
// When the listener is within the sphere of effect of one or more 3d reverbs, the listener's 3D reverb properties are a weighted combination of such 3d reverbs.
// When the listener is outside all of the reverbs, no reverb is applied.
//
// In FMOD Ex a special 'ambient' reverb setting was used when outside the influence of all reverb spheres. This function no longer exists.
// In FMOD Studio "System.SetReverbProperties" can be used to create an alternative reverb that can be used for 2D and background global reverb.
// To avoid this reverb intefering with the reverb slot used by the 3d reverb, 2d reverb should use a different slot id with "System.SetReverbProperties",
// otherwise ADVANCEDSETTINGS::reverb3Dinstance can also be used to place 3d reverb on a different physical reverb slot.
//
// Creating multiple reverb objects does not impact performance. These are 'virtual reverbs'.
// There will still be only 1 physical reverb DSP running that just morphs between the different virtual reverbs.
//
// Note about phsyical reverb DSP unit allocation.
// To remove the DSP unit and the associated CPU cost, first make sure all 3d reverb objects are released.
// Then call "System.SetReverbProperties" with the 3d reverb's slot ID (default is 0) with a property point of 0 or NULL, to signal that the physical reverb instance should be deleted.
// If a 3d reverb is still present, and "System.SetReverbProperties" function is called to free the physical reverb,
// the 3D reverb system will immediately recreate it upon the next "System.Update" call.
func (s *System) CreateReverb3D() (*Reverb3D, error) {
	// TODO Finalizer
	var reverb3d Reverb3D
	defer runtime.SetFinalizer(&reverb3d, (*Reverb3D).Release)
	res := C.FMOD_System_CreateReverb3D(s.cptr, &reverb3d.cptr)
	return &reverb3d, errs[res]
}

// Plays a sound object on a particular channel and ChannelGroup if desired.
//
// sound: Pointer to the sound to play. This is opened with "System.CreateSound".
//
// channelgroup: Pointer to a channelgroup become a member of.
// This is more efficient than using "Channel.SetChannelGroup", as it does it during the channel setup, rather than connecting to the master channel group,
// then later disconnecting and connecting to the new channelgroup when specified. Use 0/NULL to ignore.
//
// paused: true or false flag to specify whether to start the channel paused or not.
// Starting a channel paused allows the user to alter its attributes without it being audible, and unpausing with "Channel.SetPaused" actually starts the sound.
//
// When a sound is played, it will use the sound's default frequency and priority.
//
// A sound defined as FMOD_3D will by default play at the position of the listener.
// To set the 3D position of the channel before the sound is audible, start the channel paused by setting the paused flag to true,
// and calling "Channel.Set3DAttributes". Following that, unpause the channel with "Channel.SetPaused".
//
// Channels are reference counted. If a channel is stolen by the FMOD priority system, then the handle to the stolen voice becomes invalid,
// and Channel based commands will not affect the new sound playing in its place.
// If all channels are currently full playing a sound, FMOD will steal a channel with the lowest priority sound.
//
// If more channels are playing than are currently available on the soundcard/sound device or software mixer, then FMOD will 'virtualize' the channel.
// This type of channel is not heard, but it is updated as if it was playing.
// When its priority becomes high enough or another sound stops that was using a real hardware/software channel, it will start playing from where it should be.
// This technique saves CPU time (thousands of sounds can be played at once without actually being mixed or taking up resources), and also removes the need for the user to manage voices themselves.
// An example of virtual channel usage is a dungeon with 100 torches burning, all with a looping crackling sound, but with a soundcard that only supports 32 hardware voices.
// If the 3D positions and priorities for each torch are set correctly, FMOD will play all 100 sounds without any 'out of channels' errors,
// and swap the real voices in and out according to which torches are closest in 3D space.
// Priority for virtual channels can be changed in the sound's defaults, or at runtime with "Channel.SetPriority".
func (s *System) PlaySound(sound *Sound, channelgroup *ChannelGroup, paused bool) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_PlaySound(s.cptr, sound.cptr, (*C.FMOD_CHANNELGROUP)(null), getBool(paused), &channel.cptr)
	return &channel, errs[res]
}

// Plays a DSP unit object and its input network on a particular channel.
//
// dsp: Pointer to the dsp unit to play. This is opened with "System.CreateDSP", "System.CreateDSPByType", "System.CreateDSPByPlugin".
//
// channelgroup: Pointer to a channelgroup become a member of.
// This is more efficient than using "Channel.SetChannelGroup", as it does it during the channel setup,
// rather than connecting to the master channel group, then later disconnecting and connecting to the new channelgroup when specified. Use 0/NULL to ignore.
//
// paused: true or false flag to specify whether to start the channel paused or not.
// Starting a channel paused allows the user to alter its attributes without it being audible, and unpausing with "Channel.SetPaused" actually starts the dsp running.
//
// When a dsp is played, it will use the dsp's default frequency, volume, pan, levels and priority.
//
// A dsp defined as FMOD_3D will by default play at the position of the listener.
// To change channel attributes before the dsp is audible, start the channel paused by setting the paused flag to true, and calling the relevant channel based functions.
// Following that, unpause the channel with "Channel.SetPaused".
// Channels are reference counted. If a channel is stolen by the FMOD priority system, then the handle to the stolen voice becomes invalid,
// and Channel based commands will not affect the new channel playing in its place.
// If all channels are currently full playing a dsp or sound, FMOD will steal a channel with the lowest priority dsp or sound.
// If more channels are playing than are currently available on the soundcard/sound device or software mixer, then FMOD will 'virtualize' the channel.
// This type of channel is not heard, but it is updated as if it was playing. When its priority becomes high enough or another sound stops that was using a real hardware/software channel,
// it will start playing from where it should be. This technique saves CPU time (thousands of sounds can be played at once without actually being mixed or taking up resources),
// and also removes the need for the user to manage voices themselves.
// An example of virtual channel usage is a dungeon with 100 torches burning, all with a looping crackling sound, but with a soundcard that only supports 32 hardware voices.
// If the 3D positions and priorities for each torch are set correctly, FMOD will play all 100 sounds without any 'out of channels' errors, and swap the real voices in
// and out according to which torches are closest in 3D space.
// Priority for virtual channels can be changed in the sound's defaults, or at runtime with "Channel.SetPriority".
func (s *System) PlayDSP(dsp *DSP, channelgroup *ChannelGroup, paused bool) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_PlayDSP(s.cptr, dsp.cptr, (*C.FMOD_CHANNELGROUP)(null), getBool(paused), &channel.cptr)
	return &channel, errs[res]
}

// Retrieves a handle to a channel by ID.
// This function is mainly for getting handles to existing (playing) channels and setting their attributes.
func (s *System) Channel(channelid int) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_GetChannel(s.cptr, C.int(channelid), &channel.cptr)
	return &channel, errs[res]
}

// Retrieves a handle to the internal master channel group. This is the default channel group that all channels play on.
// This channel group can be used to do things like set the master volume for all playing sounds.
// See the ChannelGroup API for more functionality.
func (s *System) MasterChannelGroup() (*ChannelGroup, error) {
	var channelgroup ChannelGroup
	res := C.FMOD_System_GetMasterChannelGroup(s.cptr, &channelgroup.cptr)
	return &channelgroup, errs[res]
}

// Retrieves the default sound group, where all sounds are placed when they are created.
// If a user based soundgroup is deleted/released, the sounds will be put back into this sound group.
func (s *System) MasterSoundGroup() (*SoundGroup, error) {
	var soundgroup SoundGroup
	res := C.FMOD_System_GetMasterSoundGroup(s.cptr, &soundgroup.cptr)
	return &soundgroup, errs[res]
}

/*
   Routing to ports.
*/

// NOTE: Not implement yet
// Route the signal from a channel group into a seperate audio port on the output driver.
//
// portType: Output driver specific audio port type
//
// portIndex: Output driver specific index of the audio port
//
// channelgroup: Channel group to route away to the new port
//
// passThru: If true the signal will continue to be passed through to the main mix, if false the signal will be entirely to the designated port.
func (s *System) AttachChannelGroupToPort(portType C.FMOD_PORT_TYPE, portIndex C.FMOD_PORT_INDEX, channelgroup *C.FMOD_CHANNELGROUP, passThru C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_AttachChannelGroupToPort  (FMOD_SYSTEM *system, FMOD_PORT_TYPE portType, FMOD_PORT_INDEX portIndex, FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL passThru);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Disconnect a channel group from a and route audio back to the default port of the output driver
func (s *System) DetachChannelGroupFromPort(channelgroup *C.FMOD_CHANNELGROUP) error {
	//FMOD_RESULT F_API FMOD_System_DetachChannelGroupFromPort(FMOD_SYSTEM *system, FMOD_CHANNELGROUP *channelgroup);
	return ErrNoImpl
}

/*
   Reverb API.
*/

// NOTE: Not implement yet
// Sets parameters for the global reverb environment.
// Reverb parameters can be set manually, or automatically using the pre-defined presets given in the fmod.h header.
func (s *System) SetReverbProperties(instance C.int, prop *C.FMOD_REVERB_PROPERTIES) error {
	//FMOD_RESULT F_API FMOD_System_SetReverbProperties       (FMOD_SYSTEM *system, int instance, const FMOD_REVERB_PROPERTIES *prop);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the current reverb environment for the specified reverb instance.
// You must specify the 'Instance' value (usually 0 unless you are using multiple reverbs) before calling this function.
// Note! It is important to specify the 'Instance' value in the REVERB_PROPERTIES structure correctly, otherwise you will get an FMOD_ERR_REVERB_INSTANCE error.
func (s *System) ReverbProperties(instance C.int, prop *C.FMOD_REVERB_PROPERTIES) error {
	//FMOD_RESULT F_API FMOD_System_GetReverbProperties       (FMOD_SYSTEM *system, int instance, FMOD_REVERB_PROPERTIES *prop);
	return ErrNoImpl
}

/*
   System level DSP functionality.
*/

// Mutual exclusion function to lock the FMOD DSP engine (which runs asynchronously in another thread), so that it will not execute.
// If the FMOD DSP engine is already executing, this function will block until it has completed.
// The function may be used to synchronize DSP network operations carried out by the user.
// An example of using this function may be for when the user wants to construct a DSP sub-network,
// without the DSP engine executing in the background while the sub-network is still under construction.
//
// Once the user no longer needs the DSP engine locked, it must be unlocked with "System.UnlockDSP()".
// Note that the DSP engine should not be locked for a significant amount of time, otherwise inconsistency in the audio output may result. (audio skipping/stuttering).
func (s *System) LockDSP() error {
	res := C.FMOD_System_LockDSP(s.cptr)
	return errs[res]
}

// Mutual exclusion function to unlock the FMOD DSP engine (which runs asynchronously in another thread) and let it continue executing.
// The DSP engine must be locked with "System.LockDSP()" before this function is called.
func (s *System) UnlockDSP() error {
	res := C.FMOD_System_UnlockDSP(s.cptr)
	return errs[res]

}

/*
   Recording API.
*/

// NOTE: Not implement yet
// Retrieves the number of recording devices available for this output mode.
// Use this to enumerate all recording devices possible so that the user can select one.
func (s *System) RecordNumDrivers(numdrivers, numconnected *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordNumDrivers       (FMOD_SYSTEM *system, int *numdrivers, int *numconnected);
	return ErrNoImpl
}

// NOTE: Not implement yet
// TODO: add more docs
// Retrieves identification information about a sound device specified by its index, and specific to the output mode set with "System.SetOutput".
func (s *System) RecordDriverInfo(id C.int, name *C.char, namelen C.int, guid *C.FMOD_GUID, systemrate *C.int, speakermode *C.FMOD_SPEAKERMODE, speakermodechannels *C.int, state *C.FMOD_DRIVER_STATE) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordDriverInfo       (FMOD_SYSTEM *system, int id, char *name, int namelen, FMOD_GUID *guid, int *systemrate, FMOD_SPEAKERMODE *speakermode, int *speakermodechannels, FMOD_DRIVER_STATE *state);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the current recording position of the record buffer in PCM samples.
//
// id: Enumerated driver ID. This must be in a valid range delimited by "System.RecordNumDrivers".
func (s *System) RecordPosition(id C.int, position *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordPosition         (FMOD_SYSTEM *system, int id, unsigned int *position);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Starts the recording engine recording to the specified recording sound.
//
// id: Enumerated driver ID. This must be in a valid range delimited by "System.RecordNumDrivers".
//
// sound: User created sound for the user to record to.
//
// loop: Boolean flag to tell the recording engine whether to continue recording to the provided sound from the start again, after it has reached the end.
// If this is set to true the data will be continually be overwritten once every loop.
func (s *System) RecordStart(id C.int, sound *C.FMOD_SOUND, loop C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_RecordStart               (FMOD_SYSTEM *system, int id, FMOD_SOUND *sound, FMOD_BOOL loop);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Stops the recording engine from recording to the specified recording sound.
func (s *System) RecordStop(id C.int) error {
	//FMOD_RESULT F_API FMOD_System_RecordStop                (FMOD_SYSTEM *system, int id);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the state of the FMOD recording API, ie if it is currently recording or not.
func (s *System) IsRecording(id C.int, recording *C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_IsRecording               (FMOD_SYSTEM *system, int id, FMOD_BOOL *recording);
	return ErrNoImpl
}

/*
   Geometry API.
*/

// Geometry creation function. This function will create a base geometry object which can then have polygons added to it.
//
// maxpolygons: Maximum number of polygons within this object.
//
// maxvertices: Maximum number of vertices within this object.
//
// Polygons can be added to a geometry object using "Geometry.AddPolygon".
// A geometry object stores its list of polygons in a structure optimized for quick line intersection testing and efficient insertion and updating.
// The structure works best with regularly shaped polygons with minimal overlap.
// Many overlapping polygons, or clusters of long thin polygons may not be handled efficiently.
// Axis aligned polygons are handled most efficiently.
//
// The same type of structure is used to optimize line intersection testing with multiple geometry objects.
// It is important to set the value of maxworldsize to an appropriate value using "System.SetGeometrySettings".
// Objects or polygons outside the range of maxworldsize will not be handled efficiently.
// Conversely, if maxworldsize is excessively large, the structure may lose precision and efficiency may drop.
func (s *System) CreateGeometry(maxpolygons, maxvertices int) (Geometry, error) {
	var geom Geometry
	defer runtime.SetFinalizer(&geom, (*Geometry).Release)
	res := C.FMOD_System_CreateGeometry(s.cptr, C.int(maxpolygons), C.int(maxvertices), &geom.cptr)
	return geom, errs[res]
}

// Sets the maximum world size for the geometry engine for performance / precision reasons.
// Setting maxworldsize should be done first before creating any geometry.
// It can be done any time afterwards but may be slow in this case.
// Objects or polygons outside the range of maxworldsize will not be handled efficiently.
// Conversely, if maxworldsize is excessively large, the structure may loose precision and efficiency may drop.
func (s *System) SetGeometrySettings(maxworldsize float64) error {
	res := C.FMOD_System_SetGeometrySettings(s.cptr, C.float(maxworldsize))
	return errs[res]
}

// Retrieves the maximum world size for the geometry engine.
func (s *System) GeometrySettings() (float64, error) {
	var maxworldsize C.float
	res := C.FMOD_System_GetGeometrySettings(s.cptr, &maxworldsize)
	return float64(maxworldsize), errs[res]
}

// NOTE: Not implement yet
// Creates a geometry object from a block of memory which contains pre-saved geometry data, saved by "Geometry.Save".
func (s *System) LoadGeometry(data *interface{}, datasize C.int, geometry **C.FMOD_GEOMETRY) error {
	//FMOD_RESULT F_API FMOD_System_LoadGeometry              (FMOD_SYSTEM *system, const void *data, int datasize, FMOD_GEOMETRY **geometry);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Calculates geometry occlusion between a listener and a sound source.
// If single sided polygons have been created, it is important to get the source and listener positions round the right way,
// as the occlusion from point A to point B may not be the same as the occlusion from point B to point A.
func (s *System) GeometryOcclusion(listener, source *C.FMOD_VECTOR, direct, reverb C.float) error {
	//FMOD_RESULT F_API FMOD_System_GetGeometryOcclusion      (FMOD_SYSTEM *system, const FMOD_VECTOR *listener, const FMOD_VECTOR *source, float *direct, float *reverb);
	return ErrNoImpl
}

/*
   Network functions.
*/

// Set a proxy server to use for all subsequent internet connections.
// Basic authentication is supported. To use it, this parameter must be in
// user:password@host:port format e.g. bob:sekrit123@www.fmod.org:8888 Set this parameter to 0 / NULL if no proxy is required.
func (s *System) SetNetworkProxy(proxy string) error {
	cproxy := C.CString(proxy)
	defer C.free(unsafe.Pointer(cproxy))
	res := C.FMOD_System_SetNetworkProxy(s.cptr, cproxy)
	return errs[res]
}

// NOTE: Not implement yet
// Retrieves the URL of the proxy server used in internet streaming.
func (s *System) NetworkProxy(proxy *C.char, proxylen C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetNetworkProxy           (FMOD_SYSTEM *system, char *proxy, int proxylen);
	return ErrNoImpl
}

// Set the timeout for network streams.
// timeout: The timeout value in ms.
func (s *System) SetNetworkTimeout(timeout int) error {
	res := C.FMOD_System_SetNetworkTimeout(s.cptr, C.int(timeout))
	return errs[res]
}

// Retrieve the timeout value for network streams
func (s *System) NetworkTimeout() (int, error) {
	var timeout C.int
	res := C.FMOD_System_GetNetworkTimeout(s.cptr, &timeout)
	return int(timeout), errs[res]
}

/*
   Userdata set/get.
*/

// Sets a user value that the System object will store internally. Can be retrieved with "System.UserData".
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
// It can be useful if an FMOD callback passes an object of this type as a parameter, and the user does not know which object it is (if many of these types of objects exist).
// Using "System.UserData" would help in the identification of the object.
func (s *System) SetUserData(userdata interface{}) error {
	data := *(*[]*C.char)(unsafe.Pointer(&userdata))
	res := C.FMOD_System_SetUserData(s.cptr, unsafe.Pointer(&data))
	return errs[res]
}

//Retrieves the user value that that was set by calling the System.SetUserData function.
func (s *System) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_System_GetUserData(s.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}
