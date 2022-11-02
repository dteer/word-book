package service

/*
	策略占比
*/

// 策略一
type strategyOne struct {
	Weight struct {
		Day        float64 // 最近一次因子权重占比
		Complexity float64 // 单词复杂度因子权重占比
		ShowTime   float64 // 展示次数因子权重占比
	}
	DayWeightMap     map[int]float64 // 最近一次距离现在的天数的比重
	DefaultDayWeight float64         // 默认的天数占比 1
}

func NewStrategyOne() *strategyOne {
	s := &strategyOne{
		Weight: struct {
			Day        float64
			Complexity float64
			ShowTime   float64
		}{
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
	return
}

func (s *strategyOne) getB(day int) (b float64) {
	b, ok := s.DayWeightMap[day]
	if !ok {
		b = s.DefaultDayWeight
	}
	return
}

func (s *strategyOne) GetY(day int, complexity int, showTime int) (y float64) {
	x := s.getX(day, complexity, showTime)
	b := s.getB(day)
	y = 1/x + b
	return
}
