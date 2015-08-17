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

func (s *SoundGroup) Release() error {
	res := C.FMOD_SoundGroup_Release(s.cptr)
	return errs[res]
}

func (s *SoundGroup) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_SoundGroup_GetSystemObject(s.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   SoundGroup control functions.
*/

func (s *SoundGroup) SetMaxAudible(maxaudible int) error {
	res := C.FMOD_SoundGroup_SetMaxAudible(s.cptr, C.int(maxaudible))
	return errs[res]
}

func (s *SoundGroup) MaxAudible() (int, error) {
	var maxaudible C.int
	res := C.FMOD_SoundGroup_GetMaxAudible(s.cptr, &maxaudible)
	return int(maxaudible), errs[res]
}

func (s *SoundGroup) SetMaxAudibleBehavior(behavior C.FMOD_SOUNDGROUP_BEHAVIOR) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_SetMaxAudibleBehavior (FMOD_SOUNDGROUP *soundgroup, FMOD_SOUNDGROUP_BEHAVIOR behavior);
	return ErrNoImpl
}

func (s *SoundGroup) MaxAudibleBehavior(behavior *C.FMOD_SOUNDGROUP_BEHAVIOR) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetMaxAudibleBehavior (FMOD_SOUNDGROUP *soundgroup, FMOD_SOUNDGROUP_BEHAVIOR *behavior);
	return ErrNoImpl
}

func (s *SoundGroup) SetMuteFadeSpeed(speed float64) error {
	res := C.FMOD_SoundGroup_SetMuteFadeSpeed(s.cptr, C.float(speed))
	return errs[res]
}

func (s *SoundGroup) MuteFadeSpeed() (float64, error) {
	var speed C.float
	res := C.FMOD_SoundGroup_GetMuteFadeSpeed(s.cptr, &speed)
	return float64(speed), errs[res]
}

func (s *SoundGroup) SetVolume(volume float64) error {
	res := C.FMOD_SoundGroup_SetVolume(s.cptr, C.float(volume))
	return errs[res]
}

func (s *SoundGroup) Volume() (float64, error) {
	var volume C.float
	res := C.FMOD_SoundGroup_GetVolume(s.cptr, &volume)
	return float64(volume), errs[res]
}

func (s *SoundGroup) Stop() error {
	res := C.FMOD_SoundGroup_Stop(s.cptr)
	return errs[res]
}

/*
   Information only functions.
*/

func (s *SoundGroup) Name(name *C.char, namelen C.int) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetName(FMOD_SOUNDGROUP *soundgroup, char *name, int namelen);
	return ErrNoImpl
}

func (s *SoundGroup) NumSounds() (int, error) {
	var numsounds C.int
	res := C.FMOD_SoundGroup_GetNumSounds(s.cptr, &numsounds)
	return int(numsounds), errs[res]
}

func (s *SoundGroup) Sound(index int) (*Sound, error) {
	var sound Sound
	res := C.FMOD_SoundGroup_GetSound(s.cptr, C.int(index), &sound.cptr)
	return &sound, errs[res]
}

func (s *SoundGroup) NumPlaying() (int, error) {
	var numplaying C.int
	res := C.FMOD_SoundGroup_GetNumPlaying(s.cptr, &numplaying)
	return int(numplaying), errs[res]
}

/*
   Userdata set/get.
*/

func (s *SoundGroup) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_SetUserData           (FMOD_SOUNDGROUP *soundgroup, void *userdata);
	return ErrNoImpl
}

func (s *SoundGroup) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_SoundGroup_GetUserData           (FMOD_SOUNDGROUP *soundgroup, void **userdata);
	return ErrNoImpl
}
