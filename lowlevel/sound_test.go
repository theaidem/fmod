package lowlevel

import (
	"reflect"
	"testing"
	"time"
)

func TestSoundCreate(t *testing.T) {
	system, done, err := NewSystem(1000 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_CREATESAMPLE, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = system.PlaySound(censor, nil, false)
	if err != nil {
		t.Fatal(err)
	}

	sys, err := censor.SystemObject()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(sys, system) {
		t.Error("Systems not equals")
	}

	<-done
}

func TestSoundDefaults(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_CREATESAMPLE, nil)
	if err != nil {
		t.Fatal(err)
	}

	frequency, priority, err := censor.Defaults()
	if err != nil {
		t.Fatal(err)
	}

	if frequency != 44100 {
		t.Error("frequency expected 44100 but got", frequency)
	}

	if priority != 128 {
		t.Error("priority expected 128 but got", priority)
	}

	err = censor.SetDefaults(22050, 64)
	if err != nil {
		t.Fatal(err)
	}

	frequency, priority, err = censor.Defaults()
	if err != nil {
		t.Fatal(err)
	}

	if frequency != 22050 {
		t.Error("frequency expected 22050 but got", frequency)
	}

	if priority != 64 {
		t.Error("priority expected 64 but got", priority)
	}

	<-done
}

func TestSoundLength(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_CREATESAMPLE, nil)
	if err != nil {
		t.Fatal(err)
	}

	msLength, err := censor.Length(TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	if msLength != 1000 {
		t.Error("Length (ms) expected 1000 but got", msLength)
	}

	pcmLength, err := censor.Length(TIMEUNIT_PCM)
	if err != nil {
		t.Fatal(err)
	}

	if pcmLength != 44100 {
		t.Error("Length (pcm) expected 44100 but got", pcmLength)
	}

	pcmBytesLength, err := censor.Length(TIMEUNIT_PCMBYTES)
	if err != nil {
		t.Fatal(err)
	}

	if pcmBytesLength != 88200 {
		t.Error("Length (bytes pcm) expected 88200 but got", pcmBytesLength)
	}

	rawBytesLength, err := censor.Length(TIMEUNIT_RAWBYTES)
	if err != nil {
		t.Fatal(err)
	}

	if rawBytesLength != 88200 {
		t.Error("Length (bytes raw) expected 88200 but got", rawBytesLength)
	}

	<-done
}

func TestSound3DMinMaxDistance(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_CREATESAMPLE, nil)
	if err != nil {
		t.Fatal(err)
	}

	min, max, err := censor.Get3DMinMaxDistance()
	if err != nil {
		t.Fatal(err)
	}

	if min != 1 {
		t.Error("3DMinDistance expected 1 but got", min)
	}

	if max != 10000 {
		t.Error("3DMaxDistance expected 10000 but got", max)
	}

	err = censor.Set3DMinMaxDistance(10, 1000)
	if err != nil {
		t.Fatal(err)
	}

	min, max, err = censor.Get3DMinMaxDistance()
	if err != nil {
		t.Fatal(err)
	}

	if min != 10 {
		t.Error("3DMinDistance expected 10 but got", min)
	}

	if max != 1000 {
		t.Error("3DMaxDistance expected 1000 but got", max)
	}

	<-done
}

func TestSound3DConeSettings(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_CREATESAMPLE, nil)
	if err != nil {
		t.Fatal(err)
	}

	insideconeangle, outsideconeangle, outsidevolume, err := censor.Get3DConeSettings()
	if err != nil {
		t.Fatal(err)
	}

	if insideconeangle != 360 {
		t.Error("Inside cone angle expected 360 but got", insideconeangle)
	}

	if outsideconeangle != 360 {
		t.Error("Outside cone angle expected 360 but got", outsideconeangle)
	}

	if outsidevolume != 1 {
		t.Error("Cone outside volume expected 1 but got", outsidevolume)
	}

	err = censor.Set3DConeSettings(180, 90, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	insideconeangle, outsideconeangle, outsidevolume, err = censor.Get3DConeSettings()
	if err != nil {
		t.Fatal(err)
	}

	if insideconeangle != 180 {
		t.Error("Inside cone angle expected 180 but got", insideconeangle)
	}

	if outsideconeangle != 90 {
		t.Error("Outside cone angle expected 90 but got", outsideconeangle)
	}

	if outsidevolume != 0.5 {
		t.Error("Cone outside volume expected 0.5 but got", outsidevolume)
	}

	<-done
}

