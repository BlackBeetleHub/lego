package main

type Dictionary interface {
	GetUnlearnedWorldsAtPage(page int)
	GetAllUnlearnedWords()
	GetCountWords()
}

type Word interface {
	GetWord()
	GetTranslate()
	HasLearned()
}

type Translate interface {
	GetValue()
}

type LeoTranslateImpl struct {
	translate_value string
}

type LeoDictionaryImpl struct {
	error_msg string
	count_words string
	userdict3 []LeoWordImpl
}

func (impl *LeoDictionaryImpl)getCount()  {
	println(impl.count_words)
}

type LeoWordImpl struct {
	word_id int
	word_value string
	transcription string
	progress_percent int
	user_translates []LeoTranslateImpl
}