package json

type LeoTranslateImpl struct {
	Translate_value string
}

type LeoWordImpl struct {
	Word_id int
	Word_value string
	Transcription string
	User_translates []LeoTranslateImpl
}

type LeoWordsImpl struct {
	Words []LeoWordImpl
}

type LeoDictionaryImpl struct {
	Error_msg   string
	Count_words int
	User_dict     []LeoWordsImpl `json:"userdict3"`
}

func (impl *LeoDictionaryImpl) GetCountWords() int {
	return impl.Count_words
}

func (impl *LeoWordImpl) GetWord() string {
	return impl.Word_value
}

func (impl *LeoWordImpl) HasLearned() bool{
	return false
}

func (impl *LeoTranslateImpl) GetValue() string {
	return impl.Translate_value;
}

func(impl* LeoWordImpl) GetTranslates() []Translate {
	var result []Translate
	for _, value := range impl.User_translates {
		result = append(result, &value)
	}
	return result
}

func (impl *LeoDictionaryImpl) GetWords() []Word {
	var result []Word
	for _, v := range impl.User_dict[0].Words {
		result = append(result, &LeoWordImpl{
			v.Word_id,
			v.Word_value,
			v.Transcription,
			v.User_translates})
	}
	return result
}

