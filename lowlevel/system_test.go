package lowlevel

import (
	"testing"
	"time"
)

var system *System

func NewSystem() (chan bool, error) {
	var err error
	done := make(chan bool)

	system, err = SystemCreate()
	if err != nil {
		return done, err
	}

	err = system.SetOutput(OUTPUTTYPE_AUTODETECT)
	if err != nil {
		return done, err
	}

	//Init the System object
	err = system.Init(10, INIT_NORMAL, 0)
	if err != nil {
		return done, err
	}

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				err := system.Update()
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	return done, nil
}

func TestSystemInstance(t *testing.T) {
	done, err := NewSystem()
	if err != nil {
		t.Fatal(err)
	}

	//Get System version
	v, err := system.GetVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v <= 0 {
		t.Error("expected more than zero but got", v)
	}

	channels, err := system.GetChannelsPlaying()
	if err != nil {
		t.Fatal(err)
	}

	if channels > 0 {
		t.Error("expected ", 0, " but got", v)
	}

	cpu, err := system.GetCPUUsage()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", cpu)

	ram, err := system.GetSoundRAM()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", ram)

	done <- true
}

func TestSystemSetup(t *testing.T) {

	done, err := NewSystem()
	if err != nil {
		t.Fatal(err)
	}

	output, err := system.GetOutput()
	if err != nil {
		t.Fatal(err)
	}

	if output == OUTPUTTYPE_AUTODETECT {
		t.Error("Hmmm, output isn't detected?")
	}

	drivers, err := system.GetNumDrivers()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(drivers)

	guid, systemrate, speakermode, speakermodechannels, err := system.GetDriverInfo(0, "", 0)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", guid)
	t.Logf("%#v\n", systemrate)
	t.Logf("%#v\n", speakermode)
	t.Logf("%#v\n", speakermodechannels)

	err = system.SetDriver(0)
	if err != nil {
		t.Fatal(err)
	}

	driver, err := system.GetDriver()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(driver)

	channels, err := system.GetSoftwareChannels()
	if err != nil {
		t.Fatal(err)
	}

	if channels != 64 {
		t.Error("expected", 64, "got", channels)
	}

	samplerate, speakermode, numrawspeakers, err := system.GetSoftwareFormat()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", samplerate)
	t.Logf("%#v\n", speakermode)
	t.Logf("%#v\n", numrawspeakers)

	bufferlength, numbuffers, err := system.GetDSPBufferSize()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", bufferlength)
	t.Logf("%#v\n", numbuffers)

	done <- true
}

func TestSystemPlugins(t *testing.T) {

	done, err := NewSystem()
	if err != nil {
		t.Fatal(err)
	}

	codecplugins, err := system.GetNumPlugins(PLUGINTYPE_CODEC)
	if err != nil {
		t.Fatal(err)
	}

	dspplugins, err := system.GetNumPlugins(PLUGINTYPE_DSP)
	if err != nil {
		t.Fatal(err)
	}

	outplugins, err := system.GetNumPlugins(PLUGINTYPE_OUTPUT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Codecs: %#v\n", codecplugins)
	t.Logf("DSP: %#v\n", dspplugins)
	t.Logf("Output: %#v\n", outplugins)

	done <- true
}

func TestSystemPostInit(t *testing.T) {

	done, err := NewSystem()
	if err != nil {
		t.Fatal(err)
	}

	x, y, active, err := system.GetSpeakerPosition(SPEAKER_FRONT_LEFT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SPEAKER_FRONT_LEFT Active: %#v\n", active)
	t.Logf("SPEAKER_FRONT_LEFT X: %#v\n", x)
	t.Logf("SPEAKER_FRONT_LEFT Y: %#v\n", y)
	t.Log("")

	x, y, active, err = system.GetSpeakerPosition(SPEAKER_FRONT_RIGHT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SPEAKER_FRONT_RIGHT Active: %#v\n", active)
	t.Logf("SPEAKER_FRONT_RIGHT X: %#v\n", x)
	t.Logf("SPEAKER_FRONT_RIGHT Y: %#v\n", y)
	t.Log("")

	err = system.MixerSuspend()
	if err != nil {
		t.Fatal(err)
	}

	err = system.MixerResume()
	if err != nil {
		t.Fatal(err)
	}

	done <- true
}

func TestSystemCreation(t *testing.T) {

	done, err := NewSystem()
	if err != nil {
		t.Fatal(err)
	}

	var exinfo *CreatesSoundExInfo
	bell, err := system.CreateSound("./media/bell.mp3", MODE_DEFAULT, exinfo)
	if err != nil {
		t.Fatal(err)
	}

	guiro, err := system.CreateSound("./media/guiro.mp3", MODE_DEFAULT, exinfo)
	if err != nil {
		t.Fatal(err)
	}

	var channelgroup *ChannelGroup
	_, err = system.PlaySound(bell, channelgroup, true)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.PlaySound(guiro, channelgroup, false)
	if err != nil {
		t.Fatal(err)
	}

	playing, err := system.GetChannelsPlaying()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(playing)

	time.Sleep(100 * time.Millisecond)

	err = system.SetSpeakerPosition(SPEAKER_FRONT_LEFT, 36.5, 10, false)
	if err != nil {
		t.Fatal(err)
	}

	err = system.SetSpeakerPosition(SPEAKER_FRONT_RIGHT, 36.5, 10, false)
	if err != nil {
		t.Fatal(err)
	}

	x, y, active, err := system.GetSpeakerPosition(SPEAKER_FRONT_LEFT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SPEAKER_FRONT_LEFT Active: %#v\n", active)
	t.Logf("SPEAKER_FRONT_LEFT X: %#v\n", x)
	t.Logf("SPEAKER_FRONT_LEFT Y: %#v\n", y)
	t.Log("")

	x, y, active, err = system.GetSpeakerPosition(SPEAKER_FRONT_RIGHT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SPEAKER_FRONT_RIGHT Active: %#v\n", active)
	t.Logf("SPEAKER_FRONT_RIGHT X: %#v\n", x)
	t.Logf("SPEAKER_FRONT_RIGHT Y: %#v\n", y)
	t.Log("")

	dopplerscale, distancefactor, rolloffscale, err := system.Get3DSettings()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("dopplerscale: %#v\n", dopplerscale)
	t.Logf("distancefactor: %#v\n", distancefactor)
	t.Logf("rolloffscale: %#v\n", rolloffscale)

	_, err = system.PlaySound(guiro, channelgroup, false)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1000 * time.Millisecond)
	done <- true
}
