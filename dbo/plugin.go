package dbo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/Cyberpull/gokit/dbo/scopes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type xPluginCallback func(db *gorm.DB)

type xPlugin struct {
	tags map[string]*xPluginTag
}

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
		model := reflect.New(db.Statement.Schema.ModelType)

		// Process Tags
		for _, field := range db.Statement.Schema.Fields {
			tag := x.getFieldTag(field)

			if tag == nil {
				continue
			}

			if tag.Preload {
				args := make([]any, 0)

				method := model.MethodByName(field.Name + "Preloader")

				if method.IsValid() && !method.IsZero() {
					args = append(args, method.Interface())
				}

				db = db.Preload(field.Name, args...)
			}
		}

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
	}
}

func (x *xPlugin) onAfterQuery() xPluginCallback {
	return func(db *gorm.DB) {
		// for _, field := range db.Statement.Schema.Fields {
		// 	tag := x.getFieldTag(field)
		// }
	}
}

// ==================================

func (x *xPlugin) key(field *schema.Field) (value string) {
	return fmt.Sprintf("%v.%v", field.Schema.Name, field.Name)
}

func (x *xPlugin) getFieldTag(field *schema.Field) (tag *xPluginTag) {
	key := x.key(field)

	if val, ok := x.tags[key]; ok {
		tag = val
		return
	}

	tagString := field.Tag.Get("gokit-dbo")

	if tagString == "" {
		return
	}

	tag = &xPluginTag{}

	var isValue bool
	var kbuff, vbuff bytes.Buffer

	lastindex := len(tagString) - 1

	for i, char := range tagString {
		switch char {
		case ':':
			isValue = true

		case ';':
			x.updateTagValue(tag, &kbuff, &vbuff)
			isValue = false

		default:
			if isValue {
				vbuff.WriteRune(char)
			} else {
				kbuff.WriteRune(char)
			}
		}

		if i == lastindex {
			x.updateTagValue(tag, &kbuff, &vbuff)
			isValue = false
			continue
		}
	}

	x.tags[key] = tag

	return
}

func (x *xPlugin) updateTagValue(tag *xPluginTag, kbuff, vbuff *bytes.Buffer) {
	defer kbuff.Reset()
	defer vbuff.Reset()

	key := kbuff.String()

	switch key {
	case "preload":
		tag.Preload = true

	case "hidden":
		tag.Hidden = true
	}
}

// ==============================================

func NewPlugin() (v *xPlugin) {
	return &xPlugin{
		tags: make(map[string]*xPluginTag),
	}
}
