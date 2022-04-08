package response

type Pinyin struct {
	PinyinId  uint   `json:"pinyin_id"`
	Pinyin    string `json:"pinyin"`
	InitialId uint   `json:"initial_id"`
	FinalId   uint   `json:"final_id"`
	Tone      string `json:"tone"`
}
