package storage

import (
	"fmt"
)

type Field string

func (a Field) Eq() string {
	return fmt.Sprintf("%s = ?", a)
}

func NewField(name string) Field {
	return Field(name)
}

func (a Field) OrderAsc() Ordering {
	return Ordering(fmt.Sprintf("%s asc", a))
}

func (a Field) OrderDesc() Ordering {
	return Ordering(fmt.Sprintf("%s desc", a))
}

func (a Field) Like() string {
	return fmt.Sprintf("%s ilike ?", a)

}
