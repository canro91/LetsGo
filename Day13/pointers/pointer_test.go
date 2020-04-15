package pointers

import (
	"testing"
)

type Employee struct {
	Seniority    string
	HourlySalary int
	Due          int
}

func (e *Employee) Pay(hours int) {
	e.Due = e.HourlySalary * hours
}

func Promote(e *Employee, level string) {
	// You don't need to dereference (*e).Seniority = level
	e.Seniority = level
}

func TestPointers(t *testing.T) {
	t.Run("Pay", func(t *testing.T) {
		e := Employee{HourlySalary: 10}

		// You don't need to use (&e).Pay(15)
		e.Pay(15)

		balance := 150

		if balance != e.Due {
			t.Errorf("Expected %d but was %d", balance, e.Due)
		}
	})

	t.Run("Promote", func(t *testing.T) {
		// Notice, you need to explicitly declare a pointer here
		e := &Employee{Seniority: "Junior"}

		level := "S1"
		Promote(e, level)

		if e.Seniority != level {
			t.Errorf("Expected %q but was %q", level, e.Seniority)
		}
	})
}
