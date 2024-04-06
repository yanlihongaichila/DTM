package pkg

import "pkg/mysql"

func Migrate() error {
	top := NewGoods()

	err := mysql.Db.AutoMigrate(top)
	if err != nil {
		return err
	}

	return nil
}
