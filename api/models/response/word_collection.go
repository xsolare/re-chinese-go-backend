package response

type WordCollection struct {
	Id             uint   `json:"id"`
	Hieroglyphics  string `json:"hieroglyphics"`
	Pinyin         string `json:"pinyin"`
	PartOfSpeechId uint   `json:"partOfSpeechId"`
	Translate      string `json:"translate"`
	Hsk            uint   `json:"hsk"`
}