func TestSound3DCustomRolloff(t *testing.T) {
	// TODO: Complite this units
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_DEFAULT, nil)
	if err != nil {
		t.Fatal(err)
	}

	points := NewVector()
	points.X = 90
	numpoints := 1
	err = censor.Set3DCustomRolloff(&points, numpoints)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = censor.Get3DCustomRolloff()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

func TestSoundFormat(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_DEFAULT, nil)
	if err != nil {
		t.Fatal(err)
	}

	typ, format, channels, bits, err := censor.Format()
	if err != nil {
		t.Fatal(err)
	}

	if typ != SOUND_TYPE_WAV {
		t.Error("Sound type expected 15 but got", typ)
	}

	if format != SOUND_FORMAT_PCM16 {
		t.Error("Sound format expected 2 but got", format)
	}

	if channels != 1 {
		t.Error("Sound channels expected 1 but got", channels)
	}

	if bits != 16 {
		t.Error("Sound bits expected 16 but got", bits)
	}

	guiro, err := system.CreateSound("media/guiro.mp3", MODE_DEFAULT, nil)
	if err != nil {
		t.Fatal(err)
	}

	typ, format, channels, bits, err = guiro.Format()
	if err != nil {
		t.Fatal(err)
	}

	if typ != SOUND_TYPE_MPEG {
		t.Error("Sound type expected 9 but got", typ)
	}

	<-done
}

func TestSoundLoopCount(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_DEFAULT, nil)
	if err != nil {
		t.Fatal(err)
	}

	lc, err := censor.LoopCount()
	if err != nil {
		t.Fatal(err)
	}

	if lc != -1 {
		t.Error("Loop count expected -1 (loop forever) but got", lc)
	}

	err = censor.SetLoopCount(5)
	if err != nil {
		t.Fatal(err)
	}

	lc, err = censor.LoopCount()
	if err != nil {
		t.Fatal(err)
	}

	if lc != 5 {
		t.Error("Loop count expected 5 but got", lc)
	}

	<-done
}

func TestSoundLoopPoints(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_LOOP_NORMAL, nil)
	if err != nil {
		t.Fatal(err)
	}

	lp1, lp2, err := censor.LoopPoints(TIMEUNIT_MS, TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	if lp1 != 0 {
		t.Error("The loop start point expected 0 but got", lp1)
	}

	if lp2 != 999 {
		t.Error("The loop end point expected 999 but got", lp2)
	}

	err = censor.SetLoopPoints(10, TIMEUNIT_MS, 500, TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	lp1, lp2, err = censor.LoopPoints(TIMEUNIT_MS, TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	if lp1 != 10 {
		t.Error("The loop start point expected 0 but got", lp1)
	}

	if lp2 != 500 {
		t.Error("The loop end point expected 999 but got", lp2)
	}

	<-done
}

func TestSoundUserData(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	censor, err := system.CreateSound("media/censor.wav", MODE_LOOP_NORMAL, nil)
	if err != nil {
		t.Fatal(err)
	}

	userData := "TestData"

	err = censor.SetUserData(userData)
	if err != nil {
		t.Fatal(err)
	}

	data, err := censor.UserData()
	if err != nil {
		t.Fatal(err)
	}

	if data != userData {
		t.Error("Data is different")
	}

	<-done
}
