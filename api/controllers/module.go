package controllers

import "github.com/xsolare/re-chinese-go-backend/api/service"

type ControllerGroup struct {
	FinalController
	UserController
	HieroglyphCollectionController
	WordCollectionController
	WordController
	PinyinController
}

/// 																	///

var pinyinService = new(service.ServiceGroup).PinyinService
var wordService = new(service.ServiceGroup).WordService
var finalService = new(service.ServiceGroup).FinalService
var userService = new(service.ServiceGroup).UserService
var hieroglyphCollectionService = new(service.ServiceGroup).HieroglyphCollectionService
var wordCollectionService = new(service.ServiceGroup).WordCollectionService
