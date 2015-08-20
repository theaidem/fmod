package lowlevel

/*
#include <fmod_common.h>
*/
import "C"

const VERSION = C.FMOD_VERSION

// These callback types are used with System::setCallback.
//
// Each callback has commanddata parameters passed as void* unique to the type of callback.
// See reference to "SYSTEM_CALLBACK" to determine what they might mean for each type of callback.
//
// Note! Using "SYSTEM_CALLBACK_DEVICELISTCHANGED" (Windows only) will disable any automated device ejection/insertion handling by FMOD.
// Use this callback to control the behaviour yourself.
//
// Note! Using "SYSTEM_CALLBACK_DEVICELISTCHANGED" (on Mac only) requires the application to be running an event loop which will allow external changes to device list to be detected by FMOD.
//
// Note! The 'system' object pointer will be null for "SYSTEM_CALLBACK_MEMORYALLOCATIONFAILED" callback.
type SystemCallbackType C.FMOD_SYSTEM_CALLBACK_TYPE

const (

	// Called from System.Update when the enumerated list of devices has changed.
	SYSTEM_CALLBACK_DEVICELISTCHANGED SystemCallbackType = C.FMOD_SYSTEM_CALLBACK_DEVICELISTCHANGED

	// Called from System.Update when an output device has been lost due to control panel parameter changes and FMOD cannot automatically recover.
	SYSTEM_CALLBACK_DEVICELOST = C.FMOD_SYSTEM_CALLBACK_DEVICELOST

	// Called directly when a memory allocation fails somewhere in FMOD.
	// (NOTE - 'system' will be NULL in this callback type.)
	SYSTEM_CALLBACK_MEMORYALLOCATIONFAILED = C.FMOD_SYSTEM_CALLBACK_MEMORYALLOCATIONFAILED

	// Called directly when a thread is created.
	SYSTEM_CALLBACK_THREADCREATED = C.FMOD_SYSTEM_CALLBACK_THREADCREATED

	// Called when a bad connection was made with DSP.AddInput.
	// Usually called from mixer thread because that is where the connections are made.
	SYSTEM_CALLBACK_BADDSPCONNECTION = C.FMOD_SYSTEM_CALLBACK_BADDSPCONNECTION

	// Called each tick before a mix update happens.
	SYSTEM_CALLBACK_PREMIX = C.FMOD_SYSTEM_CALLBACK_PREMIX

	// Called each tick after a mix update happens.
	SYSTEM_CALLBACK_POSTMIX = C.FMOD_SYSTEM_CALLBACK_POSTMIX

	// Called when each API function returns an error code, including delayed async functions.
	SYSTEM_CALLBACK_ERROR = C.FMOD_SYSTEM_CALLBACK_ERROR

	// Called each tick in mix update after clocks have been updated before the main mix occurs.
	SYSTEM_CALLBACK_MIDMIX = C.FMOD_SYSTEM_CALLBACK_MIDMIX

	// Called directly when a thread is destroyed.
	SYSTEM_CALLBACK_THREADDESTROYED = C.FMOD_SYSTEM_CALLBACK_THREADDESTROYED

	// Called at start of System.Update function.
	SYSTEM_CALLBACK_PREUPDATE = C.FMOD_SYSTEM_CALLBACK_PREUPDATE

	// Called at end of System.Update function.
	SYSTEM_CALLBACK_POSTUPDATE = C.FMOD_SYSTEM_CALLBACK_POSTUPDATE
)

// TODO: add more docs
// Structure describing a globally unique identifier.
type Guid C.FMOD_GUID

type CreatesSoundExInfo C.FMOD_CREATESOUNDEXINFO

type SystemCallback C.FMOD_SYSTEM_CALLBACK

type InitFlags C.FMOD_INITFLAGS

const (

	// Initialize normally
	INIT_NORMAL InitFlags = C.FMOD_INIT_NORMAL

	// No stream thread is created internally.
	// Streams are driven from System.Update.  Mainly used with non-realtime outputs.
	INIT_STREAM_FROM_UPDATE = C.FMOD_INIT_STREAM_FROM_UPDATE

	// Win/PS3/Xbox 360 Only - FMOD Mixer thread is woken up to do a mix when System.Update is called rather than waking periodically on its own timer.
	INIT_MIX_FROM_UPDATE = C.FMOD_INIT_MIX_FROM_UPDATE

	// FMOD will treat +X as right, +Y as up and +Z as backwards (towards you).
	INIT_3D_RIGHTHANDED = C.FMOD_INIT_3D_RIGHTHANDED

	// All FMOD_3D based voices will add a software lowpass filter effect into the DSP chain which is automatically used when Channel.Set3DOcclusion is used or the geometry API.
	// This also causes sounds to sound duller when the sound goes behind the listener, as a fake HRTF style effect.
	// Use System.SetAdvancedSettings to disable or adjust cutoff frequency for this feature.
	INIT_CHANNEL_LOWPASS = C.FMOD_INIT_CHANNEL_LOWPASS

	// All FMOD_3D based voices will add a software lowpass and highpass filter effect into the DSP chain which will act as a distance-automated bandpass filter.
	// Use System.SetAdvancedSettings to adjust the center frequency.
	INIT_CHANNEL_DISTANCEFILTER = C.FMOD_INIT_CHANNEL_DISTANCEFILTER

	// Enable TCP/IP based host which allows FMOD Designer or FMOD Profiler to connect to it, and view memory, CPU and the DSP network graph in real-time.
	INIT_PROFILE_ENABLE = C.FMOD_INIT_PROFILE_ENABLE

	// Any sounds that are 0 volume will go virtual and not be processed except for having their positions updated virtually.
	// Use System.SetAdvancedSettings to adjust what volume besides zero to switch to virtual at.
	INIT_VOL0_BECOMES_VIRTUAL = C.FMOD_INIT_VOL0_BECOMES_VIRTUAL

	// With the geometry engine, only process the closest polygon rather than accumulating all polygons the sound to listener line intersects.
	INIT_GEOMETRY_USECLOSEST = C.FMOD_INIT_GEOMETRY_USECLOSEST

	// When using FMOD_SPEAKERMODE_5POINT1 with a stereo output device, use the Dolby Pro Logic II downmix algorithm instead of the SRS Circle Surround algorithm.
	INIT_PREFER_DOLBY_DOWNMIX = C.FMOD_INIT_PREFER_DOLBY_DOWNMIX

	// Disables thread safety for API calls.
	// Only use this if FMOD low level is being called from a single thread, and if Studio API is not being used!
	INIT_THREAD_UNSAFE = C.FMOD_INIT_THREAD_UNSAFE

	// Slower, but adds level metering for every single DSP unit in the graph.
	// Use DSP.SetMeteringEnabled to turn meters off individually.
	INIT_PROFILE_METER_ALL = C.FMOD_INIT_PROFILE_METER_ALL
)

