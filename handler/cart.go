package handler

import (
	"context"

	"github.com/kiven-man/cart/common"
	"github.com/kiven-man/cart/domain/model"
	"github.com/kiven-man/cart/domain/service"
	cart2 "github.com/kiven-man/cart/proto"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, request *cart2.CartInfo, response *cart2.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = c.CartDataService.AddCart(cart)
	return err
}
func (c *Cart) CleanCart(ctx context.Context, request *cart2.Clean, response *cart2.Response) (err error) {
	if err := c.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Msg = "购物车清空成功"
	return nil
}
func (c *Cart) Incr(ctx context.Context, request *cart2.Item, response *cart2.Response) (err error) {
	if err := c.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物车添加成功"
	return nil

}
func (c *Cart) Decr(ctx context.Context, request *cart2.Item, response *cart2.Response) (err error) {
	if err := c.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物车减少成功"
	return nil
}
func (c *Cart) DeleteItemByID(ctx context.Context, request *cart2.CartID, response *cart2.Response) (err error) {
	if err := c.CartDataService.DeleteCartByID(request.Id); err != nil {
		return err
	}
	response.Msg = "购物车删除成功"
	return nil
}
func (c *Cart) GetAll(ctx context.Context, request *cart2.CartFindAll, response *cart2.CartAll) (err error) {
	cartAll, err := c.CartDataService.FindAll(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart2.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
