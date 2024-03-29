package main

import "testing"

func TestFuelRequired(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Single Mass: 12", func(t *testing.T) {
		got := fuelRequired(12)
		want := 2
		assertCorrectMessage(t, got, want)
	})
	t.Run("Single Mass: 14", func(t *testing.T) {
		got := fuelRequired(14)
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("Single Mass: 1969", func(t *testing.T) {
		got := fuelRequired(1969)
		want := 654
		assertCorrectMessage(t, got, want)
	})
	t.Run("Single Mass: 100756", func(t *testing.T) {
		got := fuelRequired(100756)
		want := 33583
		assertCorrectMessage(t, got, want)
	})
}

func TestTotalFuelRequired(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Multiple Mass: 12, 14", func(t *testing.T) {
		masses := []int{12, 14}
		got := totalFuelRequired(masses)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}

func TestTotalFuelRequiredPart2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Fuel Remainder Single Mass Part II: 654", func(t *testing.T) {
		masses := []int{1969}
		got := totalFuelRequiredPartII(masses)
		want := 966
		assertCorrectMessage(t, got, want)
	})
	t.Run("Fuel Remainder Single Mass Part II: 100756", func(t *testing.T) {
		masses := []int{100756}
		got := totalFuelRequiredPartII(masses)
		want := 50346
		assertCorrectMessage(t, got, want)
	})
	t.Run("No Fuel Remainder Multiple Mass Part II: 12, 14", func(t *testing.T) {
		masses := []int{12, 14}
		got := totalFuelRequiredPartII(masses)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}
