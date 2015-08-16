package lowlevel

/*
#include <fmod.h>
*/
import "C"

type ChannelGroup struct {
	cptr *C.FMOD_CHANNELGROUP
}

/*
   'ChannelGroup' API
*/

func (c *ChannelGroup) GetSystemObject() (*System, error) {
	var system System
	res := C.FMOD_ChannelGroup_GetSystemObject(c.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   General control functionality for Channels and ChannelGroups.
*/

func (c *ChannelGroup) Stop() error {
	res := C.FMOD_ChannelGroup_Stop(c.cptr)
	return errs[res]
}

func (c *ChannelGroup) SetPaused(paused bool) error {
	res := C.FMOD_ChannelGroup_SetPaused(c.cptr, getBool(paused))
	return errs[res]
}

func (c *ChannelGroup) GetPaused() (bool, error) {
	var paused C.FMOD_BOOL
	res := C.FMOD_ChannelGroup_GetPaused(c.cptr, &paused)
	return setBool(paused), errs[res]
}

func (c *ChannelGroup) SetVolume(volume float64) error {
	res := C.FMOD_ChannelGroup_SetVolume(c.cptr, C.float(volume))
	return errs[res]
}

func (c *ChannelGroup) GetVolume() (float64, error) {
	var volume C.float
	res := C.FMOD_ChannelGroup_GetVolume(c.cptr, &volume)
	return float64(volume), errs[res]
}

func (c *ChannelGroup) SetVolumeRamp(ramp bool) error {
	res := C.FMOD_ChannelGroup_SetVolumeRamp(c.cptr, getBool(ramp))
	return errs[res]
}

func (c *ChannelGroup) GetVolumeRamp() (bool, error) {
	var ramp C.FMOD_BOOL
	res := C.FMOD_ChannelGroup_GetVolumeRamp(c.cptr, &ramp)
	return setBool(ramp), errs[res]
}

func (c *ChannelGroup) GetAudibility() (float64, error) {
	var audibility C.float
	res := C.FMOD_ChannelGroup_GetAudibility(c.cptr, &audibility)
	return float64(audibility), errs[res]
}

func (c *ChannelGroup) SetPitch(pitch float64) error {
	res := C.FMOD_ChannelGroup_SetPitch(c.cptr, C.float(pitch))
	return errs[res]
}

func (c *ChannelGroup) GetPitch() (float64, error) {
	var pitch C.float
	res := C.FMOD_ChannelGroup_GetPitch(c.cptr, &pitch)
	return float64(pitch), errs[res]
}

func (c *ChannelGroup) SetMute(mute bool) error {
	res := C.FMOD_ChannelGroup_SetMute(c.cptr, getBool(mute))
	return errs[res]
}

func (c *ChannelGroup) GetMute() (bool, error) {
	var mute C.FMOD_BOOL
	res := C.FMOD_ChannelGroup_GetMute(c.cptr, &mute)
	return setBool(mute), errs[res]
}

func (c *ChannelGroup) SetReverbProperties(instance int, wet float64) error {
	res := C.FMOD_ChannelGroup_SetReverbProperties(c.cptr, C.int(instance), C.float(wet))
	return errs[res]
}

func (c *ChannelGroup) GetReverbProperties(instance int) (float64, error) {
	var wet C.float
	res := C.FMOD_ChannelGroup_GetReverbProperties(c.cptr, C.int(instance), &wet)
	return float64(wet), errs[res]
}

func (c *ChannelGroup) SetLowPassGain(gain float64) error {
	res := C.FMOD_ChannelGroup_SetLowPassGain(c.cptr, C.float(gain))
	return errs[res]
}

func (c *ChannelGroup) GetLowPassGain() (float64, error) {
	var gain C.float
	res := C.FMOD_ChannelGroup_GetLowPassGain(c.cptr, &gain)
	return float64(gain), errs[res]
}

func (c *ChannelGroup) SetMode(mode Mode) error {
	res := C.FMOD_ChannelGroup_SetMode(c.cptr, C.FMOD_MODE(mode))
	return errs[res]
}

func (c *ChannelGroup) GetMode() (Mode, error) {
	var mode C.FMOD_MODE
	res := C.FMOD_ChannelGroup_GetMode(c.cptr, &mode)
	return Mode(mode), errs[res]
}

func (c *ChannelGroup) SetCallback(callback C.FMOD_CHANNELCONTROL_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_SetCallback         (FMOD_CHANNELGROUP *channelgroup, FMOD_CHANNELCONTROL_CALLBACK callback);
	return ErrNoImpl
}

func (c *ChannelGroup) IsPlaying() (bool, error) {
	var isplaying C.FMOD_BOOL
	res := C.FMOD_ChannelGroup_IsPlaying(c.cptr, &isplaying)
	return setBool(isplaying), errs[res]
}

/*
   Note all 'set' functions alter a final matrix, this is why the only get function is getMixMatrix, to avoid other get functions returning incorrect/obsolete values.
*/

func (c *ChannelGroup) SetPan(pan float64) error {
	res := C.FMOD_ChannelGroup_SetPan(c.cptr, C.float(pan))
	return errs[res]
}

func (c *ChannelGroup) SetMixLevelsOutput(frontleft, frontright, center, lfe, surroundleft, surroundright, backleft, backright float64) error {
	res := C.FMOD_ChannelGroup_SetMixLevelsOutput(c.cptr, C.float(frontleft), C.float(frontright), C.float(center), C.float(lfe), C.float(surroundleft), C.float(surroundright), C.float(backleft), C.float(backright))
	return errs[res]
}

func (c *ChannelGroup) SetMixLevelsInput(levels *C.float, numlevels C.int) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_SetMixLevelsInput   (FMOD_CHANNELGROUP *channelgroup, float *levels, int numlevels);
	return ErrNoImpl
}

func (c *ChannelGroup) SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_SetMixMatrix        (FMOD_CHANNELGROUP *channelgroup, float *matrix, int outchannels, int inchannels, int inchannel_hop);
	return ErrNoImpl
}

