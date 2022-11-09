package one

import (
	"sort"
	"word-book/dao/word"
	util "word-book/pkg/utils"
)

type weightWord struct {
	word   word.Word
	weight float64
}

type weightWordList []weightWord

func (s weightWordList) Len() int { return len(s) }

func (s weightWordList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s weightWordList) Less(i, j int) bool { return s[i].weight > s[j].weight }

// todo 后续加入缓存处理
var OldWords []word.Word
var recommendWrods []word.Word

type Rcommend struct {
	Today int
	Words []word.Word
}

// 设置推荐的单词
func (r *Rcommend) setWord() {
	StartTimeGt := 0
	findData := word.FindData{
		StartTimeLt: &r.Today,
		StartTimeGt: &StartTimeGt,
	}
	r.Words = word.Find(findData)
}

// 获取单词的权重
func (r *Rcommend) GetWeightWord() []weightWord {
	var sortWord []weightWord
	s := NewStrategyOne()
	for _, data := range r.Words {
		data := data
		day := util.Interval(int(data.StartTime), r.Today)
		weight := s.GetY(day, data.Complexity, int(data.ShowTime))
		obj := weightWord{
			word:   data,
			weight: weight,
		}
		sortWord = append(sortWord, obj)
	}
	sort.Sort(sort.Reverse(weightWordList(sortWord)))
	return sortWord
}

// 保存推荐单词
func (r *Rcommend) SaveReommed(datas []weightWord) {
	var result []word.Word
	for _, data := range datas {
		data := data.word
		result = append(result, data)
	}
	r.Words = result
	recommendWrods = result
}

func HandleRommend() {
	r := Rcommend{
		Today: util.GetNowDay(),
	}
	r.setWord()
	datas := r.GetWeightWord()
	r.SaveReommed(datas)
}

func GetRecommendWrods(count int) (words []word.Word) {
	allLen := len(recommendWrods)
	if allLen >= count {
		return recommendWrods[:count]
	} else {
		return recommendWrods[:allLen]
	}
}
