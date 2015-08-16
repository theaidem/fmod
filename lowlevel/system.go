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

type System struct {
	cptr *C.FMOD_SYSTEM
}

/*
   'System' API
*/

func SystemCreate() (*System, error) {
	var s System
	res := C.FMOD_System_Create(&s.cptr)
	runtime.SetFinalizer(&s, (*System).Release)
	return &s, errs[res]
}

func (s *System) Release() error {
	res := C.FMOD_System_Release(s.cptr)
	return errs[res]
}

/*
   Setup functions.
*/

func (s *System) SetOutput(output OutputType) error {
	res := C.FMOD_System_SetOutput(s.cptr, C.FMOD_OUTPUTTYPE(output))
	return errs[res]
}

func (s *System) GetOutput() (OutputType, error) {
	var output C.FMOD_OUTPUTTYPE
	res := C.FMOD_System_GetOutput(s.cptr, &output)
	return OutputType(output), errs[res]
}

func (s *System) GetNumDrivers() (int, error) {
	var numdrivers C.int
	res := C.FMOD_System_GetNumDrivers(s.cptr, &numdrivers)
	return int(numdrivers), errs[res]
}

func (s *System) GetDriverInfo(id int, name string, namelen int) (Guid, int, SpeakerMode, int, error) {
	var guid C.FMOD_GUID
	var systemrate C.int
	var speakermode C.FMOD_SPEAKERMODE
	var speakermodechannels C.int
	res := C.FMOD_System_GetDriverInfo(s.cptr, C.int(id), C.CString(name), C.int(namelen), &guid, &systemrate, &speakermode, &speakermodechannels)
	return Guid(guid), int(systemrate), SpeakerMode(speakermode), int(speakermodechannels), errs[res]
}

func (s *System) SetDriver(driver int) error {
	res := C.FMOD_System_SetDriver(s.cptr, C.int(driver))
	return errs[res]
}

func (s *System) GetDriver() (int, error) {
	var driver C.int
	res := C.FMOD_System_GetDriver(s.cptr, &driver)
	return int(driver), errs[res]
}

func (s *System) SetSoftwareChannels(numsoftwarechannels int) error {
	res := C.FMOD_System_SetSoftwareChannels(s.cptr, C.int(numsoftwarechannels))
	return errs[res]
}

func (s *System) GetSoftwareChannels() (int, error) {
	var numsoftwarechannels C.int
	res := C.FMOD_System_GetSoftwareChannels(s.cptr, &numsoftwarechannels)
	return int(numsoftwarechannels), errs[res]
}

func (s *System) SetSoftwareFormat(samplerate int, speakermode SpeakerMode, numrawspeakers int) error {
	res := C.FMOD_System_SetSoftwareFormat(s.cptr, C.int(samplerate), C.FMOD_SPEAKERMODE(speakermode), C.int(numrawspeakers))
	return errs[res]
}

func (s *System) GetSoftwareFormat() (int, SpeakerMode, int, error) {
	var samplerate C.int
	var speakermode C.FMOD_SPEAKERMODE
	var numrawspeakers C.int
	res := C.FMOD_System_GetSoftwareFormat(s.cptr, &samplerate, &speakermode, &numrawspeakers)
	return int(samplerate), SpeakerMode(speakermode), int(numrawspeakers), errs[res]
}

func (s *System) SetDSPBufferSize(bufferlength uint32, numbuffers int) error {
	res := C.FMOD_System_SetDSPBufferSize(s.cptr, C.uint(bufferlength), C.int(numbuffers))
	return errs[res]
}

func (s *System) GetDSPBufferSize() (uint32, int, error) {
	var bufferlength C.uint
	var numbuffers C.int
	res := C.FMOD_System_GetDSPBufferSize(s.cptr, &bufferlength, &numbuffers)
	return uint32(bufferlength), int(numbuffers), errs[res]
}

