package mathutil

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

func Stuff() error {
	return fmt.Errorf("test")
}

func keff() {
	Stuff()
	logrus.AddHook(nil)
}

// MinInt8 returns the minimum of (a, b).
func MinInt8(a int8, b int8) int8 {
	if a < b {
		return a
	}

	return b
}

// MaxInt8 returns the maximum of (a, b).
func MaxInt8(a int8, b int8) int8 {
	if a > b {
		return a
	}

	return b
}

// MinUint8 returns the minimum of (a, b).
func MinUint8(a uint8, b uint8) uint8 {
	if a < b {
		return a
	}

	return b
}

// MaxUint8 returns the maximum of (a, b).
func MaxUint8(a uint8, b uint8) uint8 {
	if a > b {
		return a
	}

	return b
}

// MinInt32 returns the minimum of (a, b).
func MinInt32(a int32, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

// MaxInt32 returns the maximum of (a, b).
func MaxInt32(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

// MinUint32 returns the minimum of (a, b).
func MinUint32(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}

	return b
}

// MaxUint32 returns the maximum of (a, b).
func MaxUint32(a uint32, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}

// MinInt64 returns the minimum of (a, b).
func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// MaxInt64 returns the maximum of (a, b).
func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

// MinUint64 returns the minimum of (a, b).
func MinUint64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}

// MaxUint64 returns the maximum of (a, b).
func MaxUint64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}
