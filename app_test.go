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
func Test_luckyDraw(t *testing.T) {
	got := luckyDraw(1221)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}

func Test_luckyDraw_false(t *testing.T) {
	got := luckyDraw(1234)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isLuckyDraw(t *testing.T) {
	got := isLuckyDraw(1221)
	expected := true
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_isLuckyDraw_false(t *testing.T) {
	got := isLuckyDraw(1234)
	expected := false
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade(t *testing.T) {
	got := grade(90)
	expected := ("A")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_gradeB(t *testing.T) {
	got := grade(80)
	expected := ("b")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_gradeC(t *testing.T) {
	got := grade(70)
	expected := ("C")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_gradeD(t *testing.T) {
	got := grade(60)
	expected := ("D")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_gradeE(t *testing.T) {
	got := grade(50)
	expected := "E"
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_gradeFail(t *testing.T) {
	got := grade(45)
	expected := "fail"
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2(t *testing.T) {
	got := grade(90)
	expected := ("A")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2B(t *testing.T) {
	got := grade(80)
	expected := ("b")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2C(t *testing.T) {
	got := grade(70)
	expected := ("C")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2D(t *testing.T) {
	got := grade(60)
	expected := ("D")
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2E(t *testing.T) {
	got := grade(50)
	expected := "E"
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
func Test_grade2Fail(t *testing.T) {
	got := grade(45)
	expected := "fail"
	if got != expected {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
