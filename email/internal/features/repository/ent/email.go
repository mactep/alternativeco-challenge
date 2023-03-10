// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/mactep/alternativeco-challenge/email/internal/features/repository/ent/email"
)

// Email is the model entity for the Email schema.
type Email struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Email) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case email.FieldID:
			values[i] = new(sql.NullInt64)
		case email.FieldEmail:
			values[i] = new(sql.NullString)
		case email.FieldCreatedAt, email.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Email", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Email fields.
func (e *Email) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case email.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case email.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				e.Email = value.String
			}
		case email.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case email.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Email.
// Note that you need to call Email.Unwrap() before calling this method if this Email
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Email) Update() *EmailUpdateOne {
	return (&EmailClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Email entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Email) Unwrap() *Email {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Email is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Email) String() string {
	var builder strings.Builder
	builder.WriteString("Email(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("email=")
	builder.WriteString(e.Email)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Emails is a parsable slice of Email.
type Emails []*Email

func (e Emails) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
