package scopes

import (
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type xCore struct{}

func (x *xCore) Schema(db *gorm.DB, model any) (*Schema, error) {
	return schema.Parse(model, &sync.Map{}, db.NamingStrategy)
}

func (x *xCore) Statement(db *gorm.DB, model any) *Statement {
	data, _ := x.Schema(db, model)

	return &Statement{
		DB:      db,
		Table:   data.Table,
		Schema:  data,
		Clauses: map[string]Clause{},
	}
}

func (x *xCore) TableName(table ...string) string {
	if len(table) > 0 && table[0] != "" {
		return table[0]
	}

	return CurrentTable
}

func (x *xCore) FindInSet(value any, column Column) Expression {
	return Expr{
		SQL: "FIND_IN_SET",
	}
}

func (x *xCore) Sum(column string) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(Select{Expression: Expr{
			SQL: "COALESCE(SUM(?), 0)",
			Vars: []any{Column{
				Table: CurrentTable,
				Name:  column,
			}},
			WithoutParentheses: true,
		}})
	}
}

// func (x *xCore) IN(cond ...IN) Expression {
// 	return
// }

// ===========================

var Core xCore