type OutputType C.FMOD_OUTPUTTYPE

const (
	// Picks the best output mode for the platform. This is the default.
	OUTPUTTYPE_AUTODETECT OutputType = C.FMOD_OUTPUTTYPE_AUTODETECT

	// All - 3rd party plugin, unknown.
	// This is for use with System.Output only.
	OUTPUTTYPE_UNKNOWN = C.FMOD_OUTPUTTYPE_UNKNOWN

	// All - Perform all mixing but discard the final output.
	OUTPUTTYPE_NOSOUND = C.FMOD_OUTPUTTYPE_NOSOUND

	// All - Writes output to a .wav file.
	OUTPUTTYPE_WAVWRITER = C.FMOD_OUTPUTTYPE_WAVWRITER

	// All - Non-realtime version of FMOD_OUTPUTTYPE_NOSOUND.
	// User can drive mixer with System::update at whatever rate they want.
	OUTPUTTYPE_NOSOUND_NRT = C.FMOD_OUTPUTTYPE_NOSOUND_NRT

	// All - Non-realtime version of FMOD_OUTPUTTYPE_WAVWRITER.
	// User can drive mixer with System::update at whatever rate they want.
	OUTPUTTYPE_WAVWRITER_NRT = C.FMOD_OUTPUTTYPE_WAVWRITER_NRT

	// Win                  - Direct Sound.                        (Default on Windows XP and below)
	OUTPUTTYPE_DSOUND = C.FMOD_OUTPUTTYPE_DSOUND

	// Win                  - Windows Multimedia.
	OUTPUTTYPE_WINMM = C.FMOD_OUTPUTTYPE_WINMM

	// Win/WinStore/XboxOne - Windows Audio Session API.
	// (Default on Windows Vista and above, Xbox One and Windows Store Applications)
	OUTPUTTYPE_WASAPI = C.FMOD_OUTPUTTYPE_WASAPI

	// Win                  - Low latency ASIO 2.0.
	OUTPUTTYPE_ASIO = C.FMOD_OUTPUTTYPE_ASIO

	// Linux                - Pulse Audio.
	// (Default on Linux if available)
	OUTPUTTYPE_PULSEAUDIO = C.FMOD_OUTPUTTYPE_PULSEAUDIO

	// Linux                - Advanced Linux Sound Architecture.
	//(Default on Linux if PulseAudio isn't available)
	OUTPUTTYPE_ALSA = C.FMOD_OUTPUTTYPE_ALSA

	// Mac/iOS              - Core Audio.
	// (Default on Mac and iOS)
	OUTPUTTYPE_COREAUDIO = C.FMOD_OUTPUTTYPE_COREAUDIO

	// Xbox 360             - XAudio.
	// (Default on Xbox 360)
	OUTPUTTYPE_XBOX360 = C.FMOD_OUTPUTTYPE_XBOX360

	// PS3                  - Audio Out.
	// (Default on PS3)
	OUTPUTTYPE_PS3 = C.FMOD_OUTPUTTYPE_PS3

	// Android              - Java Audio Track.
	// (Default on Android 2.2 and below)
	OUTPUTTYPE_AUDIOTRACK = C.FMOD_OUTPUTTYPE_AUDIOTRACK

	// Android              - OpenSL ES.
	// (Default on Android 2.3 and above)
	OUTPUTTYPE_OPENSL = C.FMOD_OUTPUTTYPE_OPENSL

	// Wii U                - AX.
	// (Default on Wii U)
	OUTPUTTYPE_WIIU = C.FMOD_OUTPUTTYPE_WIIU

	// PS4/PSVita           - Audio Out.
	//(Default on PS4 and PS Vita)
	OUTPUTTYPE_AUDIOOUT = C.FMOD_OUTPUTTYPE_AUDIOOUT

	// Maximum number of output types supported.
	OUTPUTTYPE_MAX = C.FMOD_OUTPUTTYPE_MAX

	// Makes sure this enum is signed 32bit.
	OUTPUTTYPE_FORCEINT = C.FMOD_OUTPUTTYPE_FORCEINT
)

