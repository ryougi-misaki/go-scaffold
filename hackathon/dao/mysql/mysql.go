package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hackathon/config"
	"hackathon/models"
)

var DB *gorm.DB

func Init() error {
	var err error
	DB, err = gorm.Open(mysql.Open(config.ReturnMySQLsetting(config.Conf.MySQLConfig)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return err
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func Create(new interface{}) error {
	err := DB.Create(new).Error
	if err != nil {
		return err
	}
	return nil
}

//根据id查data
//date必须传入引用，用于接收数据
func RetrieveByID(data interface{}, id uint) {
	DB.First(data, id)
	return
}

//tar为约束条件结构体，根据tar查第一个符合条件的data
//date必须传入引用，用于接收数据
func RetrieveByStruct(data interface{}, tar interface{}) {
	DB.Where(tar).First(data)
	return
}

//tar为约束条件结构体，根据tar查符合条件的所有data
//date必须传入引用，用于接收数据
func RetrieveArrByStruct(data interface{}, tar interface{}) {
	DB.Where(tar).Find(data)
	return
}

//tar为约束条件结构体
func Update(tar interface{}, new interface{}) error {
	err := DB.Model(tar).Updates(new).Error
	if err != nil {
		return err
	}
	return nil
}

//tar为约束条件结构体，删除符合条件的数据
func Delete(tar interface{}) error {
	err := DB.Delete(tar).Error
	if err != nil {
		return err
	}
	return nil
}
