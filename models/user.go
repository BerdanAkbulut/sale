package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"->;primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email" gorm:"index:,unique"`
}

func GetUsers(users *[]User) (err error) {
	if err = DB.Find(users).Error; err != nil {
		return err
	}

	return nil
}

func CreateUser(user *User) (err error) {
	if err = DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *User) (err error) {
	if err = DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUser(user *User, id string) (err error) {
	if err = DB.First(user, id).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(user *User, id string) (err error) {
	if err = DB.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}
