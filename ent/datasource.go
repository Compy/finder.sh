// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/compy/finder.sh/ent/datasource"
)

// DataSource is the model entity for the DataSource schema.
type DataSource struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Config holds the value of the "config" field.
	Config string `json:"config,omitempty"`
	// LastIndexed holds the value of the "last_indexed" field.
	LastIndexed time.Time `json:"last_indexed,omitempty"`
	// DateAdded holds the value of the "date_added" field.
	DateAdded    time.Time `json:"date_added,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DataSource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case datasource.FieldID:
			values[i] = new(sql.NullInt64)
		case datasource.FieldName, datasource.FieldType, datasource.FieldStatus, datasource.FieldConfig:
			values[i] = new(sql.NullString)
		case datasource.FieldLastIndexed, datasource.FieldDateAdded:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DataSource fields.
func (ds *DataSource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case datasource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ds.ID = int(value.Int64)
		case datasource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ds.Name = value.String
			}
		case datasource.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ds.Type = value.String
			}
		case datasource.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ds.Status = value.String
			}
		case datasource.FieldConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value.Valid {
				ds.Config = value.String
			}
		case datasource.FieldLastIndexed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_indexed", values[i])
			} else if value.Valid {
				ds.LastIndexed = value.Time
			}
		case datasource.FieldDateAdded:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_added", values[i])
			} else if value.Valid {
				ds.DateAdded = value.Time
			}
		default:
			ds.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DataSource.
// This includes values selected through modifiers, order, etc.
func (ds *DataSource) Value(name string) (ent.Value, error) {
	return ds.selectValues.Get(name)
}

// Update returns a builder for updating this DataSource.
// Note that you need to call DataSource.Unwrap() before calling this method if this DataSource
// was returned from a transaction, and the transaction was committed or rolled back.
func (ds *DataSource) Update() *DataSourceUpdateOne {
	return NewDataSourceClient(ds.config).UpdateOne(ds)
}

// Unwrap unwraps the DataSource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ds *DataSource) Unwrap() *DataSource {
	_tx, ok := ds.config.driver.(*txDriver)
	if !ok {
		panic("ent: DataSource is not a transactional entity")
	}
	ds.config.driver = _tx.drv
	return ds
}

// String implements the fmt.Stringer.
func (ds *DataSource) String() string {
	var builder strings.Builder
	builder.WriteString("DataSource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ds.ID))
	builder.WriteString("name=")
	builder.WriteString(ds.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ds.Type)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(ds.Status)
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(ds.Config)
	builder.WriteString(", ")
	builder.WriteString("last_indexed=")
	builder.WriteString(ds.LastIndexed.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_added=")
	builder.WriteString(ds.DateAdded.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DataSources is a parsable slice of DataSource.
type DataSources []*DataSource
