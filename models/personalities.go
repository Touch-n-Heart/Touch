package models

import db "github.com/Touch/datasource"

// personality用户推荐表
// 展示用户相关的个性，以及和个性相关的用户

type Personality struct {
	ID          string `json:"ID"`          // ID
	Username    string `json:"username"`    // 名称
	Personality string `json:"personality"` // 个性
}

func (a *Personality) TableName() string {
	return "personality"
}

func ListPersonalityByID(ctx *db.Context, condition *Personality) (*Personality, error) {
	ctx.Logger().Info("通过ID获取账户")

	personality := &Personality{ID: "YES"}
	//err := ctx.DB().Model(Account{}).First(account, condition).Error
	//if err != nil {
	//	return nil, err
	//}

	return personality, nil
}
