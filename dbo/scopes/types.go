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

type Neq = clause.Neq

type Like Eq

type OrderBy = clause.OrderBy

type JoinType = clause.JoinType

const (
	Associations string = clause.Associations

	CrossJoin JoinType = clause.CrossJoin
	InnerJoin JoinType = clause.InnerJoin
	LeftJoin  JoinType = clause.LeftJoin
	RightJoin JoinType = clause.RightJoin

	AndWithSpace string = clause.AndWithSpace
	OrWithSpace  string = clause.OrWithSpace

	CurrentTable string = clause.CurrentTable
	PrimaryKey   string = clause.PrimaryKey

	LockingOptionsNoWait     string = clause.LockingOptionsNoWait
	LockingOptionsSkipLocked string = clause.LockingOptionsSkipLocked
	LockingStrengthShare     string = clause.LockingStrengthShare
	LockingStrengthUpdate    string = clause.LockingStrengthUpdate
)

var PrimaryColumn Column

func init() {
	PrimaryColumn = clause.PrimaryColumn
}