func (c *ChannelGroup) GetMixMatrix(matrix *C.float, outchannels, inchannels *C.int, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetMixMatrix        (FMOD_CHANNELGROUP *channelgroup, float *matrix, int *outchannels, int *inchannels, int inchannel_hop);
	return ErrNoImpl
}

/*
   Clock based functionality.
*/

func (c *ChannelGroup) GetDSPClock() (uint64, uint64, error) {
	var dspclock, parentclock C.ulonglong
	res := C.FMOD_ChannelGroup_GetDSPClock(c.cptr, &dspclock, &parentclock)
	return uint64(dspclock), uint64(parentclock), errs[res]
}

func (c *ChannelGroup) SetDelay(dspclock_start, dspclock_end uint64, stopchannels bool) error {
	res := C.FMOD_ChannelGroup_SetDelay(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end), getBool(stopchannels))
	return errs[res]
}

func (c *ChannelGroup) GetDelay() (uint64, uint64, bool, error) {
	var dspclock_start, dspclock_end C.ulonglong
	var stopchannels C.FMOD_BOOL
	res := C.FMOD_ChannelGroup_GetDelay(c.cptr, &dspclock_start, &dspclock_end, &stopchannels)
	return uint64(dspclock_start), uint64(dspclock_end), setBool(stopchannels), errs[res]
}

