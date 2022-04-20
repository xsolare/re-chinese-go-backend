package service

type ServiceGroup struct {
	UserService
	FinalService
	HieroglyphCollectionService
	WordCollectionService
	WordService
	PinyinService
	HieroglyphService
	OperationRecordService
}
