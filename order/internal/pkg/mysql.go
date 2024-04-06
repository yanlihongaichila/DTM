package pkg

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"order/order"
	"pkg/mysql"
	"time"
)

type Order struct {
	gorm.Model
	UserID  int64  `gorm:"INDEX,type:varchar(20)"`
	OrderNO string `gorm:"INDEX,type:text"`
	Amount  string `gorm:"type:decimal(10,2)"`
	State   int64  `gorm:"type:int"`
}

func NewOrder() *Order {
	return new(Order)
}

type OrderGoods struct {
	gorm.Model
	OrderID   int64  `gorm:"INDEX,type:int"`
	GoodsID   int64  `gorm:"INDEX,type:int"`
	UnitPrice string `gorm:"type:decimal(10,2)"`
	GoodName  string `gorm:"INDEX,type:varchar(50)"`
	Num       int64  `gorm:"type:int"`
}

func NewOrderGoods() *OrderGoods {
	return new(OrderGoods)
}

func CreateOrder(ids *order.OrderInfo) (Order, error) {

	orderInfo := NewOrder()
	info := Order{
		UserID:  ids.UserID,
		OrderNO: fmt.Sprintf("%v%v%v%v%v", time.Now().UnixNano(), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)),
		Amount:  ids.Amount,
		State:   ids.State,
	}
	err := mysql.Db.Model(orderInfo).Create(&info).Error

	if err != nil {
		fmt.Println(err)
		return Order{}, err
	}
	return info, nil
}

func CreateOrderRollback(rollbackId int64) (int64, error) {

	orderInfo := NewOrder()

	err := mysql.Db.Model(orderInfo).Delete("id = ?", rollbackId).Error

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return rollbackId, nil
}

func GetOrder(id int64) (*Order, error) {
	orderInfo := NewOrder()
	err := mysql.Db.Model(orderInfo).Where("id = ?", id).First(&orderInfo).Error

	if err != nil {
		return nil, err
	}
	return orderInfo, nil
}

func UpdateOrder(updateInfo *order.OrderInfo) error {
	orderInfo := NewOrder()

	err := mysql.Db.Model(orderInfo).Updates(&updateInfo).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateOrderGoods(createInfos []*order.OrderGoodsInfo) error {
	orderInfoReq := NewOrderGoods()

	for _, val := range createInfos {
		info := OrderGoods{
			OrderID:   val.OrderID,
			GoodsID:   val.GoodsID,
			UnitPrice: val.UnitPrice,
			GoodName:  val.GoodName,
			Num:       val.Num,
		}
		err := mysql.Db.Model(orderInfoReq).Create(&info).Error

		if err != nil {
			return fmt.Errorf("%v号商品添加失败", info.GoodsID)
		}
	}
	return nil
}

func GetOrderGoodsByOrderID(id int64) ([]OrderGoods, error) {
	orderInfoReq := NewOrderGoods()
	infoRes := []OrderGoods{}
	err := mysql.Db.Model(orderInfoReq).Where("order_id = ?", id).Find(&infoRes).Error

	if err != nil {
		return nil, err
	}

	return infoRes, nil
}

//TCC

func TCCTradingOrder(ids *order.OrderInfo) (Order, error) {
	orderInfo := NewOrder()

	info := Order{
		UserID:  ids.UserID,
		OrderNO: ids.OrderNO,
		Amount:  ids.Amount,
		State:   ids.State,
	}
	err := mysql.Db.Model(orderInfo).Create(&info).Error
	if err != nil {
		fmt.Println(err)
		return Order{}, err
	}

	return info, nil
}

func TCCTradingLockOrder(ids *order.OrderInfo) (Order, error) {
	//orderInfo := NewOrder()
	//info := Order{
	//	UserID:  ids.UserID,
	//	OrderNO: fmt.Sprintf("%v%v%v%v%v", time.Now().UnixNano(), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)),
	//	Amount:  ids.Amount,
	//	State:   ids.State,
	//}
	//err := mysql.Db.Model(orderInfo).Create(&info).Error
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return Order{}, err
	//}
	return Order{
		Model:   gorm.Model{ID: uint(ids.ID)},
		UserID:  ids.UserID,
		OrderNO: ids.OrderNO,
		Amount:  ids.Amount,
		State:   ids.State,
	}, nil
}

func TCCTradingRollbackOrder(orderNo string) (string, error) {
	orderInfo := NewOrder()

	err := mysql.Db.Model(orderInfo).Delete("order_no = ?", orderNo).Error

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return orderNo, nil
}
