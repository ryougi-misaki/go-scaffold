package service

import (
	"golang.org/x/crypto/bcrypt"
	"hackathon/dao/mysql"
	"hackathon/models"
	"hackathon/response"
	"hackathon/util"
)

func Register(p *models.ParamRegister) int {
	//数据验证
	if len(p.Telephone) != 11 {
		return response.CodePhoneLength
	}
	if len(p.Password) < 6 {
		return response.CodePwdLength
	}

	//如果名称没有传，就给名称一个随机的十位字符串
	if len(p.Name) == 0 {
		p.Name = util.RandomString(10)
	}
	//判断手机号是否存在
	DB := mysql.GetDB()
	if mysql.IsTelephoneExist(DB, p.Telephone) {
		//response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return response.CodePhoneExist
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		//response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"加密错误")
		return response.CodeEncryptError
	}
	newUser := &models.User{
		Name:      p.Name,
		Telephone: p.Telephone,
		Password:  string(hasedPassword),
	}
	err = mysql.Create(newUser)
	if err != nil {
		return response.Error
	}
	//返回结果
	return response.OK
}

func Login(p *models.ParamLogin) (string, int) {
	//手机号是否存在
	DB := mysql.GetDB()
	var user models.User
	DB.Where("telephone = ?", p.Telephone).First(&user)
	if user.ID == 0 {

		return "", response.CodePhoneExist
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password)); err != nil {
		return "", response.CodePwdWrong
	}
	//发放token
	token, err := util.ReleaseToken(user)
	if err != nil {
		return "", response.Error
	}
	return token, response.OK
}