func (s *System) SetFileSystem(useropen C.FMOD_FILE_OPEN_CALLBACK, userclose C.FMOD_FILE_CLOSE_CALLBACK, userread C.FMOD_FILE_READ_CALLBACK, userseek C.FMOD_FILE_SEEK_CALLBACK, userasyncread C.FMOD_FILE_ASYNCREAD_CALLBACK, userasynccancel C.FMOD_FILE_ASYNCCANCEL_CALLBACK, blockalign C.int) error {
	//FMOD_RESULT F_API FMOD_System_SetFileSystem             (FMOD_SYSTEM *system, FMOD_FILE_OPEN_CALLBACK useropen, FMOD_FILE_CLOSE_CALLBACK userclose, FMOD_FILE_READ_CALLBACK userread, FMOD_FILE_SEEK_CALLBACK userseek, FMOD_FILE_ASYNCREAD_CALLBACK userasyncread, FMOD_FILE_ASYNCCANCEL_CALLBACK userasynccancel, int blockalign);
	return ErrNoImpl
}

func (s *System) AttachFileSystem(useropen C.FMOD_FILE_OPEN_CALLBACK, userclose C.FMOD_FILE_CLOSE_CALLBACK, userread C.FMOD_FILE_READ_CALLBACK, userseek C.FMOD_FILE_SEEK_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_System_AttachFileSystem          (FMOD_SYSTEM *system, FMOD_FILE_OPEN_CALLBACK useropen, FMOD_FILE_CLOSE_CALLBACK userclose, FMOD_FILE_READ_CALLBACK userread, FMOD_FILE_SEEK_CALLBACK userseek);
	return ErrNoImpl
}

func (s *System) SetAdvancedSettings(settings *AdvancedSettings) error {
	var csettings = settings.toC()
	res := C.FMOD_System_SetAdvancedSettings(s.cptr, &csettings)
	return errs[res]
}

func (s *System) GetAdvancedSettings() (*AdvancedSettings, error) {
	var settings C.FMOD_ADVANCEDSETTINGS
	settings.cbSize = C.int(unsafe.Sizeof(settings))
	res := C.FMOD_System_GetAdvancedSettings(s.cptr, &settings)
	/*ॐ*/
	as := new(AdvancedSettings)
	as.fromC(settings)
	return as, errs[res]
}

func (s *System) SetCallback(callback C.FMOD_SYSTEM_CALLBACK, callbackmask C.FMOD_SYSTEM_CALLBACK_TYPE) error {
	//FMOD_RESULT F_API FMOD_System_SetCallback               (FMOD_SYSTEM *system, FMOD_SYSTEM_CALLBACK callback, FMOD_SYSTEM_CALLBACK_TYPE callbackmask);
	return ErrNoImpl
}

/*
   Plug-in support.
*/

func (s *System) SetPluginPath(path *C.char) error {
	//FMOD_RESULT F_API FMOD_System_SetPluginPath             (FMOD_SYSTEM *system, const char *path);
	return ErrNoImpl
}

func (s *System) LoadPlugin(filename *C.char, handle *C.uint, priority C.uint) error {
	//FMOD_RESULT F_API FMOD_System_LoadPlugin                (FMOD_SYSTEM *system, const char *filename, unsigned int *handle, unsigned int priority);
	return ErrNoImpl
}

func (s *System) UnloadPlugin(handle C.uint) error {
	//FMOD_RESULT F_API FMOD_System_UnloadPlugin              (FMOD_SYSTEM *system, unsigned int handle);
	return ErrNoImpl
}

func (s *System) GetNumPlugins(plugintype PluginType) (int, error) {
	var numplugins C.int
	res := C.FMOD_System_GetNumPlugins(s.cptr, C.FMOD_PLUGINTYPE(plugintype), &numplugins)
	return int(numplugins), errs[res]
}

func (s *System) GetPluginHandle(plugintype C.FMOD_PLUGINTYPE, index C.int, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetPluginHandle           (FMOD_SYSTEM *system, FMOD_PLUGINTYPE plugintype, int index, unsigned int *handle);
	return ErrNoImpl
}

func (s *System) GetPluginInfo(handle C.uint, plugintype *C.FMOD_PLUGINTYPE, name *C.char, version *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetPluginInfo             (FMOD_SYSTEM *system, unsigned int handle, FMOD_PLUGINTYPE *plugintype, char *name, int namelen, unsigned int *version);
	return ErrNoImpl
}

func (s *System) SetOutputByPlugin(handle C.uint) error {
	//FMOD_RESULT F_API FMOD_System_SetOutputByPlugin         (FMOD_SYSTEM *system, unsigned int handle);
	return ErrNoImpl
}

