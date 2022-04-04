package controllers

import "github.com/xsolare/re-chinese-go-backend/api/service"

type ControllerGroup struct {
	FinalController
	UserController
	HieroglyphCollectionController
}

var finalService = new(service.ServiceGroup).FinalService
