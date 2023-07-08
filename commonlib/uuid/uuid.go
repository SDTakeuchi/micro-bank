package uuid

import (
	"google/uuid"
)

type UUID struct {
	uuid uuid.UUID
}

func NewUUID() UUID {
	return uuid.New()
}

func (u UUID) String() string {
	return u.uuid.String()
}

func ParseUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}