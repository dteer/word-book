package service

import (
	"sort"
	"word-book/dao/word"
	util "word-book/pkg/utils"
)

func handleAddWord(count int) (words []word.Word) {
	StartTime := 0
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		StartTime: &StartTime,
		Order:     "id asc",
	}
	words = word.Find(findData)
	return
}

func getNewWord(count int, tody int) (words []word.Word) {
	// 获取今天需要学习的数据
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		StartTime: &tody,
	}
	words = word.Find(findData)
	// 如果今天还没学习，需要录入信息
	if len(words) < count {
		addWords := handleAddWord(count - len(words))
		word.SetTodayStartTime(addWords, tody)
		words = append(words, addWords...)
	}
	return
}

// todo 后续加入缓存处理
var OldWords []word.Word
var recommendWrods []word.Word

func getAllOldWord(tody int) (words []word.Word) {
	StartTimeGt := 0
	findData := word.FindData{
		StartTimeLt: &tody,
		StartTimeGt: &StartTimeGt,
	}
	if OldWords != nil {
		return OldWords
	}
	words = word.Find(findData)
	OldWords = words
	return
}

type weightWord struct {
	word   word.Word
	weight float64
}

type weightWordList []weightWord

func (s weightWordList) Len() int { return len(s) }

func (s weightWordList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s weightWordList) Less(i, j int) bool { return s[i].weight > s[j].weight }

func recommendSort() {
	var sortWord []weightWord
	today := util.GetNowDay()
	allData := getAllOldWord(today)
	s := NewStrategyOne()
	for _, data := range allData {
		data := data
		day := util.Interval(int(data.StartTime), today)
		weight := s.GetY(day, data.Complexity, int(data.ShowTime))
		obj := weightWord{
			word:   data,
			weight: weight,
		}
		sortWord = append(sortWord, obj)
	}
	sort.Sort(sort.Reverse(weightWordList(sortWord)))
	// 恢复原数据
	var result []word.Word
	for _, data := range sortWord {
		data := data.word
		result = append(result, data)
	}
	recommendWrods = result
}

func getRecommendWrods(count int) (words []word.Word) {
	recommendSort()
	allLen := len(recommendWrods)
	if allLen >= count {
		return recommendWrods[:count]
	} else {
		return recommendWrods[:allLen]
	}
}
