package uuidx

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func UUID4() string {
	return uuid.New().String()
}

func RearrangeStrUUID4(uuid string) (string, error) {
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		return "", errors.New("invalid uuid")
	}
	return parts[2] + parts[1] + parts[0] + parts[3] + parts[4], nil
}

func RearrangeUUID4() string {
	rearrangeStrUUID4, _ := RearrangeStrUUID4(UUID4())
	return rearrangeStrUUID4
}
