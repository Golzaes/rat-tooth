package idgen

import "github.com/gofrs/uuid"

type UuidArgs struct {
	u    uuid.UUID
	name string
}

func NewUuidArgs(u byte, name string) *UuidArgs {
	return &UuidArgs{uuid.UUID{u}, name}
}

// IsValidUUID determine a UUID validity
func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}

// GenerateUUID1 Generates a UUID1.
func (a *UuidArgs) GenerateUUID1() (string, error) {
	v1, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return v1.String(), nil
}

// GenerateUUID3 Generates a UUID3.
func (a *UuidArgs) GenerateUUID3() (string, error) {
	v3 := uuid.NewV3(a.u, a.name)
	return v3.String(), nil
}

// GenerateUUID4 Generates a UUID4.
func (a *UuidArgs) GenerateUUID4() (string, error) {
	v4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return v4.String(), nil
}

// GenerateUUID5 Generates a UUID5.
func (a *UuidArgs) GenerateUUID5() (string, error) {
	v5 := uuid.NewV5(a.u, a.name)
	return v5.String(), nil
}

// GenerateUUID6 Generates a UUID6.
func (a *UuidArgs) GenerateUUID6() (string, error) {
	v6, err := uuid.NewV6()
	if err != nil {
		return "", err
	}
	return v6.String(), nil
}

// GenerateUUID7 Generates a UUID7.
func (a *UuidArgs) GenerateUUID7() (string, error) {
	v7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return v7.String(), nil
}
