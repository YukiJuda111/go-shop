package model

type RoomGoods struct {
	BaseModel
	RoomId    int64
	GoodsId   int64
	Weight    int64
	IsCurrent int8
}

// TableName 声明表名
func (RoomGoods) TableName() string {
	return "yuk1_room_goods"
}
