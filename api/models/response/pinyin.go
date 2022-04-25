package response

type Pinyin struct {
	Id          uint   `json:"id"`
	Pinyin      string `json:"pinyin"`
	InitialId   uint   `json:"initialId"`
	FinalId     uint   `json:"finalId"`
	FinalToneId uint   `json:"finalToneId"`
	Tone        string `json:"tone"`
	InitialPos  uint   `json:"initialPos"`
	FinalPos    uint   `json:"finalPos"`
}