type SpeakerMode C.FMOD_SPEAKERMODE

const (
	// Default speaker mode based on operating system/output mode.
	// Windows = control panel setting, Xbox = 5.1, PS3 = 7.1 etc.
	SPEAKERMODE_DEFAULT SpeakerMode = C.FMOD_SPEAKERMODE_DEFAULT

	// There is no specific speakermode.  Sound channels are mapped in order of input to output.
	// Use System.SetSoftwareFormat to specify speaker count. See remarks for more information.
	SPEAKERMODE_RAW = C.FMOD_SPEAKERMODE_RAW

	// The speakers are monaural.
	SPEAKERMODE_MONO = C.FMOD_SPEAKERMODE_MONO

	// The speakers are stereo.
	SPEAKERMODE_STEREO = C.FMOD_SPEAKERMODE_STEREO

	// 4 speaker setup.    This includes front left, front right, surround left, surround right.
	SPEAKERMODE_QUAD = C.FMOD_SPEAKERMODE_QUAD

	// 5 speaker setup.    This includes front left, front right, center, surround left, surround right.
	SPEAKERMODE_SURROUND = C.FMOD_SPEAKERMODE_SURROUND

	// 5.1 speaker setup.  This includes front left, front right, center, surround left, surround right and an LFE speaker.
	SPEAKERMODE_5POINT1 = C.FMOD_SPEAKERMODE_5POINT1

	// 7.1 speaker setup.  This includes front left, front right, center, surround left, surround right, back left, back right and an LFE speaker.
	SPEAKERMODE_7POINT1 = C.FMOD_SPEAKERMODE_7POINT1

	// Maximum number of speaker modes supported.
	SPEAKERMODE_MAX = C.FMOD_SPEAKERMODE_MAX

	// Makes sure this enum is signed 32bit.
	SPEAKERMODE_FORCEINT = C.FMOD_SPEAKERMODE_FORCEINT
)

type PluginType C.FMOD_PLUGINTYPE

const (
	// The plugin type is an output module.
	// FMOD mixed audio will play through one of these devices
	PLUGINTYPE_OUTPUT PluginType = C.FMOD_PLUGINTYPE_OUTPUT

	// The plugin type is a file format codec.
	// FMOD will use these codecs to load file formats for playback.
	PLUGINTYPE_CODEC = C.FMOD_PLUGINTYPE_CODEC

	// The plugin type is a DSP unit.
	// FMOD will use these plugins as part of its DSP network to apply effects to output or generate sound in realtime.
	PLUGINTYPE_DSP = C.FMOD_PLUGINTYPE_DSP

	// Maximum number of plugin types supported.
	PLUGINTYPE_MAX = C.FMOD_PLUGINTYPE_MAX

	// Makes sure this enum is signed 32bit.
	PLUGINTYPE_FORCEINT = C.FMOD_PLUGINTYPE_FORCEINT
)

type Speaker C.FMOD_SPEAKER

const (
	SPEAKER_FRONT_LEFT     Speaker = C.FMOD_SPEAKER_FRONT_LEFT
	SPEAKER_FRONT_RIGHT            = C.FMOD_SPEAKER_FRONT_RIGHT
	SPEAKER_FRONT_CENTER           = C.FMOD_SPEAKER_FRONT_CENTER
	SPEAKER_LOW_FREQUENCY          = C.FMOD_SPEAKER_LOW_FREQUENCY
	SPEAKER_SURROUND_LEFT          = C.FMOD_SPEAKER_SURROUND_LEFT
	SPEAKER_SURROUND_RIGHT         = C.FMOD_SPEAKER_SURROUND_RIGHT
	SPEAKER_BACK_LEFT              = C.FMOD_SPEAKER_BACK_LEFT
	SPEAKER_BACK_RIGHT             = C.FMOD_SPEAKER_BACK_RIGHT

	// Maximum number of speaker types supported.
	SPEAKER_MAX = C.FMOD_SPEAKER_MAX

	// Makes sure this enum is signed 32bit.
	SPEAKER_FORCEINT = C.FMOD_SPEAKER_FORCEINT
)

type Mode C.FMOD_MODE

