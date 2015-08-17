package lowlevel

import (
	"runtime"
	"testing"
	"time"
)

func init() {
	runtime.LockOSThread()
}

func TestSoundCreate(t *testing.T) {
	done, err := NewSystem()
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

	t.Logf("censorFrequency: %#v\n", frequency)
	t.Logf("censorPriority: %#v\n", priority)

	msLength, err := censor.Length(TIMEUNIT_MS)
	if err != nil {
		t.Fatal(err)
	}

	pcmLength, err := censor.Length(TIMEUNIT_PCM)
	if err != nil {
		t.Fatal(err)
	}

	pcmBytesLength, err := censor.Length(TIMEUNIT_PCMBYTES)
	if err != nil {
		t.Fatal(err)
	}

	rawBytesLength, err := censor.Length(TIMEUNIT_RAWBYTES)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("censorLength (ms): %d\n", msLength)
	t.Logf("censorLength (pcm): %d\n", pcmLength)
	t.Logf("censorLength (pcm bytes): %d\n", pcmBytesLength)
	t.Logf("censorLength (raw bytes): %d\n", rawBytesLength)

	_, err = system.PlaySound(censor, nil, false)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Duration(msLength) * time.Millisecond)

	done <- true
}
