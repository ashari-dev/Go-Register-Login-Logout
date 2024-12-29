package controllers

import (
	"fmt"
	"trial/dtos"
	"trial/lib"
	"trial/models"
	"trial/repository"

	"github.com/gin-gonic/gin"
)

func RegisterController(ctx *gin.Context) {
	var formReg dtos.Register
	err := ctx.Bind(&formReg)
	if err != nil {
		lib.HandlerBadReq(ctx, "Confirm password wrong")
		return
	}

	user, err := repository.CreateUser(models.Users{
		Email:    formReg.Email,
		FullName: formReg.FullName,
		Password: formReg.Password,
	})
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(ctx, "Email already in use")
		return
	}

	lib.HandlerOK(ctx, "Register success", user)

}

func LoginController(ctx *gin.Context) {
	var formLogin dtos.Login
	err := ctx.Bind(&formLogin)
	if err != nil {
		lib.HandlerBadReq(ctx, "bad request")
		return
	}

	found, _ := repository.GetUserByEmail(formLogin.Email)
	if found == (models.Users{}) {
		lib.HandlerUnauthorized(ctx, "Email is wrong")
		return
	}

	isVerified := lib.Verify(formLogin.Password, found.Password)
	if !isVerified {
		lib.HandlerUnauthorized(ctx, "Password is wrong")
		return
	} else {
		JWToken := lib.GenerateUserTokenById(found.Id)
		lib.HandlerOK(ctx, "Login success", dtos.Token{Token: JWToken})
	}
}

func LogoutController(ctx *gin.Context) {
	lib.HandlerOK(ctx, "Logout success", dtos.Token{Token: ""})
}
