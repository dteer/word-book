package service

import context "context"

type wordService struct {
}

func (w *wordService) GetRecommendWord(context.Context, *ShowRequest) (*ShowResponse, error) {
		// 刷新推荐
		var funcList = make([]func(), 1)
		funcList[0] = one.HandleRommend
		reFresh := util.TimingRefresh{
			F:     funcList,
			Times: config.C.Common.RemmandInterval,
		}
		go reFresh.Run()
	
		newCount := config.C.Common.New
		oldCount := config.C.Common.Old
		Today := util.GetNowDay()
		newWords := getNewWord(newCount, Today)
		wordChanel := make(chan word.Word, 20)
		go handle(wordChanel)
		for {
			words := append(newWords, one.GetRecommendWrods(oldCount)...)
			for _, w := range words {
				popWord(w)
				wordChanel <- w
				time.Sleep(time.Duration(config.C.Common.Interval) * time.Second)
			}
		}
	}	
}

// 通过队列进行处理相关信息
func handle(wordLi chan word.Word) {
	data := <-wordLi
	word.UdateShowTime(data)
}

func mustEmbedUnimplementedProdServiceServer() {

}
