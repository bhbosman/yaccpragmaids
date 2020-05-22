package yaccpragmaids

type LocalVersion struct {
	value string
}

func NewLocalVersion(value string) ILocalVersion {
	return &LocalVersion{value: value}
}

func (self LocalVersion) IsITypeVersion() bool {
	return true
}

func (self LocalVersion) IsLocalVersion() bool {
	return true
}

func (self LocalVersion) VersionType() string {
	return "Local"
}
