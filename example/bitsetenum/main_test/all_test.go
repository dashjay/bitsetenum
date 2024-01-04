package main_test

type Status uint64

const (
	Creating       Status = 1 << 0
	Normal         Status = 1 << 1
	Updating       Status = 1 << 2
	CanaryUpdating Status = 1 << 3
)

func (s *Status) SetStatusCreating(val bool) {
	if val {
		*s |= Creating
	} else {
		*s &= ^Creating
	}
}
func (s *Status) GetStatusCreating() bool {
	return (*s)&Creating != 0
}
func (s *Status) SetStatusNormal(val bool) {
	if val {
		*s |= Normal
	} else {
		*s &= ^Normal
	}
}
func (s *Status) GetStatusNormal() bool {
	return (*s)&Normal != 0
}
func (s *Status) SetStatusUpdating(val bool) {
	if val {
		*s |= Updating
	} else {
		*s &= ^Updating
	}
}
func (s *Status) GetStatusUpdating() bool {
	return (*s)&Updating != 0
}
func (s *Status) SetStatusCanaryUpdating(val bool) {
	if val {
		*s |= CanaryUpdating
	} else {
		*s &= ^CanaryUpdating
	}
}
func (s *Status) GetStatusCanaryUpdating() bool {
	return (*s)&CanaryUpdating != 0
}

type StatusWithUnset uint64

const (
	Creating1       StatusWithUnset = 1 << 0
	Normal1         StatusWithUnset = 1 << 1
	Updating1       StatusWithUnset = 1 << 2
	CanaryUpdating1 StatusWithUnset = 1 << 3
)

func (s *StatusWithUnset) SetStatusWithUnsetCreating1() {
	*s |= Creating1
}
func (s *StatusWithUnset) UnSetStatusWithUnsetCreating1() {
	*s &= ^Creating1
}
func (s *StatusWithUnset) GetStatusWithUnsetCreating1() bool {
	return (*s)&Creating1 != 0
}
func (s *StatusWithUnset) SetStatusWithUnsetNormal1() {
	*s |= Normal1
}
func (s *StatusWithUnset) UnSetStatusWithUnsetNormal1() {
	*s &= ^Normal1
}
func (s *StatusWithUnset) GetStatusWithUnsetNormal1() bool {
	return (*s)&Normal1 != 0
}
func (s *StatusWithUnset) SetStatusWithUnsetUpdating1() {
	*s |= Updating1
}
func (s *StatusWithUnset) UnSetStatusWithUnsetUpdating1() {
	*s &= ^Updating1
}
func (s *StatusWithUnset) GetStatusWithUnsetUpdating1() bool {
	return (*s)&Updating1 != 0
}
func (s *StatusWithUnset) SetStatusWithUnsetCanaryUpdating1() {
	*s |= CanaryUpdating1
}
func (s *StatusWithUnset) UnSetStatusWithUnsetCanaryUpdating1() {
	*s &= ^CanaryUpdating1
}
func (s *StatusWithUnset) GetStatusWithUnsetCanaryUpdating1() bool {
	return (*s)&CanaryUpdating1 != 0
}
