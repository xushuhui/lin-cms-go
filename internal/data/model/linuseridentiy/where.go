// Code generated by entc, DO NOT EDIT.

package linuseridentiy

import (
	"lin-cms-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// IdentityType applies equality check predicate on the "identity_type" field. It's identical to IdentityTypeEQ.
func IdentityType(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentityType), v))
	})
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// Credential applies equality check predicate on the "credential" field. It's identical to CredentialEQ.
func Credential(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredential), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// IdentityTypeEQ applies the EQ predicate on the "identity_type" field.
func IdentityTypeEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeNEQ applies the NEQ predicate on the "identity_type" field.
func IdentityTypeNEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeIn applies the In predicate on the "identity_type" field.
func IdentityTypeIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdentityType), v...))
	})
}

// IdentityTypeNotIn applies the NotIn predicate on the "identity_type" field.
func IdentityTypeNotIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdentityType), v...))
	})
}

// IdentityTypeGT applies the GT predicate on the "identity_type" field.
func IdentityTypeGT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeGTE applies the GTE predicate on the "identity_type" field.
func IdentityTypeGTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeLT applies the LT predicate on the "identity_type" field.
func IdentityTypeLT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeLTE applies the LTE predicate on the "identity_type" field.
func IdentityTypeLTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeContains applies the Contains predicate on the "identity_type" field.
func IdentityTypeContains(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeHasPrefix applies the HasPrefix predicate on the "identity_type" field.
func IdentityTypeHasPrefix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeHasSuffix applies the HasSuffix predicate on the "identity_type" field.
func IdentityTypeHasSuffix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeEqualFold applies the EqualFold predicate on the "identity_type" field.
func IdentityTypeEqualFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentityType), v))
	})
}

// IdentityTypeContainsFold applies the ContainsFold predicate on the "identity_type" field.
func IdentityTypeContainsFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentityType), v))
	})
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdentifier), v...))
	})
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdentifier), v...))
	})
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentifier), v))
	})
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentifier), v))
	})
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentifier), v))
	})
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentifier), v))
	})
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentifier), v))
	})
}

// CredentialEQ applies the EQ predicate on the "credential" field.
func CredentialEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredential), v))
	})
}

// CredentialNEQ applies the NEQ predicate on the "credential" field.
func CredentialNEQ(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCredential), v))
	})
}

// CredentialIn applies the In predicate on the "credential" field.
func CredentialIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCredential), v...))
	})
}

// CredentialNotIn applies the NotIn predicate on the "credential" field.
func CredentialNotIn(vs ...string) predicate.LinUserIdentiy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCredential), v...))
	})
}

// CredentialGT applies the GT predicate on the "credential" field.
func CredentialGT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCredential), v))
	})
}

// CredentialGTE applies the GTE predicate on the "credential" field.
func CredentialGTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCredential), v))
	})
}

// CredentialLT applies the LT predicate on the "credential" field.
func CredentialLT(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCredential), v))
	})
}

// CredentialLTE applies the LTE predicate on the "credential" field.
func CredentialLTE(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCredential), v))
	})
}

// CredentialContains applies the Contains predicate on the "credential" field.
func CredentialContains(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCredential), v))
	})
}

// CredentialHasPrefix applies the HasPrefix predicate on the "credential" field.
func CredentialHasPrefix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCredential), v))
	})
}

// CredentialHasSuffix applies the HasSuffix predicate on the "credential" field.
func CredentialHasSuffix(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCredential), v))
	})
}

// CredentialEqualFold applies the EqualFold predicate on the "credential" field.
func CredentialEqualFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCredential), v))
	})
}

// CredentialContainsFold applies the ContainsFold predicate on the "credential" field.
func CredentialContainsFold(v string) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCredential), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LinUserIdentiy) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LinUserIdentiy) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LinUserIdentiy) predicate.LinUserIdentiy {
	return predicate.LinUserIdentiy(func(s *sql.Selector) {
		p(s.Not())
	})
}
