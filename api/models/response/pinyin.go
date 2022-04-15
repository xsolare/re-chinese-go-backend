package response

type Pinyin struct {
	Id          uint   `json:"id"`
	Pinyin      string `json:"pinyin"`
	InitialId   uint   `json:"initial_id"`
	FinalId     uint   `json:"final_id"`
	FinalToneId uint   `json:"final_tone_id"`
	Tone        string `json:"tone"`
	InitialPos  uint   `json:"initial_pos"`
	FinalPos    uint   `json:"final_pos"`
}
