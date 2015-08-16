package lowlevel

import "testing"

func TestAdvancedSettings(t *testing.T) {
	system, err := SystemCreate()
	if err != nil {
		t.Fatal(err)
	}

	//Init the System object
	err = system.Init(10, INIT_NORMAL, 0)
	if err != nil {
		t.Fatal(err)
	}

	set_settings := NewAdvancedSettings()
	set_settings.RandomSeed = 1000
	set_settings.ReSamplerMethod = DSP_RESAMPLER_SPLINE
	err = system.SetAdvancedSettings(&set_settings)
	if err != nil {
		t.Fatal(err)
	}

	get_settings, err := system.GetAdvancedSettings()
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(set_settings)
	//t.Log(get_settings)

	if set_settings.RandomSeed != get_settings.RandomSeed {
		t.Fatal("Some Settings arent equals")
	}

	if set_settings.ReSamplerMethod != get_settings.ReSamplerMethod {
		t.Fatal("Some Settings arent equals")
	}

	//Close the System object
	err = system.Close()
	if err != nil {
		t.Fatal(err)
	}

	//Free System
	err = system.Release()
	if err != nil {
		t.Fatal(err)
	}

}
