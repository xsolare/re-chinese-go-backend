package router

type RouterGroup struct {
	UserRouter
	FinalRouter
	HieroglyphCollectionRouter
	WordCollectionRouter
	WordRouter
	PinyinRouter
}
