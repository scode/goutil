package mathutil

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMinInt8(t *testing.T) {
	assert.Equal(t, int8(1), MinInt8(1, 2))
	assert.Equal(t, int8(1), MinInt8(2, 1))
	assert.Equal(t, int8(1), MinInt8(1, 1))
}

func TestMaxInt8(t *testing.T) {
	assert.Equal(t, int8(2), MaxInt8(1, 2))
	assert.Equal(t, int8(2), MaxInt8(2, 1))
	assert.Equal(t, int8(2), MaxInt8(2, 2))
}

func TestMinUint8(t *testing.T) {
	assert.Equal(t, uint8(1), MinUint8(1, 2))
	assert.Equal(t, uint8(1), MinUint8(2, 1))
	assert.Equal(t, uint8(1), MinUint8(1, 1))
}

func TestMaxUint8(t *testing.T) {
	assert.Equal(t, uint8(2), MaxUint8(1, 2))
	assert.Equal(t, uint8(2), MaxUint8(2, 1))
	assert.Equal(t, uint8(2), MaxUint8(2, 2))
}

func TestMinInt32(t *testing.T) {
	assert.Equal(t, int32(1), MinInt32(1, 2))
	assert.Equal(t, int32(1), MinInt32(2, 1))
	assert.Equal(t, int32(1), MinInt32(1, 1))
}

func TestMaxInt32(t *testing.T) {
	assert.Equal(t, int32(2), MaxInt32(1, 2))
	assert.Equal(t, int32(2), MaxInt32(2, 1))
	assert.Equal(t, int32(2), MaxInt32(2, 2))
}

func TestMinUint32(t *testing.T) {
	assert.Equal(t, uint32(1), MinUint32(1, 2))
	assert.Equal(t, uint32(1), MinUint32(2, 1))
	assert.Equal(t, uint32(1), MinUint32(1, 1))
}

func TestMaxUint32(t *testing.T) {
	assert.Equal(t, uint32(2), MaxUint32(1, 2))
	assert.Equal(t, uint32(2), MaxUint32(2, 1))
	assert.Equal(t, uint32(2), MaxUint32(2, 2))
}

func TestMinInt64(t *testing.T) {
	assert.Equal(t, int64(1), MinInt64(1, 2))
	assert.Equal(t, int64(1), MinInt64(2, 1))
	assert.Equal(t, int64(1), MinInt64(1, 1))
}

func TestMaxInt64(t *testing.T) {
	assert.Equal(t, int64(2), MaxInt64(1, 2))
	assert.Equal(t, int64(2), MaxInt64(2, 1))
	assert.Equal(t, int64(2), MaxInt64(2, 2))
}

func TestMinUint64(t *testing.T) {
	assert.Equal(t, uint64(1), MinUint64(1, 2))
	assert.Equal(t, uint64(1), MinUint64(2, 1))
	assert.Equal(t, uint64(1), MinUint64(1, 1))
}

func TestMaxUint64(t *testing.T) {
	assert.Equal(t, uint64(2), MaxUint64(1, 2))
	assert.Equal(t, uint64(2), MaxUint64(2, 1))
	assert.Equal(t, uint64(2), MaxUint64(2, 2))
}
