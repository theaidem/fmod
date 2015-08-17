package lowlevel

/*
#include <fmod.h>
*/
import "C"

type ReverbProperties struct {

	// [r/w] 0.0    20000.0 1500.0  Reverberation decay time in ms
	DecayTime float64

	// [r/w] 0.0    300.0   7.0     Initial reflection delay time
	EarlyDelay float64

	// [r/w] 0.0    100     11.0    Late reverberation delay time relative to initial reflection
	LateDelay float64

	// [r/w] 20.0   20000.0 5000    Reference high frequency (hz)
	HFReference float64

	// [r/w] 10.0   100.0   50.0    High-frequency to mid-frequency decay time ratio
	HFDecayRatio float64

	// [r/w] 0.0    100.0   100.0   Value that controls the echo density in the late reverberation decay.
	Diffusion float64

	// [r/w] 0.0    100.0   100.0   Value that controls the modal density in the late reverberation decay
	Density float64

	// [r/w] 20.0   1000.0  250.0   Reference low frequency (hz)
	LowShelfFrequency float64

	// [r/w] -36.0  12.0    0.0     Relative room effect level at low frequencies
	LowShelfGain float64

	// [r/w] 20.0   20000.0 20000.0 Relative room effect level at high frequencies
	HighCut float64

	// [r/w] 0.0    100.0   50.0    Early reflections level relative to room effect
	EarlyLateMix float64

	// [r/w] -80.0  20.0    -6.0    Room effect level (at mid frequencies)
	WetLevel float64
}

func NewReverbProperties() ReverbProperties {
	return ReverbProperties{
		DecayTime:         0,
		EarlyDelay:        0,
		LateDelay:         0,
		HFReference:       20,
		HFDecayRatio:      10,
		Diffusion:         0,
		Density:           0,
		LowShelfFrequency: 20,
		LowShelfGain:      -36,
		HighCut:           20,
		EarlyLateMix:      0,
		WetLevel:          -80,
	}
}

func (r *ReverbProperties) fromC(cr C.FMOD_REVERB_PROPERTIES) {
	r.DecayTime = float64(cr.DecayTime)
	r.EarlyDelay = float64(cr.EarlyDelay)
	r.LateDelay = float64(cr.LateDelay)
	r.HFReference = float64(cr.HFReference)
	r.HFDecayRatio = float64(cr.HFDecayRatio)
	r.Diffusion = float64(cr.Diffusion)
	r.Density = float64(cr.Density)
	r.LowShelfFrequency = float64(cr.LowShelfFrequency)
	r.LowShelfGain = float64(cr.LowShelfGain)
	r.HighCut = float64(cr.HighCut)
	r.EarlyLateMix = float64(cr.EarlyLateMix)
	r.WetLevel = float64(cr.WetLevel)
}

func (r *ReverbProperties) toC() C.FMOD_REVERB_PROPERTIES {
	var cr C.FMOD_REVERB_PROPERTIES
	cr.DecayTime = C.float(r.DecayTime)
	cr.EarlyDelay = C.float(r.EarlyDelay)
	cr.LateDelay = C.float(r.LateDelay)
	cr.HFReference = C.float(r.HFReference)
	cr.HFDecayRatio = C.float(r.HFDecayRatio)
	cr.Diffusion = C.float(r.Diffusion)
	cr.Density = C.float(r.Density)
	cr.LowShelfFrequency = C.float(r.LowShelfFrequency)
	cr.LowShelfGain = C.float(r.LowShelfGain)
	cr.HighCut = C.float(r.HighCut)
	cr.EarlyLateMix = C.float(r.EarlyLateMix)
	cr.WetLevel = C.float(r.WetLevel)
	return cr
}
