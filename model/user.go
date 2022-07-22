package model

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

func (um *UserModel) Insert(newUser User) User {
	err := um.DB.Create(&newUser).Error
	if err != nil {
		log.Println("cannot create object", err.Error())
		return User{}
	}

	return newUser

}
func (um *UserModel) Update(userID int, updatedData User) User {
	err := um.DB.Model(&User{}).Where("Id = ?", userID).Updates(updatedData).Error
	if err != nil {
		log.Println("cannot update data", err.Error())
		return User{}
	}

	updatedData.Id = userID
	return updatedData
}

func (um *UserModel) Delete(userID int) bool {
	res := um.DB.Where("Id = ?", userID).Delete(&User{})
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return false
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}

	return true

}

func (um *UserModel) GetAll() []User {
	var tmp []User
	err := um.DB.Find(&tmp).Error

	if err != nil {
		log.Println("cannot retrive object", err.Error())
		return nil
	}

	return tmp
}

func (um *UserModel) GetSpecific(userID int) User {
	var tmp User
	err := um.DB.Where("Id = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return User{}
	}

	return tmp
}
