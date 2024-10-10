package main

import (
	"github.com/gin-gonic/gin"
	"software/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()

	//
	//db := database.ConnectMySQL()
	//defer db.Close()
	//
	//db.AutoMigrate(&data.UserRegister{})
	//
	//a := data.UserRegister{"管理员", "123456", "111111", "1"}
	//db.Create(&a)
	//var u []data.UserRegister
	//db.Find(&u)
	//fmt.Printf("u:%+v\n", u)
	//b := data.UserRegister{"管理员", "123456", "123456", "111111"}
	////更新
	//db.Table(data.TableUserregister).Save(&b)
	//db.Find(&u)
	//fmt.Printf("u:%+v\n", u)

	//查询
	//var userregister []data.UserRegister
	//err := db.Table(data.TableUserregister).Where("username=?", "123456").First(&userregister).Error
	//fmt.Println(err)
	//fmt.Printf("userregister:%+v\n", userregister)
	//
	//var userinformation data.UserInformation
	//db.AutoMigrate(&data.UserInformation{})
	//b := data.UserInformation{"刘稼宇", "410305200109250531", "193401050119", "15837991078", "沈阳", "一号社区", "否", "否", "否", "否", "", "36.1", "36.1", "36.1", "2023.2.9", "1"}
	//db.Create(&b)
	//var uu []data.UserInformation
	//db.Find(&uu)
	//fmt.Printf("u:%+v\n", uu)
	//db.Table(data.TableUserInformation).Where("name=?", "孙健").First(&userinformation)
	//fmt.Printf("u:%+v\n", userinformation)

	//var u []data.UserInformation
	//var a []string
	//db := database.ConnectMySQL()
	//defer db.Close()
	//if err := db.Table(data.TableUserInformation).Where("write=?", "已填报").Find(&u).Error; err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("old:%+v\n", u)
	//for i := 0; i < len(u); i++ {
	//
	//	a = append(a, u[i].Name)
	//}
	//fmt.Printf("old:%+v\n", a)

}
