package utils

import (
	"time"
)

func GenerateId() string {
	var Id string

	t := time.Now()
	Id = t.UTC().Format("20060102150405.000000000")

	return Id
}

