package lowlevel

/*
#include <fmod_common.h>
*/
import "C"

type DSPMeteringInfo struct {
	// [r] The number of samples considered for this metering info.
	Numsamples int

	// [r] The peak level per channel.
	Peaklevel float64

	// [r] The rms level per channel.
	Rmslevel float64

	// [r] Number of channels.
	Numchannels int16
}

func NewDSPMeteringInfo() DSPMeteringInfo {
	return DSPMeteringInfo{}
}

func (d *DSPMeteringInfo) fromC(cd C.FMOD_DSP_METERING_INFO) {
	d.Numsamples = int(cd.numsamples)
	//d.Peaklevel = float64(cd.peaklevel)
	//d.Rmslevel = float64(cd.rmslevel)
	d.Numchannels = int16(cd.numchannels)
}

func (d *DSPMeteringInfo) toC() C.FMOD_DSP_METERING_INFO {
	var cd C.FMOD_DSP_METERING_INFO
	cd.numsamples = C.int(d.Numsamples)
	//cd.peaklevel = C.float(d.Peaklevel)
	//cd.rmslevel = C.float(d.Rmslevel)
	cd.numchannels = C.short(d.Numchannels)
	return cd
}
