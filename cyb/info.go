package cyb

import "github.com/google/uuid"

type Info struct {
	Name        string
	Description string
	UUID        string
}

func (x *Info) GenerateUUID() {
	x.UUID = uuid.NewString()
}
