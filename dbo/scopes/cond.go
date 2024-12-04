package scopes

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type xCond struct{}

func (x *xCond) PrimaryKey(value any) Scope {
	return x.Equal(PrimaryKey, value)
}

func (x *xCond) Equal(field string, value any) Scope {
	return x.Where(Eq{
		Column: Column{Table: CurrentTable, Name: field},
		Value:  value,
	})
}

func (x *xCond) NotEqual(field string, value any) Scope {
	return x.Where(Neq{
		Column: Column{Table: CurrentTable, Name: field},
		Value:  value,
	})
}

func (x *xCond) Where(expr ...Expression) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(Where{Exprs: expr})
	}
}

func (x *xCond) Not(expr ...Expression) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Not(expr...))
	}
}

func (x *xCond) WhereIn(expr ...IN) Scope {
	return x.Where(ex(expr)...)
}

var Cond xCond
