package user

import (
	"gitea.ysicing.net/cloud/pangu/internal/models/user"
	"gitea.ysicing.net/cloud/pangu/pkg/util"
	"github.com/ergoapi/util/expass"
	"github.com/sirupsen/logrus"
)

func Init() (err error) {
	users := user.FindAll()

	if len(users) == 0 {
		initPassword := util.GetKeyFromYaml("admin.password", expass.PwGenAlphaNumSymbols(16))
		var pass []byte
		if pass, err = expass.GenerateHash([]byte(initPassword)); err != nil {
			return err
		}

		inituser := user.User{
			Username: util.GetKeyFromYaml("admin.username", "admin"),
			Password: string(pass),
			Nickname: "超级管理员",
			Token:    expass.PwGenAlphaNum(32),
		}
		if err := user.Create(&inituser); err != nil {
			return err
		}
		logrus.Infof("初始用户创建成功，账号:「%v」密码:「%v」", inituser.Username, initPassword)
	} else {
		logrus.Warn("用户已经初始化过")
	}
	return nil
}
