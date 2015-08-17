package lowlevel

/*
#include <stdlib.h>
#include <fmod_common.h>
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type AdvancedSettings struct {
	// [w]   Size of this structure.
	// Use sizeof(FMOD_ADVANCEDSETTINGS)
	// NOTE: This must be set before calling System.AdvancedSettings or System.SetAdvancedSettings!
	CbSize int

	// [r/w] Optional. Specify 0 to ignore. Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// MPEG   codecs consume 30,528 bytes per instance and this number will determine how many MPEG   channels can be played simultaneously.
	MaxMPEGCodecs int

	// [r/w] Optional. Specify 0 to ignore.  Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// ADPCM  codecs consume  3,128 bytes per instance and this number will determine how many ADPCM  channels can be played simultaneously.
	MaxADPCMCodecs int

	// [r/w] Optional. Specify 0 to ignore.  Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// XMA    codecs consume 14,836 bytes per instance and this number will determine how many XMA    channels can be played simultaneously.
	MaxXMACodecs int

	// [r/w] Optional. Specify 0 to ignore.  Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// Vorbis codecs consume 23,256 bytes per instance and this number will determine how many Vorbis channels can be played simultaneously.
	MaxVorbisCodecs int

	// [r/w] Optional. Specify 0 to ignore. Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// AT9    codecs consume  8,720 bytes per instance and this number will determine how many AT9    channels can be played simultaneously.
	MaxAT9Codecs int

	// [r/w] Optional. Specify 0 to ignore.  Default = 32.
	// For use with FMOD_CREATECOMPRESSEDSAMPLE only.
	// This number will determine how many FADPCM channels can be played simultaneously.
	MaxFADPCMCodecs int

	// [r/w] Optional. Specify 0 to ignore.  Default = 16.
	// For use with PS3 only.
	// PCM    codecs consume 12,672 bytes per instance and this number will determine how many streams and PCM voices can be played simultaneously.
	MaxPCMCodecs int

	// [r/w] Optional. Specify 0 to ignore.
	// Number of channels available on the ASIO device.
	ASIONumChannels int

	// [r/w] Optional. Specify 0 to ignore.
	// Pointer to an array of strings (number of entries defined by ASIONumChannels) with ASIO channel names.
	ASIOChannelList []string

	// [r/w] Optional. Specify 0 to ignore.
	// Pointer to a list of speakers that the ASIO channels map to.  This can be called after System.Init to remap ASIO output.
	ASIOSpeakerList *Speaker //FMOD_SPEAKER

	// [r/w] Optional. Default = 180.0.
	// For use with FMOD_INIT_HRTF_LOWPASS.
	// The angle range (0-360) of a 3D sound in relation to the listener, at which the HRTF function begins to have an effect.
	// 0 = in front of the listener.
	// 180 = from 90 degrees to the left of the listener to 90 degrees to the right.
	// 360 = behind the listener.
	HRTFMinAngle float64

	// [r/w] Optional. Default = 360.0.
	// For use with FMOD_INIT_HRTF_LOWPASS.
	// The angle range (0-360) of a 3D sound in relation to the listener, at which the HRTF function has maximum effect.
	// 0 = front of the listener.
	// 180 = from 90 degrees to the left of the listener to 90 degrees to the right.
	// 360 = behind the listener.
	HRTFMaxAngle float64

	// [r/w] Optional. Specify 0 to ignore.  Default = 4000.0.
	// For use with FMOD_INIT_HRTF_LOWPASS.
	// The cutoff frequency of the HRTF's lowpass filter function when at maximum effect. (i.e. at HRTFMaxAngle).
	HRTFFreq float64

	// [r/w] Optional. Specify 0 to ignore.
	// For use with FMOD_INIT_VOL0_BECOMES_VIRTUAL.
	// If this flag is used, and the volume is below this, then the sound will become virtual.
	// Use this value to raise the threshold to a different point where a sound goes virtual.
	Vol0virtualvol float64

	// [r/w] Optional. Specify 0 to ignore.   Default = 400ms
	// For streams. This determines the default size of the double buffer (in milliseconds) that a stream uses.
	DefaultDecodeBufferSize uint32

	// [r/w] Optional. Specify 0 to ignore.
	// For use with FMOD_INIT_PROFILE_ENABLE.
	// Specify the port to listen on for connections by the profiler application.
	ProfilePort uint16

	// [r/w] Optional. Specify 0 to ignore.
	// The maximum time in miliseconds it takes for a channel to fade to the new level when its occlusion changes.
	GeometryMaxFadeTime uint32

	// [r/w] Optional. Specify 0 to ignore.  Default = 1500.0.
	// For use with FMOD_INIT_DISTANCE_FILTERING.
	// The default center frequency in Hz for the distance filtering effect.
	DistanceFilterCenterFreq float64

	// [r/w] Optional. Specify 0 to ignore.
	// Out of 0 to 3, 3d reverb spheres will create a phyical reverb unit on this instance slot.
	// See FMOD_REVERB_PROPERTIES.
	Reverb3Dinstance int

	// [r/w] Optional. Specify 0 to ignore. Default = 8.
	// Number of buffers in DSP buffer pool.
	// Each buffer will be DSPBlockSize * sizeof(float) * SpeakerModeChannelCount.  ie 7.1 @ 1024 DSP block size = 8 * 1024 * 4 = 32kb.
	DSPBufferPoolSize int

	// [r/w] Optional. Specify 0 to ignore. Default 49,152 (48kb)
	// Specify the stack size for the FMOD Stream thread in bytes.
	// Useful for custom codecs that use excess stack.
	StackSizeStream uint32

	// [r/w] Optional. Specify 0 to ignore.   Default 65,536 (64kb)
	// Specify the stack size for the FMOD_NONBLOCKING loading thread.
	// Useful for custom codecs that use excess stack.
	StackSizeNonBlocking uint32

	// [r/w] Optional. Specify 0 to ignore.   Default 49,152 (48kb)
	// Specify the stack size for the FMOD mixer thread.
	// Useful for custom dsps that use excess stack.
	StackSizeMixer uint32

	// [r/w] Optional. Specify 0 to ignore.
	// Resampling method used with fmod's software mixer.
	// See FMOD_DSP_RESAMPLER for details on methods.
	ReSamplerMethod DSPReSampler

	// [r/w] Optional. Specify 0 to ignore. Default 2048 (2kb)
	// Specify the command queue size for thread safe processing.
	CommandQueueSize uint32

	// [r/w] Optional. Specify 0 to ignore.
	// Seed value that FMOD will use to initialize its internal random number generators.
	RandomSeed uint32
}

func NewAdvancedSettings() AdvancedSettings {
	s := AdvancedSettings{}
	var csettings C.FMOD_ADVANCEDSETTINGS
	s.CbSize = int(unsafe.Sizeof(csettings))
	return s
}

func (a *AdvancedSettings) fromC(as C.FMOD_ADVANCEDSETTINGS) {
	a.CbSize = int(as.cbSize)
	a.MaxMPEGCodecs = int(as.maxMPEGCodecs)
	a.MaxADPCMCodecs = int(as.maxADPCMCodecs)
	a.MaxXMACodecs = int(as.maxXMACodecs)
	a.MaxVorbisCodecs = int(as.maxVorbisCodecs)
	a.MaxAT9Codecs = int(as.maxAT9Codecs)
	a.MaxFADPCMCodecs = int(as.maxFADPCMCodecs)
	a.MaxPCMCodecs = int(as.maxPCMCodecs)
	a.ASIONumChannels = int(as.ASIONumChannels)

	channels := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(as.ASIOChannelList)),
		Len:  a.ASIONumChannels,
		Cap:  a.ASIONumChannels,
	}

	a.ASIOChannelList = *((*[]string)(unsafe.Pointer(channels)))
	//a.ASIOSpeakerList = Speaker(as.ASIOSpeakerList) //TODO
	a.HRTFMinAngle = float64(as.HRTFMinAngle)
	a.HRTFMaxAngle = float64(as.HRTFMaxAngle)
	a.HRTFFreq = float64(as.HRTFFreq)
	a.Vol0virtualvol = float64(as.vol0virtualvol)
	a.DefaultDecodeBufferSize = uint32(as.defaultDecodeBufferSize)
	a.ProfilePort = uint16(as.profilePort)
	a.GeometryMaxFadeTime = uint32(as.geometryMaxFadeTime)
	a.DistanceFilterCenterFreq = float64(as.distanceFilterCenterFreq)
	a.Reverb3Dinstance = int(as.reverb3Dinstance)
	a.DSPBufferPoolSize = int(as.DSPBufferPoolSize)
	a.StackSizeStream = uint32(as.stackSizeStream)
	a.StackSizeNonBlocking = uint32(as.stackSizeNonBlocking)
	a.StackSizeMixer = uint32(as.stackSizeMixer)
	a.ReSamplerMethod = DSPReSampler(as.resamplerMethod)
	a.CommandQueueSize = uint32(as.commandQueueSize)
	a.RandomSeed = uint32(as.randomSeed)
}

func (a *AdvancedSettings) toC() C.FMOD_ADVANCEDSETTINGS {
	var as C.FMOD_ADVANCEDSETTINGS
	as.cbSize = C.int(a.CbSize)
	as.maxMPEGCodecs = C.int(a.MaxMPEGCodecs)
	as.maxADPCMCodecs = C.int(a.MaxADPCMCodecs)
	as.maxXMACodecs = C.int(a.MaxXMACodecs)
	as.maxVorbisCodecs = C.int(a.MaxVorbisCodecs)
	as.maxAT9Codecs = C.int(a.MaxAT9Codecs)
	as.maxFADPCMCodecs = C.int(a.MaxFADPCMCodecs)
	as.maxPCMCodecs = C.int(a.MaxPCMCodecs)
	as.ASIONumChannels = C.int(a.ASIONumChannels)

	outer := make([]*C.char, a.ASIONumChannels+1)
	for i, inner := range a.ASIOChannelList {
		outer[i] = C.CString(string(inner))
	}

	as.ASIOChannelList = (**C.char)(unsafe.Pointer(&outer[0]))

	//as.ASIOSpeakerList = C.FMOD_SPEAKER(a.ASIOSpeakerList) //TODO
	as.HRTFMinAngle = C.float(a.HRTFMinAngle)
	as.HRTFMaxAngle = C.float(a.HRTFMaxAngle)
	as.HRTFFreq = C.float(a.HRTFFreq)
	as.vol0virtualvol = C.float(a.Vol0virtualvol)
	as.defaultDecodeBufferSize = C.uint(a.DefaultDecodeBufferSize)
	as.profilePort = C.ushort(a.ProfilePort)
	as.geometryMaxFadeTime = C.uint(a.GeometryMaxFadeTime)
	as.distanceFilterCenterFreq = C.float(a.DistanceFilterCenterFreq)
	as.reverb3Dinstance = C.int(a.Reverb3Dinstance)
	as.DSPBufferPoolSize = C.int(a.DSPBufferPoolSize)
	as.stackSizeStream = C.uint(a.StackSizeStream)
	as.stackSizeNonBlocking = C.uint(a.StackSizeNonBlocking)
	as.stackSizeMixer = C.uint(a.StackSizeMixer)
	as.resamplerMethod = C.FMOD_DSP_RESAMPLER(a.ReSamplerMethod)
	as.commandQueueSize = C.uint(a.CommandQueueSize)
	as.randomSeed = C.uint(a.RandomSeed)

	return as
}
