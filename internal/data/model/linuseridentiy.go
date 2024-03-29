// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"lin-cms-go/internal/data/model/linuseridentiy"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// LinUserIdentiy is the model entity for the LinUserIdentiy schema.
type LinUserIdentiy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	// 用户id
	UserID int `json:"user_id,omitempty"`
	// IdentityType holds the value of the "identity_type" field.
	IdentityType string `json:"identity_type,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier string `json:"identifier,omitempty"`
	// Credential holds the value of the "credential" field.
	Credential                string `json:"credential,omitempty"`
	lin_user_lin_user_identiy *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LinUserIdentiy) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case linuseridentiy.FieldID, linuseridentiy.FieldUserID:
			values[i] = new(sql.NullInt64)
		case linuseridentiy.FieldIdentityType, linuseridentiy.FieldIdentifier, linuseridentiy.FieldCredential:
			values[i] = new(sql.NullString)
		case linuseridentiy.FieldCreateTime, linuseridentiy.FieldUpdateTime, linuseridentiy.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		case linuseridentiy.ForeignKeys[0]: // lin_user_lin_user_identiy
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LinUserIdentiy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LinUserIdentiy fields.
func (lui *LinUserIdentiy) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case linuseridentiy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lui.ID = int(value.Int64)
		case linuseridentiy.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				lui.CreateTime = value.Time
			}
		case linuseridentiy.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				lui.UpdateTime = value.Time
			}
		case linuseridentiy.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				lui.DeleteTime = value.Time
			}
		case linuseridentiy.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				lui.UserID = int(value.Int64)
			}
		case linuseridentiy.FieldIdentityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identity_type", values[i])
			} else if value.Valid {
				lui.IdentityType = value.String
			}
		case linuseridentiy.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				lui.Identifier = value.String
			}
		case linuseridentiy.FieldCredential:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field credential", values[i])
			} else if value.Valid {
				lui.Credential = value.String
			}
		case linuseridentiy.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field lin_user_lin_user_identiy", value)
			} else if value.Valid {
				lui.lin_user_lin_user_identiy = new(int)
				*lui.lin_user_lin_user_identiy = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LinUserIdentiy.
// Note that you need to call LinUserIdentiy.Unwrap() before calling this method if this LinUserIdentiy
// was returned from a transaction, and the transaction was committed or rolled back.
func (lui *LinUserIdentiy) Update() *LinUserIdentiyUpdateOne {
	return (&LinUserIdentiyClient{config: lui.config}).UpdateOne(lui)
}

// Unwrap unwraps the LinUserIdentiy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lui *LinUserIdentiy) Unwrap() *LinUserIdentiy {
	tx, ok := lui.config.driver.(*txDriver)
	if !ok {
		panic("model: LinUserIdentiy is not a transactional entity")
	}
	lui.config.driver = tx.drv
	return lui
}

// String implements the fmt.Stringer.
func (lui *LinUserIdentiy) String() string {
	var builder strings.Builder
	builder.WriteString("LinUserIdentiy(")
	builder.WriteString(fmt.Sprintf("id=%v", lui.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(lui.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(lui.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(lui.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", lui.UserID))
	builder.WriteString(", identity_type=")
	builder.WriteString(lui.IdentityType)
	builder.WriteString(", identifier=")
	builder.WriteString(lui.Identifier)
	builder.WriteString(", credential=")
	builder.WriteString(lui.Credential)
	builder.WriteByte(')')
	return builder.String()
}

// LinUserIdentiys is a parsable slice of LinUserIdentiy.
type LinUserIdentiys []*LinUserIdentiy

func (lui LinUserIdentiys) config(cfg config) {
	for _i := range lui {
		lui[_i].config = cfg
	}
}
