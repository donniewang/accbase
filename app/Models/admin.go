package Models

import (
	. "accbase/database"
	"github.com/jinzhu/gorm"
)

type Admins struct {
	gorm.Model
	UserName string `json:"user_name"  binding:"required"`
	NickName string `json:"nick_name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Email string `json:"email" binding:"required"`
	Qq string `json:"qq" binding:"required"`
	Department int `json:"department" binding:"required"`
}

func init(){
	DB.AutoMigrate(&Admins{})
}


// Insert 新增admin用户
func (admin *Admins) Insert() (userID uint, err error) {

	result := DB.Create(&admin)
	userID = admin.ID // 返回值
	if result.Error != nil {
		err = result.Error
	}
	return
}

// Destroy 删除admin用户
func (admin *Admins) Destroy(id int) (err error) {
	result := DB.Where(`id = ?`, id).Delete(&admin) // 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
	if result.Error != nil {
		err = result.Error
	}
	return
}


// Update 修改admin用户
func (admin *Admins) Update(id int64) (user Admins, err error) {
	result := DB.Model(&admin).Where("id = ?", id).Updates(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询指定id admin用户
func (admin *Admins) FindOne(id int) (user Admins, err error) {

	result := DB.First(&admin, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAll 查询所有admin用户
func (admin *Admins) FindAll() (admins []Admins, err error) {

	result := DB.Find(&admins) // 这里的 &admins 跟返回参数要一致

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}