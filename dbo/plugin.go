package dbo

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Cyberpull/gokit/dbo/scopes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type xPluginCallback func(db *gorm.DB)

type xPlugin struct{}

func (x *xPlugin) Name() string {
	return "gokit-dbo"
}

func (x *xPlugin) Initialize(db *gorm.DB) (err error) {
	db.Callback().Query().Before("*").Register("gokit-dbo:before_query", x.onBeforeQuery())
	db.Callback().Query().After("*").Register("gokit-dbo:after_query", x.onAfterQuery())

	return
}

// ==================================

func (x *xPlugin) onBeforeQuery() xPluginCallback {
	return func(db *gorm.DB) {
		defer func() {
			_ = recover()
		}()

		model := reflect.New(db.Statement.Schema.ModelType)

		// Process Scopes
		for i := 0; i < model.NumMethod(); i++ {
			name := model.Type().Method(i).Name

			if strings.HasPrefix(name, "Scope") {
				method, ok := model.Method(i).Interface().(scopes.Scope)

				if ok {
					db = method(db)
				}
			}
		}

		// Process Tags
		for _, field := range db.Statement.Schema.Fields {
			method := model.MethodByName("Preload" + field.Name)

			if method.IsValid() && !method.IsZero() {
				db = db.Preload(field.Name, method.Interface())
			}
		}
	}
}

func (x *xPlugin) onAfterQuery() xPluginCallback {
	return func(db *gorm.DB) {
		// for _, field := range db.Statement.Schema.Fields {
		// 	tag := x.getFieldTag(field)
		// }
	}
}

func (x *xPlugin) key(field *schema.Field) (value string) {
	return fmt.Sprintf("%v.%v", field.Schema.Name, field.Name)
}

// ==============================================

func NewPlugin() (v *xPlugin) {
	return &xPlugin{}
}
