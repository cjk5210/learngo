package main

import (
	"fmt"

	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"database/sql"
)
type Product struct{
	gorm.Model
	Code string
	Price uint
}
type Test struct{
	id int
	name string
}
type User struct{
	gorm.Model
	Birthday time.Time
	Age int
	Name string `gorm:"size:255"`
	Num int `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard
	Emails []Email

	BillingAddress Address
	BillingAddressID sql.NullInt64

	ShippingAddress Address
	ShippingAddressID int

	IgnoreMe int `gorm:"-"`
	Languages []Language `gorm:"many2many:user_languages;"`

}
type Email struct{
	ID int
	UserID int `gorm:"index"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Subcribed bool
}
type Address struct{
	ID int
	Address1 string
	Address2 string
	Post sql.NullString
}
type Language struct{
	ID int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}
type CreditCard struct{
	gorm.Model
	UserID uint
	Number string
}
func main() {
	fmt.Println(time.Now())
	db,err:=gorm.Open("mysql","mysqluser:77585210@(59.110.242.165:3306)/gorm?charset=utf8")
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
	}

	//db.AutoMigrate(&Product{})
	//db.AutoMigrate(&User{},&Email{},&Address{},&Language{},&CreditCard{})
	fmt.Println("USER table has been builded",db.HasTable(&User{}))
	fmt.Println("Email table has been builded",db.HasTable(&Email{}))
	fmt.Println("Address table has been builded",db.HasTable(&Address{}))
	fmt.Println("Language table has been builded",db.HasTable(&Language{}))
	fmt.Println("CreditCard table has been builded",db.HasTable(&CreditCard{}))
	user:=User{Name:"Jinzhu",Age:18,Birthday:time.Now()}
	count:=0;
	db.Table("users").Count(&count)
	fmt.Println("init,now User has record :",count)
	db.NewRecord(user)
	db.Table("users").Count(&count)
	fmt.Println("newrecord,now User has record :",count)
	db.Create(&user)
	db.Table("users").Count(&count)
	fmt.Println("create,now User has record :",count)
}