func (s *System) GetOutputByPlugin(handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetOutputByPlugin         (FMOD_SYSTEM *system, unsigned int *handle);
	return ErrNoImpl
}

func (s *System) CreateDSPByPlugin(handle C.uint, dsp **C.FMOD_DSP) error {
	//FMOD_RESULT F_API FMOD_System_CreateDSPByPlugin         (FMOD_SYSTEM *system, unsigned int handle, FMOD_DSP **dsp);
	return ErrNoImpl
}

func (s *System) GetDSPInfoByPlugin(handle C.uint, description **C.FMOD_DSP_DESCRIPTION) error {
	//FMOD_RESULT F_API FMOD_System_GetDSPInfoByPlugin        (FMOD_SYSTEM *system, unsigned int handle, const FMOD_DSP_DESCRIPTION **description);
	return ErrNoImpl
}

func (s *System) RegisterCodec(description *C.FMOD_CODEC_DESCRIPTION, handle *C.uint, priority C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterCodec             (FMOD_SYSTEM *system, FMOD_CODEC_DESCRIPTION *description, unsigned int *handle, unsigned int priority);
	return ErrNoImpl
}

func (s *System) RegisterDSP(description *C.FMOD_DSP_DESCRIPTION, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterDSP               (FMOD_SYSTEM *system, const FMOD_DSP_DESCRIPTION *description, unsigned int *handle);
	return ErrNoImpl
}

func (s *System) RegisterOutput(description *C.FMOD_OUTPUT_DESCRIPTION, handle *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_RegisterOutput            (FMOD_SYSTEM *system, const FMOD_OUTPUT_DESCRIPTION *description, unsigned int *handle);
	return ErrNoImpl
}

/*
   Init/Close.
*/

func (s *System) Init(maxchannels int, flags InitFlags, extradriverdata interface{}) error {
	res := C.FMOD_System_Init(s.cptr, C.int(maxchannels), C.FMOD_INITFLAGS(flags), unsafe.Pointer(uintptr(extradriverdata.(int))))
	return errs[res]
}

func (s *System) Close() error {
	res := C.FMOD_System_Close(s.cptr)
	return errs[res]
}

/*
   General post-init system functions.
*/

func (s *System) Update() error {
	res := C.FMOD_System_Update(s.cptr)
	return errs[res]
}

func (s *System) SetSpeakerPosition(speaker Speaker, x, y float32, active bool) error {
	res := C.FMOD_System_SetSpeakerPosition(s.cptr, C.FMOD_SPEAKER(speaker), C.float(x), C.float(y), getBool(active))
	return errs[res]
}

func (s *System) GetSpeakerPosition(speaker Speaker) (float32, float32, bool, error) {
	var x, y C.float
	var active C.FMOD_BOOL
	res := C.FMOD_System_GetSpeakerPosition(s.cptr, C.FMOD_SPEAKER(speaker), &x, &y, &active)
	return float32(x), float32(y), setBool(active), errs[res]
}

func (s *System) SetStreamBufferSize(filebuffersize uint32, filebuffersizetype TimeUnit) error {
	res := C.FMOD_System_SetStreamBufferSize(s.cptr, C.uint(filebuffersize), C.FMOD_TIMEUNIT(filebuffersizetype))
	return errs[res]
}

func (s *System) GetStreamBufferSize() (uint32, TimeUnit, error) {
	var filebuffersize C.uint
	var filebuffersizetype C.FMOD_TIMEUNIT
	res := C.FMOD_System_GetStreamBufferSize(s.cptr, &filebuffersize, &filebuffersizetype)
	return uint32(filebuffersize), TimeUnit(filebuffersizetype), errs[res]
}

func (s *System) Set3DSettings(dopplerscale, distancefactor, rolloffscale float32) error {
	res := C.FMOD_System_Set3DSettings(s.cptr, C.float(dopplerscale), C.float(distancefactor), C.float(rolloffscale))
	return errs[res]
}

func (s *System) Get3DSettings() (float32, float32, float32, error) {
	var dopplerscale, distancefactor, rolloffscale C.float
	res := C.FMOD_System_Get3DSettings(s.cptr, &dopplerscale, &distancefactor, &rolloffscale)
	return float32(dopplerscale), float32(distancefactor), float32(rolloffscale), errs[res]
}

