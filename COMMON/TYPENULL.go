package COMMON

import (
	"database/sql"
	"encoding/json"
)

//use when read data from DB and it may be have null value
type MyNullString struct {
	sql.NullString
}

//define value when change to json type
func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

// Unmarshalling into a pointer will let us detect null
func (v *MyNullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

// Unmarshalling into a pointer will let us detect null
func (v *MyNullInt) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

//use when read data from DB and it may be have null value
type MyNullTime struct {
	sql.NullTime
}

//define value when change to json type
func (s MyNullTime) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Time)
	}
	return []byte(`null`), nil
}

//use when read data from DB and it may be have null value
type MyNullInt struct {
	sql.NullInt64
}

//define value when change to json type
func (s MyNullInt) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Int64)
	}
	return []byte(`null`), nil
}
