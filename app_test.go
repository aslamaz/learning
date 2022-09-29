package main

import "testing"

func Test_countDigits(t *testing.T) {
	got := countDigits(1234)
	expected := 4
	if got != expected {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
func Test_reverse(t *testing.T) {
	got := reverse(1234)
	expected := 4321
	if got != expected {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
func Test_isPalindrome_false(t *testing.T) {
	got := isPalindrome(1234)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isPalindrome_true(t *testing.T) {
	got := isPalindrome(121)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isArmstrong_true(t *testing.T) {
	got := isArmstrong(153)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isArmstrong_false(t *testing.T) {
	got := isArmstrong(123)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isNarcissistic_true(t *testing.T) {
	got := isNarcissistic(9474)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isNarcissistic_false(t *testing.T) {
	got := isArmstrong(123)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isPrime(t *testing.T) {
	got := isPrime(2)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isPrime_false(t *testing.T) {
	got := isPrime(6)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
