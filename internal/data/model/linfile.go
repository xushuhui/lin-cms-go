// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"lin-cms-go/internal/data/model/linfile"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// LinFile is the model entity for the LinFile schema.
type LinFile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Type holds the value of the "type" field.
	// 1 LOCAL 本地，2 REMOTE 远程
	Type int8 `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Extension holds the value of the "extension" field.
	Extension string `json:"extension,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Md5 holds the value of the "md5" field.
	// md5值，防止上传重复文件
	Md5 string `json:"md5,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LinFile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case linfile.FieldID, linfile.FieldType, linfile.FieldSize:
			values[i] = new(sql.NullInt64)
		case linfile.FieldPath, linfile.FieldName, linfile.FieldExtension, linfile.FieldMd5:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LinFile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LinFile fields.
func (lf *LinFile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case linfile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lf.ID = int(value.Int64)
		case linfile.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				lf.Path = value.String
			}
		case linfile.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				lf.Type = int8(value.Int64)
			}
		case linfile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				lf.Name = value.String
			}
		case linfile.FieldExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension", values[i])
			} else if value.Valid {
				lf.Extension = value.String
			}
		case linfile.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				lf.Size = int(value.Int64)
			}
		case linfile.FieldMd5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5", values[i])
			} else if value.Valid {
				lf.Md5 = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LinFile.
// Note that you need to call LinFile.Unwrap() before calling this method if this LinFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (lf *LinFile) Update() *LinFileUpdateOne {
	return (&LinFileClient{config: lf.config}).UpdateOne(lf)
}

// Unwrap unwraps the LinFile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lf *LinFile) Unwrap() *LinFile {
	tx, ok := lf.config.driver.(*txDriver)
	if !ok {
		panic("model: LinFile is not a transactional entity")
	}
	lf.config.driver = tx.drv
	return lf
}

// String implements the fmt.Stringer.
func (lf *LinFile) String() string {
	var builder strings.Builder
	builder.WriteString("LinFile(")
	builder.WriteString(fmt.Sprintf("id=%v", lf.ID))
	builder.WriteString(", path=")
	builder.WriteString(lf.Path)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", lf.Type))
	builder.WriteString(", name=")
	builder.WriteString(lf.Name)
	builder.WriteString(", extension=")
	builder.WriteString(lf.Extension)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", lf.Size))
	builder.WriteString(", md5=")
	builder.WriteString(lf.Md5)
	builder.WriteByte(')')
	return builder.String()
}

// LinFiles is a parsable slice of LinFile.
type LinFiles []*LinFile

func (lf LinFiles) config(cfg config) {
	for _i := range lf {
		lf[_i].config = cfg
	}
}
