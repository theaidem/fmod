package lowlevel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSoundGroupCreate(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	setName := "SoundGroupTest"
	sg, err := system.CreateSoundGroup(setName)
	if err != nil {
		t.Fatal(err)
	}

	sgSystem, err := sg.SystemObject()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(sgSystem, system) {
		t.Error("Systems not equals")
	}

	gotName, err := sg.Name()
	if err != nil {
		t.Fatal(err)
	}

	if gotName != setName {
		t.Error("Names isn't equals")
	}

	<-done
}

func TestSoundGroupNCreate(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	num := 10000
	var soundGroups []*SoundGroup
	for i := 1; i <= num; i++ {
		setName := fmt.Sprintf("SoundGroupTest:%d", i)
		sg, err := system.CreateSoundGroup(setName)
		if err != nil {
			t.Fatal(err)
		}
		if sg != nil {
			gotName, err := sg.Name()
			if err != nil {
				t.Fatal(err)
			}

			if gotName != setName {
				t.Error("Names isn't equals")
			}

			soundGroups = append(soundGroups, sg)
		}
	}

	if len(soundGroups) != num {
		t.Error("Wrong groups num/len ")
	}

	<-done
}

func TestSoundGroupMaxAudible(t *testing.T) {
	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	sg, err := system.CreateSoundGroup("SoundGroupTest")
	if err != nil {
		t.Fatal(err)
	}

	max, err := sg.MaxAudible()
	if err != nil {
		t.Fatal(err)
	}

	if max != -1 {
		t.Error("MaxAudible expected -1 (unlimited) but got", max)
	}

	err = sg.SetMaxAudible(5)
	if err != nil {
		t.Fatal(err)
	}

	max, err = sg.MaxAudible()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Error("MaxAudible expected 5 but got", max)
	}

	<-done
}

func TestSoundGroupUserData(t *testing.T) {

	system, done, err := NewSystem(0)
	if err != nil {
		t.Fatal(err)
	}

	setName := "SoundGroupTest"
	sg, err := system.CreateSoundGroup(setName)
	if err != nil {
		t.Fatal(err)
	}

	userData := "TestData"
	err = sg.SetUserData(userData)
	if err != nil {
		t.Fatal(err)
	}

	data, err := sg.UserData()
	if err != nil {
		t.Fatal(err)
	}

	if data != userData {
		t.Error("Data is different")
	}

	<-done
}
