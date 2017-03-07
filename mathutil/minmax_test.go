package mathutil

import "testing"

func TestMinInt8(t *testing.T) {
	if MinInt8(1, 2) != 1 {
		t.Fail()
	}
	if MinInt8(2, 1) != 1 {
		t.Fail()
	}
	if MinInt8(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxInt8(t *testing.T) {
	if MaxInt8(1, 2) != 2 {
		t.Fail()
	}
	if MaxInt8(2, 1) != 2 {
		t.Fail()
	}
	if MaxInt8(2, 2) != 2 {
		t.Fail()
	}
}

func TestMinUint8(t *testing.T) {
	if MinUint8(1, 2) != 1 {
		t.Fail()
	}
	if MinUint8(2, 1) != 1 {
		t.Fail()
	}
	if MinUint8(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxUint8(t *testing.T) {
	if MaxUint8(1, 2) != 2 {
		t.Fail()
	}
	if MaxUint8(2, 1) != 2 {
		t.Fail()
	}
	if MaxUint8(2, 2) != 2 {
		t.Fail()
	}
}

func TestMinInt32(t *testing.T) {
	if MinInt32(1, 2) != 1 {
		t.Fail()
	}
	if MinInt32(2, 1) != 1 {
		t.Fail()
	}
	if MinInt32(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxInt32(t *testing.T) {
	if MaxInt32(1, 2) != 2 {
		t.Fail()
	}
	if MaxInt32(2, 1) != 2 {
		t.Fail()
	}
	if MaxInt32(2, 2) != 2 {
		t.Fail()
	}
}

func TestMinUint32(t *testing.T) {
	if MinUint32(1, 2) != 1 {
		t.Fail()
	}
	if MinUint32(2, 1) != 1 {
		t.Fail()
	}
	if MinUint32(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxUint32(t *testing.T) {
	if MaxUint32(1, 2) != 2 {
		t.Fail()
	}
	if MaxUint32(2, 1) != 2 {
		t.Fail()
	}
	if MaxUint32(2, 2) != 2 {
		t.Fail()
	}
}

func TestMinInt64(t *testing.T) {
	if MinInt64(1, 2) != 1 {
		t.Fail()
	}
	if MinInt64(2, 1) != 1 {
		t.Fail()
	}
	if MinInt64(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxInt64(t *testing.T) {
	if MaxInt64(1, 2) != 2 {
		t.Fail()
	}
	if MaxInt64(2, 1) != 2 {
		t.Fail()
	}
	if MaxInt64(2, 2) != 2 {
		t.Fail()
	}
}

func TestMinUint64(t *testing.T) {
	if MinUint64(1, 2) != 1 {
		t.Fail()
	}
	if MinUint64(2, 1) != 1 {
		t.Fail()
	}
	if MinUint64(1, 1) != 1 {
		t.Fail()
	}
}

func TestMaxUint64(t *testing.T) {
	if MaxUint64(1, 2) != 2 {
		t.Fail()
	}
	if MaxUint64(2, 1) != 2 {
		t.Fail()
	}
	if MaxUint64(2, 2) != 2 {
		t.Fail()
	}
}
