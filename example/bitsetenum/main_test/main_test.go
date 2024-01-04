package main_test

import "testing"

func TestStatus(t *testing.T) {
	var status Status

	shouldBeTrue := func(in bool) {
		if in == false {
			t.Fatalf("input should be true")
		}
	}
	shouldBeFalse := func(in bool) {
		if in == true {
			t.Fatalf("input should be false")
		}
	}
	shouldBeFalse(status.GetStatusCreating())
	shouldBeFalse(status.GetStatusNormal())
	shouldBeFalse(status.GetStatusUpdating())
	shouldBeFalse(status.GetStatusCanaryUpdating())

	status.SetStatusCreating(true)
	status.SetStatusNormal(true)
	status.SetStatusUpdating(true)
	status.SetStatusCanaryUpdating(true)

	shouldBeTrue(status.GetStatusCreating())
	shouldBeTrue(status.GetStatusNormal())
	shouldBeTrue(status.GetStatusUpdating())
	shouldBeTrue(status.GetStatusCanaryUpdating())

	status.SetStatusCreating(false)
	status.SetStatusNormal(false)
	status.SetStatusUpdating(false)
	status.SetStatusCanaryUpdating(false)

	shouldBeFalse(status.GetStatusCreating())
	shouldBeFalse(status.GetStatusNormal())
	shouldBeFalse(status.GetStatusUpdating())
	shouldBeFalse(status.GetStatusCanaryUpdating())
}

func TestStatusWithUnset(t *testing.T) {
	var status StatusWithUnset

	shouldBeTrue := func(in bool) {
		if in == false {
			t.Fatalf("input should be true")
		}
	}
	shouldBeFalse := func(in bool) {
		if in == true {
			t.Fatalf("input should be false")
		}
	}
	shouldBeFalse(status.GetStatusWithUnsetCreating1())
	shouldBeFalse(status.GetStatusWithUnsetNormal1())
	shouldBeFalse(status.GetStatusWithUnsetUpdating1())
	shouldBeFalse(status.GetStatusWithUnsetCanaryUpdating1())

	status.SetStatusWithUnsetCreating1()
	status.SetStatusWithUnsetNormal1()
	status.SetStatusWithUnsetUpdating1()
	status.SetStatusWithUnsetCanaryUpdating1()

	shouldBeTrue(status.GetStatusWithUnsetCreating1())
	shouldBeTrue(status.GetStatusWithUnsetNormal1())
	shouldBeTrue(status.GetStatusWithUnsetUpdating1())
	shouldBeTrue(status.GetStatusWithUnsetCanaryUpdating1())

	status.UnSetStatusWithUnsetCreating1()
	status.UnSetStatusWithUnsetNormal1()
	status.UnSetStatusWithUnsetUpdating1()
	status.UnSetStatusWithUnsetCanaryUpdating1()

	shouldBeFalse(status.GetStatusWithUnsetCreating1())
	shouldBeFalse(status.GetStatusWithUnsetNormal1())
	shouldBeFalse(status.GetStatusWithUnsetUpdating1())
	shouldBeFalse(status.GetStatusWithUnsetCanaryUpdating1())

}