const (
	// Default for all modes listed below.
	// LOOP_OFF, MODE_2D, MODE_3D_WORLDRELATIVE, MODE_3D_INVERSEROLLOFF
	MODE_DEFAULT Mode = C.FMOD_DEFAULT

	// For non looping sounds. (DEFAULT).
	// Overrides MODE_LOOP_NORMAL / MODE_LOOP_BIDI.
	MODE_LOOP_OFF = C.FMOD_LOOP_OFF

	// For forward looping sounds.
	MODE_LOOP_NORMAL = C.FMOD_LOOP_NORMAL

	// For bidirectional looping sounds. (only works on software mixed static sounds).
	MODE_LOOP_BIDI = C.FMOD_LOOP_BIDI

	// Ignores any 3d processing. (DEFAULT).
	MODE_2D = C.FMOD_2D

	// Makes the sound positionable in 3D.
	// Overrides MODE_2D.
	MODE_3D = C.FMOD_3D

	// Decompress at runtime, streaming from the source provided (ie from disk).
	// Overrides MODE_CREATESAMPLE and MODE_CREATECOMPRESSEDSAMPLE.
	// Note a stream can only be played once at a time due to a stream only having 1 stream buffer and file handle.
	// Open multiple streams to have them play concurrently.
	MODE_CREATESTREAM = C.FMOD_CREATESTREAM

	// Decompress at loadtime, decompressing or decoding whole file into memory as the target sample format (ie PCM).
	// Fastest for playback and most flexible.
	MODE_CREATESAMPLE = C.FMOD_CREATESAMPLE

	// Load MP2/MP3/IMAADPCM/Vorbis/AT9 or XMA into memory and leave it compressed.
	// Vorbis/AT9 encoding only supported in the FSB file format.
	// During playback the FMOD software mixer will decode it in realtime as a 'compressed sample'.
	// Overrides MODE_CREATESAMPLE.
	// If the sound data is not one of the supported formats, it will behave as if it was created with MODE_CREATESAMPLE and decode the sound into PCM.
	MODE_CREATECOMPRESSEDSAMPLE = C.FMOD_CREATECOMPRESSEDSAMPLE

	// Opens a user created static sample or stream.
	// Use FMOD_CREATESOUNDEXINFO to specify format and/or read callbacks.
	// If a user created 'sample' is created with no read callback, the sample will be empty.
	// Use Sound.Lock and Sound.Unlock to place sound data into the sound if this is the case.
	MODE_OPENUSER = C.FMOD_OPENUSER

	// "name_or_data" will be interpreted as a pointer to memory instead of filename for creating sounds.
	// Use FMOD_CREATESOUNDEXINFO to specify length.
	// If used with FMOD_CREATESAMPLE or MODE_CREATECOMPRESSEDSAMPLE, FMOD duplicates the memory into its own buffers.
	// Your own buffer can be freed after open.
	// If used with FMOD_CREATESTREAM, FMOD will stream out of the buffer whose pointer you passed in.
	// In this case, your own buffer should not be freed until you have finished with and released the stream.
	MODE_OPENMEMORY = C.FMOD_OPENMEMORY

	// "name_or_data" will be interpreted as a pointer to memory instead of filename for creating sounds.
	// Use FMOD_CREATESOUNDEXINFO to specify length.
	// This differs to MODE_OPENMEMORY in that it uses the memory as is, without duplicating the memory into its own buffers.
	// Cannot be freed after open, only after Sound.Release.
	// Will not work if the data is compressed and MODE_CREATECOMPRESSEDSAMPLE is not used.
	MODE_OPENMEMORY_POINT = C.FMOD_OPENMEMORY_POINT

	// Will ignore file format and treat as raw pcm.
	// Use FMOD_CREATESOUNDEXINFO to specify format.
	// Requires at least defaultfrequency, numchannels and format to be specified before it will open.
	// Must be little endian data.
	MODE_OPENRAW = C.FMOD_OPENRAW

	// Just open the file, dont prebuffer or read.
	// Good for fast opens for info, or when Sound.ReadData is to be used.
	MODE_OPENONLY = C.FMOD_OPENONLY

	// For System.CreateSound - for accurate Sound.Length/Channel.SetPosition on VBR MP3, and MOD/S3M/XM/IT/MIDI files.
	// Scans file first, so takes longer to open. MODE_OPENONLY does not affect this.
	MODE_ACCURATETIME = C.FMOD_ACCURATETIME

	// For corrupted / bad MP3 files.
	// This will search all the way through the file until it hits a valid MPEG header.
	// Normally only searches for 4k.
	MODE_MPEGSEARCH = C.FMOD_MPEGSEARCH

	// For opening sounds and getting streamed subsounds (seeking) asyncronously.
	// Use Sound.GetOpenState to poll the state of the sound as it opens or retrieves the subsound in the background.
	MODE_NONBLOCKING = C.FMOD_NONBLOCKING

	// Unique sound, can only be played one at a time
	MODE_UNIQUE = C.FMOD_UNIQUE

	// Make the sound's position, velocity and orientation relative to the listener.
	MODE_3D_HEADRELATIVE = C.FMOD_3D_HEADRELATIVE

	// Make the sound's position, velocity and orientation absolute (relative to the world). (DEFAULT)
	MODE_3D_WORLDRELATIVE = C.FMOD_3D_WORLDRELATIVE

	// This sound will follow the inverse rolloff model where mindistance = full volume, maxdistance = where sound stops attenuating, and rolloff is fixed according to the global rolloff factor.  (DEFAULT)
	MODE_3D_INVERSEROLLOFF = C.FMOD_3D_INVERSEROLLOFF

	// This sound will follow a linear rolloff model where mindistance = full volume, maxdistance = silence.
	MODE_3D_LINEARROLLOFF = C.FMOD_3D_LINEARROLLOFF

	// This sound will follow a linear-square rolloff model where mindistance = full volume, maxdistance = silence.
	MODE_3D_LINEARSQUAREROLLOFF = C.FMOD_3D_LINEARSQUAREROLLOFF

	// This sound will follow the inverse rolloff model at distances close to mindistance and a linear-square rolloff close to maxdistance.
	MODE_3D_INVERSETAPEREDROLLOFF = C.FMOD_3D_INVERSETAPEREDROLLOFF

	// This sound will follow a rolloff model defined by Sound.Set3DCustomRolloff / Channel.Set3DCustomRolloff.
	MODE_3D_CUSTOMROLLOFF = C.FMOD_3D_CUSTOMROLLOFF

	// Is not affect by geometry occlusion.  If not specified in Sound.SetMode, or Channel.SetMode, the flag is cleared and it is affected by geometry again.
	MODE_3D_IGNOREGEOMETRY = C.FMOD_3D_IGNOREGEOMETRY

	// Skips id3v2/asf/etc tag checks when opening a sound, to reduce seek/read overhead when opening files (helps with CD performance).
	MODE_IGNORETAGS = C.FMOD_IGNORETAGS

	// Removes some features from samples to give a lower memory overhead, like Sound.Name.  See remarks.
	MODE_LOWMEM = C.FMOD_LOWMEM

	// Load sound into the secondary RAM of supported platform. On PS3, sounds will be loaded into RSX/VRAM.
	MODE_LOADSECONDARYRAM = C.FMOD_LOADSECONDARYRAM

	// For sounds that start virtual (due to being quiet or low importance), instead of swapping back to audible, and playing at the correct offset according to time, this flag makes the sound play from the start.
	MODE_VIRTUAL_PLAYFROMSTART = C.FMOD_VIRTUAL_PLAYFROMSTART
)

