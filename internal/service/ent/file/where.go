// Code generated by entc, DO NOT EDIT.

package file

import (
	"go-rpc-happysooner-file/internal/service/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func IDGT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
