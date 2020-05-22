package yaccpragmaids

type Version struct {
	major uint64
	minor uint64
}

func NewVersion(major uint64, minor uint64) IVersion {
	return &Version{major: major, minor: minor}
}

func (self Version) IsIVersion() bool {
	return true
}

func (self Version) Major() uint64 {
	return self.major
}

func (self Version) Minor() uint64 {
	return self.minor
}
