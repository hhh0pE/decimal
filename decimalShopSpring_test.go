package decimal

import (
	"testing"
	"fmt"
)

func TestDecimalShopspringWrapper(t *testing.T) {
	var d Decimal
	//d := NewDecimal()
	fmt.Printf("%v", d)
}

func TestDecimalFromFloat(t *testing.T) {
	d := NewFromFloat(1.0000004)
	fmt.Println(d)

	d2, _ := NewFromString("1.0000004")
	fmt.Println(d2)
}

func TestDecimalIsZero(t *testing.T) {
	d := NewDecimal()
	if !d.IsZero() {
		t.Error("!")
	}

	d2, _ := NewFromString("0")
	if !d2.IsZero() {
		t.Error("!")
	}

	d3 := NewFromFloat(0)
	if !d3.IsZero() {
		t.Error("!")
	}
}

func TestDivRem(t *testing.T) {

	//var rest Decimal
	var b = NewFromFloat(30)
	if !b.Modulo(NewFromFloat(10)).IsZero() {
		t.Error("!")
	}

}

func TestToString(t *testing.T) {
	var d1 = NewFromFloat(0.00000001)
	var d2, _ = NewFromString("0.00000001")

	d1.b.Context.MaxScale = 20
	fmt.Println(d1.b.Context.RoundingMode)
	d1.b.String()
	fmt.Println(d1)
	fmt.Println(d2)
}