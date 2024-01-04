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
