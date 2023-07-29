package excel

var textMap map[TextMapHash]string

func init() {
	load("TextMap/TextMapEN.json", &textMap)
}

type TextMapHash uint32

func (h TextMapHash) String() string {
	return textMap[h]
}

type CurveData struct {
	Level      uint32
	CurveInfos []*CurveInfo
}

type CurveInfo struct {
	Type  string
	Arith string
	Value float32
}

func (c *CurveData) Type(typ string) (string, float32) {
	for _, v := range c.CurveInfos {
		if v.Type == typ {
			return v.Arith, v.Value
		}
	}
	return "", 0
}

type PropGrowCurves struct {
	Type      FightProp
	GrowCurve string
}

type FightPropData struct {
	PropType FightProp
	Value    float32
}

func FindCurveData(list []*CurveData, level uint32) *CurveData {
	for _, v := range list {
		if v.Level != level {
			continue
		}
		return v
	}
	return nil
}
