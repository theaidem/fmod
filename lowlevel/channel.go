package lowlevel

/*
#include <fmod.h>
*/
import "C"

type Channel struct {
	cptr *C.FMOD_CHANNEL
}

/*
   'Channel' API
*/

func (c *Channel) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_Channel_GetSystemObject(c.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   General control functionality for Channels and ChannelGroups.
*/

func (c *Channel) Stop() error {
	res := C.FMOD_Channel_Stop(c.cptr)
	return errs[res]
}

func (c *Channel) SetPaused(paused bool) error {
	res := C.FMOD_Channel_SetPaused(c.cptr, getBool(paused))
	return errs[res]
}

func (c *Channel) IsPaused() (bool, error) {
	var paused C.FMOD_BOOL
	res := C.FMOD_Channel_GetPaused(c.cptr, &paused)
	return setBool(paused), errs[res]
}

func (c *Channel) SetVolume(volume float64) error {
	res := C.FMOD_Channel_SetVolume(c.cptr, C.float(volume))
	return errs[res]
}

func (c *Channel) Volume() (float64, error) {
	var volume C.float
	res := C.FMOD_Channel_GetVolume(c.cptr, &volume)
	return float64(volume), errs[res]
}

func (c *Channel) SetVolumeRamp(ramp bool) error {
	res := C.FMOD_Channel_SetVolumeRamp(c.cptr, getBool(ramp))
	return errs[res]
}

func (c *Channel) VolumeRamp() (bool, error) {
	var ramp C.FMOD_BOOL
	res := C.FMOD_Channel_GetVolumeRamp(c.cptr, &ramp)
	return setBool(ramp), errs[res]
}

func (c *Channel) Audibility() (float64, error) {
	var audibility C.float
	res := C.FMOD_Channel_GetAudibility(c.cptr, &audibility)
	return float64(audibility), errs[res]
}

func (c *Channel) SetPitch(pitch float64) error {
	res := C.FMOD_Channel_SetPitch(c.cptr, C.float(pitch))
	return errs[res]
}

func (c *Channel) Pitch() (float64, error) {
	var pitch C.float
	res := C.FMOD_Channel_GetPitch(c.cptr, &pitch)
	return float64(pitch), errs[res]
}

func (c *Channel) SetMute(mute bool) error {
	res := C.FMOD_Channel_SetMute(c.cptr, getBool(mute))
	return errs[res]
}

func (c *Channel) Mute() (bool, error) {
	var mute C.FMOD_BOOL
	res := C.FMOD_Channel_GetMute(c.cptr, &mute)
	return setBool(mute), errs[res]
}

func (c *Channel) SetReverbProperties(instance int, wet float64) error {
	res := C.FMOD_Channel_SetReverbProperties(c.cptr, C.int(instance), C.float(wet))
	return errs[res]
}

func (c *Channel) ReverbProperties(instance int) (float64, error) {
	var wet C.float
	res := C.FMOD_Channel_GetReverbProperties(c.cptr, C.int(instance), &wet)
	return float64(wet), errs[res]
}

func (c *Channel) SetLowPassGain(gain float64) error {
	res := C.FMOD_Channel_SetLowPassGain(c.cptr, C.float(gain))
	return errs[res]
}

func (c *Channel) LowPassGain() (float64, error) {
	var gain C.float
	res := C.FMOD_Channel_GetLowPassGain(c.cptr, &gain)
	return float64(gain), errs[res]
}

func (c *Channel) SetMode(mode Mode) error {
	res := C.FMOD_Channel_SetMode(c.cptr, C.FMOD_MODE(mode))
	return errs[res]
}

func (c *Channel) Mode() (Mode, error) {
	var mode C.FMOD_MODE
	res := C.FMOD_Channel_GetMode(c.cptr, &mode)
	return Mode(mode), errs[res]
}

// NOTE: Not implement yet
func (c *Channel) SetCallback(callback C.FMOD_CHANNELCONTROL_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_Channel_SetCallback(FMOD_CHANNEL *channel, FMOD_CHANNELCONTROL_CALLBACK callback);
	return ErrNoImpl
}

func (c *Channel) IsPlaying() (bool, error) {
	var isplaying C.FMOD_BOOL
	res := C.FMOD_Channel_IsPlaying(c.cptr, &isplaying)
	return setBool(isplaying), errs[res]
}

/*
   Note all 'set' functions alter a final matrix, this is why the only get function is getMixMatrix, to avoid other get functions returning incorrect/obsolete values.
*/

func (c *Channel) SetPan(pan float64) error {
	res := C.FMOD_Channel_SetPan(c.cptr, C.float(pan))
	return errs[res]
}

func (c *Channel) SetMixLevelsOutput(frontleft, frontright, center, lfe, surroundleft, surroundright, backleft, backright float64) error {
	res := C.FMOD_Channel_SetMixLevelsOutput(c.cptr, C.float(frontleft), C.float(frontright), C.float(center), C.float(lfe), C.float(surroundleft), C.float(surroundright), C.float(backleft), C.float(backright))
	return errs[res]
}

// NOTE: Not implement yet
func (c *Channel) SetMixLevelsInput(levels *C.float, numlevels C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_SetMixLevelsInput(FMOD_CHANNEL *channel, float *levels, int numlevels);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (c *Channel) SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_SetMixMatrix(FMOD_CHANNEL *channel, float *matrix, int outchannels, int inchannels, int inchannel_hop);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (c *Channel) MixMatrix(matrix *C.float, outchannels, inchannels *C.int, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_GetMixMatrix(FMOD_CHANNEL *channel, float *matrix, int *outchannels, int *inchannels, int inchannel_hop);
	return ErrNoImpl
}

/*
   Clock based functionality.
*/

func (c *Channel) DSPClock() (uint64, uint64, error) {
	var dspclock, parentclock C.ulonglong
	res := C.FMOD_Channel_GetDSPClock(c.cptr, &dspclock, &parentclock)
	return uint64(dspclock), uint64(parentclock), errs[res]
}

func (c *Channel) SetDelay(dspclock_start, dspclock_end uint64, stopchannels bool) error {
	res := C.FMOD_Channel_SetDelay(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end), getBool(stopchannels))
	return errs[res]
}

func (c *Channel) Delay() (uint64, uint64, bool, error) {
	var dspclock_start, dspclock_end C.ulonglong
	var stopchannels C.FMOD_BOOL
	res := C.FMOD_Channel_GetDelay(c.cptr, &dspclock_start, &dspclock_end, &stopchannels)
	return uint64(dspclock_start), uint64(dspclock_end), setBool(stopchannels), errs[res]
}

func (c *Channel) AddFadePoint(dspclock uint64, volume float64) error {
	res := C.FMOD_Channel_AddFadePoint(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

func (c *Channel) SetFadePointRamp(dspclock uint64, volume float64) error {
	res := C.FMOD_Channel_SetFadePointRamp(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

func (c *Channel) RemoveFadePoints(dspclock_start, dspclock_end uint64) error {
	res := C.FMOD_Channel_RemoveFadePoints(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end))
	return errs[res]
}

func (c *Channel) FadePoints() (uint32, uint64, float64, error) {
	var numpoints C.uint
	var point_dspclock C.ulonglong
	var point_volume C.float
	res := C.FMOD_Channel_GetFadePoints(c.cptr, &numpoints, &point_dspclock, &point_volume)
	return uint32(numpoints), uint64(point_dspclock), float64(point_dspclock), errs[res]
}

/*
   DSP effects.
*/

func (c *Channel) DSP(index int) (DSP, error) {
	var dsp DSP
	res := C.FMOD_Channel_GetDSP(c.cptr, C.int(index), &dsp.cptr)
	return dsp, errs[res]
}

func (c *Channel) AddDSP(index int, dsp DSP) error {
	res := C.FMOD_Channel_AddDSP(c.cptr, C.int(index), dsp.cptr)
	return errs[res]
}

func (c *Channel) RemoveDSP(dsp DSP) error {
	res := C.FMOD_Channel_RemoveDSP(c.cptr, dsp.cptr)
	return errs[res]
}

func (c *Channel) NumDSPs() (int, error) {
	var numdsps C.int
	res := C.FMOD_Channel_GetNumDSPs(c.cptr, &numdsps)
	return int(numdsps), errs[res]
}

func (c *Channel) SetDSPIndex(dsp DSP, index int) error {
	res := C.FMOD_Channel_SetDSPIndex(c.cptr, dsp.cptr, C.int(index))
	return errs[res]
}

func (c *Channel) DSPIndex(dsp DSP) (int, error) {
	var index C.int
	res := C.FMOD_Channel_GetDSPIndex(c.cptr, dsp.cptr, &index)
	return int(index), errs[res]
}

// NOTE: Not implement yet
func (c *Channel) OverridePanDSP(pan *C.FMOD_DSP) error {
	//FMOD_RESULT F_API FMOD_Channel_OverridePanDSP           (FMOD_CHANNEL *channel, FMOD_DSP *pan);
	return ErrNoImpl
}

/*
   3D functionality.
*/

func (c *Channel) Set3DAttributes(pos, vel, alt_pan_pos Vector) error {
	cpos := pos.toC()
	cvel := vel.toC()
	calt_pan_pos := alt_pan_pos.toC()
	res := C.FMOD_Channel_Set3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	return errs[res]
}

func (c *Channel) Get3DAttributes() (Vector, Vector, Vector, error) {
	var pos, vel, alt_pan_pos Vector
	var cpos, cvel, calt_pan_pos C.FMOD_VECTOR
	res := C.FMOD_Channel_Get3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	pos.fromC(cpos)
	vel.fromC(cvel)
	alt_pan_pos.fromC(calt_pan_pos)
	return pos, vel, alt_pan_pos, errs[res]
}

func (c *Channel) Set3DMinMaxDistance(mindistance, maxdistance float64) error {
	res := C.FMOD_Channel_Set3DMinMaxDistance(c.cptr, C.float(mindistance), C.float(maxdistance))
	return errs[res]
}

func (c *Channel) Get3DMinMaxDistance() (float64, float64, error) {
	var mindistance, maxdistance C.float
	res := C.FMOD_Channel_Get3DMinMaxDistance(c.cptr, &mindistance, &maxdistance)
	return float64(mindistance), float64(maxdistance), errs[res]
}

func (c *Channel) Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error {
	res := C.FMOD_Channel_Set3DConeSettings(c.cptr, C.float(insideconeangle), C.float(outsideconeangle), C.float(outsidevolume))
	return errs[res]
}

func (c *Channel) Get3DConeSettings() (float64, float64, float64, error) {
	var insideconeangle, outsideconeangle, outsidevolume C.float
	res := C.FMOD_Channel_Get3DConeSettings(c.cptr, &insideconeangle, &outsideconeangle, &outsidevolume)
	return float64(insideconeangle), float64(outsideconeangle), float64(outsidevolume), errs[res]
}

func (c *Channel) Set3DConeOrientation(orientation Vector) error {
	corientation := orientation.toC()
	res := C.FMOD_Channel_Set3DConeOrientation(c.cptr, &corientation)
	return errs[res]
}

func (c *Channel) Get3DConeOrientation() (Vector, error) {
	var corientation C.FMOD_VECTOR
	res := C.FMOD_Channel_Get3DConeOrientation(c.cptr, &corientation)
	orientation := NewVector()
	orientation.fromC(corientation)
	return orientation, errs[res]
}

func (c *Channel) Set3DCustomRolloff(points Vector, numpoints int) error {
	cpoints := points.toC()
	res := C.FMOD_Channel_Set3DCustomRolloff(c.cptr, &cpoints, C.int(numpoints))
	return errs[res]
}

func (c *Channel) Get3DCustomRolloff() (Vector, int, error) {
	var cpoints *C.FMOD_VECTOR
	var numpoints C.int
	res := C.FMOD_Channel_Get3DCustomRolloff(c.cptr, &cpoints, &numpoints)
	points := NewVector()
	points.fromC(*cpoints)
	return points, int(numpoints), errs[res]
}

func (c *Channel) Set3DOcclusion(directocclusion, reverbocclusion float64) error {
	res := C.FMOD_Channel_Set3DOcclusion(c.cptr, C.float(directocclusion), C.float(reverbocclusion))
	return errs[res]
}

func (c *Channel) Get3DOcclusion() (float64, float64, error) {
	var directocclusion, reverbocclusion C.float
	res := C.FMOD_Channel_Get3DOcclusion(c.cptr, &directocclusion, &reverbocclusion)
	return float64(directocclusion), float64(reverbocclusion), errs[res]
}

func (c *Channel) Set3DSpread(angle float64) error {
	res := C.FMOD_Channel_Set3DSpread(c.cptr, C.float(angle))
	return errs[res]
}

func (c *Channel) Get3DSpread() (float64, error) {
	var angle C.float
	res := C.FMOD_Channel_Get3DSpread(c.cptr, &angle)
	return float64(angle), errs[res]
}

func (c *Channel) Set3DLevel(level float64) error {
	res := C.FMOD_Channel_Set3DLevel(c.cptr, C.float(level))
	return errs[res]
}

func (c *Channel) Get3DLevel() (float64, error) {
	var level C.float
	res := C.FMOD_Channel_Get3DLevel(c.cptr, &level)
	return float64(level), errs[res]
}

func (c *Channel) Set3DDopplerLevel(level float64) error {
	res := C.FMOD_Channel_Set3DDopplerLevel(c.cptr, C.float(level))
	return errs[res]
}

func (c *Channel) Get3DDopplerLevel() (float64, error) {
	var level C.float
	res := C.FMOD_Channel_Get3DDopplerLevel(c.cptr, &level)
	return float64(level), errs[res]
}

func (c *Channel) Set3DDistanceFilter(custom bool, customLevel, centerFreq float64) error {
	res := C.FMOD_Channel_Set3DDistanceFilter(c.cptr, getBool(custom), C.float(customLevel), C.float(centerFreq))
	return errs[res]
}

func (c *Channel) Get3DDistanceFilter() (bool, float64, float64, error) {
	var custom C.FMOD_BOOL
	var customLevel, centerFreq C.float
	res := C.FMOD_Channel_Get3DDistanceFilter(c.cptr, &custom, &customLevel, &centerFreq)
	return setBool(custom), float64(customLevel), float64(centerFreq), errs[res]
}

/*
   Userdata set/get.
*/

// NOTE: Not implement yet
func (c *Channel) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_Channel_SetUserData              (FMOD_CHANNEL *channel, void *userdata);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (c *Channel) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_Channel_GetUserData              (FMOD_CHANNEL *channel, void **userdata);
	return ErrNoImpl
}

/*
   Channel specific control functionality.
*/

func (c *Channel) SetFrequency(frequency float64) error {
	res := C.FMOD_Channel_SetFrequency(c.cptr, C.float(frequency))
	return errs[res]
}

func (c *Channel) Frequency() (float64, error) {
	var frequency C.float
	res := C.FMOD_Channel_GetFrequency(c.cptr, &frequency)
	return float64(frequency), errs[res]
}

func (c *Channel) SetPriority(priority int) error {
	res := C.FMOD_Channel_SetPriority(c.cptr, C.int(priority))
	return errs[res]
}

func (c *Channel) Priority() (int, error) {
	var priority C.int
	res := C.FMOD_Channel_GetPriority(c.cptr, &priority)
	return int(priority), errs[res]
}

func (c *Channel) SetPosition(position uint32, postype TimeUnit) error {
	res := C.FMOD_Channel_SetPosition(c.cptr, C.uint(position), C.FMOD_TIMEUNIT(postype))
	return errs[res]
}

func (c *Channel) Position(postype TimeUnit) (uint32, error) {
	var position C.uint
	res := C.FMOD_Channel_GetPosition(c.cptr, &position, C.FMOD_TIMEUNIT(postype))
	return uint32(position), errs[res]
}

func (c *Channel) SetChannelGroup(channelgroup ChannelGroup) error {
	res := C.FMOD_Channel_SetChannelGroup(c.cptr, channelgroup.cptr)
	return errs[res]
}

func (c *Channel) ChannelGroup() (ChannelGroup, error) {
	var channelgroup ChannelGroup
	res := C.FMOD_Channel_GetChannelGroup(c.cptr, &channelgroup.cptr)
	return channelgroup, errs[res]
}

func (c *Channel) SetLoopCount(loopcount int) error {
	res := C.FMOD_Channel_SetLoopCount(c.cptr, C.int(loopcount))
	return errs[res]
}

func (c *Channel) LoopCount() (int, error) {
	var loopcount C.int
	res := C.FMOD_Channel_GetLoopCount(c.cptr, &loopcount)
	return int(loopcount), errs[res]
}

func (c *Channel) SetLoopPoints(loopstart uint32, loopstarttype TimeUnit, loopend uint32, loopendtype TimeUnit) error {
	res := C.FMOD_Channel_SetLoopPoints(c.cptr, C.uint(loopstart), C.FMOD_TIMEUNIT(loopstarttype), C.uint(loopend), C.FMOD_TIMEUNIT(loopendtype))
	return errs[res]
}

func (c *Channel) LoopPoints(loopstarttype, loopendtype TimeUnit) (uint32, uint32, error) {
	var loopstart, loopend C.uint
	res := C.FMOD_Channel_GetLoopPoints(c.cptr, &loopstart, C.FMOD_TIMEUNIT(loopstarttype), &loopend, C.FMOD_TIMEUNIT(loopendtype))
	return uint32(loopstart), uint32(loopend), errs[res]
}

/*
   Information only functions.
*/

func (c *Channel) IsVirtual() (bool, error) {
	var isvirtual C.FMOD_BOOL
	res := C.FMOD_Channel_IsVirtual(c.cptr, &isvirtual)
	return setBool(isvirtual), errs[res]
}

func (c *Channel) CurrentSound() (Sound, error) {
	var sound Sound
	res := C.FMOD_Channel_GetCurrentSound(c.cptr, &sound.cptr)
	return sound, errs[res]
}

func (c *Channel) Index() (int, error) {
	var index C.int
	res := C.FMOD_Channel_GetIndex(c.cptr, &index)
	return int(index), errs[res]
}
