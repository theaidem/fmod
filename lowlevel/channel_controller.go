package lowlevel

/*
#include <fmod.h>
*/
import "C"

type ChannelController interface {
	SystemObject() (*System, error)
	Stop() error
	Audibility() (float64, error)
	SetCallback(callback C.FMOD_CHANNELCONTROL_CALLBACK) error
	IsPlaying() (bool, error)
	SetPan(pan float64) error
	SetMixLevelsOutput(frontleft, frontright, center, lfe, surroundleft, surroundright, backleft, backright float64) error
	SetMixLevelsInput(levels *C.float, numlevels C.int) error
	SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error
	MixMatrix(matrix *C.float, outchannels, inchannels *C.int, inchannel_hop C.int) error

	ThreeDimensionalAccessor
	DSPAccessor
	FadePointAccessor
	DelayAccessor
	LPGainAccessor
	ReverbPropsAccessor
	MuteAccessor
	PitchAccessor
	VolumeAccessor
	PauseAccessor
	ModeAccessor
	UserDataAccessor
}

type ThreeDimensionalAccessor interface {
	Set3DAttributes(pos, vel, alt_pan_pos Vector) error
	Get3DAttributes() (Vector, Vector, Vector, error)
	Set3DMinMaxDistance(mindistance, maxdistance float64) error
	Get3DMinMaxDistance() (float64, float64, error)
	Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error
	Get3DConeSettings() (float64, float64, float64, error)
	Set3DConeOrientation(orientation Vector) error
	Get3DConeOrientation() (Vector, error)
	Set3DCustomRolloff(points Vector, numpoints int) error
	Get3DCustomRolloff() (Vector, int, error)
	Set3DOcclusion(directocclusion, reverbocclusion float64) error
	Get3DOcclusion() (float64, float64, error)
	Set3DSpread(angle float64) error
	Get3DSpread() (float64, error)
	Set3DLevel(level float64) error
	Get3DLevel() (float64, error)
	Set3DDopplerLevel(level float64) error
	Get3DDopplerLevel() (float64, error)
	Set3DDistanceFilter(custom bool, customLevel, centerFreq float64) error
	Get3DDistanceFilter() (bool, float64, float64, error)
}

type DSPAccessor interface {
	DSPClock() (uint64, uint64, error)
	DSP(index int) (DSP, error)
	AddDSP(index int, dsp DSP) error
	RemoveDSP(dsp DSP) error
	NumDSPs() (int, error)
	SetDSPIndex(dsp DSP, index int) error
	DSPIndex(dsp DSP) (int, error)
	OverridePanDSP(pan *C.FMOD_DSP) error
}

type FadePointAccessor interface {
	AddFadePoint(dspclock uint64, volume float64) error
	SetFadePointRamp(dspclock uint64, volume float64) error
	RemoveFadePoints(dspclock_start, dspclock_end uint64) error
	FadePoints() (uint32, uint64, float64, error)
}

type DelayAccessor interface {
	SetDelay(dspclock_start, dspclock_end uint64, stopchannels bool) error
	Delay() (uint64, uint64, bool, error)
}

type LPGainAccessor interface {
	SetLowPassGain(gain float64) error
	LowPassGain() (float64, error)
}

type ReverbPropsAccessor interface {
	SetReverbProperties(instance int, wet float64) error
	ReverbProperties(instance int) (float64, error)
}

type MuteAccessor interface {
	SetMute(mute bool) error
	Mute() (bool, error)
}

type PitchAccessor interface {
	SetPitch(pitch float64) error
	Pitch() (float64, error)
}

type VolumeAccessor interface {
	SetVolume(volume float64) error
	Volume() (float64, error)
	SetVolumeRamp(ramp bool) error
	VolumeRamp() (bool, error)
}

type PauseAccessor interface {
	SetPaused(paused bool) error
	IsPaused() (bool, error)
}

type ModeAccessor interface {
	SetMode(mode Mode) error
	Mode() (Mode, error)
}

type UserDataAccessor interface {
	SetUserData(userdata interface{}) error
	UserData() (interface{}, error)
}
