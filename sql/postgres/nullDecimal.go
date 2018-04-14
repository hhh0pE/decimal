package postgres

import (
	"database/sql/driver"
)

type NullDecimal struct {
	Decimal
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullDecimal) Scan(value interface{}) error {
	if value == nil {
		n.Valid = false
		return nil
	}
	if err := n.Decimal.Scan(value); err != nil {
		return err
	} else {
		n.Valid = true
		return nil
	}
}

// Value implements the driver Valuer interface.
func (n NullDecimal) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Decimal.Value()
}



