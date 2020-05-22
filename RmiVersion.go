package yaccpragmaids

type RmiVersion struct {
	className               string
	hashCode                uint64
	serializationVersionUID uint64
}

func (self RmiVersion) SerializationVersionUID() uint64 {
	return self.serializationVersionUID
}

func (self RmiVersion) Hash() uint64 {
	return self.hashCode
}

func (self RmiVersion) ClassName() string {
	return self.className
}

func NewRmiVersion(className string, hashCode uint64, serializationVersionUID uint64) IRmiVersion {
	return &RmiVersion{
		className:               className,
		hashCode:                hashCode,
		serializationVersionUID: serializationVersionUID,
	}
}

func (self RmiVersion) IsITypeVersion() bool {
	return true
}

func (self RmiVersion) VersionType() string {
	return "RMI"
}

func (self RmiVersion) IsRmiVersion() bool {
	return true
}
