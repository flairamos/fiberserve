package tool

import (
	"github.com/google/uuid"
)

// UUID 生成没有破折号的 UUID
func Uuid() string {
	return uuid.NewString()
}
