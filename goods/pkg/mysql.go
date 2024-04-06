package pkg

import (
	"fmt"
	"gorm.io/gorm"
	"pkg/mysql"
)

type Goods struct {
	gorm.Model
	Name         string `gorm:"INDEX,type:varchar(20)"`
	Price        string `gorm:"type:decimal(10,2)"`
	Stock        int64  `gorm:"type:int"`
	TradingStock int64  `gorm:"type:int"` //冻结库存字段
}

func NewGoods() *Goods {
	return new(Goods)
}

func GetGoodsByIDs(ids []int64) ([]Goods, error) {

	goods := NewGoods()
	goodsInfos := []Goods{}
	err := mysql.Db.Model(goods).Where("deleted_at IS NULL").Where("id IN ?", ids).Find(&goodsInfos).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return goodsInfos, nil
}

func UpdateGoodsStocks(updateGoods map[int64]int64) error {
	fmt.Println(updateGoods)
	goods := NewGoods()
	for key, val := range updateGoods {
		err := mysql.Db.Model(goods).Where("id = ?", key).Update("stock", gorm.Expr("stock  + ?", val)).Error
		if err != nil {
			return fmt.Errorf("%v商品库存修改失败", key)
		}
	}

	return nil
}

func UpdateGoodsStocksRollback(updateGoods map[int64]int64) error {
	fmt.Println(updateGoods)
	goods := NewGoods()
	for key, val := range updateGoods {
		err := mysql.Db.Model(goods).Where("id = ?", key).Update("stock", gorm.Expr("stock  - ?", val)).Error
		if err != nil {
			return fmt.Errorf("%v商品库存修改失败", key)
		}
	}

	return nil
}

// 冻结
func TCCTradingNum(TCCInfos map[int64]int64) error {
	fmt.Println(TCCInfos)
	//开启mysql事务
	begin := mysql.Db.Begin()
	goods := NewGoods()
	for key, val := range TCCInfos {
		err := begin.Model(goods).Where("id = ?", key).Where("trading_stock + ? + stock", val).Update("stock", gorm.Expr("stock  + ?", val)).Error
		if err != nil {
			begin.Rollback()
			return fmt.Errorf("%v商品库存修改失败", key)
		}
	}

	//提交mysql事务
	begin.Commit()

	return nil
}

// 冻结确定
func TCCTradingLockNum(TCCInfos map[int64]int64) error {
	fmt.Println(TCCInfos)
	//开启mysql事务
	begin := mysql.Db.Begin()
	goods := NewGoods()
	for key, val := range TCCInfos {

		err := begin.Model(goods).Where("id = ?", key).Update("trading_stock", gorm.Expr("trading_stock  + ?", val)).Update("stock", gorm.Expr("stock  + ?", val)).Error
		if err != nil {
			begin.Rollback()
			return fmt.Errorf("%v商品库存确认修改失败", key)
		}
	}

	//提交mysql事务
	begin.Commit()

	return nil
}

// 回滚
func TCCBalanceNum(TCCInfos map[int64]int64) error {
	fmt.Println(TCCInfos)
	//开启mysql事务
	begin := mysql.Db.Begin()
	goods := NewGoods()
	for key, val := range TCCInfos {
		err := begin.Model(goods).Where("id = ?", key).Update("trading_stock", gorm.Expr("trading_stock  - ?", val)).Update("stock", gorm.Expr("stock  - ?", val)).Error
		if err != nil {
			begin.Rollback()
			return fmt.Errorf("%v商品库存归还失败", key)
		}
	}

	//提交mysql事务
	begin.Commit()

	return nil
}
