package random

import (
	"math/rand"
	"time"
)

const (
	readableString = "ABCDEFHJLMNQRTUVWXYZabcefghijkmnopqrtuvwxyz23479"
	alnumString    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Returns non-negative integer in range of 0 to maxValue
func Int(maxValue int) int {
	return rand.Intn(maxValue + 1)
}

// Returns [length] long alpha-numeric random string
func String(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = alnumString[rand.Intn(len(alnumString)-1)]
	}
	return string(buf)
}

// Returns [length] long human readable random string
func StringReadable(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = readableString[rand.Intn(len(readableString)-1)]
	}
	return string(buf)
}
