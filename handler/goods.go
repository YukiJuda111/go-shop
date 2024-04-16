package handler

import (
	"context"
	"go-shop/proto"
)

// 满足proto.GoodsSrv接口
type GoodsSrv struct {
	proto.UnimplementedGoodsServer
}

func (s *GoodsSrv) GetGoodsByRoom(context.Context, *proto.GetGoodsByRoomReq) (*proto.GoodsListResp, error) {
	return &proto.GoodsListResp{}, nil
}

func (s *GoodsSrv) GetGoodsDetail(context.Context, *proto.GetGoodsDetailReq) (*proto.GoodsDetail, error) {
	return &proto.GoodsDetail{}, nil
}
