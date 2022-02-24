package service

import (
	"github.com/kiven-man/cart/domain/model"
	"github.com/kiven-man/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindCarByID(int64) (*model.Cart, error)
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

func (c *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(cart)
}
func (c *CartDataService) DeleteCartByID(cartID int64) error {
	return c.CartRepository.DeleteCartByID(cartID)
}
func (c *CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}
func (c *CartDataService) FindCarByID(cartID int64) (*model.Cart, error) {
	return c.CartRepository.FindCarByID(cartID)
}
func (c *CartDataService) FindAll(userID int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(userID)
}

func (c *CartDataService) CleanCart(userID int64) error {
	return c.CartRepository.CleanCart(userID)
}
func (c *CartDataService) IncrNum(cartID int64, num int64) error {
	return c.CartRepository.IncrNum(cartID, num)
}
func (c *CartDataService) DecrNum(cartID int64, num int64) error {
	return c.CartRepository.DecrNum(cartID, num)
}
