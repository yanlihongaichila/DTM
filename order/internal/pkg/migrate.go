package pkg

import "pkg/mysql"

func Migrate() error {
	order := NewOrder()
	err := mysql.Db.AutoMigrate(order)
	if err != nil {
		return err
	}
	orderGoods := NewOrderGoods()
	err = mysql.Db.AutoMigrate(orderGoods)
	if err != nil {
		return err
	}
	return nil
}
