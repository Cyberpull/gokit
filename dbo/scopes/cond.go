package scopes

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type xCond struct{}

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
