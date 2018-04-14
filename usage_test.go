package decimal

import (
	"testing"
	"fmt"
)

type TestStruct struct {
	ID int64
	Value Decimal
}

func TestUsage(t *testing.T) {
	d, _ := NewFromString("123.123879")
	fmt.Println(d.Value())
	dd := Decimal{}
	fmt.Println(dd.Value())
	fmt.Println(d)
	fmt.Println(d.Scale())
	fmt.Println(d.Precision())
	d.b.Context.RoundingMode = ToZero
	fmt.Println(d.Round(6))

	d2 := New(0, 0)
	d2.add(d2, New(2, 0))
	fmt.Println(d2)
	fmt.Println(d2.Scale())

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var d3 NullDecimal
	d3.Scan(nil)
	fmt.Println(d3.Decimal, d3.Valid)

	var d4 Decimal
	d4.Scan([]byte("0.111"))
	fmt.Println(d4)
	fmt.Println(d4.b.Scale())
	fmt.Println(d4.IsZero())
	fmt.Println(Decimal{}.IsZero())
	fmt.Println(NewDecimal().IsZero())
}
