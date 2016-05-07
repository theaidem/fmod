package lowlevel

/*
#include <fmod.h>
*/
import "C"
import "unsafe"

type Sound struct {
	cptr *C.FMOD_SOUND
}

/*
   'Sound' API
*/

// Frees a sound object.
// This will free the sound object and everything created under it.
// If this is a stream that is playing as a subsound of another parent stream, then if this is the currently playing subsound, the whole stream will stop.
// Note - This function will block if it was opened with NONBLOCKING and hasn't finished opening yet.
func (s *Sound) Release() error {
	res := C.FMOD_Sound_Release(s.cptr)
	return errs[res]
}

// Retrieves the parent System object that was used to create this object.
func (s *Sound) SystemObject() (*System, error) {
	var system System
	res := C.FMOD_Sound_GetSystemObject(s.cptr, &system.cptr)
	return &system, errs[res]
}

/*
   Standard sound manipulation functions.
*/

// NOTE: Not implement yet
// Returns a pointer to the beginning of the sample data for a sound.
//
// offset: Offset in bytes to the position you want to lock in the sample buffer.
//
//length: Number of bytes you want to lock in the sample buffer.
//
// You must always unlock the data again after you have finished with it, using Sound::unlock.
// With this function you get access to the RAW audio data, for example 8, 16, 24 or 32bit PCM data, mono or stereo data, and on consoles, vag, xadpcm or gcadpcm compressed data.
// You must take this into consideration when processing the data within the pointer.
func (s *Sound) Lock(offset, length C.uint, ptr1, ptr2 **interface{}, len1, len2 *C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_Lock(FMOD_SOUND *sound, unsigned int offset, unsigned int length, void **ptr1, void **ptr2, unsigned int *len1, unsigned int *len2);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Releases previous sample data lock from "Sound.Lock".
func (s *Sound) Unlock(ptr1, ptr2 *interface{}, len1, len2 C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_Unlock(FMOD_SOUND *sound, void *ptr1, void *ptr2, unsigned int len1, unsigned int len2);
	return ErrNoImpl
}

// Sets a sounds's default attributes, so when it is played it uses these values without having to specify them later for each channel each time the sound is played.
// There are no 'ignore' values for these parameters. Use "Sound.Defaults" if you want to change only 1 and leave others unaltered.
func (s *Sound) SetDefaults(frequency float64, priority int) error {
	res := C.FMOD_Sound_SetDefaults(s.cptr, C.float(frequency), C.int(priority))
	return errs[res]
}

// Retrieves a sound's default attributes for when it is played on a channel with "System.PlaySound".
func (s *Sound) Defaults() (float64, int, error) {
	var frequency C.float
	var priority C.int
	res := C.FMOD_Sound_GetDefaults(s.cptr, &frequency, &priority)
	return float64(frequency), int(priority), errs[res]
}

// Sets the minimum and maximum audible distance for a sound.
//
// MinDistance is the minimum distance that the sound emitter will cease to continue growing louder at (as it approaches the listener).
// Within the mindistance it stays at the constant loudest volume possible. Outside of this mindistance it begins to attenuate.
// MaxDistance is the distance a sound stops attenuating at. Beyond this point it will stay at the volume it would be at maxdistance units from the listener and will not attenuate any more.
// MinDistance is useful to give the impression that the sound is loud or soft in 3d space. An example of this is a small quiet object, such as a bumblebee,
// which you could set a mindistance of to 0.1 for example, which would cause it to attenuate quickly and dissapear when only a few meters away from the listener.
// Another example is a jumbo jet, which you could set to a mindistance of 100.0, which would keep the sound volume at max until the listener was 100 meters away,
// then it would be hundreds of meters more before it would fade out.
//
// In summary, increase the mindistance of a sound to make it 'louder' in a 3d world, and decrease it to make it 'quieter' in a 3d world.
// Maxdistance is effectively obsolete unless you need the sound to stop fading out at a certain point. Do not adjust this from the default if you dont need to.
// Some people have the confusion that maxdistance is the point the sound will fade out to, this is not the case.
//
// min: The sound's minimum volume distance in "units". See remarks for more on units.
//
// max: The sound's maximum volume distance in "units". See remarks for more on units.
//
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
// See "System.Set3DSettings" for more on this.
// The default units for minimum and maximum distances are 1.0 and 10,000.0f.
func (s *Sound) Set3DMinMaxDistance(min, max float64) error {
	res := C.FMOD_Sound_Set3DMinMaxDistance(s.cptr, C.float(min), C.float(max))
	return errs[res]
}

// Retrieve the minimum and maximum audible distance for a sound.
//
// A 'distance unit' is specified by "System.Set3DSettings". By default this is set to meters which is a distance scale of 1.0.
// See "System.Set3DSettings" for more on this.
// The default units for minimum and maximum distances are 1.0 and 10,000.0f.
func (s *Sound) Get3DMinMaxDistance() (float64, float64, error) {
	var min, max C.float
	res := C.FMOD_Sound_Get3DMinMaxDistance(s.cptr, &min, &max)
	return float64(min), float64(max), errs[res]
}

// Sets the inside and outside angles of the sound projection cone, as well as the volume of the sound outside the outside angle of the sound projection cone.
//
// insideconeangle: Inside cone angle, in degrees, from 0 to 360. This is the angle within which the sound is at its normal volume.
// Must not be greater than outsideconeangle. Default = 360.
//
// outsideconeangle: Outside cone angle, in degrees, from 0 to 360. This is the angle outside of which the sound is at its outside volume.
// Must not be less than insideconeangle. Default = 360.
//
// outsidevolume: Cone outside volume, from 0 to 1.0. Default = 1.0.
func (s *Sound) Set3DConeSettings(insideconeangle, outsideconeangle, outsidevolume float64) error {
	res := C.FMOD_Sound_Set3DConeSettings(s.cptr, C.float(insideconeangle), C.float(outsideconeangle), C.float(outsidevolume))
	return errs[res]
}

// Retrieves the inside and outside angles of the sound projection cone.
func (s *Sound) Get3DConeSettings() (float64, float64, float64, error) {
	var insideconeangle, outsideconeangle, outsidevolume C.float
	res := C.FMOD_Sound_Get3DConeSettings(s.cptr, &insideconeangle, &outsideconeangle, &outsidevolume)
	return float64(insideconeangle), float64(outsideconeangle), float64(outsidevolume), errs[res]
}

// TODO: add more docs
// Point a sound to use a custom rolloff curve. Must be used in conjunction with FMOD_3D_CUSTOMROLLOFF flag to be activated.
//
// points: An array of "Vector" structures where x = distance and y = volume from 0.0 to 1.0. z should be set to 0.
//
// numpoints: The number of points in the array.
//
// Note! This function does not duplicate the memory for the points internally. The pointer you pass to FMOD must remain valid until there is no more use for it.
// Do not free the memory while in use, or use a local variable that goes out of scope while in use.
//
// Points must be sorted by distance! Passing an unsorted list to FMOD will result in an error.
//
//Set the points parameter to 0 or NULL to disable the points. If FMOD_3D_CUSTOMROLLOFF is set and the rolloff curve is 0, FMOD will revert to inverse curve rolloff.
//
// Min and maxdistance are meaningless when FMOD_3D_CUSTOMROLLOFF is used and the values are ignored.
func (s *Sound) Set3DCustomRolloff(points *Vector, numpoints int) error {
	cpoints := points.toC()
	res := C.FMOD_Sound_Set3DCustomRolloff(s.cptr, &cpoints, C.int(numpoints))
	return errs[res]
}

// Retrieves a pointer to the sound's current custom rolloff curve.
func (s *Sound) Get3DCustomRolloff() (*Vector, int, error) {
	var points = NewVector()
	var cpoints *C.FMOD_VECTOR = points.toCp()
	var numpoints C.int
	res := C.FMOD_Sound_Get3DCustomRolloff(s.cptr, unsafe.Pointer(&cpoints), &numpoints)
	if cpoints != nil {
		points.fromC(*cpoints)
	}
	return &points, int(numpoints), errs[res]
}

// Assigns a sound as a 'subsound' of another sound. A sound can contain other sounds.
// The sound object that is issuing the command will be the 'parent' sound.
//
// index: Index within the sound to set the new sound to as a 'subsound'.
//
// subsound: Sound object to set as a subsound within this sound.
/*
func (s *Sound) SetSubSound(index int, subsound *Sound) error {
	res := C.FMOD_Sound_SetSubSound(s.cptr, C.int(index), subsound.cptr)
	return errs[res]
}
*/

// Retrieves a handle to a Sound object that is contained within the parent sound.
// If the sound is a stream and FMOD_NONBLOCKING was not used, then this call will perform a blocking seek/flush to the specified subsound.
//
// If FMOD_NONBLOCKING was used to open this sound and the sound is a stream, FMOD will do a non blocking seek/flush and set the state of the subsound to FMOD_OPENSTATE_SEEKING.
// The sound won't be ready to be used in this case until the state of the sound becomes FMOD_OPENSTATE_READY (or FMOD_OPENSTATE_ERROR).
func (s *Sound) SubSound(index int) (*Sound, error) {
	var sound Sound
	res := C.FMOD_Sound_GetSubSound(s.cptr, C.int(index), &sound.cptr)
	return &sound, errs[res]
}

// Retrieves a handle to the parent Sound object that contains our subsound.
// If the sound is not a subsound, the parentsound will be returned as NULL.
func (s *Sound) SubSoundParent() (*Sound, error) {
	var parentsound Sound
	res := C.FMOD_Sound_GetSubSoundParent(s.cptr, &parentsound.cptr)
	return &parentsound, errs[res]
}

// NOTE: Not implement yet
// Retrieves the name of a sound.
// if FMOD_LOWMEM has been specified in "System.CreateSound", this function will return "(null)".
func (s *Sound) Name(name *C.char, namelen C.int) error {
	//FMOD_RESULT F_API FMOD_Sound_GetName(FMOD_SOUND *sound, char *name, int namelen);
	return ErrNoImpl
}

// Retrieves the length of the sound using the specified time unit.
//
// Certain timeunits do not work depending on the file format. For example TIMEUNIT_MODORDER will not work with an mp3 file.
// A length of 0xFFFFFFFF usually means it is of unlimited length, such as an internet radio stream or MOD/S3M/XM/IT file which may loop forever.
//
// Warning! Using a VBR source that does not have an associated length information in milliseconds or pcm samples (such as MP3 or MOD/S3M/XM/IT) may return
// inaccurate lengths specify TIMEUNIT_MS or TIMEUNIT_PCM.
// If you want FMOD to retrieve an accurate length it will have to pre-scan the file first in this case.
// You will have to specify FMOD_ACCURATETIME when loading or opening the sound. This means there is a slight delay as FMOD scans the whole file when loading the sound to find
// the right length in millseconds or pcm samples, and this also creates a seek table as it does this for seeking purposes.
func (s *Sound) Length(lengthtype TimeUnit) (uint32, error) {
	var length C.uint
	res := C.FMOD_Sound_GetLength(s.cptr, &length, C.FMOD_TIMEUNIT(lengthtype))
	return uint32(length), errs[res]
}

// Returns format information about the sound.
func (s *Sound) Format() (SoundType, SoundFormat, int, int, error) {
	var typ C.FMOD_SOUND_TYPE
	var format C.FMOD_SOUND_FORMAT
	var channels, bits C.int
	res := C.FMOD_Sound_GetFormat(s.cptr, &typ, &format, &channels, &bits)
	return SoundType(typ), SoundFormat(format), int(channels), int(bits), errs[res]
}

// Retrieves the number of subsounds stored within a sound.
// A format that has subsounds is usually a container format, such as FSB, DLS, MOD, S3M, XM, IT.
func (s *Sound) NumSubSounds() (int, error) {
	var numsubsounds C.int
	res := C.FMOD_Sound_GetNumSubSounds(s.cptr, &numsubsounds)
	return int(numsubsounds), errs[res]
}

// Retrieves the number of tags belonging to a sound.
// The 'numtagsupdated' parameter can be used to check if any tags have been updated since last calling this function.
// This can be useful to update tag fields, for example from internet based streams, such as shoutcast or icecast where the name of the song might change.
func (s *Sound) NumTags() (int, int, error) {
	var numtags, numtagsupdated C.int
	res := C.FMOD_Sound_GetNumTags(s.cptr, &numtags, &numtagsupdated)
	return int(numtags), int(numtagsupdated), errs[res]
}

// NOTE: Not implement yet
// TODO: add more docs
// Retrieves a descriptive tag stored by the sound, to describe things like the song name, author etc.
func (s *Sound) Tag(name *C.char, index C.int, tag *C.FMOD_TAG) error {
	//FMOD_RESULT F_API FMOD_Sound_GetTag(FMOD_SOUND *sound, const char *name, int index, FMOD_TAG *tag);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieves the state a sound is in after FMOD_NONBLOCKING has been used to open it, or the state of the streaming buffer.
// When a sound is opened with FMOD_NONBLOCKING, it is opened and prepared in the background, or asynchronously.
// This allows the main application to execute without stalling on audio loads.
// This function will describe the state of the asynchronous load routine i.e. whether it has succeeded, failed or is still in progress.
//
// If 'starving' is true, then you will most likely hear a stuttering/repeating sound as the decode buffer loops on itself and replays old data.
// Now that this variable exists, you can detect buffer underrun and use something like Channel::setMute to keep it quiet until it is not starving any more.
//
// Note: Always check 'openstate' to determine the state of the sound. Do not assume that if this function returns FMOD_OK then the sound has finished loading.
func (s *Sound) OpenState(openstate *C.FMOD_OPENSTATE, percentbuffered *C.uint, starving, diskbusy *C.FMOD_BOOL) error {
	//FMOD_RESULT F_API FMOD_Sound_GetOpenState(FMOD_SOUND *sound, FMOD_OPENSTATE *openstate, unsigned int *percentbuffered, FMOD_BOOL *starving, FMOD_BOOL *diskbusy);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Reads data from an opened sound to a specified pointer, using the FMOD codec created internally.
// This can be used for decoding data offline in small pieces (or big pieces), rather than playing and capturing it, or loading the whole file at once and having to lock / unlock the data.
//
// If too much data is read, it is possible FMOD_ERR_FILE_EOF will be returned, meaning it is out of data. The 'read' parameter will reflect this by returning a smaller number of bytes read than was requested.
// As a sound already reads the whole file then closes it upon calling "System.CreateSound" (unless "System.CreateStream" or FMOD_CREATESTREAM is used),
// this function will not work because the file is no longer open.
// Note that opening a stream makes it read a chunk of data and this will advance the read cursor.
// You need to either use FMOD_OPENONLY to stop the stream pre-buffering or call "Sound.SeekData" to reset the read cursor.
// If FMOD_OPENONLY flag is used when opening a sound, it will leave the file handle open, and FMOD will not read any data internally,
// so the read cursor will be at position 0. This will allow the user to read the data from the start.
// As noted previously, if a sound is opened as a stream and this function is called to read some data, then you will 'miss the start' of the sound.
// "Channel.SetPosition" will have the same result. These function will flush the stream buffer and read in a chunk of audio internally.
// This is why if you want to read from an absolute position you should use Sound::seekData and not the previously mentioned functions.
// Remember if you are calling readData and seekData on a stream it is up to you to cope with the side effects that may occur.
// Information functions such as "Channel.Position" may give misleading results. Calling "Channel.SetPosition" will reset and flush the stream, leading to the time values returning to their correct position.
//
// NOTE! Thread safety. If you call this from another stream callback, or any other thread besides the main thread, make sure to put a criticalsection around the call,
// and another around Sound::release in case the sound is still being read from while releasing.
// This function is thread safe to call from a stream callback or different thread as long as it doesnt conflict with a call to "Sound.Release".
func (s *Sound) ReadData(buffer *interface{}, lenbytes C.uint, read C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_ReadData(FMOD_SOUND *sound, void *buffer, unsigned int lenbytes, unsigned int *read);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Seeks a sound for use with data reading. This is not a function to 'seek a sound' for normal use.
// This is for use in conjunction with "Sound.ReadData".
//
// Note. If a stream is opened and this function is called to read some data, then it will advance the internal file pointer, so data will be skipped if you play the stream.
// Also calling position / time information functions will lead to misleading results.
// A stream can be reset before playing by setting the position of the channel (ie using "Channel.SetPosition"), which will make it seek, reset and flush the stream buffer.
// This will make it sound correct again.
// Remember if you are calling readData and seekData on a stream it is up to you to cope with the side effects that may occur.
func (s *Sound) SeekData(pcm C.uint) error {
	//FMOD_RESULT F_API FMOD_Sound_SeekData(FMOD_SOUND *sound, unsigned int pcm);
	return ErrNoImpl
}

// Moves the sound from its existing SoundGroup to the specified sound group.
// By default a sound is located in the 'master sound group'. This can be retrieved with "System.MasterSoundGroup".
// Putting a sound in a sound group (or just using the master sound group) allows for functionality like limiting a group of sounds to a certain number of playbacks (see "SoundGroup.SetMaxAudible").
func (s *Sound) SetSoundGroup(soundgroup *SoundGroup) error {
	res := C.FMOD_Sound_SetSoundGroup(s.cptr, soundgroup.cptr)
	return errs[res]
}

// Retrieves the sound's current soundgroup.
func (s *Sound) SoundGroup() (*SoundGroup, error) {
	var soundgroup SoundGroup
	res := C.FMOD_Sound_GetSoundGroup(s.cptr, &soundgroup.cptr)
	return &soundgroup, errs[res]
}

/*
   Synchronization point API.  These points can come from markers embedded in wav files, and can also generate channel callbacks.
*/

// NOTE: Not implement yet
// Retrieves the number of sync points stored within a sound. These points can be user generated or can come from a wav file with embedded markers.
// In sound forge, a marker can be added a wave file by clicking on the timeline / ruler, and right clicking then selecting 'Insert Marker/Region'.
// Riff wrapped mp3 files are also supported.
func (s *Sound) NumSyncPoints(numsyncpoints *C.int) error {
	//FMOD_RESULT F_API FMOD_Sound_GetNumSyncPoints           (FMOD_SOUND *sound, int *numsyncpoints);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Retrieve a handle to a sync point. These points can be user generated or can come from a wav file with embedded markers.
//
// index: Index of the sync point to retrieve. Use "Sound.NumSyncPoints" to determine the number of syncpoints.
//
// In sound forge, a marker can be added a wave file by clicking on the timeline / ruler, and right clicking then selecting 'Insert Marker/Region'.
// Riff wrapped mp3 files are also supported.
func (s *Sound) SyncPoint(index C.int, point **C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_GetSyncPoint               (FMOD_SOUND *sound, int index, FMOD_SYNCPOINT **point);
	return ErrNoImpl
}

// NOTE: Not implement yet
// TODO: add more docs
// Retrieves information on an embedded sync point. These points can be user generated or can come from a wav file with embedded markers.
// In sound forge, a marker can be added a wave file by clicking on the timeline / ruler, and right clicking then selecting 'Insert Marker/Region'.
// Riff wrapped mp3 files are also supported.
func (s *Sound) SyncPointInfo(point *C.FMOD_SYNCPOINT, name *C.char, namelen C.int, offset *C.uint, offsettype C.FMOD_TIMEUNIT) error {
	//FMOD_RESULT F_API FMOD_Sound_GetSyncPointInfo           (FMOD_SOUND *sound, FMOD_SYNCPOINT *point, char *name, int namelen, unsigned int *offset, FMOD_TIMEUNIT offsettype);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Adds a sync point at a specific time within the sound. These points can be user generated or can come from a wav file with embedded markers.
// offset: Offset in units specified by offsettype to add the callback syncpoint for a sound.
//
// offsettype: offset type to describe the offset provided. Could be PCM samples or milliseconds for example.
//
// name: A name character string to be stored with the sync point. This will be provided via the sync point callback.
//
// In sound forge, a marker can be added a wave file by clicking on the timeline / ruler, and right clicking then selecting 'Insert Marker/Region'.
// Riff wrapped mp3 files are also supported.
func (s *Sound) AddSyncPoint(offset C.uint, offsettype C.FMOD_TIMEUNIT, name *C.char, point **C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_AddSyncPoint               (FMOD_SOUND *sound, unsigned int offset, FMOD_TIMEUNIT offsettype, const char *name, FMOD_SYNCPOINT **point);
	return ErrNoImpl
}

// NOTE: Not implement yet
// Deletes a syncpoint within the sound. These points can be user generated or can come from a wav file with embedded markers.
// In sound forge, a marker can be added a wave file by clicking on the timeline / ruler, and right clicking then selecting 'Insert Marker/Region'.
// Riff wrapped mp3 files are also supported.
func (s *Sound) DeleteSyncPoint(point *C.FMOD_SYNCPOINT) error {
	//FMOD_RESULT F_API FMOD_Sound_DeleteSyncPoint            (FMOD_SOUND *sound, FMOD_SYNCPOINT *point);
	return ErrNoImpl
}

/*
   Functions also in Channel class but here they are the 'default' to save having to change it in Channel all the time.
*/

// TODO: add more docs
// Sets or alters the mode of a sound.
// When calling this function, note that it will only take effect when the sound is played again with "System.PlaySound".
// Consider this mode the 'default mode' for when the sound plays, not a mode that will suddenly change all currently playing instances of this sound.
func (s *Sound) SetMode(mode Mode) error {
	res := C.FMOD_Sound_SetMode(s.cptr, C.FMOD_MODE(mode))
	return errs[res]
}

// Retrieves the mode bits set by the codec and the user when opening the sound.
func (s *Sound) Mode() (Mode, error) {
	var mode C.FMOD_MODE
	res := C.FMOD_Sound_GetMode(s.cptr, &mode)
	return Mode(mode), errs[res]
}

// Sets a sound, by default, to loop a specified number of times before stopping if its mode is set to LOOP_NORMAL or LOOP_BIDI.
//
// Issues with streamed audio. (Sounds created with with "System.CreateStream" or FMOD_CREATESTREAM).
// When changing the loop count, sounds created with "System.CreateStream" or FMOD_CREATESTREAM may already have been pre-buffered and
// executed their loop logic ahead of time, before this call was even made.
// This is dependant on the size of the sound versus the size of the stream decode buffer. See FMOD_CREATESOUNDEXINFO.
// If this happens, you may need to reflush the stream buffer. To do this, you can call "Channel.SetPosition" which forces a reflush of the stream buffer.
// Note this will usually only happen if you have sounds or looppoints that are smaller than the stream decode buffer size. Otherwise you will not normally encounter any problems.
func (s *Sound) SetLoopCount(loopcount int) error {
	res := C.FMOD_Sound_SetLoopCount(s.cptr, C.int(loopcount))
	return errs[res]
}

// Retrieves the current loop count value for the specified sound.
//
// loopcount: A variable that receives the number of times a sound will loop by default before stopping. 0 = oneshot. 1 = loop once then stop. -1 = loop forever. Default = -1
//
// Unlike the channel loop count function, this function simply returns the value set with "Sound.SetLoopCount".
// It does not decrement as it plays (especially seeing as one sound can be played multiple times).
func (s *Sound) LoopCount() (int, error) {
	var loopcount C.int
	res := C.FMOD_Sound_GetLoopCount(s.cptr, &loopcount)
	return int(loopcount), errs[res]
}

// Sets the loop points within a sound.
//
// loopstart: The loop start point. This point in time is played, so it is inclusive.
//
// loopstarttype: The time format used for the loop start point. See "TimeUnit".
//
//loopend: The loop end point. This point in time is played, so it is inclusive.
//
// loopendtype: The time format used for the loop end point. See "TimeUnit".
//
// If a sound was 44100 samples long and you wanted to loop the whole sound, loopstart would be 0, and loopend would be 44099, not 44100.
// You wouldn't use milliseconds in this case because they are not sample accurate.
// If loop end is smaller or equal to loop start, it will result in an error.
// If loop start or loop end is larger than the length of the sound, it will result in an error.
//
// Issues with streamed audio. (Sounds created with with "System.CreateStream" or FMOD_CREATESTREAM).
// When changing the loop points, sounds created with "System.CreateStream" or FMOD_CREATESTREAM may already have been pre-buffered and
// executed their loop logic ahead of time, before this call was even made.
// This is dependant on the size of the sound versus the size of the stream decode buffer. See FMOD_CREATESOUNDEXINFO.
// If this happens, you may need to reflush the stream buffer. To do this, you can call "Channel.SetPosition" which forces a reflush of the stream buffer.
// Note this will usually only happen if you have sounds or looppoints that are smaller than the stream decode buffer size.
// Otherwise you will not normally encounter any problems.
func (s *Sound) SetLoopPoints(loopstart uint32, loopstarttype TimeUnit, loopend uint32, loopendtype TimeUnit) error {
	res := C.FMOD_Sound_SetLoopPoints(s.cptr, C.uint(loopstart), C.FMOD_TIMEUNIT(loopstarttype), C.uint(loopend), C.FMOD_TIMEUNIT(loopendtype))
	return errs[res]
}

// Retrieves the loop points for a sound.
func (s *Sound) LoopPoints(loopstarttype, loopendtype TimeUnit) (uint32, uint32, error) {
	var loopstart, loopend C.uint
	res := C.FMOD_Sound_GetLoopPoints(s.cptr, &loopstart, C.FMOD_TIMEUNIT(loopstarttype), &loopend, C.FMOD_TIMEUNIT(loopendtype))
	return uint32(loopstart), uint32(loopend), errs[res]
}

/*
   For MOD/S3M/XM/IT/MID sequenced formats only.
*/

// Gets the number of music channels inside a MOD/S3M/XM/IT/MIDI file.
func (s *Sound) MusicNumChannels() (int, error) {
	var numchannels C.int
	res := C.FMOD_Sound_GetMusicNumChannels(s.cptr, &numchannels)
	return int(numchannels), errs[res]
}

// Sets the volume of a MOD/S3M/XM/IT/MIDI music channel volume.
//
// channel: MOD/S3M/XM/IT/MIDI music subchannel to set a linear volume for.
//
// volume: Volume of the channel from 0.0 to 1.0. Default = 1.0.
//
// Use "Sound.MusicNumChannels" to get the maximum number of music channels in the song.
func (s *Sound) SetMusicChannelVolume(channel int, volume float64) error {
	res := C.FMOD_Sound_SetMusicChannelVolume(s.cptr, C.int(channel), C.float(volume))
	return errs[res]
}

// Retrieves the volume of a MOD/S3M/XM/IT/MIDI music channel volume.
// Use "Sound.MusicNumChannels" to get the maximum number of music channels in the song.
func (s *Sound) MusicChannelVolume(channel int) (float64, error) {
	var volume C.float
	res := C.FMOD_Sound_GetMusicChannelVolume(s.cptr, C.int(channel), &volume)
	return float64(volume), errs[res]
}

// Sets the relative speed of MOD/S3M/XM/IT/MIDI music.
//
// speed: Relative speed of the song from 0.01 to 100.0. 0.5 = half speed, 2.0 = double speed. Default = 1.0.
//
// Setting a speed outside the bounds of 0.01 to 100.0 will not return an error, it will clamp the value.
func (s *Sound) SetMusicSpeed(speed float64) error {
	res := C.FMOD_Sound_SetMusicSpeed(s.cptr, C.float(speed))
	return errs[res]
}

// Retrieves the relative speed of MOD/S3M/XM/IT/MIDI music.
func (s *Sound) MusicSpeed() (float64, error) {
	var speed C.float
	res := C.FMOD_Sound_GetMusicSpeed(s.cptr, &speed)
	return float64(speed), errs[res]
}

/*
   Userdata set/get.
*/

// Sets a user value that the Sound object will store internally. Can be retrieved with "Sound.UserData".
// This function is primarily used in case the user wishes to 'attach' data to an FMOD object.
// It can be useful if an FMOD callback passes an object of this type as a parameter, and the user does not know which object it is
// (if many of these types of objects exist). Using "Sound.UserData" would help in the identification of the object.
func (s *Sound) SetUserData(userdata interface{}) error {
	data := *(*[]*C.char)(unsafe.Pointer(&userdata))
	res := C.FMOD_Sound_SetUserData(s.cptr, unsafe.Pointer(&data))
	return errs[res]
}

// Retrieves the user value that that was set by calling the "Sound.SetUserData" function.
func (s *Sound) UserData() (interface{}, error) {
	var userdata *interface{}
	cUserdata := unsafe.Pointer(userdata)
	res := C.FMOD_Sound_GetUserData(s.cptr, &cUserdata)
	return *(*interface{})(cUserdata), errs[res]
}