func (s *System) Set3DNumListeners(numlisteners int) error {
	res := C.FMOD_System_Set3DNumListeners(s.cptr, C.int(numlisteners))
	return errs[res]
}

func (s *System) Get3DNumListeners() (int, error) {
	var numlisteners C.int
	res := C.FMOD_System_Get3DNumListeners(s.cptr, &numlisteners)
	return int(numlisteners), errs[res]
}

func (s *System) Set3DListenerAttributes(listener int, pos, vel, forward, up Vector) error {
	var cpos C.FMOD_VECTOR = pos.toC()
	var cvel C.FMOD_VECTOR = vel.toC()
	var cforward C.FMOD_VECTOR = forward.toC()
	var cup C.FMOD_VECTOR = up.toC()
	res := C.FMOD_System_Set3DListenerAttributes(s.cptr, C.int(listener), &cpos, &cvel, &cforward, &cup)
	return errs[res]
}

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

func (s *System) Set3DRolloffCallback(callback C.FMOD_3D_ROLLOFF_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_System_Set3DRolloffCallback      (FMOD_SYSTEM *system, FMOD_3D_ROLLOFF_CALLBACK callback);
	return ErrNoImpl
}

func (s *System) MixerSuspend() error {
	res := C.FMOD_System_MixerSuspend(s.cptr)
	return errs[res]
}

func (s *System) MixerResume() error {
	res := C.FMOD_System_MixerResume(s.cptr)
	return errs[res]
}

func (s *System) GetDefaultMixMatrix(sourcespeakermode, targetspeakermode C.FMOD_SPEAKERMODE, matrix *C.float, matrixhop C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetDefaultMixMatrix       (FMOD_SYSTEM *system, FMOD_SPEAKERMODE sourcespeakermode, FMOD_SPEAKERMODE targetspeakermode, float *matrix, int matrixhop);
	return ErrNoImpl
}

func (s *System) GetSpeakerModeChannels(mode C.FMOD_SPEAKERMODE, channels *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetSpeakerModeChannels    (FMOD_SYSTEM *system, FMOD_SPEAKERMODE mode, int *channels);
	return ErrNoImpl
}

/*
   System information functions.
*/

func (s *System) GetVersion() (uint32, error) {
	var version C.uint
	res := C.FMOD_System_GetVersion(s.cptr, &version)
	return uint32(version), errs[res]
}

func (s *System) GetOutputHandle(handle **interface{}) error {
	//FMOD_RESULT F_API FMOD_System_GetOutputHandle           (FMOD_SYSTEM *system, void **handle);
	return ErrNoImpl
}

func (s *System) GetChannelsPlaying() (int, error) {
	var channels C.int
	res := C.FMOD_System_GetChannelsPlaying(s.cptr, &channels)
	return int(channels), errs[res]
}

func (s *System) GetCPUUsage() (map[string]float32, error) {
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

func (s *System) GetSoundRAM() (map[string]int, error) {
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

func (s *System) CreateSound(name_or_data string, mode Mode, exinfo *CreatesSoundExInfo) (*Sound, error) {
	var sound Sound
	defer runtime.SetFinalizer(&sound, (*Sound).Release)
	// FIX me
	res := C.FMOD_System_CreateSound(s.cptr, C.CString(name_or_data), C.FMOD_MODE(mode), (*C.FMOD_CREATESOUNDEXINFO)(null), &sound.cptr)
	return &sound, errs[res]
}

func (s *System) CreateStream(name_or_data string, mode Mode, exinfo *CreatesSoundExInfo) (*Sound, error) {
	var sound Sound
	defer runtime.SetFinalizer(&sound, (*Sound).Release)
	// FIX me
	res := C.FMOD_System_CreateStream(s.cptr, C.CString(name_or_data), C.FMOD_MODE(mode), (*C.FMOD_CREATESOUNDEXINFO)(null), &sound.cptr)
	return &sound, errs[res]
}

func (s *System) CreateDSP(description *DSPDesc) (*DSP, error) {
	var dsp DSP
	res := C.FMOD_System_CreateDSP(s.cptr, (*C.FMOD_DSP_DESCRIPTION)(description), &dsp.cptr)
	return &dsp, errs[res]
}

func (s *System) CreateDSPByType(typ DSPType) (*DSP, error) {
	var dsp DSP
	res := C.FMOD_System_CreateDSPByType(s.cptr, C.FMOD_DSP_TYPE(typ), &dsp.cptr)
	return &dsp, errs[res]
}

func (s *System) CreateChannelGroup(name string) (*ChannelGroup, error) {
	var channelgroup ChannelGroup
	res := C.FMOD_System_CreateChannelGroup(s.cptr, C.CString(name), &channelgroup.cptr)
	return &channelgroup, errs[res]
}

func (s *System) CreateSoundGroup(name string) (*SoundGroup, error) {
	var soundgroup SoundGroup
	res := C.FMOD_System_CreateSoundGroup(s.cptr, C.CString(name), &soundgroup.cptr)
	return &soundgroup, errs[res]
}

func (s *System) CreateReverb3D() (*Reverb3D, error) {
	var reverb3d Reverb3D
	res := C.FMOD_System_CreateReverb3D(s.cptr, &reverb3d.cptr)
	return &reverb3d, errs[res]
}

func (s *System) PlaySound(sound *Sound, channelgroup *ChannelGroup, paused bool) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_PlaySound(s.cptr, sound.cptr, (*C.FMOD_CHANNELGROUP)(null), getBool(paused), &channel.cptr)
	return &channel, errs[res]
}

func (s *System) PlayDSP(dsp *DSP, channelgroup *ChannelGroup, paused bool) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_PlayDSP(s.cptr, dsp.cptr, (*C.FMOD_CHANNELGROUP)(null), getBool(paused), &channel.cptr)
	return &channel, errs[res]
}

