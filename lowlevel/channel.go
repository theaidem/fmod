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

// Retrieves the parent System object that created the channel or channel group.
func (c *Channel) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_Channel_GetSystemObject(c.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   General control functionality for Channels and ChannelGroups.
*/

// Stops the channel (or all channels in the channel group) from playing. Makes it available for re-use by the priority system.
func (c *Channel) Stop() error {
	res := C.FMOD_Channel_Stop(c.cptr)
	return errs[res]
}

// Sets the paused state.
// Each channel and channel group has its own paused state, pausing a channel group will pause all contained channels but will not affect their individual setting.
func (c *Channel) SetPaused(paused bool) error {
	res := C.FMOD_Channel_SetPaused(c.cptr, getBool(paused))
	return errs[res]
}

// Retrieves the paused state.
func (c *Channel) IsPaused() (bool, error) {
	var paused C.FMOD_BOOL
	res := C.FMOD_Channel_GetPaused(c.cptr, &paused)
	return setBool(paused), errs[res]
}

// Sets the volume level linearly.
//
// volume: Linear volume level, default = 1.0.
//
// Volume level can be below 0 to invert a signal and above 1 to amplify the signal. Note that increasing the signal level too far may cause audible distortion.
// "Sound.SetDefaults" can be used to change the default volume for any channels played using that sound.
func (c *Channel) SetVolume(volume float64) error {
	res := C.FMOD_Channel_SetVolume(c.cptr, C.float(volume))
	return errs[res]
}

// Retrieves the volume level.
func (c *Channel) Volume() (float64, error) {
	var volume C.float
	res := C.FMOD_Channel_GetVolume(c.cptr, &volume)
	return float64(volume), errs[res]
}

// Sets whether the channel automatically ramps when setting volumes.
//
// ramp: Whether to enable volume ramping.
//
// When changing volumes on a non-paused channel, FMOD normally adds a small ramp to avoid a pop sound.
// This function allows that setting to be overriden and volume changes to be applied immediately.
func (c *Channel) SetVolumeRamp(ramp bool) error {
	res := C.FMOD_Channel_SetVolumeRamp(c.cptr, getBool(ramp))
	return errs[res]
}

// Retrieves whether volume ramp is enabled.
func (c *Channel) VolumeRamp() (bool, error) {
	var ramp C.FMOD_BOOL
	res := C.FMOD_Channel_GetVolumeRamp(c.cptr, &ramp)
	return setBool(ramp), errs[res]
}

// Retrieves the combined volume after 3D spatialization and geometry occlusion calculations including any volumes set via the API.
// This does not represent the waveform, just the calculated result of all volume modifiers. This value is used by the virtual channel system to order its channels between real and virtual.
func (c *Channel) Audibility() (float64, error) {
	var audibility C.float
	res := C.FMOD_Channel_GetAudibility(c.cptr, &audibility)
	return float64(audibility), errs[res]
}

// Sets the pitch value.
//
// pitch: Pitch value, 0.5 = half pitch, 2.0 = double pitch, etc default = 1.0.
//
// This function scales existing frequency values by the pitch.
func (c *Channel) SetPitch(pitch float64) error {
	res := C.FMOD_Channel_SetPitch(c.cptr, C.float(pitch))
	return errs[res]
}

// Retrieves the pitch value.
func (c *Channel) Pitch() (float64, error) {
	var pitch C.float
	res := C.FMOD_Channel_GetPitch(c.cptr, &pitch)
	return float64(pitch), errs[res]
}

// Sets the mute state effectively silencing it or returning it to its normal volume.
//
// mute: Mute state, true = mute (silent), false = normal volume.
//
// Each channel and channel group has its own mute state, muting a channel group will mute all child channels but will not affect their individual setting.
// Calling "Channel.Mute" will always return the value you set.
func (c *Channel) SetMute(mute bool) error {
	res := C.FMOD_Channel_SetMute(c.cptr, getBool(mute))
	return errs[res]
}

// Retrieves the mute state.
func (c *Channel) Mute() (bool, error) {
	var mute C.FMOD_BOOL
	res := C.FMOD_Channel_GetMute(c.cptr, &mute)
	return setBool(mute), errs[res]
}

// Sets the wet level (or send level) of a particular reverb instance for the Channel
//
// instance: Index of the particular reverb instance to target, from 0 to 3.
//
// wet: Send level for the signal to the reverb, from 0 (none) to 1.0 (full), default = 1.0 for Channels, 0.0 for ChannelGroups..
// A Channel is automatically connected to a reverb instance 0 if it exists. A ChannelGroup though is a dynamic object and will not send to a reverb by default.
// To make a ChannelGroup send to a reverb, setReverbProperties must be called.
// A ChannelGroup reverb is optimal for the case where you want to send 1 mixed signal to the reverb, rather than a lot of indivual channel reverb sends.
// It is advisable to do this to reduce CPU if you have many channels inside a channelgroup.
//
// With A ChannelGroup that has reverb activated on it, and if the ChannelGroup also has Channels that have sends to the same reverb,
// the behaviour will be that the reverb signal will be sent twice for the channel.
// It is better to disable the reverb for channels if you are sending reverb from the parent ChannelGroup.
func (c *Channel) SetReverbProperties(instance int, wet float64) error {
	res := C.FMOD_Channel_SetReverbProperties(c.cptr, C.int(instance), C.float(wet))
	return errs[res]
}

// Retrieves the wet level (or send level) for a particular reverb instance.
func (c *Channel) ReverbProperties(instance int) (float64, error) {
	var wet C.float
	res := C.FMOD_Channel_GetReverbProperties(c.cptr, C.int(instance), &wet)
	return float64(wet), errs[res]
}

// Sets the gain of the dry signal when lowpass filtering is applied.
//
// gain: Linear gain level, from 0 (silent, full filtering) to 1.0 (full volume, no filtering), default = 1.0.
//
// Requires the built in lowpass to be created with "INIT_CHANNEL_LOWPASS" or "INIT_CHANNEL_DISTANCEFILTER".
func (c *Channel) SetLowPassGain(gain float64) error {
	res := C.FMOD_Channel_SetLowPassGain(c.cptr, C.float(gain))
	return errs[res]
}

// Retrieves the gain of the dry signal when lowpass filtering is applied.
func (c *Channel) LowPassGain() (float64, error) {
	var gain C.float
	res := C.FMOD_Channel_GetLowPassGain(c.cptr, &gain)
	return float64(gain), errs[res]
}

// Changes some attributes for a channel based on the mode passed in.
//
// Issues with streamed audio:
//
// When changing the loop mode, sounds created with "System.CreateStream" or CREATESTREAM may have already been pre-buffered and executed their loop logic ahead of
// time before this call was even made. This is dependant on the size of the sound versus the size of the stream decode buffer (see CREATESOUNDEXINFO).
// If this happens, you may need to reflush the stream buffer by calling "Channel.SetPosition".
// Note this will usually only happen if you have sounds or loop points that are smaller than the stream decode buffer size.
//
// Issues with PCM samples:
//
// When changing the loop mode of sounds created with with "System.CreateSound" or CREATESAMPLE, if the sound was set up as LOOP_OFF,
// then set to LOOP_NORMAL with this function, the sound may click when playing the end of the sound.
// This is because the sound needs to be pre-prepared for looping using "Sound.SetMode", by modifying the content of the PCM data
// (i.e. data past the end of the actual sample data) to allow the interpolators to read ahead without clicking.
// If you use "Channel.SetMode" it will not do this (because different channels may have different loop modes for the same sound) and may click if you try to set it to looping
// on an unprepared sound. If you want to change the loop mode at runtime it may be better to load the sound as looping first (or use "Sound.SetMode"),
// to let it pre-prepare the data as if it was looping so that it does not click whenever "Channel.SetMode" is used to turn looping on.
//
// If 3D_IGNOREGEOMETRY or VIRTUAL_PLAYFROMSTART is not specified, the flag will be cleared if it was specified previously.
func (c *Channel) SetMode(mode Mode) error {
	res := C.FMOD_Channel_SetMode(c.cptr, C.FMOD_MODE(mode))
	return errs[res]
}

// Retrieves the mode bit flags for the channel.
func (c *Channel) Mode() (Mode, error) {
	var mode C.FMOD_MODE
	res := C.FMOD_Channel_GetMode(c.cptr, &mode)
	return Mode(mode), errs[res]
}

// NOTE: Not implement yet
// Sets a callback to perform action for a specific event.
// Currently callbacks are driven by "System.Update" and will only occur when this function is called. This has the main advantage of far less complication due to thread issues,
// and allows all FMOD commands, including loading sounds and playing new sounds from the callback.
// It also allows any type of sound to have an end callback, no matter what it is. The only disadvantage is that callbacks are not asynchronous and are bound by the latency caused by
// the rate the user calls the update command.
func (c *Channel) SetCallback(callback C.FMOD_CHANNELCONTROL_CALLBACK) error {
	//FMOD_RESULT F_API FMOD_Channel_SetCallback(FMOD_CHANNEL *channel, FMOD_CHANNELCONTROL_CALLBACK callback);
	return ErrNoImpl
}

// Retrieves the playing state.
func (c *Channel) IsPlaying() (bool, error) {
	var isplaying C.FMOD_BOOL
	res := C.FMOD_Channel_IsPlaying(c.cptr, &isplaying)
	return setBool(isplaying), errs[res]
}

/*
   Note all 'set' functions alter a final matrix, this is why the only get function is getMixMatrix, to avoid other get functions returning incorrect/obsolete values.
*/

// Sets the pan level, this is a helper to avoid calling "Channel.SetMixMatrix".
//
//pan: Pan level, from -1.0 (left) to 1.0 (right), default = 0 (center).
//
// Mono sounds are panned from left to right using constant power panning (non linear fade).
// This means when pan = 0.0, the balance for the sound in each speaker is 71% left and 71% right, not 50% left and 50% right. This gives (audibly) smoother pans.
//
// Stereo sounds heave each left/right value faded up and down according to the specified pan position.
// This means when pan = 0.0, the balance for the sound in each speaker is 100% left and 100% right. When pan = -1.0, only the left channel of the stereo sound is audible,
// when pan = 1.0, only the right channel of the stereo sound is audible.
//
// Panning does not work if the speaker mode is "SPEAKERMODE_RAW".
func (c *Channel) SetPan(pan float64) error {
	res := C.FMOD_Channel_SetPan(c.cptr, C.float(pan))
	return errs[res]
}

// Sets the speaker volume levels for each speaker individually, this is a helper to avoid calling "Channel.SetMixMatrix".
//
// frontleft: Volume level for the front left speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// frontright: Volume level for the front right speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// center:Volume level for the center speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// lfe: Volume level for the subwoofer speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// surroundleft: Volume level for the surround left speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// surroundright: Volume level for the surround right speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// backleft: Volume level for the back left speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// backright: Volume level for the back right speaker of a multichannel speaker setup, 0.0 (silent), 1.0 (normal volume).
//
// Levels can be below 0 to invert a signal and above 1 to amplify the signal. Note that increasing the signal level too far may cause audible distortion.
// Speakers specified that don't exist will simply be ignored.
// For more advanced speaker control, including sending the different channels of a stereo sound to arbitrary speakers, see Channel.SetMixMatrix.
func (c *Channel) SetMixLevelsOutput(frontleft, frontright, center, lfe, surroundleft, surroundright, backleft, backright float64) error {
	res := C.FMOD_Channel_SetMixLevelsOutput(c.cptr, C.float(frontleft), C.float(frontright), C.float(center), C.float(lfe), C.float(surroundleft), C.float(surroundright), C.float(backleft), C.float(backright))
	return errs[res]
}

// NOTE: Not implement yet
// Sets the incoming volume level for each channel, this is a helper to avoid calling "Channel.SetMixMatrix".
// This means if you have multichannel audio you can turn channels on and off, a mono signal has 1 input channel, stereo has 2, etc.
//
// levels: Array of volume levels for each incoming channel.
//
// numlevels: Number of levels in the array, from 0 to 32 inclusive.
//
// Levels can be below 0 to invert a signal and above 1 to amplify the signal. Note that increasing the signal level too far may cause audible distortion.
func (c *Channel) SetMixLevelsInput(levels *C.float, numlevels C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_SetMixLevelsInput(FMOD_CHANNEL *channel, float *levels, int numlevels);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Sets a 2D pan matrix that maps input channels (columns) to output speakers (rows).
//
// matrix: Address of a 2 dimensional array of volume levels in row-major order. Each row represents an output speaker, each column represents an input channel.
//
// outchannels: Number of output channels (rows) in the matrix being passed in, from 0 to MAX_CHANNEL_WIDTH inclusive.
//
// inchannels: Number of input channels (columns) in the matrix being passed in, from 0 to MAX_CHANNEL_WIDTH inclusive.
//
// matrixhop: The width (total number of columns) of the matrix. Optional. If this is 0, inchannels will be taken as the width of the matrix. Maximum of MAX_CHANNEL_WIDTH.
//
// The gain for input channel 's' to output channel 't' is matrix[t * matrixhop + s].
//
// Levels can be below 0 to invert a signal and above 1 to amplify the signal. Note that increasing the signal level too far may cause audible distortion.
// The matrix size will generally be the size of the number of channels in the current speaker mode. Use "System.SoftwareFormat" to determine this.
// If a matrix already exists then the matrix passed in will applied over the top of it. The input matrix can be smaller than the existing matrix.
func (c *Channel) SetMixMatrix(matrix *C.float, outchannels, inchannels, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_SetMixMatrix(FMOD_CHANNEL *channel, float *matrix, int outchannels, int inchannels, int inchannel_hop);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a 2D pan matrix that maps input channels (columns) to output speakers (rows).
//
// The gain for input channel 's' to output channel 't' is matrix[t * matrixhop + s].
//
// Levels can be below 0 to invert a signal and above 1 to amplify the signal. Note that increasing the signal level too far may cause audible distortion.
// The matrix size will generally be the size of the number of channels in the current speaker mode. Use "System.SoftwareFormat" to determine this.
// Passing NULL for 'matrix' will allow you to query 'outchannels' and 'inchannels' without copying any data.
func (c *Channel) MixMatrix(matrix *C.float, outchannels, inchannels *C.int, inchannel_hop C.int) error {
	//FMOD_RESULT F_API FMOD_Channel_GetMixMatrix(FMOD_CHANNEL *channel, float *matrix, int *outchannels, int *inchannels, int inchannel_hop);
	return ErrNoImpl
}

/*
   Clock based functionality.
*/

// Retrieves the DSP clock values which count up by the number of samples per second in the software mixer,
// i.e. if the default sample rate is 48KHz, the DSP clock increments by 48000 per second.
//
// Use result with "Channel.SetDelay" to play a sound on an exact tick in the future, or stop it in the future.
// Note that when delaying a channel or channel group you want to sync it to the parent channel group DSP clock value, not its own DSP clock value.
func (c *Channel) DSPClock() (uint64, uint64, error) {
	var dspclock, parentclock C.ulonglong
	res := C.FMOD_Channel_GetDSPClock(c.cptr, &dspclock, &parentclock)
	return uint64(dspclock), uint64(parentclock), errs[res]
}

// Sets a start (and/or stop) time relative to the parent channel group DSP clock, with sample accuracy.
//
// dspclock_start: DSP clock of the parent channel group to audibly start playing sound at, a value of 0 indicates no delay.
//
// dspclock_end: DSP clock of the parent channel group to audibly stop playing sound at, a value of 0 indicates no delay.
//
// stopchannels: TRUE = stop according to "Channel.IsPlaying". FALSE = remain 'active' and a new start delay could start playback again at a later time.
//
// Every channel and channel group has its own DSP Clock. A channel or channel group can be delayed relatively against its parent, with sample accurate positioning.
// To delay a sound, use the 'parent' channel group DSP clock to reference against when passing values into this function.
// If a parent channel group changes its pitch, the start and stop times will still be correct as the parent clock is rate adjusted by that pitch.
func (c *Channel) SetDelay(dspclock_start, dspclock_end uint64, stopchannels bool) error {
	res := C.FMOD_Channel_SetDelay(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end), getBool(stopchannels))
	return errs[res]
}

// Retrieves a start (and/or stop) time relative to the parent channel group DSP clock, with sample accuracy.
func (c *Channel) Delay() (uint64, uint64, bool, error) {
	var dspclock_start, dspclock_end C.ulonglong
	var stopchannels C.FMOD_BOOL
	res := C.FMOD_Channel_GetDelay(c.cptr, &dspclock_start, &dspclock_end, &stopchannels)
	return uint64(dspclock_start), uint64(dspclock_end), setBool(stopchannels), errs[res]
}

// Add a volume point to fade from or towards, using a clock offset and 0 to 1 volume level.
//
// dspclock: DSP clock of the parent channel group to set the fade point volume.
//
// volume: Volume level where 0 is silent and 1.0 is normal volume. Amplification is supported.
//
// For every fade point, FMOD will do a per sample volume ramp between them. It will scale with the current Channel or ChannelGroup's volume.
func (c *Channel) AddFadePoint(dspclock uint64, volume float64) error {
	res := C.FMOD_Channel_AddFadePoint(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

// Add an volume ramp at the specified time using fade points.
//
// dspclock: DSP clock of the parent channel group when the volume will be ramped to.
//
// volume: Volume level where 0 is silent and 1.0 is normal volume. Amplification is supported.
//
// This is a helper function that automatically ramps from the current fade volume to the newly provided volume. It will clear any fade points set after this time.
func (c *Channel) SetFadePointRamp(dspclock uint64, volume float64) error {
	res := C.FMOD_Channel_SetFadePointRamp(c.cptr, C.ulonglong(dspclock), C.float(volume))
	return errs[res]
}

// Remove volume fade points on the timeline. This function will remove multiple fade points with a single call if the points lay between the 2 specified clock values (inclusive).
func (c *Channel) RemoveFadePoints(dspclock_start, dspclock_end uint64) error {
	res := C.FMOD_Channel_RemoveFadePoints(c.cptr, C.ulonglong(dspclock_start), C.ulonglong(dspclock_end))
	return errs[res]
}

// Retrieve information about fade points stored within a Channel.
//
// To first get the number of points for memory purposes, and not store any data, call this function with point_dpsclock and point_volume parameters being 0 or NULL.
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

// Retrieve the DSP unit at the specified index.
func (c *Channel) DSP(index int) (DSP, error) {
	var dsp DSP
	res := C.FMOD_Channel_GetDSP(c.cptr, C.int(index), &dsp.cptr)
	return dsp, errs[res]
}

// Add a pre-created DSP unit to the specified index in the DSP chain.
func (c *Channel) AddDSP(index int, dsp DSP) error {
	res := C.FMOD_Channel_AddDSP(c.cptr, C.int(index), dsp.cptr)
	return errs[res]
}

// Remove a particular DSP unit from the DSP chain.
func (c *Channel) RemoveDSP(dsp DSP) error {
	res := C.FMOD_Channel_RemoveDSP(c.cptr, dsp.cptr)
	return errs[res]
}

// Retrieves the number of DSP units in the DSP chain.
func (c *Channel) NumDSPs() (int, error) {
	var numdsps C.int
	res := C.FMOD_Channel_GetNumDSPs(c.cptr, &numdsps)
	return int(numdsps), errs[res]
}

// Moves the position in the DSP chain of a specified DSP unit.
//
// This function is useful for reordering DSP units inside a Channel or ChannelGroup so that processing can happen in the desired order.
// You can verify the order of the DSP chain using iteration via "Channel.NumDSPs" and "Channel.DSP" or with the FMOD Profiler tool.
func (c *Channel) SetDSPIndex(dsp DSP, index int) error {
	res := C.FMOD_Channel_SetDSPIndex(c.cptr, dsp.cptr, C.int(index))
	return errs[res]
}

// Retrieves the index in the DSP chain of the provided DSP.
func (c *Channel) DSPIndex(dsp DSP) (int, error) {
	var index C.int
	res := C.FMOD_Channel_GetDSPIndex(c.cptr, dsp.cptr, &index)
	return int(index), errs[res]
}

// NOTE: Not implement yet
// Replaces the built in panner unit FMOD uses per ChannelGroup and Channel, with a user selected panner.
// Can also be used to revert the panner back to the built in panner.
//
// panner: Pointer to a "DSP_TYPE_PANNER" dsp unit. Specify 0 or NULL to make FMOD use the built in panner.
//
// A panner is a unit that a channel uses to control the pan from setMixMatrix, setPan, and from parameters like 3D position.
// A panner unit must support the enumerations specified by the DSP_PAN enumeration list.
// Use "Channel.DSP" with "DSP_INDEX" enum "DSP_PANNER" to get access to the built in panner. (todo: fix me)
func (c *Channel) OverridePanDSP(pan *C.FMOD_DSP) error {
	//FMOD_RESULT F_API FMOD_Channel_OverridePanDSP           (FMOD_CHANNEL *channel, FMOD_DSP *pan);
	return ErrNoImpl
}

/*
   3D functionality.
*/

// Sets the position and velocity used to apply panning, attenuation and doppler.
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
// For a stereo 3D sound, you can set the spread of the left/right parts in speaker space by using "Channel.Set3DSpread".
func (c *Channel) Set3DAttributes(pos, vel, alt_pan_pos Vector) error {
	cpos := pos.toC()
	cvel := vel.toC()
	calt_pan_pos := alt_pan_pos.toC()
	res := C.FMOD_Channel_Set3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	return errs[res]
}

// Retrieves the position and velocity used to apply panning, attenuation and doppler.
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
func (c *Channel) Get3DAttributes() (Vector, Vector, Vector, error) {
	var pos, vel, alt_pan_pos Vector
	var cpos, cvel, calt_pan_pos C.FMOD_VECTOR
	res := C.FMOD_Channel_Get3DAttributes(c.cptr, &cpos, &cvel, &calt_pan_pos)
	pos.fromC(cpos)
	vel.fromC(cvel)
	alt_pan_pos.fromC(calt_pan_pos)
	return pos, vel, alt_pan_pos, errs[res]
}

// Sets the minimum and maximum audible distance.
//
// mindistance: Minimum volume distance in 'units', default = 1.0.
//
// maxdistance: Maximum volume distance in 'units', default = 10000.0.
//
//When the listener is in-between the minimum distance and the sound source the volume will be at its maximum.
// As the listener moves from the minimum distance to the maximum distance the sound will attenuate following the rolloff curve set.
// When outside the maximum distance the sound will no longer attenuate.
//
// Minimum distance is useful to give the impression that the sound is loud or soft in 3D space.
// An example of this is a small quiet object, such as a bumblebee, which you could set a small mindistance such as 0.1.
// This would cause it to attenuate quickly and dissapear when only a few meters away from the listener.
// Another example is a jumbo jet, which you could set to a mindistance of 100.0 causing the volume to stay at its loudest until the listener was 100 meters away,
// then it would be hundreds of meters more before it would fade out.
//
// Maximum distance is effectively obsolete unless you need the sound to stop fading out at a certain point.
// Do not adjust this from the default if you dont need to.
// Some people have the confusion that maxdistance is the point the sound will fade out to zero, this is not the case.
//
// A 'distance unit' is specified by System::set3DSettings. By default this is set to meters which is a distance scale of 1.0.
//
// To define the min and max distance per sound use "Sound.Set3DMinMaxDistance".
//
// If FMOD_3D_CUSTOMROLLOFF is used, then these values are stored, but ignored in 3D processing.
func (c *Channel) Set3DMinMaxDistance(mindistance, maxdistance float64) error {
	res := C.FMOD_Channel_Set3DMinMaxDistance(c.cptr, C.float(mindistance), C.float(maxdistance))
	return errs[res]
}

// Retrieves the minimum and maximum audible distance.
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
func (c *Channel) Get3DMinMaxDistance() (float64, float64, error) {
	var mindistance, maxdistance C.float
	res := C.FMOD_Channel_Get3DMinMaxDistance(c.cptr, &mindistance, &maxdistance)
	return float64(mindistance), float64(maxdistance), errs[res]
}

// Sets the angles that define the sound projection cone including the volume when outside the cone.
//
// insideconeangle: Inside cone angle, in degrees. This is the angle within which the sound is at its normal volume. Must not be greater than 'outsideconeangle'. Default = 360.
//
// outsideconeangle: Outside cone angle, in degrees. This is the angle outside of which the sound is at its outside volume. Must not be less than 'insideconeangle'. Default = 360.
//
// outsidevolume: Cone outside volume, from 0.0 to 1.0, default = 1.0.
//
// To define the parameters per sound use "Sound.Set3DConeSettings".
func (c *Channel) Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error {
	res := C.FMOD_Channel_Set3DConeSettings(c.cptr, C.float(insideconeangle), C.float(outsideconeangle), C.float(outsidevolume))
	return errs[res]
}

// Retrieves the angles that define the sound projection cone including the volume when outside the cone.
func (c *Channel) Get3DConeSettings() (float64, float64, float64, error) {
	var insideconeangle, outsideconeangle, outsidevolume C.float
	res := C.FMOD_Channel_Get3DConeSettings(c.cptr, &insideconeangle, &outsideconeangle, &outsidevolume)
	return float64(insideconeangle), float64(outsideconeangle), float64(outsidevolume), errs[res]
}

// Sets the orientation of the sound projection cone.
//
// orientation: Coordinates of the sound cone orientation vector, the vector information represents the center of the sound cone.
//
// This function has no effect unless the cone angle and cone outside volume have also been set to values other than the default.
func (c *Channel) Set3DConeOrientation(orientation Vector) error {
	corientation := orientation.toC()
	res := C.FMOD_Channel_Set3DConeOrientation(c.cptr, &corientation)
	return errs[res]
}

// Retrieves the orientation of the sound projection cone.
func (c *Channel) Get3DConeOrientation() (Vector, error) {
	var corientation C.FMOD_VECTOR
	res := C.FMOD_Channel_Get3DConeOrientation(c.cptr, &corientation)
	orientation := NewVector()
	orientation.fromC(corientation)
	return orientation, errs[res]
}

// TODO: add more docs
// Sets a custom rolloff curve to define how audio will attenuate over distance.
// Must be used in conjunction with FMOD_3D_CUSTOMROLLOFF flag to be activated.
//
// points: Array of "Vector" structures where x = distance and y = volume from 0.0 to 1.0. z should be set to 0.
//
// numpoints: Number of points in the array.
//
// Note! This function does not duplicate the memory for the points internally.
// The pointer you pass to FMOD must remain valid until there is no more use for it.
// Do not free the memory while in use, or use a local variable that goes out of scope while in use.
//
// Points must be sorted by distance! Passing an unsorted list to FMOD will result in an error.
//
// Set the points parameter to 0 or NULL to disable the points. If FMOD_3D_CUSTOMROLLOFF is set and the rolloff curve is 0, FMOD will revert to inverse curve rolloff.
//
// Values set with "Channel.SetMinMaxDistance" are meaningless when FMOD_3D_CUSTOMROLLOFF is used, their values are ignored.
//
// Distances between points are linearly interpolated.
// Note that after the highest distance specified, the volume in the last entry is used from that distance onwards.
// To define the parameters per sound use "Sound.Set3DCustomRolloff".
func (c *Channel) Set3DCustomRolloff(points Vector, numpoints int) error {
	cpoints := points.toC()
	res := C.FMOD_Channel_Set3DCustomRolloff(c.cptr, &cpoints, C.int(numpoints))
	return errs[res]
}

// Retrieves a pointer to the current custom rolloff curve.
func (c *Channel) Get3DCustomRolloff() (Vector, int, error) {
	var cpoints *C.FMOD_VECTOR
	var numpoints C.int
	res := C.FMOD_Channel_Get3DCustomRolloff(c.cptr, &cpoints, &numpoints)
	points := NewVector()
	points.fromC(*cpoints)
	return points, int(numpoints), errs[res]
}

// Sets the occlusion factors manually for when the FMOD geometry engine is not being used.
//
// directocclusion: Occlusion factor for the direct path, from 0.0 (not occluded) to 1.0 (fully occluded), default = 0.0.
//
// reverbocclusion: Occlusion factor for the reverb mix, from 0.0 (not occluded) to 1.0 (fully occluded), default = 0.0.
//
// Normally the volume is simply attenuated by the 'directocclusion' factor however if "INIT_CHANNEL_LOWPASS" is specified frequency filtering will be used with a very small CPU hit.
func (c *Channel) Set3DOcclusion(directocclusion, reverbocclusion float64) error {
	res := C.FMOD_Channel_Set3DOcclusion(c.cptr, C.float(directocclusion), C.float(reverbocclusion))
	return errs[res]
}

// Retrieves the occlusion factors.
func (c *Channel) Get3DOcclusion() (float64, float64, error) {
	var directocclusion, reverbocclusion C.float
	res := C.FMOD_Channel_Get3DOcclusion(c.cptr, &directocclusion, &reverbocclusion)
	return float64(directocclusion), float64(reverbocclusion), errs[res]
}

// Sets the spread of a 3D sound in speaker space.
//
// angle: Speaker spread angle. 0 = all sound channels are located at the same speaker location and is 'mono'.
// 360 = all sound channels are located at the opposite speaker location to the speaker location that it should be according to 3D position. Default = 0.
//
// Normally a 3D sound is aimed at one position in a speaker array depending on the 3D position to give it direction.
// Left and right parts of a stereo sound for example are consequently summed together and become 'mono'.
// When increasing the 'spread' of a sound, the left and right parts of a stereo sound rotate away from their original position, to give it more 'stereoness'.
// The rotation of the sound channels are done in 'speaker space'.
//
// Multichannel sounds with channel counts greater than stereo have their sub-channels spread evently through the specified angle.
// For example a 6 channel sound over a 90 degree spread has each channel located 15 degrees apart from each other in the speaker array.
//
// Mono sounds are spread as if they were a stereo signal, i.e. the signal is split into 2.
// The power will remain the same as it spreads around the speakers.
//
// To summarize (for a stereo sound).
//
// 1. A spread angle of 0 makes the stereo sound mono at the point of the 3D emitter.
//
// 2. A spread angle of 90 makes the left part of the stereo sound place itself at 45 degrees to the left and the right part 45 degrees to the right.
//
// 3. A spread angle of 180 makes the left part of the stero sound place itself at 90 degrees to the left and the right part 90 degrees to the right.
//
// 4. A spread angle of 360 makes the stereo sound mono at the opposite speaker location to where the 3D emitter should be located (by moving the left part 180 degrees left and
// the right part 180 degrees right). So in this case, behind you when the sound should be in front of you!
func (c *Channel) Set3DSpread(angle float64) error {
	res := C.FMOD_Channel_Set3DSpread(c.cptr, C.float(angle))
	return errs[res]
}

// Retrieves the spread of a 3D sound in speaker space.
func (c *Channel) Get3DSpread() (float64, error) {
	var angle C.float
	res := C.FMOD_Channel_Get3DSpread(c.cptr, &angle)
	return float64(angle), errs[res]
}

// Sets how much the 3D engine has an effect on the channel, versus that set by 2D panning functions.
//
// level: 3D pan level from 0.0 (attenuation is ignored and panning as set by 2D panning functions) to 1.0 (pan and attenuate according to 3D position), default = 1.0.
//
// Only affects sounds created FMOD_3D.
//
// 2D panning functions include "Channel.SetPan", "Channel.SetMixLevelsOutput", "Channel.SetMixLevelsInput", "Channel.SetMixMatrix", etc
//
// Useful for morhping a sound between 3D and 2D.
// This is most common in volumetric sound, when the sound goes from directional, to 'all around you' (and doesn't pan according to listener position / direction).
func (c *Channel) Set3DLevel(level float64) error {
	res := C.FMOD_Channel_Set3DLevel(c.cptr, C.float(level))
	return errs[res]
}

// Retrieves the current 3D mix level set by "Channel.Set3DLevel".
func (c *Channel) Get3DLevel() (float64, error) {
	var level C.float
	res := C.FMOD_Channel_Get3DLevel(c.cptr, &level)
	return float64(level), errs[res]
}

// Sets the amount by which doppler is scaled.
//
// level: Doppler scale from 0.0 (none), to 1.0 (normal) to 5.0 (exaggerated), default = 1.0.
func (c *Channel) Set3DDopplerLevel(level float64) error {
	res := C.FMOD_Channel_Set3DDopplerLevel(c.cptr, C.float(level))
	return errs[res]
}

// Retrieves the amount by which doppler is scaled.
func (c *Channel) Get3DDopplerLevel() (float64, error) {
	var level C.float
	res := C.FMOD_Channel_Get3DDopplerLevel(c.cptr, &level)
	return float64(level), errs[res]
}

// Control the behaviour of a 3D distance filter, whether to enable or disable it, and frequency characteristics.
//
// custom: Specify true to disable FMOD distance rolloff calculation. Default = false.
//
// customLevel: Specify a attenuation factor manually here, where 1.0 = no attenuation and 0 = complete attenuation. Default = 1.0.
//
// centerFreq: Specify a center frequency in hz for the high-pass filter used to simulate distance attenuation, from 10.0 to 22050.0. Default = 1500.0.
func (c *Channel) Set3DDistanceFilter(custom bool, customLevel, centerFreq float64) error {
	res := C.FMOD_Channel_Set3DDistanceFilter(c.cptr, getBool(custom), C.float(customLevel), C.float(centerFreq))
	return errs[res]
}

// Retrieve the settings for the 3D distance filter properties for a Channel.
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
// Sets a user value that can be retrieved with "Channel.UserData".
func (c *Channel) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_Channel_SetUserData              (FMOD_CHANNEL *channel, void *userdata);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves a user value that can be set with "Channel.SetUserData".
func (c *Channel) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_Channel_GetUserData              (FMOD_CHANNEL *channel, void **userdata);
	return ErrNoImpl
}

/*
   Channel specific control functionality.
*/

// Sets the channel frequency or playback rate, in Hz.
//
// frequency: Frequency value in Hz. This value can also be negative to play the sound backwards (negative frequencies allowed with non-stream sounds only).
//
// When a sound is played, it plays at the default frequency of the sound which can be set by "Sound.SetDefaults".
// For most file formats, the default frequency is determined by the audio format.
func (c *Channel) SetFrequency(frequency float64) error {
	res := C.FMOD_Channel_SetFrequency(c.cptr, C.float(frequency))
	return errs[res]
}

// Retrieves the channel frequency or playback rate, in Hz.
func (c *Channel) Frequency() (float64, error) {
	var frequency C.float
	res := C.FMOD_Channel_GetFrequency(c.cptr, &frequency)
	return float64(frequency), errs[res]
}

// Sets the priority for the channel after it has been played.
//
// priority: Priority for the channel, from 0 (most important) to 256 (least important), default = 128.
//
// When more channels than available are played the virtual channel system will choose existing channels to steal.
// Lower priority sounds will always be stolen before higher priority sounds.
// For channels of equal priority, that with the quietest "Channel.Audibility" value will be stolen.
func (c *Channel) SetPriority(priority int) error {
	res := C.FMOD_Channel_SetPriority(c.cptr, C.int(priority))
	return errs[res]
}

// Retrieves the priority for the channel.
func (c *Channel) Priority() (int, error) {
	var priority C.int
	res := C.FMOD_Channel_GetPriority(c.cptr, &priority)
	return int(priority), errs[res]
}

// Sets the playback position for the currently playing sound to the specified offset.
//
// position: Position of the channel to set in units specified in the 'postype' parameter.
//
// postype: Time unit to set the channel position by. See "TimeUnit".
//
// Certain timeunits do not work depending on the file format. For example "TIMEUNIT_MODORDER" will not work with an MP3 file.
//
// If you are calling this function on a stream, it has to possibly reflush its buffer to get zero latency playback when it resumes playing, therefore it could potentially cause a stall
// or take a small amount of time to do this.
//
// If you are using "NONBLOCKING", note that a stream will go into "OPENSTATE_SETPOSITION" state (see "Sound.OpenState") and sound commands will return ERR_NOTREADY.
// "Channel.Position" will also not update until this non-blocking setposition operation has completed.
//
// Warning! Using a VBR source that does not have an associated seek table or seek information (such as MP3 or MOD/S3M/XM/IT) may cause inaccurate seeking if you specify
// TIMEUNIT_MS or TIMEUNIT_PCM. If you want FMOD to create a PCM vs bytes seek table so that seeking is accurate, you will have to specify "ACCURATETIME"
// when loading or opening the sound. This means there is a slight delay as FMOD scans the whole file when loading the sound to create this table.
func (c *Channel) SetPosition(position uint32, postype TimeUnit) error {
	res := C.FMOD_Channel_SetPosition(c.cptr, C.uint(position), C.FMOD_TIMEUNIT(postype))
	return errs[res]
}

// Returns the current playback position for the specified channel.
//
// Certain timeunits do not work depending on the file format. For example "TIMEUNIT_MODORDER" will not work with an MP3 file.
func (c *Channel) Position(postype TimeUnit) (uint32, error) {
	var position C.uint
	res := C.FMOD_Channel_GetPosition(c.cptr, &position, C.FMOD_TIMEUNIT(postype))
	return uint32(position), errs[res]
}

// Sets a channel to belong to a specified channel group. A channel group can contain many channels.
//
// channelgroup: Pointer to a ChannelGroup object.
//
// Setting a channel to a channel group removes it from any previous group, it does not allow sharing of channel groups.
func (c *Channel) SetChannelGroup(channelgroup ChannelGroup) error {
	res := C.FMOD_Channel_SetChannelGroup(c.cptr, channelgroup.cptr)
	return errs[res]
}

// Retrieves the currently assigned channel group for the channel.
func (c *Channel) ChannelGroup() (ChannelGroup, error) {
	var channelgroup ChannelGroup
	res := C.FMOD_Channel_GetChannelGroup(c.cptr, &channelgroup.cptr)
	return channelgroup, errs[res]
}

// Sets a channel to loop a specified number of times before stopping.
//
// loopcount: Number of times to loop before stopping. 0 = oneshot, 1 = loop once then stop, -1 = loop forever, default = -1.
//
// Issues with streamed audio:
//
// When changing the loop count, sounds created with "System.CreateStream" or "CREATESTREAM" may have already been pre-buffered
// and executed their loop logic ahead of time before this call was even made.
// This is dependant on the size of the sound versus the size of the stream decode buffer (see CREATESOUNDEXINFO).
// If this happens, you may need to reflush the stream buffer by calling "Channel.SetPosition".
// Note this will usually only happen if you have sounds or loop points that are smaller than the stream decode buffer size.
func (c *Channel) SetLoopCount(loopcount int) error {
	res := C.FMOD_Channel_SetLoopCount(c.cptr, C.int(loopcount))
	return errs[res]
}

// Retrieves the current loop count for the specified channel.
//
// This function retrieves the current loop countdown value for the channel being played.
// This means it will decrement until reaching 0, as it plays. To reset the value, use "Channel.SetLoopCount".
func (c *Channel) LoopCount() (int, error) {
	var loopcount C.int
	res := C.FMOD_Channel_GetLoopCount(c.cptr, &loopcount)
	return int(loopcount), errs[res]
}

// Sets the loop points within the channel.
//
// loopstart: Loop start point, this point in time is played so it is inclusive.
//
// loopstarttype: Time format used for the loop start point (see "TimeUnit").
//
// loopend: Loop end point, this point in time is played so it is inclusive.
//
// loopendtype: Time format used for the loop end point (see "TimeUnit").
//
// If a sound was 44100 samples long and you wanted to loop the whole sound, loopstart would be 0, and loopend would be 44099, not 44100.
// You wouldn't use milliseconds in this case because they are not sample accurate.
//
// Issues with streamed audio:
//
// When changing the loop count, sounds created with "System.CreateStream" or "CREATESTREAM" may have already been pre-buffered and executed their loop logic
// ahead of time before this call was even made.
// This is dependant on the size of the sound versus the size of the stream decode buffer (see CREATESOUNDEXINFO).
// If this happens, you may need to reflush the stream buffer by calling "Channel.SetPosition".
// Note this will usually only happen if you have sounds or loop points that are smaller than the stream decode buffer size.
func (c *Channel) SetLoopPoints(loopstart uint32, loopstarttype TimeUnit, loopend uint32, loopendtype TimeUnit) error {
	res := C.FMOD_Channel_SetLoopPoints(c.cptr, C.uint(loopstart), C.FMOD_TIMEUNIT(loopstarttype), C.uint(loopend), C.FMOD_TIMEUNIT(loopendtype))
	return errs[res]
}

// Retrieves the loop points for the channel.
func (c *Channel) LoopPoints(loopstarttype, loopendtype TimeUnit) (uint32, uint32, error) {
	var loopstart, loopend C.uint
	res := C.FMOD_Channel_GetLoopPoints(c.cptr, &loopstart, C.FMOD_TIMEUNIT(loopstarttype), &loopend, C.FMOD_TIMEUNIT(loopendtype))
	return uint32(loopstart), uint32(loopend), errs[res]
}

/*
   Information only functions.
*/

// Retrieves whether the channel is virtual (emulated) or not due to the virtual channel management system.
func (c *Channel) IsVirtual() (bool, error) {
	var isvirtual C.FMOD_BOOL
	res := C.FMOD_Channel_IsVirtual(c.cptr, &isvirtual)
	return setBool(isvirtual), errs[res]
}

// Retrieves the currently playing sound for this channel.
func (c *Channel) CurrentSound() (Sound, error) {
	var sound Sound
	res := C.FMOD_Channel_GetCurrentSound(c.cptr, &sound.cptr)
	return sound, errs[res]
}

// Retrieves the internal channel index for a channel.
func (c *Channel) Index() (int, error) {
	var index C.int
	res := C.FMOD_Channel_GetIndex(c.cptr, &index)
	return int(index), errs[res]
}
