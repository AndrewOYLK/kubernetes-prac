package ServicesImpl

import (
	"context"
	"fmt"
	"proj15-valid/AppInit"
	"proj15-valid/DBModels"
	"proj15-valid/Services"
)

type UserService struct {
}

func (this *UserService) UserReg(ctx context.Context, user *Services.UserModel, rsp *Services.RegResponse) error {
	users := DBModels.Users{
		UserName: user.UserName,
		UserPwd:  user.UserPwd,
	}

	fmt.Println("====")
	fmt.Println(users)

	db := AppInit.GetDB()
	fmt.Println(db)
	if err := db.Create(&users).Error; err != nil {
		rsp.Message = err.Error()
		rsp.Status = "error"
	} else {
		rsp.Message = ""
		rsp.Status = "success"
	}
	return nil
}
