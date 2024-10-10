package database

import (
	"fmt"
	"software/data"
	"time"
)

func CreateUser(user data.UserRegister) int { //注册用
	db := ConnectMySQL()
	defer db.Close()
	db.AutoMigrate(&data.UserRegister{})
	if err := db.Create(user).Error; err == nil {
		return 1
	} else {
		return 0
	}
}

func SearchUser(user data.UserSign) int { //登陆用
	var userregister data.UserRegister
	db := ConnectMySQL()
	defer db.Close()
	if err := db.Table(data.TableUserregister).Where("username=?", user.Username).Take(&userregister).Error; err == nil {
		if user.Password == userregister.Password {
			if userregister.Identity == "用户" {
				return 1
			} else {
				return 0
			}
		} else {
			return 2
		}
	} else {
		return 3
	}
}

func Init(u data.UserRegister) {
	db := ConnectMySQL()
	defer db.Close()
	b := data.UserInformation{"", "", u.Username, "", "", "", "", "", "", "", "", "", "", "", "", ""}
	db.Create(&b)
}

func GetUserInformation(people data.People) int { //填报信息用
	var userinformation data.UserInformation
	db := ConnectMySQL()
	defer db.Close()
	t := time.Now()
	var time string
	time = t.Format("2006.1.2")
	fmt.Println(time)
	if err := db.Table(data.TableUserInformation).Where("id=?", people.Id).Take(&userinformation).Error; err == nil {
		fmt.Printf("old:%+v\n", userinformation)
		if time == userinformation.T {
			return 1
		} else {
			userinformationnew := data.FormatInformation(people)
			fmt.Printf("new:%+v\n", userinformationnew)
			db.Table(data.TableUserInformation).Save(&userinformationnew)
			return 0
		}

	} else {
		return 2
	}
}

func GetInformation(id string) data.UserInformation { //管理员查询信息
	var information data.UserInformation
	db := ConnectMySQL()
	defer db.Close()
	if err := db.Table(data.TableUserInformation).Where("id=?", id).First(&information).Error; err == nil {
		fmt.Printf("old:%+v\n", information)
	}
	return information
}

func IncreaseInformation(in data.UserInformation) int { //管理员添加修改信息
	db := ConnectMySQL()
	defer db.Close()
	if err := db.Table(data.TableUserInformation).Save(&in).Error; err == nil {
		fmt.Printf("old:%+v\n", in)
		return 1
	}
	return 0
}

func DeleteInformation(userinformation data.UserInformation) int {
	db := ConnectMySQL()
	defer db.Close()

	userinformation.High_region = ""
	userinformation.Isolation = ""
	userinformation.Typical_symptoms = ""
	userinformation.Fever = ""
	userinformation.Other = ""
	userinformation.Temperature_morning = ""
	userinformation.Temperature_moon = ""
	userinformation.Temperature_aftermoon = ""
	userinformation.T = ""
	userinformation.Write = ""
	if err := db.Table(data.TableUserInformation).Save(&userinformation).Error; err == nil {
		fmt.Printf("old:%+v\n", userinformation)
		return 1
	}
	return 0
}
