package yaccpragmaids

type IdlVersion struct {
	identifier string
	version    IVersion
}

func (self IdlVersion) Value() IVersion {
	return self.version
}

func (self IdlVersion) Identifier() string {
	return self.identifier
}

func (self IdlVersion) VersionType() string {
	return "IDL"
}

func NewIdlVersion(identifier string, version IVersion) IIdlVersion {
	return &IdlVersion{identifier: identifier, version: version}
}

func (self IdlVersion) IsITypeVersion() bool {
	return true
}

func (self IdlVersion) IsIIdlVersion() bool {
	return true
}
