package godash_test

import (
	"fmt"
	"testing"

	"github.com/DinnieJ/godash"
	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	a := 1
	b := 2

	if got := godash.If(true, a, b); got != a {
		t.Errorf("If(true, %d, %d) = %d; want %d", a, b, got, a)
	}

	if got := godash.If(false, a, b); got != b {
		t.Errorf("If(false, %d, %d) = %d; want %d", a, b, got, b)
	}
}

func TestContain(t *testing.T) {
	a := []int{1, 2, 3}
	b := 2
	if !godash.Contain(a, b) {
		t.Errorf("Contain(%v, %d) = false; want true", a, b)
	}

	a = []int{1, 2, 3}
	b = 4
	if godash.Contain(a, b) {
		t.Errorf("Contain(%v, %d) = true; want false", a, b)
	}
}

func TestSprf(t *testing.T) {
	if got := godash.Sprf("%d", 1); got != "1" {
		t.Errorf("Sprf(%d) = %s; want 1", 1, got)
	}

	if got := godash.Sprf("%d %d", 1, 2); got != "1 2" {
		t.Errorf("Sprf(%d %d) = %s; want 1 2", 1, 2, got)
	}
}

func TestMust(t *testing.T) {
	a := 1
	b := 2
	got := godash.Must(a, nil)
	if got != a {
		t.Errorf("Must(%d, %d) = %d; want %d", a, b, got, a)
	}

	assert.Panics(t, func() { godash.Must(got, fmt.Errorf("error")) })
}