type DSPType C.FMOD_DSP_TYPE

const (

	// This unit was created via a non FMOD plugin so has an unknown purpose.
	DSP_TYPE_UNKNOWN DSPType = C.FMOD_DSP_TYPE_UNKNOWN

	// This unit does nothing but take inputs and mix them together then feed the result to the soundcard unit.
	DSP_TYPE_MIXER = C.FMOD_DSP_TYPE_MIXER

	// This unit generates sine/square/saw/triangle or noise tones.
	DSP_TYPE_OSCILLATOR = C.FMOD_DSP_TYPE_OSCILLATOR

	// This unit filters sound using a high quality, resonant lowpass filter algorithm but consumes more CPU time.
	DSP_TYPE_LOWPASS = C.FMOD_DSP_TYPE_LOWPASS

	// This unit filters sound using a resonant lowpass filter algorithm that is used in Impulse Tracker, but with limited cutoff range (0 to 8060hz).
	DSP_TYPE_ITLOWPASS = C.FMOD_DSP_TYPE_ITLOWPASS

	// This unit filters sound using a resonant highpass filter algorithm.
	DSP_TYPE_HIGHPASS = C.FMOD_DSP_TYPE_HIGHPASS

	// This unit produces an echo on the sound and fades out at the desired rate.
	DSP_TYPE_ECHO = C.FMOD_DSP_TYPE_ECHO

	// This unit pans and scales the volume of a unit.
	DSP_TYPE_FADER = C.FMOD_DSP_TYPE_FADER

	// This unit produces a flange effect on the sound.
	DSP_TYPE_FLANGE = C.FMOD_DSP_TYPE_FLANGE

	// This unit distorts the sound.
	DSP_TYPE_DISTORTION = C.FMOD_DSP_TYPE_DISTORTION

	// This unit normalizes or amplifies the sound to a certain level.
	DSP_TYPE_NORMALIZE = C.FMOD_DSP_TYPE_NORMALIZE

	// This unit limits the sound to a certain level
	DSP_TYPE_LIMITER = C.FMOD_DSP_TYPE_LIMITER

	// This unit attenuates or amplifies a selected frequency range.
	DSP_TYPE_PARAMEQ = C.FMOD_DSP_TYPE_PARAMEQ

	// This unit bends the pitch of a sound without changing the speed of playback.
	DSP_TYPE_PITCHSHIFT = C.FMOD_DSP_TYPE_PITCHSHIFT

	// This unit produces a chorus effect on the sound.
	DSP_TYPE_CHORUS = C.FMOD_DSP_TYPE_CHORUS

	// This unit allows the use of Steinberg VST plugins
	DSP_TYPE_VSTPLUGIN = C.FMOD_DSP_TYPE_VSTPLUGIN

	// This unit allows the use of Nullsoft Winamp plugins
	DSP_TYPE_WINAMPPLUGIN = C.FMOD_DSP_TYPE_WINAMPPLUGIN

	// This unit produces an echo on the sound and fades out at the desired rate as is used in Impulse Tracker.
	DSP_TYPE_ITECHO = C.FMOD_DSP_TYPE_ITECHO

	// This unit implements dynamic compression (linked/unlinked multichannel, wideband)
	DSP_TYPE_COMPRESSOR = C.FMOD_DSP_TYPE_COMPRESSOR

	// This unit implements SFX reverb
	DSP_TYPE_SFXREVERB = C.FMOD_DSP_TYPE_SFXREVERB

	// This unit filters sound using a simple lowpass with no resonance, but has flexible cutoff and is fast.
	DSP_TYPE_LOWPASS_SIMPLE = C.FMOD_DSP_TYPE_LOWPASS_SIMPLE

	// This unit produces different delays on individual channels of the sound.
	DSP_TYPE_DELAY = C.FMOD_DSP_TYPE_DELAY

	// This unit produces a tremolo / chopper effect on the sound.
	DSP_TYPE_TREMOLO = C.FMOD_DSP_TYPE_TREMOLO

	// Unsupported / Deprecated.
	DSP_TYPE_LADSPAPLUGIN = C.FMOD_DSP_TYPE_LADSPAPLUGIN

	// This unit sends a copy of the signal to a return DSP anywhere in the DSP tree.
	DSP_TYPE_SEND = C.FMOD_DSP_TYPE_SEND

	// This unit receives signals from a number of send DSPs.
	DSP_TYPE_RETURN = C.FMOD_DSP_TYPE_RETURN

	// This unit filters sound using a simple highpass with no resonance, but has flexible cutoff and is fast.
	DSP_TYPE_HIGHPASS_SIMPLE = C.FMOD_DSP_TYPE_HIGHPASS_SIMPLE

	// This unit pans the signal, possibly upmixing or downmixing as well.
	DSP_TYPE_PAN = C.FMOD_DSP_TYPE_PAN

	// This unit is a three-band equalizer.
	DSP_TYPE_THREE_EQ = C.FMOD_DSP_TYPE_THREE_EQ

	// This unit simply analyzes the signal and provides spectrum information back through getParameter.
	DSP_TYPE_FFT = C.FMOD_DSP_TYPE_FFT

	// This unit analyzes the loudness and true peak of the signal.
	DSP_TYPE_LOUDNESS_METER = C.FMOD_DSP_TYPE_LOUDNESS_METER

	// This unit tracks the envelope of the input/sidechain signal. Format to be publicly disclosed soon.
	DSP_TYPE_ENVELOPEFOLLOWER = C.FMOD_DSP_TYPE_ENVELOPEFOLLOWER

	// This unit implements convolution reverb.
	DSP_TYPE_CONVOLUTIONREVERB = C.FMOD_DSP_TYPE_CONVOLUTIONREVERB

	// Maximum number of pre-defined DSP types.
	DSP_TYPE_MAX = C.FMOD_DSP_TYPE_MAX

	// Makes sure this enum is signed 32bit.
	DSP_TYPE_FORCEINT = C.FMOD_DSP_TYPE_FORCEINT
)