func (s *System) GetChannel(channelid int) (*Channel, error) {
	var channel Channel
	res := C.FMOD_System_GetChannel(s.cptr, C.int(channelid), &channel.cptr)
	return &channel, errs[res]
}

func (s *System) GetMasterChannelGroup() (*ChannelGroup, error) {
	var channelgroup ChannelGroup
	res := C.FMOD_System_GetMasterChannelGroup(s.cptr, &channelgroup.cptr)
	return &channelgroup, errs[res]
}

func (s *System) GetMasterSoundGroup() (*SoundGroup, error) {
	var soundgroup SoundGroup
	res := C.FMOD_System_GetMasterSoundGroup(s.cptr, &soundgroup.cptr)
	return &soundgroup, errs[res]
}

/*
   Routing to ports.
*/

func (s *System) AttachChannelGroupToPort(portType C.FMOD_PORT_TYPE, portIndex C.FMOD_PORT_INDEX, channelgroup *C.FMOD_CHANNELGROUP, passThru C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_AttachChannelGroupToPort  (FMOD_SYSTEM *system, FMOD_PORT_TYPE portType, FMOD_PORT_INDEX portIndex, FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL passThru);
	return ErrNoImpl
}

func (s *System) DetachChannelGroupFromPort(channelgroup *C.FMOD_CHANNELGROUP) error {
	//FMOD_RESULT F_API FMOD_System_DetachChannelGroupFromPort(FMOD_SYSTEM *system, FMOD_CHANNELGROUP *channelgroup);
	return ErrNoImpl
}

/*
   Reverb API.
*/

func (s *System) SetReverbProperties(instance C.int, prop *C.FMOD_REVERB_PROPERTIES) error {
	//FMOD_RESULT F_API FMOD_System_SetReverbProperties       (FMOD_SYSTEM *system, int instance, const FMOD_REVERB_PROPERTIES *prop);
	return ErrNoImpl
}

func (s *System) GetReverbProperties(instance C.int, prop *C.FMOD_REVERB_PROPERTIES) error {
	//FMOD_RESULT F_API FMOD_System_GetReverbProperties       (FMOD_SYSTEM *system, int instance, FMOD_REVERB_PROPERTIES *prop);
	return ErrNoImpl
}

/*
   System level DSP functionality.
*/

func (s *System) LockDSP() error {
	res := C.FMOD_System_LockDSP(s.cptr)
	return errs[res]
}

func (s *System) UnlockDSP() error {
	res := C.FMOD_System_UnlockDSP(s.cptr)
	return errs[res]

}

/*
   Recording API.
*/

func (s *System) GetRecordNumDrivers(numdrivers, numconnected *C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordNumDrivers       (FMOD_SYSTEM *system, int *numdrivers, int *numconnected);
	return ErrNoImpl
}

