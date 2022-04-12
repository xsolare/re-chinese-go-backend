package response

type WordCollection struct {
	WordId         uint   `json:"word_id"`
	Name           string `json:"name"`
	Pinyin         string `json:"pinyin"`
	PartOfSpeechId uint   `json:"part_of_speech_id"`
	Translate      string `json:"translate"`
}