func (c *ChannelGroup) AddFadePoint(dspclock uint64, volume float64) error {
	res := C.FMOD_ChannelGroup_AddFadePoint(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

func (c *ChannelGroup) SetFadePointRamp(dspclock uint64, volume float64) error {
	res := C.FMOD_ChannelGroup_SetFadePointRamp(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

func (c *ChannelGroup) RemoveFadePoints(dspclock_start, dspclock_end uint64) error {
	res := C.FMOD_ChannelGroup_RemoveFadePoints(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end))
	return errs[res]
}

func (c *ChannelGroup) GetFadePoints() (uint32, uint64, float64, error) {
	var numpoints C.uint
	var point_dspclock C.ulonglong
	var point_volume C.float
	res := C.FMOD_ChannelGroup_GetFadePoints(c.cptr, &numpoints, &point_dspclock, &point_volume)
	return uint32(numpoints), uint64(point_dspclock), float64(point_dspclock), errs[res]
}

/*
   DSP effects.
*/

func (c *ChannelGroup) GetDSP(index int) (DSP, error) {
	var dsp DSP
	res := C.FMOD_ChannelGroup_GetDSP(c.cptr, C.int(index), &dsp.cptr)
	return dsp, errs[res]
}

func (c *ChannelGroup) AddDSP(index int, dsp DSP) error {
	res := C.FMOD_ChannelGroup_AddDSP(c.cptr, C.int(index), dsp.cptr)
	return errs[res]
}

func (c *ChannelGroup) RemoveDSP(dsp DSP) error {
	res := C.FMOD_ChannelGroup_RemoveDSP(c.cptr, dsp.cptr)
	return errs[res]
}

func (c *ChannelGroup) GetNumDSPs() (int, error) {
	var numdsps C.int
	res := C.FMOD_ChannelGroup_GetNumDSPs(c.cptr, &numdsps)
	return int(numdsps), errs[res]
}

func (c *ChannelGroup) SetDSPIndex(dsp DSP, index int) error {
	res := C.FMOD_ChannelGroup_SetDSPIndex(c.cptr, dsp.cptr, C.int(index))
	return errs[res]
}

func (c *ChannelGroup) GetDSPIndex(dsp DSP) (int, error) {
	var index C.int
	res := C.FMOD_ChannelGroup_GetDSPIndex(c.cptr, dsp.cptr, &index)
	return int(index), errs[res]
}

func (c *ChannelGroup) OverridePanDSP(pan *C.FMOD_DSP) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_OverridePanDSP      (FMOD_CHANNELGROUP *channelgroup, FMOD_DSP *pan);
	return ErrNoImpl
}

/*
   3D functionality.
*/

func (c *ChannelGroup) Set3DAttributes(pos, vel, alt_pan_pos Vector) error {
	cpos := pos.toC()
	cvel := vel.toC()
	calt_pan_pos := alt_pan_pos.toC()
	res := C.FMOD_ChannelGroup_Set3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	return errs[res]
}

func (c *ChannelGroup) Get3DAttributes() (Vector, Vector, Vector, error) {
	var pos, vel, alt_pan_pos Vector
	var cpos, cvel, calt_pan_pos C.FMOD_VECTOR
	res := C.FMOD_ChannelGroup_Get3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	pos.fromC(cpos)
	vel.fromC(cvel)
	alt_pan_pos.fromC(calt_pan_pos)
	return pos, vel, alt_pan_pos, errs[res]
}

func (c *ChannelGroup) Set3DMinMaxDistance(mindistance, maxdistance float64) error {
	res := C.FMOD_ChannelGroup_Set3DMinMaxDistance(c.cptr, C.float(mindistance), C.float(maxdistance))
	return errs[res]
}

func (c *ChannelGroup) Get3DMinMaxDistance() (float64, float64, error) {
	var mindistance, maxdistance C.float
	res := C.FMOD_ChannelGroup_Get3DMinMaxDistance(c.cptr, &mindistance, &maxdistance)
	return float64(mindistance), float64(maxdistance), errs[res]
}

func (c *ChannelGroup) Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error {
	res := C.FMOD_ChannelGroup_Set3DConeSettings(c.cptr, C.float(insideconeangle), C.float(outsideconeangle), C.float(outsidevolume))
	return errs[res]
}

func (c *ChannelGroup) Get3DConeSettings() (float64, float64, float64, error) {
	var insideconeangle, outsideconeangle, outsidevolume C.float
	res := C.FMOD_ChannelGroup_Get3DConeSettings(c.cptr, &insideconeangle, &outsideconeangle, &outsidevolume)
	return float64(insideconeangle), float64(outsideconeangle), float64(outsidevolume), errs[res]
}

func (c *ChannelGroup) Set3DConeOrientation(orientation Vector) error {
	corientation := orientation.toC()
	res := C.FMOD_ChannelGroup_Set3DConeOrientation(c.cptr, &corientation)
	return errs[res]
}

func (c *ChannelGroup) Get3DConeOrientation() (Vector, error) {
	var corientation C.FMOD_VECTOR
	res := C.FMOD_ChannelGroup_Get3DConeOrientation(c.cptr, &corientation)
	orientation := NewVector()
	orientation.fromC(corientation)
	return orientation, errs[res]
}

func (c *ChannelGroup) Set3DCustomRolloff(points Vector, numpoints int) error {
	cpoints := points.toC()
	res := C.FMOD_ChannelGroup_Set3DCustomRolloff(c.cptr, &cpoints, C.int(numpoints))
	return errs[res]
}

func (c *ChannelGroup) Get3DCustomRolloff() (Vector, int, error) {
	var cpoints *C.FMOD_VECTOR
	var numpoints C.int
	res := C.FMOD_ChannelGroup_Get3DCustomRolloff(c.cptr, &cpoints, &numpoints)
	points := NewVector()
	points.fromC(*cpoints)
	return points, int(numpoints), errs[res]
}

func (c *ChannelGroup) Set3DOcclusion(directocclusion, reverbocclusion float64) error {
	res := C.FMOD_ChannelGroup_Set3DOcclusion(c.cptr, C.float(directocclusion), C.float(reverbocclusion))
	return errs[res]
}

func (c *ChannelGroup) Get3DOcclusion() (float64, float64, error) {
	var directocclusion, reverbocclusion C.float
	res := C.FMOD_ChannelGroup_Get3DOcclusion(c.cptr, &directocclusion, &reverbocclusion)
	return float64(directocclusion), float64(reverbocclusion), errs[res]
}

func (c *ChannelGroup) Set3DSpread(angle float64) error {
	res := C.FMOD_ChannelGroup_Set3DSpread(c.cptr, C.float(angle))
	return errs[res]
}

func (c *ChannelGroup) Get3DSpread() (float64, error) {
	var angle C.float
	res := C.FMOD_ChannelGroup_Get3DSpread(c.cptr, &angle)
	return float64(angle), errs[res]
}

func (c *ChannelGroup) Set3DLevel(level float64) error {
	res := C.FMOD_ChannelGroup_Set3DLevel(c.cptr, C.float(level))
	return errs[res]
}

func (c *ChannelGroup) Get3DLevel() (float64, error) {
	var level C.float
	res := C.FMOD_ChannelGroup_Get3DLevel(c.cptr, &level)
	return float64(level), errs[res]
}

func (c *ChannelGroup) Set3DDopplerLevel(level float64) error {
	res := C.FMOD_ChannelGroup_Set3DDopplerLevel(c.cptr, C.float(level))
	return errs[res]
}

func (c *ChannelGroup) Get3DDopplerLevel() (float64, error) {
	var level C.float
	res := C.FMOD_ChannelGroup_Get3DDopplerLevel(c.cptr, &level)
	return float64(level), errs[res]
}

func (c *ChannelGroup) Set3DDistanceFilter(custom bool, customLevel, centerFreq float64) error {
	res := C.FMOD_ChannelGroup_Set3DDistanceFilter(c.cptr, getBool(custom), C.float(customLevel), C.float(centerFreq))
	return errs[res]
}

func (c *ChannelGroup) Get3DDistanceFilter() (bool, float64, float64, error) {
	var custom C.FMOD_BOOL
	var customLevel, centerFreq C.float
	res := C.FMOD_ChannelGroup_Get3DDistanceFilter(c.cptr, &custom, &customLevel, &centerFreq)
	return setBool(custom), float64(customLevel), float64(centerFreq), errs[res]
}

/*
   Userdata set/get.
*/

func (c *ChannelGroup) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_SetUserData(FMOD_CHANNELGROUP *channelgroup, void *userdata);
	return ErrNoImpl
}

func (c *ChannelGroup) GetUserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetUserData(FMOD_CHANNELGROUP *channelgroup, void **userdata);
	return ErrNoImpl
}

