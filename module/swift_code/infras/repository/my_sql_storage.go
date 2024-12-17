package repository

import (
	"data_crawler/module/swift_code/model"
	"data_crawler/shared/gormc"
)

type MySQLStorage struct{}

func (s *MySQLStorage) CreateInfoSwiftCode(payload []model.Swiftcode) error {
	// Insert Swift code here
	db := gormc.Init()

	if err := db.Table("info").Create(&payload).Error; err != nil {
		return err
	}
	return nil
}
