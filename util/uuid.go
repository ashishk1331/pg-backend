package util

import (
	"github.com/google/uuid"
 	"strings"
)

func UUID() string {
	id := uuid.New().String()
	return strings.ReplaceAll(id, "-", "")
}