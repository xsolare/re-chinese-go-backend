package response

type Word struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Pinyin string `json:"pinyin"`
	Hsk    uint   `json:"hsk"`
}
