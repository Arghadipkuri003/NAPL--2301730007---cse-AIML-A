package main

import (
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("Divide(10, 2) returned unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("Divide(10, 2) = %v; want 5", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Fatal("Divide(10, 0) should return an error")
		}
		if !strings.Contains(err.Error(), "division by zero") {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}

func TestSquareRoot(t *testing.T) {
	t.Run("valid square root", func(t *testing.T) {
		got, err := SquareRoot(25)
		if err != nil {
			t.Fatalf("SquareRoot(25) returned unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("SquareRoot(25) = %v; want 5", got)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		_, err := SquareRoot(-4)
		if err == nil {
			t.Fatal("SquareRoot(-4) should return an error")
		}
		if !strings.Contains(err.Error(), "negative square root") {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}

func TestParseInput(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		got, err := ParseInput(" 12.5 ")
		if err != nil {
			t.Fatalf("ParseInput returned unexpected error: %v", err)
		}
		if got != 12.5 {
			t.Errorf("ParseInput() = %v; want 12.5", got)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		_, err := ParseInput("   ")
		if err == nil {
			t.Fatal("ParseInput with empty input should return an error")
		}
		if !strings.Contains(err.Error(), "empty input") {
			t.Errorf("unexpected error message: %v", err)
		}
	})

	t.Run("non-numeric input", func(t *testing.T) {
		_, err := ParseInput("hello")
		if err == nil {
			t.Fatal("ParseInput with non-numeric input should return an error")
		}
		if !strings.Contains(err.Error(), "invalid input") {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}
