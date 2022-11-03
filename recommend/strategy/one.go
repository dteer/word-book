package strategy

type Weight struct {
	Day        float64 `json:"day" text:"离第一次学习的天数差  权重占比"`
	Lately     float64 `json:"lately" text:"离最近一次学习的天数差 权重占比"`
	Complexity float64 `json:"complexity" text:"单词复杂度 权重占比"`
	ShowTime   float64 `json:"ShowTime" text:"展示次数   权重占比"`
}
type strategyOne struct {
	Weight
	DayWeightMap     map[int]float64 `json:"day_weight_map" text:"离第一次学习的天数差  权重占比(加权)"`
	DefaultDayWeight float64         `json:"default_day_weight" text:"默认的天数占比 权重占比(加权)"`
}

func NewStrategyOne() *strategyOne {
	s := &strategyOne{
		Weight: Weight{
			Day:        0.3,
			Complexity: 0.6,
			ShowTime:   0.01,
		},
		DayWeightMap: map[int]float64{
			1:  100,
			2:  80,
			5:  73,
			7:  75,
			10: 79,
		},
	}
	return s
}

/*
:func 获取X轴的值
:param day 距离当前学习的天数
:param complexity 单词复杂度
:param showTime 展示次数
*/
func (s *strategyOne) getX(day int, complexity int, showTime int) (x float64) {
	x = s.Weight.Day*float64(day) + s.Weight.Complexity*float64(complexity) + s.Weight.ShowTime*float64(showTime)
	// 添加决定因素
	b, ok := s.DayWeightMap[day]
	if !ok {
		b = s.DefaultDayWeight
	}
	x += b
	return
}

func (s *strategyOne) GetY(day int, complexity int, showTime int) (y float64) {
	x := s.getX(day, complexity, showTime)
	y = 1 / x
	return
}
