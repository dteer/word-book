package service

import (
	context "context"
	"word-book/config"
	util "word-book/service/pkg/utils"
	"word-book/service/recommend/strategy/one"
)

type wordService struct {
}

var ProdService = &wordService{}

func (w *wordService) GetRecommendWord(context.Context, *ShowRequest) (*ShowResponse, error) {
	newCount := config.C.Common.New
	oldCount := config.C.Common.Old
	Today := util.GetNowDay()
	newWords := getNewWord(newCount, Today)
	words := append(newWords, one.GetRecommendWrods(oldCount)...)
	var datas []*Data
	for _, word := range words {
		tmp := Data{
			ID:              word.ID,
			Title:           word.Title,
			PhoneticStymbol: word.PhoneticStymbol,
			Description:     word.Description,
		}
		datas = append(datas, &tmp)
	}
	data := &ShowResponse{
		Code: true,
		Data: datas,
	}
	return data, nil

}

func (w *wordService) mustEmbedUnimplementedProdServiceServer() {}