type DSPReSampler C.FMOD_DSP_RESAMPLER

const (
	// Default interpolation method.
	// Currently equal to DSP_RESAMPLER_LINEAR.
	DSP_RESAMPLER_DEFAULT DSPReSampler = C.FMOD_DSP_RESAMPLER_DEFAULT

	// No interpolation.
	// High frequency aliasing hiss will be audible depending on the sample rate of the sound.
	DSP_RESAMPLER_NOINTERP = C.FMOD_DSP_RESAMPLER_NOINTERP

	// Linear interpolation (default method).
	// Fast and good quality, causes very slight lowpass effect on low frequency sounds.
	DSP_RESAMPLER_LINEAR = C.FMOD_DSP_RESAMPLER_LINEAR

	// Cubic interpolation.
	// Slower than linear interpolation but better quality.
	DSP_RESAMPLER_CUBIC = C.FMOD_DSP_RESAMPLER_CUBIC

	// 5 point spline interpolation.
	// Slowest resampling method but best quality.
	DSP_RESAMPLER_SPLINE = C.FMOD_DSP_RESAMPLER_SPLINE

	// Maximum number of resample methods supported.
	DSP_RESAMPLER_MAX = C.FMOD_DSP_RESAMPLER_MAX

	// Makes sure this enum is signed 32bit.
	DSP_RESAMPLER_FORCEINT = C.FMOD_DSP_RESAMPLER_FORCEINT
)

type TimeUnit C.FMOD_TIMEUNIT

const (

	// Milliseconds.
	TIMEUNIT_MS TimeUnit = C.FMOD_TIMEUNIT_MS

	// PCM samples, related to milliseconds * samplerate / 1000.
	TIMEUNIT_PCM = C.FMOD_TIMEUNIT_PCM

	// Bytes, related to PCM samples * channels * datawidth (ie 16bit = 2 bytes).
	TIMEUNIT_PCMBYTES = C.FMOD_TIMEUNIT_PCMBYTES

	// Raw file bytes of (compressed) sound data (does not include headers).
	// Only used by Sound.Length and Channel.Position.
	TIMEUNIT_RAWBYTES = C.FMOD_TIMEUNIT_RAWBYTES

	// Fractions of 1 PCM sample.  Unsigned int range 0 to 0xFFFFFFFF.
	// Used for sub-sample granularity for DSP purposes.
	TIMEUNIT_PCMFRACTION = C.FMOD_TIMEUNIT_PCMFRACTION

	// MOD/S3M/XM/IT.
	// Order in a sequenced module format.
	// Use Sound.Format to determine the PCM format being decoded to.
	TIMEUNIT_MODORDER = C.FMOD_TIMEUNIT_MODORDER

	// MOD/S3M/XM/IT.
	// Current row in a sequenced module format.
	// Sound.Length will return the number of rows in the currently playing or seeked to pattern.
	TIMEUNIT_MODROW = C.FMOD_TIMEUNIT_MODROW

	// MOD/S3M/XM/IT.
	// Current pattern in a sequenced module format.
	// Sound.Length will return the number of patterns in the song and Channel.Position will return the currently playing pattern.
	TIMEUNIT_MODPATTERN = C.FMOD_TIMEUNIT_MODPATTERN

	// Time value as seen by buffered stream.
	// This is always ahead of audible time, and is only used for processing.
	TIMEUNIT_BUFFERED = C.FMOD_TIMEUNIT_BUFFERED
)

type SoundType C.FMOD_SOUND_TYPE

