package data

import (
	"fmt"
	"time"
)

type People struct {
	Name                      string `json:"name"`                  //姓名
	IDnumber                  string `json:"idnumber"`              //身份证号
	Id                        string `json:"id" gorm:"primary_key"` //用户名
	Phone_number              string `json:"phone_number"`          //手机号
	City                      string `json:"city"`                  //城市
	Academy                   string `json:"academy"`               //社区选择
	High_region_sub           string `json:"high_region"`           //七天内受否有高低风险旅居史、接触史
	Isolation_sub             string `json:"isolation"`             //是否隔离观察
	Typical_symptoms_sub      string `json:"typical_symptoms"`      //是否身体有疑似典型症状
	Fever_sub                 string `json:"fever"`                 //是否发热
	Other                     string `json:"other"`                 //其他信息
	Temperature_morning_sub   string `json:"temperature_morning"`   //体温（早）
	Temperature_moon_sub      string `json:"temperature_moon"`      //体温（午）
	Temperature_aftermoon_sub string `json:"temperature_aftermoon"` //体温（晚）
	//High_region               bool    //七天内受否有高低风险旅居史、接触史
	//Isolation                 bool    //是否隔离观察
	//Typical_symptoms          bool    //是否身体有疑似典型症状
	//Fever                     bool    //是否发热

}

type UserInformation struct {
	Name                  string `json:"name"`                  //姓名
	IDnumber              string `json:"idnumber"`              //身份证号
	Id                    string `json:"id" gorm:"primary_key"` //用户民
	Phone_number          string `json:"phone_number"`          //手机号
	City                  string `json:"city"`                  //城市
	Academy               string `json:"academy"`               //社区选择
	High_region           string `json:"high_region"`           //七天内受否有高低风险旅居史、接触史
	Isolation             string `json:"isolation"`             //是否隔离观察
	Typical_symptoms      string `json:"typical_symptoms"`      //是否身体有疑似典型症状
	Fever                 string `json:"fever"`                 //是否发热
	Other                 string `json:"other"`                 //其他信息
	Temperature_morning   string `json:"temperature_morning"`   //体温（早）
	Temperature_moon      string `json:"temperature_moon"`      //体温（午）
	Temperature_aftermoon string `json:"temperature_aftermoon"` //体温（晚）
	T                     string `json:"date"`                  //日期
	Write                 string `json:"write"`                 //当日是否填报   0表示未填写，1表示已填写
}

func Initpeople(sign UserSign) {
	fmt.Println("init\n")
}

func FormatInformation(people People) UserInformation {
	var userinformation UserInformation
	userinformation.Name = people.Name
	userinformation.IDnumber = people.IDnumber
	userinformation.Id = people.Id
	userinformation.Phone_number = people.Phone_number
	userinformation.City = people.City
	userinformation.Academy = people.Academy
	userinformation.High_region = people.High_region_sub
	userinformation.Isolation = people.Isolation_sub
	userinformation.Typical_symptoms = people.Typical_symptoms_sub
	userinformation.Fever = people.Fever_sub
	userinformation.Other = people.Other
	//userinformation.Temperature_morning, _ = strconv.ParseFloat(people.Temperature_morning_sub, 64)
	//userinformation.Temperature_moon, _ = strconv.ParseFloat(people.Temperature_moon_sub, 64)
	//userinformation.Temperature_aftermoon, _ = strconv.ParseFloat(people.Temperature_aftermoon_sub, 64)
	userinformation.Temperature_morning = people.Temperature_morning_sub
	userinformation.Temperature_moon = people.Temperature_moon_sub
	userinformation.Temperature_aftermoon = people.Temperature_aftermoon_sub
	t := time.Now()
	userinformation.T = t.Format("2006.1.2")
	userinformation.Write = "已填报"
	return userinformation
	//people.T = time.Now()
	//t := people.T
	//fmt.Printf("%4d.%02d.%02d\n", t.Year(), t.Month(), t.Day())
}
