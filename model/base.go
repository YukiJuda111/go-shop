package model

import (
	"time"
)

type BaseModel struct {
	ID       uint
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string
	Version  int16
	IsDel    int8
}
