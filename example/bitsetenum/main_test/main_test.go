package main_test

import "testing"

func TestAll(t *testing.T) {
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
}
