package main

import (
	"context"
	"testing"
)

func TestBoilSuccess(t *testing.T) {
	cups := 1 * CupsCoffee
	ctx := context.Background()
	water := cups.Water()

	expect := HotWater(water)
	result, err := boil(ctx, water)

	if err != nil {
		t.Errorf("failed boil test %v", err)
	}

	if result != expect {
		t.Errorf("expected: %v, got: %v", expect, result)
	}
}

func TestBoilFail(t *testing.T) {
	cups2 := 4 * CupsCoffee
	ctx2 := context.Background()
	water := cups2.Water()

	expect := HotWater(0)

	result, _ := boil(ctx2, water)
	if result != expect {
		t.Errorf("expected: %v, got: %v", expect, result)
	}
}