const (
	// 3rd party / unknown plugin format.
	SOUND_TYPE_UNKNOWN SoundType = C.FMOD_SOUND_TYPE_UNKNOWN

	// AIFF.
	SOUND_TYPE_AIFF = C.FMOD_SOUND_TYPE_AIFF

	// Microsoft Advanced Systems Format (ie WMA/ASF/WMV).
	SOUND_TYPE_ASF = C.FMOD_SOUND_TYPE_ASF

	// Sony ATRAC 3 format
	SOUND_TYPE_AT3 = C.FMOD_SOUND_TYPE_AT3

	// Sound font / downloadable sound bank.
	SOUND_TYPE_DLS = C.FMOD_SOUND_TYPE_DLS

	// FLAC lossless codec.
	SOUND_TYPE_FLAC = C.FMOD_SOUND_TYPE_FLAC

	// FMOD Sample Bank.
	SOUND_TYPE_FSB = C.FMOD_SOUND_TYPE_FSB

	// Nintendo GameCube/Wii ADPCM
	SOUND_TYPE_GCADPCM = C.FMOD_SOUND_TYPE_GCADPCM

	// Impulse Tracker.
	SOUND_TYPE_IT = C.FMOD_SOUND_TYPE_IT

	// MIDI. extracodecdata is a pointer to an FMOD_MIDI_EXTRACODECDATA structure.
	SOUND_TYPE_MIDI = C.FMOD_SOUND_TYPE_MIDI

	// Protracker / Fasttracker MOD.
	SOUND_TYPE_MOD = C.FMOD_SOUND_TYPE_MOD

	// MP2/MP3 MPEG.
	SOUND_TYPE_MPEG = C.FMOD_SOUND_TYPE_MPEG

	// Ogg vorbis.
	SOUND_TYPE_OGGVORBIS = C.FMOD_SOUND_TYPE_OGGVORBIS

	// Information only from ASX/PLS/M3U/WAX playlists
	SOUND_TYPE_PLAYLIST = C.FMOD_SOUND_TYPE_PLAYLIST

	// Raw PCM data.
	SOUND_TYPE_RAW = C.FMOD_SOUND_TYPE_RAW

	// ScreamTracker 3.
	SOUND_TYPE_S3M = C.FMOD_SOUND_TYPE_S3M

	// User created sound.
	SOUND_TYPE_USER = C.FMOD_SOUND_TYPE_USER

	// Microsoft WAV.
	SOUND_TYPE_WAV = C.FMOD_SOUND_TYPE_WAV

	// FastTracker 2 XM.
	SOUND_TYPE_XM = C.FMOD_SOUND_TYPE_XM

	// Xbox360 XMA
	SOUND_TYPE_XMA = C.FMOD_SOUND_TYPE_XMA

	// PlayStation Portable ADPCM VAG format.
	SOUND_TYPE_VAG = C.FMOD_SOUND_TYPE_VAG

	// iPhone hardware decoder, supports AAC, ALAC and MP3. extracodecdata is a pointer to an FMOD_AUDIOQUEUE_EXTRACODECDATA structure.
	SOUND_TYPE_AUDIOQUEUE = C.FMOD_SOUND_TYPE_AUDIOQUEUE

	// Xbox360 XWMA
	SOUND_TYPE_XWMA = C.FMOD_SOUND_TYPE_XWMA

	// 3DS BCWAV container format for DSP ADPCM and PCM
	SOUND_TYPE_BCWAV = C.FMOD_SOUND_TYPE_BCWAV

	// PS4 / PSVita ATRAC 9 format
	SOUND_TYPE_AT9 = C.FMOD_SOUND_TYPE_AT9

	// Vorbis
	SOUND_TYPE_VORBIS = C.FMOD_SOUND_TYPE_VORBIS

	// Windows Store Application built in system codecs
	SOUND_TYPE_MEDIA_FOUNDATION = C.FMOD_SOUND_TYPE_MEDIA_FOUNDATION

	// Android MediaCodec
	SOUND_TYPE_MEDIACODEC = C.FMOD_SOUND_TYPE_MEDIACODEC

	// FMOD Adaptive Differential Pulse Code Modulation
	SOUND_TYPE_FADPCM = C.FMOD_SOUND_TYPE_FADPCM

	// Maximum number of sound types supported.
	SOUND_TYPE_MAX = C.FMOD_SOUND_TYPE_MAX

	// Makes sure this enum is signed 32bit.
	SOUND_TYPE_FORCEINT = C.FMOD_SOUND_TYPE_FORCEINT
)

type SoundFormat C.FMOD_SOUND_FORMAT