func (s *System) GetRecordDriverInfo(id C.int, name *C.char, namelen C.int, guid *C.FMOD_GUID, systemrate *C.int, speakermode *C.FMOD_SPEAKERMODE, speakermodechannels *C.int, state *C.FMOD_DRIVER_STATE) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordDriverInfo       (FMOD_SYSTEM *system, int id, char *name, int namelen, FMOD_GUID *guid, int *systemrate, FMOD_SPEAKERMODE *speakermode, int *speakermodechannels, FMOD_DRIVER_STATE *state);
	return ErrNoImpl
}

func (s *System) GetRecordPosition(id C.int, position *C.uint) error {
	//FMOD_RESULT F_API FMOD_System_GetRecordPosition         (FMOD_SYSTEM *system, int id, unsigned int *position);
	return ErrNoImpl
}

func (s *System) RecordStart(id C.int, sound *C.FMOD_SOUND, loop C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_RecordStart               (FMOD_SYSTEM *system, int id, FMOD_SOUND *sound, FMOD_BOOL loop);
	return ErrNoImpl
}

func (s *System) RecordStop(id C.int) error {
	//FMOD_RESULT F_API FMOD_System_RecordStop                (FMOD_SYSTEM *system, int id);
	return ErrNoImpl
}

func (s *System) IsRecording(id C.int, recording *C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_System_IsRecording               (FMOD_SYSTEM *system, int id, FMOD_BOOL *recording);
	return ErrNoImpl
}

/*
   Geometry API.
*/

func (s *System) CreateGeometry(maxpolygons, maxvertices int) (Geometry, error) {
	var geom Geometry
	res := C.FMOD_System_CreateGeometry(s.cptr, C.int(maxpolygons), C.int(maxvertices), &geom.cptr)
	return geom, errs[res]
}

func (s *System) SetGeometrySettings(maxworldsize float64) error {
	res := C.FMOD_System_SetGeometrySettings(s.cptr, C.float(maxworldsize))
	return errs[res]
}

func (s *System) GetGeometrySettings() (float64, error) {
	var maxworldsize C.float
	res := C.FMOD_System_GetGeometrySettings(s.cptr, &maxworldsize)
	return float64(maxworldsize), errs[res]

}

func (s *System) LoadGeometry(data *interface{}, datasize C.int, geometry **C.FMOD_GEOMETRY) error {
	//FMOD_RESULT F_API FMOD_System_LoadGeometry              (FMOD_SYSTEM *system, const void *data, int datasize, FMOD_GEOMETRY **geometry);
	return ErrNoImpl
}

func (s *System) GetGeometryOcclusion(listener, source *C.FMOD_VECTOR, direct, reverb C.float) error {
	//FMOD_RESULT F_API FMOD_System_GetGeometryOcclusion      (FMOD_SYSTEM *system, const FMOD_VECTOR *listener, const FMOD_VECTOR *source, float *direct, float *reverb);
	return ErrNoImpl
}

/*
   Network functions.
*/

func (s *System) SetNetworkProxy(proxy string) error {
	cproxy := C.CString(proxy)
	defer C.free(unsafe.Pointer(cproxy))
	res := C.FMOD_System_SetNetworkProxy(s.cptr, cproxy)
	return errs[res]
}

func (s *System) GetNetworkProxy(proxy *C.char, proxylen C.int) error {
	//FMOD_RESULT F_API FMOD_System_GetNetworkProxy           (FMOD_SYSTEM *system, char *proxy, int proxylen);
	return ErrNoImpl
}

func (s *System) SetNetworkTimeout(timeout int) error {
	res := C.FMOD_System_SetNetworkTimeout(s.cptr, C.int(timeout))
	return errs[res]
}

func (s *System) GetNetworkTimeout() (int, error) {
	var timeout C.int
	res := C.FMOD_System_GetNetworkTimeout(s.cptr, &timeout)
	return int(timeout), errs[res]
}

/*
   Userdata set/get.
*/

func (s *System) SetUserData(userdata interface{}) error {
	//FMOD_RESULT F_API FMOD_System_SetUserData               (FMOD_SYSTEM *system, void *userdata);
	return ErrNoImpl
}

func (s *System) GetUserData() error {
	//FMOD_RESULT F_API FMOD_System_GetUserData               (FMOD_SYSTEM *system, void **userdata);
	return ErrNoImpl
}
