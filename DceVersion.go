package yaccpragmaids

import "github.com/google/uuid"

type DceVersion struct {
	uid     uuid.UUID
	version IVersion
}

func (self DceVersion) Value() IVersion {
	return self.version
}

func (self DceVersion) Identifier() string {
	return self.uid.String()
}

func NewDceVersion(uid string, version IVersion) (IDceVersion, error) {
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}
	return &DceVersion{uid: uuid, version: version}, nil
}

func (self DceVersion) IsITypeVersion() bool {
	return true
}

func (self DceVersion) VersionType() string {
	return "DCE"
}

func (self DceVersion) IsDceVersion() bool {
	return true
}
