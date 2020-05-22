package yaccpragmaids

type IVersion interface {
	IsIVersion() bool
	Major() uint64
	Minor() uint64
}

type ITypeVersion interface {
	IsITypeVersion() bool
	VersionType() string
}

type IIdlVersion interface {
	ITypeVersion
	IsIIdlVersion() bool
	Identifier() string
	Value() IVersion
}

type ILocalVersion interface {
	ITypeVersion
	IsLocalVersion() bool
}

type IDceVersion interface {
	ITypeVersion
	IsDceVersion() bool
	Identifier() string
	Value() IVersion
}

type IRmiVersion interface {
	ITypeVersion
	IsRmiVersion() bool
	ClassName() string
	Hash() uint64
	SerializationVersionUID() uint64
}

type ISetTypeVersion interface {
	SetTypeValue(version ITypeVersion)
}
type IGetTypeVersion interface {
	GetTypeValue() ITypeVersion
}
