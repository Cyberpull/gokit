package dbo

import (
	"math"
	"reflect"

	"github.com/Cyberpull/gokit/errors"

	"gorm.io/gorm"
)

type Pagination[T any] struct {
	CurrentPage uint `json:"current_page"`
	From        uint `json:"from"`
	LastPage    uint `json:"last_page"`
	PerPage     int  `json:"per_page"`
	To          uint `json:"to"`
	Total       uint `json:"total"`
	Data        []T  `json:"data"`
}

func Paginate[T any](tx *gorm.DB, page uint, limit ...uint) (value *Pagination[T], err error) {
	if len(limit) == 0 {
		limit = append(limit, 20)
	}

	if tx, err = parseModel[T](tx); err != nil {
		return
	}

	tx = tx.Offset(-1).Limit(-1).Session(&gorm.Session{})

	tmpValue := &Pagination[T]{}
	tmpValue.Data = make([]T, 0)
	tmpValue.CurrentPage = uint(math.Max(float64(page), 1))
	tmpValue.PerPage = int(math.Max(float64(limit[0]), 1))

	offset := int(0)

	if tmpValue.CurrentPage > 1 {
		offset = int(tmpValue.CurrentPage) - 1
		offset *= tmpValue.PerPage
	}

	if offset > 0 {
		tmpValue.From = uint(offset) + 1
	} else {
		tmpValue.From = 1
	}

	tx = tx.Offset(offset).Limit(tmpValue.PerPage)

	if err = tx.Find(&tmpValue.Data).Error; err != nil {
		return
	}

	tmpValue.To = uint(offset) + uint(len(tmpValue.Data))

	var total int64

	tx = tx.Offset(-1).Limit(-1)

	if err = tx.Count(&total).Error; err != nil {
		return
	}

	lastPage := float64(total) / float64(tmpValue.PerPage)
	tmpValue.LastPage = uint(math.Ceil(lastPage))
	tmpValue.Total = uint(total)

	value = tmpValue

	return
}

func parseModel[T any](tx *gorm.DB) (vtx *gorm.DB, err error) {
	var model T

	vType := reflect.TypeOf(model)

	if vType.Kind() == reflect.Pointer {
		vType = vType.Elem()
		model = reflect.New(vType).Interface().(T)
		vtx = tx.Model(model)
	} else {
		vtx = tx.Model(&model)
	}

	if vType.Kind() != reflect.Struct {
		err = errors.New("Model should be a struct")
	}

	return
}
