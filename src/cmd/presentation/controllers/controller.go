package controllers

import "github.com/vashp/ms-user/src/cmd/domain/usecases"

type Controller struct {
	userService usecases.UserService
}

func NewController(userService usecases.UserService) *Controller {
	return &Controller{userService: userService}
}
