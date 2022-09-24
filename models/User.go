package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []Photo `gorm:"OnDelete:CASCADE;"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Where("id = ?", uid).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(&u).Where("id = ?", uid).Updates(User{Name: u.Name, Email: u.Email}).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) UpdatePassword(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(&u).Where("id = ?", uid).Update("password", u.Password).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint) (int64, error) {
	db = db.Where("id = ?", uid).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
