package lowlevel

/*
#include <fmod.h>
*/
import "C"

type SoundGroup struct {
	cptr *C.FMOD_SOUNDGROUP
}

/*
   'SoundGroup' API
*/

// Releases a soundgroup object and returns all sounds back to the master sound group.
// You cannot release the master sound group.
func (s *SoundGroup) Release() error {
	res := C.FMOD_SoundGroup_Release(s.cptr)
	return errs[res]
}

// Retrieves the parent System object that was used to create this object.
func (s *SoundGroup) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_SoundGroup_GetSystemObject(s.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   SoundGroup control functions.
*/

// Limits the number of concurrent playbacks of sounds in a sound group to the specified value.
// After this, if the sounds in the sound group are playing this many times, any attepts to play more of the sounds in the sound group will by default fail with FMOD_ERR_MAXAUDIBLE.
// Use "SoundGroup.SetMaxAudibleBehavior" to change the way the sound playback behaves when too many sounds are playing. Muting, failing and stealing behaviors can be specified.
// "SoundGroup.NumPlaying" can be used to determine how many instances of the sounds in the sound group are currently playing.
//
// maxaudible: A variable to receive the number of playbacks to be audible at once. -1 = unlimited. 0 means no sounds in this group will succeed. Default = -1.
func (s *SoundGroup) SetMaxAudible(maxaudible int) error {
	res := C.FMOD_SoundGroup_SetMaxAudible(s.cptr, C.int(maxaudible))
	return errs[res]
}

// Retrieves the number of concurrent playbacks of sounds in a sound group to the specified value.
// If the sounds in the sound group are playing this many times, any attepts to play more of the sounds in the sound group will fail with FMOD_ERR_MAXAUDIBLE.
func (s *SoundGroup) MaxAudible() (int, error) {
	var maxaudible C.int
	res := C.FMOD_SoundGroup_GetMaxAudible(s.cptr, &maxaudible)
	return int(maxaudible), errs[res]
}

// NOTE: Not implement yet
// This function changes the way the sound playback behaves when too many sounds are playing in a soundgroup. Muting, failing and stealing behaviors can be specified.
// behavior: Specify a behavior determined with a FMOD_SOUNDGROUP_BEHAVIOR flag. Default is FMOD_SOUNDGROUP_BEHAVIOR_FAIL.
func (s *SoundGroup) SetMaxAudibleBehavior(behavior C.FMOD_SOUNDGROUP_BEHAVIOR) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_SetMaxAudibleBehavior (FMOD_SOUNDGROUP *soundgroup, FMOD_SOUNDGROUP_BEHAVIOR behavior);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the current max audible behavior method.
func (s *SoundGroup) MaxAudibleBehavior(behavior *C.FMOD_SOUNDGROUP_BEHAVIOR) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetMaxAudibleBehavior (FMOD_SOUNDGROUP *soundgroup, FMOD_SOUNDGROUP_BEHAVIOR *behavior);
	return ErrNoImpl
}

// Specify a time in seconds for FMOD_SOUNDGROUP_BEHAVIOR_MUTE behavior to fade with. By default there is no fade.
//
// speed: Fade time in seconds (1.0 = 1 second). Default = 0.0. (no fade).
//
// When more sounds are playing in a SoundGroup than are specified with "SoundGroup.SetMaxAudible", the least important sound (ie lowest priority / lowest audible volume due to 3d position, volume etc)
// will fade to silence if FMOD_SOUNDGROUP_BEHAVIOR_MUTE is used, and any previous sounds that were silent because of this rule will fade in if they are more important.
//
// If a mode besides FMOD_SOUNDGROUP_BEHAVIOR_MUTE is used, the fade speed is ignored.
func (s *SoundGroup) SetMuteFadeSpeed(speed float64) error {
	res := C.FMOD_SoundGroup_SetMuteFadeSpeed(s.cptr, C.float(speed))
	return errs[res]
}

// Retrieves the current time in seconds for FMOD_SOUNDGROUP_BEHAVIOR_MUTE behavior to fade with.
// If a mode besides FMOD_SOUNDGROUP_BEHAVIOR_MUTE is used, the fade speed is ignored.
func (s *SoundGroup) MuteFadeSpeed() (float64, error) {
	var speed C.float
	res := C.FMOD_SoundGroup_GetMuteFadeSpeed(s.cptr, &speed)
	return float64(speed), errs[res]
}

// Sets the volume for a sound group, affecting all channels playing the sounds in this soundgroup.
//
// volume: A linear volume level. 0.0 = silent, 1.0 = full volume. Default = 1.0. Negative volumes and amplification (> 1.0) are supported.
func (s *SoundGroup) SetVolume(volume float64) error {
	res := C.FMOD_SoundGroup_SetVolume(s.cptr, C.float(volume))
	return errs[res]
}

// Retrieves the volume for the sounds within a soundgroup.
func (s *SoundGroup) Volume() (float64, error) {
	var volume C.float
	res := C.FMOD_SoundGroup_GetVolume(s.cptr, &volume)
	return float64(volume), errs[res]
}

// Stops all sounds within this soundgroup.
func (s *SoundGroup) Stop() error {
	res := C.FMOD_SoundGroup_Stop(s.cptr)
	return errs[res]
}

/*
   Information only functions.
*/

// NOTE: Not implement yet
// Retrieves the name of the sound group.
func (s *SoundGroup) Name(name *C.char, namelen C.int) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetName(FMOD_SOUNDGROUP *soundgroup, char *name, int namelen);
	return ErrNoImpl
}

// Retrieves the current number of sounds in this sound group.
func (s *SoundGroup) NumSounds() (int, error) {
	var numsounds C.int
	res := C.FMOD_SoundGroup_GetNumSounds(s.cptr, &numsounds)
	return int(numsounds), errs[res]
}

// Retrieves a pointer to a sound from within a sound group.
//
// index: Index of the sound that is to be retrieved.
//
// Use "SoundGroup.NumSounds" in conjunction with this function to enumerate all sounds in a sound group.
func (s *SoundGroup) Sound(index int) (*Sound, error) {
	var sound Sound
	res := C.FMOD_SoundGroup_GetSound(s.cptr, C.int(index), &sound.cptr)
	return &sound, errs[res]
}

// Retrieves the number of currently playing channels for the sound group.
// This routine returns the number of channels playing. If the sound group only has 1 sound, and that sound is playing twice, the figure returned will be 2.
func (s *SoundGroup) NumPlaying() (int, error) {
	var numplaying C.int
	res := C.FMOD_SoundGroup_GetNumPlaying(s.cptr, &numplaying)
	return int(numplaying), errs[res]
}

/*
   Userdata set/get.
*/

// NOTE: Not implement yet
// Sets a user value that the SoundGroup object will store internally. Can be retrieved with "SoundGroup.UserData".
//
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
// It can be useful if an FMOD callback passes an object of this type as a parameter, and the user does not know which object it is (if many of these types of objects exist).
// Using "SoundGroup.UserData" would help in the identification of the object.
func (s *SoundGroup) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_SetUserData           (FMOD_SOUNDGROUP *soundgroup, void *userdata);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the user value that that was set by calling the "SoundGroup.SetUserData" function.
func (s *SoundGroup) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetUserData           (FMOD_SOUNDGROUP *soundgroup, void **userdata);
	return ErrNoImpl
}