func (c *ChannelGroup) Release() error {
	res := C.FMOD_ChannelGroup_Release(c.cptr)
	return errs[res]
}

/*
   Nested channel groups.
*/

func (c *ChannelGroup) AddGroup(group ChannelGroup, propagatedspclock bool) (DspConnection, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_AddGroup(FMOD_CHANNELGROUP *channelgroup, FMOD_CHANNELGROUP *group, FMOD_BOOL propagatedspclock, FMOD_DSPCONNECTION **connection);
	var connection DspConnection
	res := C.FMOD_ChannelGroup_AddGroup(c.cptr, group.cptr, getBool(propagatedspclock), &connection.cptr)
	return connection, errs[res]
}

func (c *ChannelGroup) GetNumGroups() (int, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetNumGroups(FMOD_CHANNELGROUP *channelgroup, int *numgroups);
	var numgroups C.int
	res := C.FMOD_ChannelGroup_GetNumGroups(c.cptr, &numgroups)
	return int(numgroups), errs[res]
}

func (c *ChannelGroup) GetGroup(index int) (ChannelGroup, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetGroup(FMOD_CHANNELGROUP *channelgroup, int index, FMOD_CHANNELGROUP **group);
	var group ChannelGroup
	res := C.FMOD_ChannelGroup_GetGroup(c.cptr, C.int(index), &group.cptr)
	return group, errs[res]
}

func (c *ChannelGroup) GetParentGroup() (ChannelGroup, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetParentGroup(FMOD_CHANNELGROUP *channelgroup, FMOD_CHANNELGROUP **group);
	var group ChannelGroup
	res := C.FMOD_ChannelGroup_GetParentGroup(c.cptr, &group.cptr)
	return group, errs[res]
}

/*
   Information only functions.
*/

func (c *ChannelGroup) GetName(name *C.char, namelen C.int) error {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetName(FMOD_CHANNELGROUP *channelgroup, char *name, int namelen);
	return ErrNoImpl

}

func (c *ChannelGroup) GetNumChannels() (int, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetNumChannels(FMOD_CHANNELGROUP *channelgroup, int *numchannels);
	var numchannels C.int
	res := C.FMOD_ChannelGroup_GetNumChannels(c.cptr, &numchannels)
	return int(numchannels), errs[res]
}

func (c *ChannelGroup) GetChannel(index int) (Channel, error) {
	//FMOD_RESULT F_API FMOD_ChannelGroup_GetChannel(FMOD_CHANNELGROUP *channelgroup, int index, FMOD_CHANNEL **channel);
	var channel Channel
	res := C.FMOD_ChannelGroup_GetChannel(c.cptr, C.int(index), &channel.cptr)
	return channel, errs[res]
}
