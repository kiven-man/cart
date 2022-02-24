package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kiven-man/cart/domain/model"
)

type ICartRepository interface {
	InitTable() error
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindCarByID(int64) (*model.Cart, error)
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDB: db}
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (c *CartRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Cart{}).Error
}

// 创建购物车
func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDB.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

// 根据ID删除购物车商品
func (c *CartRepository) DeleteCartByID(cartID int64) error {
	return c.mysqlDB.Where("id=?", cartID).Delete(&model.Cart{}).Error
}

// 更新购物车信息
func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDB.Model(cart).Update(cart).Error
}

// 根据ID查找购物车
func (c *CartRepository) FindCarByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDB.First(cart, cartID).Error
}

// 获取结果集
func (c *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDB.Where("user_id=?", userID).Find(&cartAll).Error
}

// 清空购物车
func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysqlDB.Where("user_id=?", userID).Delete(&model.Cart{}).Error
}

// 增加购物车商品数量
func (c *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return c.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num+?", num)).Error
}

// 减少购物车商品数量
func (c *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := c.mysqlDB.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
