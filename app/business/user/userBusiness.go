package user

import (
	"fmt"
	"ginapi/app/global"
	"ginapi/app/model/test1"

	"gorm.io/gorm"
)

func CheckAccountIsExist(account string) bool {

	// 查询Resource
	db := global.App.Dbs["test1"]

	// result := db.First(&ppclibrary.Resource{}, "resource_id = ?", 1)
	var user test1.TestUser
	user.Account = account

	result := db.Where(&user).Take(&user)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		fmt.Println(result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}
	}
	// fmt.Println(result)
	fmt.Println(user)
	// fmt.Println(json.Marshal(user))
	return true
}

// 创建用户
func CreateUser(param *test1.TestUser) {
	// 查询Resource
	db := global.App.Dbs["test1"]

	db.Create(param)
}

func CheckPassword(account string, password string) bool {
	db := global.App.Dbs["test1"]
	condition := map[string]interface{}{
		"account": account,
	}
	testUser := test1.TestUser{}
	result := db.Where(condition).Take(&testUser)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}
	}
	if testUser.Password == password {
		return true
	}
	return false
}
