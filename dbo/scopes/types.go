package scopes

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Scope = func(*gorm.DB) *gorm.DB

type Statement = gorm.Statement

type Schema = schema.Schema

type Clause = clause.Clause

type NamedExpr = clause.NamedExpr

type Expression = clause.Expression

type Column = clause.Column

type Values = clause.Values

type Table = clause.Table

type Join = clause.Join

type Where = clause.Where

type Expr = clause.Expr

type IN = clause.IN

type Eq = clause.Eq

type OrderBy = clause.OrderBy

const CurrentTable string = clause.CurrentTable
