package mod1

import (
	"github.com/google/uuid"
)

func Foo() string {
	return uuid.New().String()
}

