package fizzbuzz

import "testing"


func TestSay1Should1(t *testing.T) {
	r := Say(1)

	if r != "1" {
		t.Error("Expect 1 but got", r)	
	}
}

func TestSay2Should2(t *testing.T) {
	r := Say(2)

	if r != "2" {
		t.Error("Expect 2 but got", r)
	}
}

func TestSay3ShouldFizz(t *testing.T) {
	r := Say(3)

	if r != "Fizz" {
		t.Error("Expect Fizz but got", r)
	}
}

func TestSay5ShouldBuzz(t *testing.T) {
	r := Say(5)

	if r != "Buzz" {
		t.Error("Expect Buzz but got", r)
	}
}

func TestSay6ShouldFizz(t *testing.T) {
	r := Say(6)

	if r != "Fizz" {
		t.Error("Expect Fizz but got", r)
	}
}

func TestSay10ShouldBuzz(t *testing.T) {
	r := Say(10)

	if r != "Buzz" {
		t.Error("Expect Buzz but got", r)
	}
}

func TestSay15ShouldFizzBuzz(t *testing.T) {
	r := Say(15)

	if r != "FizzBuzz" {
		t.Error("Expect FizzBuzz but got", r)
	}
}