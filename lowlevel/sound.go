package lowlevel

/*
#include <fmod.h>
*/
import "C"

type Sound struct {
	cptr *C.FMOD_SOUND
}

/*
   'Sound' API
*/

func (s *Sound) Release() error {
	res := C.FMOD_Sound_Release(s.cptr)
	return errs[res]
}

func (s *Sound) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_Sound_GetSystemObject(s.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   Standard sound manipulation functions.
*/

// NOTE: Not implement yet
func (s *Sound) Lock(offset, length C.uint, ptr1, ptr2 **interface{}, len1, len2 *C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_Lock(FMOD_SOUND *sound, unsigned int offset, unsigned int length, void **ptr1, void **ptr2, unsigned int *len1, unsigned int *len2);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) Unlock(ptr1, ptr2 *interface{}, len1, len2 C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_Unlock(FMOD_SOUND *sound, void *ptr1, void *ptr2, unsigned int len1, unsigned int len2);
	return ErrNoImpl
}

func (s *Sound) SetDefaults(frequency float64, priority int) error {
	res := C.FMOD_Sound_SetDefaults(s.cptr, C.float(frequency), C.int(priority))
	return errs[res]
}

func (s *Sound) Defaults() (float64, int, error) {
	var frequency C.float
	var priority C.int
	res := C.FMOD_Sound_GetDefaults(s.cptr, &frequency, &priority)
	return float64(frequency), int(priority), errs[res]
}

func (s *Sound) Set3DMinMaxDistance(min, max float64) error {
	res := C.FMOD_Sound_Set3DMinMaxDistance(s.cptr, C.float(min), C.float(max))
	return errs[res]
}

func (s *Sound) Get3DMinMaxDistance() (float64, float64, error) {
	var min, max C.float
	res := C.FMOD_Sound_Get3DMinMaxDistance(s.cptr, &min, &max)
	return float64(min), float64(max), errs[res]
}

func (s *Sound) Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error {
	res := C.FMOD_Sound_Set3DConeSettings(s.cptr, C.float(insideconeangle), C.float(outsideconeangle), C.float(outsidevolume))
	return errs[res]
}

func (s *Sound) Get3DConeSettings() (float64, float64, float64, error) {
	var insideconeangle, outsideconeangle, outsidevolume C.float
	res := C.FMOD_Sound_Get3DConeSettings(s.cptr, &insideconeangle, &outsideconeangle, &outsidevolume)
	return float64(insideconeangle), float64(outsideconeangle), float64(outsidevolume), errs[res]
}

func (s *Sound) Set3DCustomRolloff(points *Vector, numpoints int) error {
	cpoints := points.toC()
	res := C.FMOD_Sound_Set3DCustomRolloff(s.cptr, &cpoints, C.int(numpoints))
	return errs[res]
}

func (s *Sound) Get3DCustomRolloff() (*Vector, int, error) {
	var points Vector
	var cpoints *C.FMOD_VECTOR
	var numpoints C.int
	res := C.FMOD_Sound_Get3DCustomRolloff(s.cptr, &cpoints, &numpoints)
	points.fromC(*cpoints)
	return &points, int(numpoints), errs[res]
}

func (s *Sound) SetSubSound(index int, subsound *Sound) error {
	res := C.FMOD_Sound_SetSubSound(s.cptr, C.int(index), subsound.cptr)
	return errs[res]
}

func (s *Sound) SubSound(index int) (*Sound, error) {
	var sound Sound
	res := C.FMOD_Sound_GetSubSound(s.cptr, C.int(index), &sound.cptr)
	return &sound, errs[res]
}

func (s *Sound) SubSoundParent() (*Sound, error) {
	var parentsound Sound
	res := C.FMOD_Sound_GetSubSoundParent(s.cptr, &parentsound.cptr)
	return &parentsound, errs[res]
}

// NOTE: Not implement yet
func (s *Sound) Name(name *C.char, namelen C.int) error {
	//FMOD_RESULT F_API FMOD_Sound_GetName(FMOD_SOUND *sound, char *name, int namelen);
	return ErrNoImpl
}

func (s *Sound) Length(lengthtype TimeUnit) (uint32, error) {
	var length C.uint
	res := C.FMOD_Sound_GetLength(s.cptr, &length, C.FMOD_TIMEUNIT(lengthtype))
	return uint32(length), errs[res]
}

func (s *Sound) Format() (SoundType, SoundFormat, int, int, error) {
	var typ C.FMOD_SOUND_TYPE
	var format C.FMOD_SOUND_FORMAT
	var channels, bits C.int
	res := C.FMOD_Sound_GetFormat(s.cptr, &typ, &format, &channels, &bits)
	return SoundType(typ), SoundFormat(format), int(channels), int(bits), errs[res]
}

func (s *Sound) NumSubSounds() (int, error) {
	var numsubsounds C.int
	res := C.FMOD_Sound_GetNumSubSounds(s.cptr, &numsubsounds)
	return int(numsubsounds), errs[res]
}

func (s *Sound) NumTags() (int, int, error) {
	var numtags, numtagsupdated C.int
	res := C.FMOD_Sound_GetNumTags(s.cptr, &numtags, &numtagsupdated)
	return int(numtags), int(numtagsupdated), errs[res]
}

// NOTE: Not implement yet
func (s *Sound) Tag(name *C.char, index C.int, tag *C.FMOD_TAG) error {
	//FMOD_RESULT F_API FMOD_Sound_GetTag(FMOD_SOUND *sound, const char *name, int index, FMOD_TAG *tag);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) OpenState(openstate *C.FMOD_OPENSTATE, percentbuffered *C.uint, starving, diskbusy *C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_Sound_GetOpenState(FMOD_SOUND *sound, FMOD_OPENSTATE *openstate, unsigned int *percentbuffered, FMOD_BOOL *starving, FMOD_BOOL *diskbusy);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) ReadData(buffer *interface{}, lenbytes C.uint, read C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_ReadData(FMOD_SOUND *sound, void *buffer, unsigned int lenbytes, unsigned int *read);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) SeekData(pcm C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_SeekData(FMOD_SOUND *sound, unsigned int pcm);
	return ErrNoImpl
}

func (s *Sound) SetSoundGroup(soundgroup *SoundGroup) error {
	res := C.FMOD_Sound_SetSoundGroup(s.cptr, soundgroup.cptr)
	return errs[res]
}

func (s *Sound) SoundGroup() (*SoundGroup, error) {
	var soundgroup SoundGroup
	res := C.FMOD_Sound_GetSoundGroup(s.cptr, &soundgroup.cptr)
	return &soundgroup, errs[res]
}

/*
   Synchronization point API.  These points can come from markers embedded in wav files, and can also generate channel callbacks.
*/

// NOTE: Not implement yet
func (s *Sound) NumSyncPoints(numsyncpoints *C.int) error {
	//FMOD_RESULT F_API FMOD_Sound_GetNumSyncPoints           (FMOD_SOUND *sound, int *numsyncpoints);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) SyncPoint(index C.int, point **C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_GetSyncPoint               (FMOD_SOUND *sound, int index, FMOD_SYNCPOINT **point);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) SyncPointInfo(point *C.FMOD_SYNCPOINT, name *C.char, namelen C.int, offset *C.uint, offsettype C.FMOD_TIMEUNIT) error {
	//FMOD_RESULT F_API FMOD_Sound_GetSyncPointInfo           (FMOD_SOUND *sound, FMOD_SYNCPOINT *point, char *name, int namelen, unsigned int *offset, FMOD_TIMEUNIT offsettype);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) AddSyncPoint(offset C.uint, offsettype C.FMOD_TIMEUNIT, name *C.char, point **C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_AddSyncPoint               (FMOD_SOUND *sound, unsigned int offset, FMOD_TIMEUNIT offsettype, const char *name, FMOD_SYNCPOINT **point);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) DeleteSyncPoint(point *C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_DeleteSyncPoint            (FMOD_SOUND *sound, FMOD_SYNCPOINT *point);
	return ErrNoImpl
}

/*
   Functions also in Channel class but here they are the 'default' to save having to change it in Channel all the time.
*/

func (s *Sound) SetMode(mode Mode) error {
	res := C.FMOD_Sound_SetMode(s.cptr, C.FMOD_MODE(mode))
	return errs[res]
}

func (s *Sound) Mode() (Mode, error) {
	var mode C.FMOD_MODE
	res := C.FMOD_Sound_GetMode(s.cptr, &mode)
	return Mode(mode), errs[res]
}

func (s *Sound) SetLoopCount(loopcount int) error {
	res := C.FMOD_Sound_SetLoopCount(s.cptr, C.int(loopcount))
	return errs[res]
}

func (s *Sound) LoopCount() (int, error) {
	var loopcount C.int
	res := C.FMOD_Sound_GetLoopCount(s.cptr, &loopcount)
	return int(loopcount), errs[res]
}

func (s *Sound) SetLoopPoints(loopstart uint32, loopstarttype TimeUnit, loopend uint32, loopendtype TimeUnit) error {
	res := C.FMOD_Sound_SetLoopPoints(s.cptr, C.uint(loopstart), C.FMOD_TIMEUNIT(loopstarttype), C.uint(loopend), C.FMOD_TIMEUNIT(loopendtype))
	return errs[res]
}

func (s *Sound) LoopPoints(loopstarttype, loopendtype TimeUnit) (uint32, uint32, error) {
	var loopstart, loopend C.uint
	res := C.FMOD_Sound_GetLoopPoints(s.cptr, &loopstart, C.FMOD_TIMEUNIT(loopstarttype), &loopend, C.FMOD_TIMEUNIT(loopendtype))
	return uint32(loopstart), uint32(loopend), errs[res]
}

/*
   For MOD/S3M/XM/IT/MID sequenced formats only.
*/

func (s *Sound) MusicNumChannels() (int, error) {
	var numchannels C.int
	res := C.FMOD_Sound_GetMusicNumChannels(s.cptr, &numchannels)
	return int(numchannels), errs[res]
}

func (s *Sound) SetMusicChannelVolume(channel int, volume float64) error {
	res := C.FMOD_Sound_SetMusicChannelVolume(s.cptr, C.int(channel), C.float(volume))
	return errs[res]
}

func (s *Sound) MusicChannelVolume(channel int) (float64, error) {
	var volume C.float
	res := C.FMOD_Sound_GetMusicChannelVolume(s.cptr, C.int(channel), &volume)
	return float64(volume), errs[res]
}

func (s *Sound) SetMusicSpeed(speed float64) error {
	res := C.FMOD_Sound_SetMusicSpeed(s.cptr, C.float(speed))
	return errs[res]
}

func (s *Sound) MusicSpeed() (float64, error) {
	var speed C.float
	res := C.FMOD_Sound_GetMusicSpeed(s.cptr, &speed)
	return float64(speed), errs[res]
}

/*
   Userdata set/get.
*/

// NOTE: Not implement yet
func (s *Sound) SetUserData(userdata *interface{}) error {
	//FMOD_RESULT F_API FMOD_Sound_SetUserData                (FMOD_SOUND *sound, void *userdata);
	return ErrNoImpl
}

// NOTE: Not implement yet
func (s *Sound) UserData(userdata **interface{}) error {
	//FMOD_RESULT F_API FMOD_Sound_GetUserData                (FMOD_SOUND *sound, void **userdata);
	return ErrNoImpl
}
