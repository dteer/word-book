package word

import "word-book/run"

func Find() Word {
	db := run.GetDefaultGorm()
	var word Word
	db.First(&word, 1)
	return word
}
