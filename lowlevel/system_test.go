package lowlevel

import (
	"testing"
	"time"
)

//var system *System

func NewSystem(duration time.Duration) (*System, chan bool, error) {
	//var err error
	done := make(chan bool)

	system, err := SystemCreate()
	if err != nil {
		return nil, done, err
	}

	err = system.SetOutput(OUTPUTTYPE_AUTODETECT)
	if err != nil {
		return nil, done, err
	}

	//Init the System object
	err = system.Init(10, INIT_NORMAL, 0)
	if err != nil {
		return nil, done, err
	}

	go func() {

		defer func() {
			// Manualy Release System
			err := system.Release()
			if err != nil {
				panic(err)
			}
		}()

		for {
			select {
			case <-time.After(duration):
				done <- true
				return
			}

			err := system.Update()
			if err != nil {
				panic(err)
			}
		}
	}()

	return system, done, nil
}

func TestSystemInstance(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	//Get System version
	v, err := system.Version()
	if err != nil {
		t.Fatal(err)
	}

	if v <= 0 {
		t.Error("expected more than zero but got", v)
	}

	channels, err := system.ChannelsPlaying()
	if err != nil {
		t.Fatal(err)
	}

	if channels > 0 {
		t.Error("expected ", 0, " but got", v)
	}

	_, err = system.CPUUsage()
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.SoundRAM()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemSetup(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	output, err := system.Output()
	if err != nil {
		t.Fatal(err)
	}

	if output == OUTPUTTYPE_AUTODETECT {
		t.Error("Hmmm, output isn't detected?")
	}

	drivers, err := system.NumDrivers()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(drivers)

	guid, systemrate, speakermode, speakermodechannels, err := system.DriverInfo(0, "")
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

	driver, err := system.Driver()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(driver)

	channels, err := system.SoftwareChannels()
	if err != nil {
		t.Fatal(err)
	}

	if channels != 64 {
		t.Error("expected", 64, "got", channels)
	}

	samplerate, speakermode, numrawspeakers, err := system.SoftwareFormat()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", samplerate)
	t.Logf("%#v\n", speakermode)
	t.Logf("%#v\n", numrawspeakers)

	bufferlength, numbuffers, err := system.DSPBufferSize()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", bufferlength)
	t.Logf("%#v\n", numbuffers)

	<-done
}

func TestSystemPlugins(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	codecplugins, err := system.NumPlugins(PLUGINTYPE_CODEC)
	if err != nil {
		t.Fatal(err)
	}

	dspplugins, err := system.NumPlugins(PLUGINTYPE_DSP)
	if err != nil {
		t.Fatal(err)
	}

	outplugins, err := system.NumPlugins(PLUGINTYPE_OUTPUT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Codecs: %#v\n", codecplugins)
	t.Logf("DSP: %#v\n", dspplugins)
	t.Logf("Output: %#v\n", outplugins)

	<-done
}

func TestSystemPostInit(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	x, y, active, err := system.SpeakerPosition(SPEAKER_FRONT_LEFT)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SPEAKER_FRONT_LEFT Active: %#v\n", active)
	t.Logf("SPEAKER_FRONT_LEFT X: %#v\n", x)
	t.Logf("SPEAKER_FRONT_LEFT Y: %#v\n", y)
	t.Log("")

	x, y, active, err = system.SpeakerPosition(SPEAKER_FRONT_RIGHT)
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

	<-done
}

func TestSystemCreateSound(t *testing.T) {

	system, done, err := NewSystem(600 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}

	bell, err := system.CreateSound("./media/bell.mp3", MODE_DEFAULT, nil)
	if err != nil {
		t.Fatal(err)
	}

	lms, err := bell.Length(TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	if lms != 576 {
		t.Error("expected 576 but got", lms)
	}

	_, err = system.PlaySound(bell, nil, false)
	if err != nil {
		t.Fatal(err)
	}

	playing, err := system.ChannelsPlaying()
	if err != nil {
		t.Fatal(err)
	}

	if playing != 1 {
		t.Error("expected 1 but got", playing)
	}

	<-done
}

func TestSystemCreateStream(t *testing.T) {

	system, done, err := NewSystem(600 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}

	bell, err := system.CreateStream("./media/bell.mp3", MODE_CREATESTREAM, nil)
	if err != nil {
		t.Fatal(err)
	}

	lms, err := bell.Length(TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	if lms != 576 {
		t.Error("expected 576 but got", lms)
	}

	_, err = system.PlaySound(bell, nil, false)
	if err != nil {
		t.Fatal(err)
	}

	playing, err := system.ChannelsPlaying()
	if err != nil {
		t.Fatal(err)
	}

	if playing != 1 {
		t.Error("expected 1 but got", playing)
	}

	<-done
}

func TestSystemCreateChannelGroup(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.CreateChannelGroup("TestChannelGroup")
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemCreateSoundGroup(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.CreateSoundGroup("TestSoundGroup")
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemCreateReverb3D(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.CreateReverb3D()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemCreateGeometry(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.CreateGeometry(1, 1)
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemMasterChannelGroup(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.MasterChannelGroup()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemMasterSoundGroup(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.MasterSoundGroup()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSystemUserData(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	userData := "TestData"
	err = system.SetUserData(userData)
	if err != nil {
		t.Fatal(err)
	}

	data, err := system.UserData()
	if err != nil {
		t.Fatal(err)
	}

	if data != userData {
		t.Error("Data is different")
	}

	<-done
}

func TestSystemGlobalReverbProperties(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	propsIn := NewReverbProperties()
	err = system.SetReverbProperties(propsIn)
	if err != nil {
		t.Fatal(err)
	}

	propsOut, err := system.ReverbProperties()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", propsIn.DecayTime)
	t.Logf("%#v\n", propsOut.DecayTime)

	<-done
}

/*
func TestSystemAdvancedSettings(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	setSettings := NewAdvancedSettings()
	setSettings.RandomSeed = 1000
	setSettings.ReSamplerMethod = DSP_RESAMPLER_SPLINE
	err = system.SetAdvancedSettings(setSettings)
	if err != nil {
		t.Fatal(err)
	}

	getSettings, err := system.AdvancedSettings()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v\n", setSettings)
	t.Logf("%v\n", getSettings)

	if setSettings.RandomSeed != getSettings.RandomSeed {
		t.Fatal("Some Settings arent equals")
	}

	if setSettings.ReSamplerMethod != getSettings.ReSamplerMethod {
		t.Fatal("Some Settings arent equals")
	}

	<-done
}
*/
