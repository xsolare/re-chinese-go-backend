package response

type HieroglyphKey struct {
	Id                    uint                     `json:"id"`
	Index                 uint                     `json:"index"`
	Priority              uint8                    `json:"priority"`
	AlternativeHieroglyph string                   `json:"alternativeHieroglyph"`
	HieroglyphId          uint                     `json:"-"`
	Hieroglyph            Hieroglyph               `json:"hieroglyph"`
	Translate             []HieroglyphKeyTranslate `json:"translate"`
}

type Hieroglyph struct {
	Id         uint   `json:"id"`
	Hieroglyph string `json:"hieroglyph"`
	Hsk        uint8  `json:"hsk"`
	PinyinId   uint   `json:"pinyinId"`
}

func (s *Hieroglyph) TableName() string {
	return "hieroglyphic"
}

type HieroglyphKeyTranslate struct {
	Translate       string `json:"translate"`
	HieroglyphKeyId uint   `json:"-"`
	LanguageId      uint   `json:"languageId"`
}

func (s *HieroglyphKeyTranslate) TableName() string {
	return "hieroglyphic_keys_translates"
}
