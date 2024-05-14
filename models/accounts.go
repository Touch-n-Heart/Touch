package models

import (
	"fmt"
	db "github.com/Touch/datasource"
	"time"
)

type Account struct {
	ID            string    `json:"id" gorm:"column:id;primaryKey;"`       // 账户id
	DisplayName   string    `gorm:"display_name"`                          // 展示用户名
	UserName      string    `gorm:"user_name"`                             // 用户名
	Email         string    `gorm:"email"`                                 // 邮箱
	Password      string    `gorm:"password"`                              // 密码
	Gender        string    `gorm:"gender"`                                // 性别
	BirthDay      string    `gorm:"birthday"`                              //生日
	Height        string    `gorm:"height"`                                // 高度
	Weight        string    `gorm:"weight"`                                // 体重
	Personality   string    `gorm:"personality"`                           // 个性
	Hoppy         string    `gorm:"hoppy"`                                 // 爱好
	WalletAddress string    `gorm:"wallet_address"`                        // 钱包地址
	HomeAddress   string    `gorm:"home_address"`                          // 家庭住址
	City          string    `gorm:"city"`                                  // 城市
	Religion      string    `gorm:"religon"`                               // 宗教
	Telegram      string    `gorm:"telegram"`                              // 电报
	PersonalityID string    `gorm:"personality_id"`                        // 个性
	CreatedAt     time.Time `json:"createTime" gorm:"column:create_time" ` // 创建时间
	CreatedBy     string    `json:"create_by" gorm:"column:create_by"`     // 创建人
	UpdatedTime   time.Time `json:"updateTime" gorm:"column:update_time"`  // 更新时间
	UpdatedBy     string    `json:"update_by" gorm:"column:update_by"`     // 更新人
}

type AccountCondition struct {
	ID          string
	UserName    string
	DisplayName string
	Password    string
}

func (a *Account) TableName() string {
	return "account"
}

func ListAccountByID(ctx *db.Context, condition *AccountCondition) (*Account, error) {
	ctx.Logger().Info("通过ID获取账户")

	account := &Account{ID: "YES"}
	//err := ctx.DB().Model(Account{}).First(account, condition).Error
	//if err != nil {
	//	return nil, err
	//}

	return account, nil
}

func CreateAccount(ctx *db.Context, account Account) error {
	ctx.Logger().Info("创建账户")

	err := ctx.DB().Create(&account).Error
	if err != nil {
		return fmt.Errorf("创建账户失败: %w", err)
	}
	return nil
}

func UpdateAccount(ctx *db.Context, account Account) error {
	ctx.Logger().Info("更新账户")

	err := ctx.DB().Save(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(ctx *db.Context, id string) error {
	ctx.Logger().Info("删除账户")
	err := ctx.DB().Where("id = ?", id).Delete(&Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
