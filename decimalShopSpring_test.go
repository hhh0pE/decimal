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