const (

	// Unitialized / unknown.
	SOUND_FORMAT_NONE SoundFormat = C.FMOD_SOUND_FORMAT_NONE

	// 8bit integer PCM data.
	SOUND_FORMAT_PCM8 = C.FMOD_SOUND_FORMAT_PCM8

	// 16bit integer PCM data.
	SOUND_FORMAT_PCM16 = C.FMOD_SOUND_FORMAT_PCM16

	// 24bit integer PCM data.
	SOUND_FORMAT_PCM24 = C.FMOD_SOUND_FORMAT_PCM24

	// 32bit integer PCM data.
	SOUND_FORMAT_PCM32 = C.FMOD_SOUND_FORMAT_PCM32

	// 32bit floating point PCM data.
	SOUND_FORMAT_PCMFLOAT = C.FMOD_SOUND_FORMAT_PCMFLOAT

	// Compressed Nintendo 3DS/Wii DSP data.
	SOUND_FORMAT_GCADPCM = C.FMOD_SOUND_FORMAT_GCADPCM

	// Compressed IMA ADPCM data.
	SOUND_FORMAT_IMAADPCM = C.FMOD_SOUND_FORMAT_IMAADPCM

	// Compressed PlayStation Portable ADPCM data.
	SOUND_FORMAT_VAG = C.FMOD_SOUND_FORMAT_VAG

	// Compressed PSVita ADPCM data.
	SOUND_FORMAT_HEVAG = C.FMOD_SOUND_FORMAT_HEVAG

	// Compressed Xbox360 XMA data.
	SOUND_FORMAT_XMA = C.FMOD_SOUND_FORMAT_XMA

	// Compressed MPEG layer 2 or 3 data.
	SOUND_FORMAT_MPEG = C.FMOD_SOUND_FORMAT_MPEG

	// Not supported.
	SOUND_FORMAT_CELT = C.FMOD_SOUND_FORMAT_CELT

	// Compressed PSVita ATRAC9 data.
	SOUND_FORMAT_AT9 = C.FMOD_SOUND_FORMAT_AT9

	// Compressed Xbox360 xWMA data.
	SOUND_FORMAT_XWMA = C.FMOD_SOUND_FORMAT_XWMA

	// Compressed Vorbis data.
	SOUND_FORMAT_VORBIS = C.FMOD_SOUND_FORMAT_VORBIS

	// Compressed FADPCM data.
	SOUND_FORMAT_FADPCM = C.FMOD_SOUND_FORMAT_FADPCM

	// Maximum number of sound formats supported.
	SOUND_FORMAT_MAX = C.FMOD_SOUND_FORMAT_MAX

	// Makes sure this enum is signed 32bit.
	SOUND_FORMAT_FORCEINT = C.FMOD_SOUND_FORMAT_FORCEINT
)

type DSPConnectionType C.FMOD_DSPCONNECTION_TYPE

const (

	// Default connection type.
	// Audio is mixed from the input to the output DSP's audible buffer.
	DSPCONNECTION_TYPE_STANDARD DSPConnectionType = C.FMOD_DSPCONNECTION_TYPE_STANDARD

	// Sidechain connection type.
	// Audio is mixed from the input to the output DSP's sidechain buffer.
	DSPCONNECTION_TYPE_SIDECHAIN = C.FMOD_DSPCONNECTION_TYPE_SIDECHAIN

	// Send connection type.
	// Audio is mixed from the input to the output DSP's audible buffer, but the input is NOT executed, only copied from.
	// A standard connection or sidechain needs to make an input execute to generate data.
	DSPCONNECTION_TYPE_SEND = C.FMOD_DSPCONNECTION_TYPE_SEND

	// Send sidechain connection type.
	// Audio is mixed from the input to the output DSP's sidechain buffer, but the input is NOT executed, only copied from.
	// A standard connection or sidechain needs to make an input execute to generate data.
	DSPCONNECTION_TYPE_SEND_SIDECHAIN = C.FMOD_DSPCONNECTION_TYPE_SEND_SIDECHAIN

	// Maximum number of DSP connection types supported.
	DSPCONNECTION_TYPE_MAX = C.FMOD_DSPCONNECTION_TYPE_MAX

	// Makes sure this enum is signed 32bit.
	DSPCONNECTION_TYPE_FORCEINT = C.FMOD_DSPCONNECTION_TYPE_FORCEINT
)

type ChannelMask C.FMOD_CHANNELMASK

const (
	CHANNELMASK_FRONT_LEFT     ChannelMask = C.FMOD_CHANNELMASK_FRONT_LEFT
	CHANNELMASK_FRONT_RIGHT                = C.FMOD_CHANNELMASK_FRONT_RIGHT
	CHANNELMASK_FRONT_CENTER               = C.FMOD_CHANNELMASK_FRONT_CENTER
	CHANNELMASK_LOW_FREQUENCY              = C.FMOD_CHANNELMASK_LOW_FREQUENCY
	CHANNELMASK_SURROUND_LEFT              = C.FMOD_CHANNELMASK_SURROUND_LEFT
	CHANNELMASK_SURROUND_RIGHT             = C.FMOD_CHANNELMASK_SURROUND_RIGHT
	CHANNELMASK_BACK_LEFT                  = C.FMOD_CHANNELMASK_BACK_LEFT
	CHANNELMASK_BACK_RIGHT                 = C.FMOD_CHANNELMASK_BACK_RIGHT
	CHANNELMASK_BACK_CENTER                = C.FMOD_CHANNELMASK_BACK_CENTER

	CHANNELMASK_MONO          = C.FMOD_CHANNELMASK_MONO
	CHANNELMASK_STEREO        = C.FMOD_CHANNELMASK_STEREO
	CHANNELMASK_LRC           = C.FMOD_CHANNELMASK_LRC
	CHANNELMASK_QUAD          = C.FMOD_CHANNELMASK_QUAD
	CHANNELMASK_SURROUND      = C.FMOD_CHANNELMASK_SURROUND
	CHANNELMASK_5POINT1       = C.FMOD_CHANNELMASK_5POINT1
	CHANNELMASK_5POINT1_REARS = C.FMOD_CHANNELMASK_5POINT1_REARS
	CHANNELMASK_7POINT0       = C.FMOD_CHANNELMASK_7POINT0
	CHANNELMASK_7POINT1       = C.FMOD_CHANNELMASK_7POINT1
)
