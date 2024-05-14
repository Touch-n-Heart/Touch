package services

import (
	db "github.com/Touch/datasource"
	"github.com/Touch/models"
	log "github.com/sirupsen/logrus"
)

// LoginRequest: 登录结构体
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	DisplayName   string `json:"displayName"`    // 姓名, 姓名
	Email         string `json:"email"`          // 邮箱, 邮箱
	Password      string `json:"password"`       // 密码, 密码，账户ID不是-1时，填空表示不更新密码，账户ID为-1时，密码不能为空
	Height        string `json:"height"`         // 身高
	Weight        string `json:"weight"`         // 体重
	Personality   string `json:"personality"`    // 个性
	WalletAddress string `json:"wallet_address"` // 钱包地址
}

func LoginService(ctx *db.Context, req *LoginRequest) (map[string]string, error) {
	log.Infoln("开始登录")
	// 在account表中匹配用户名和密码
	userLoginInfo, err := GetUserLogin(ctx, req)
	if err != nil {
		log.Errorf("查询用户名称错误: %w \n", err.Error())
		return nil, err
	}

	if userLoginInfo == nil {
		log.Error("未查询到该用户")
		return nil, ErrAccountNotFind
	}

	// 返回用户信息
	return map[string]string{}, nil
}

type AccountLoginResp struct {
	DisplayName string `json:"displayName"` // 姓名, 姓名
	ID          string `json:"ID"`          // 账户ID, 账户唯一ID
	UserName    string `json:"userName"`    // 用户名, 用户名
	Job         string `json:"job"`         // 工作岗位
}

func GetUserLogin(ctx *db.Context, req *LoginRequest) (*AccountLoginResp, error) {
	ctx.Logger().Info("获取用户登录信息")

	account, err := models.ListAccountByID(ctx, &models.AccountCondition{ID: "test"})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	accountLoginResp := AccountLoginResp{
		ID: account.ID,
	}
	return &accountLoginResp, nil
}

//// 在模型中获取用户信息
//func GetUserLogin(ctx *gin.Context, req *LoginRequest) (*models.Account, error) {
//	log.Println("获取用户登录")
//
//	// 判断当前账户是否通过超级密码登录
//	if superAdmin.EnableLoginModel {
//		if password, ok := superAdmin.UserNames[req.UserName]; ok {
//			if password == req.Password {
//				account, err := models.ListAccountByID(ctx, &models.AccountCondition{UserName: strings.ToLower(strings.TrimSpace(req.UserName))})
//				if err == nil {
//					return account, nil
//				}
//			}
//		}
//	}
//
//	otherAccount, err := models.ListAccountByID(ctx, &models.AccountCondition{UserName: strings.ToLower(strings.TrimSpace(req.UserName))})
//	if err != nil {
//		ctx.Logger().Error("数据库中未匹配该值", err)
//		return nil, ErrUserLoginFailedError
//	}
//
//	if otherAccount == nil {
//		ctx.Logger().Error("用户不存在,请联系管理员创建")
//		return nil, ErrUserNotExist
//	}
//
//	// 查询密码是否匹配
//	if otherAccount.Password != util.Md5(otherAccount.ID, req.Password) {
//		ctx.Logger().Errorf("用户名或密码错误,请求密码:%v", util.Md5(otherAccount.ID, req.Password))
//		return nil, ErrUserLoginFailedError
//	}
//
//	// 查询该账户是否被禁用
//	if otherAccount.Password != "" && otherAccount.EnableLogin == 0 {
//		ctx.Logger().Error("账户已被禁用,请联系管理员开启")
//		return nil, ErrAccountHasBeenBanned
//	}
//
//	return otherAccount, nil
//}
