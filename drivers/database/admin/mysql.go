package admin

import (
	"ca-reservaksin/businesses/admin"

	"gorm.io/gorm"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) admin.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlAdminRepository) Register(dataAdmin *admin.Domain) (admin.Domain, error) {
	recAdmin := fromDomain(*dataAdmin)

	err := mysqlRepo.Conn.Create(&recAdmin).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return recAdmin.toDomain(), nil
}

func (mysqlRepo *MysqlAdminRepository) GetByUsername(username string) (admin.Domain, error) {
	rec := Admin{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (mysqlRepo *MysqlAdminRepository) GetByID(id int) (admin.Domain, error) {
	rec := Admin{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return rec.toDomain(), nil
}
