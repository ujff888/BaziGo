package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/Days"
)

// 补充五行
func supplement(zhu TZhu) TZhu {
	// 转换字符串
	zhu.GanZhiStr = GetGanZhiFromNumber(zhu.GanZhi)
	zhu.GanStr = GetTianGanFromNumber(zhu.Gan)
	zhu.ZhiStr = GetDiZhiFromNumber(zhu.Zhi)

	// 五行
	zhu.G5X = Get5XingFromGan(zhu.Gan)
	zhu.Z5X = Get5XingFromZhi(zhu.Zhi)
	zhu.G5XStr = GetWuXingFromNumber(zhu.G5X)
	zhu.Z5XStr = GetWuXingFromNumber(zhu.Z5X)

	return zhu
}

// 从八字年获得年柱
func GetZhuFromYear(nYear int) TZhu {
	var zhu TZhu
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	zhu.GanZhi = GetGanZhiFromYear(nYear)
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	zhu.Gan, zhu.Zhi = ExtractGanZhi(zhu.GanZhi)

	return supplement(zhu)
}

// 从八字月 和 年干 获得月柱
func GetZhuFromMonth(nMonth int, nGan int) TZhu {
	var zhu TZhu
	// 根据口诀从本年干数计算本年首月的干数
	switch nGan {
	case 0, 5:
		// 甲己 丙佐首
		nGan = 2
	case 1, 6:
		// 乙庚 戊为头
		nGan = 4
	case 2, 7:
		// 丙辛 寻庚起
		nGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nGan = 0
	}

	// 计算本月干数
	nGan += ((nMonth - 1) % 10)

	zhu.Gan = nGan % 10
	zhu.Zhi = (nMonth - 1 + 2) % 12
	zhu.GanZhi = CombineGanZhi(zhu.Gan, zhu.Zhi)

	return supplement(zhu)
}

// 从公历天 获得日柱
func GetZhuFromDay(nYear int, nMonth int, nDay int) TZhu {
	var zhu TZhu

	zhu.GanZhi = Days.GetGanZhiFromDay(Days.GetAllDays(nYear, nMonth, nDay))
	// 获得八字日的干0-9 对应 甲到癸
	// 获得八字日的支0-11 对应 子到亥
	zhu.Gan, zhu.Zhi = ExtractGanZhi(zhu.GanZhi)

	return supplement(zhu)
}

// 从公历小时,  获得日柱天干获取时柱
func GetZhuFromHour(nHour int, nGan int) TZhu {
	var zhu TZhu

	zhu.Gan, zhu.Zhi = Days.GetGanZhiFromHour(nHour, nGan)
	zhu.GanZhi = CombineGanZhi(zhu.Gan, zhu.Zhi)

	return supplement(zhu)
}