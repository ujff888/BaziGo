package SiZhu

// 合化冲

import (
	. "github.com/warrially/BaziGo/Common"
)

// 天干五合
// 五合是专指天干而言的，地支为六合。
// 命局中干支五行力量的对比与变化，会充分表现在天干的五行上，
// 因此天干五合情况的出现，也会进一步影响或改变命局中的五行力量强弱关系。

// 甲己合化土， 乙庚合化金， 丙辛合化水， 丁壬合化木， 戊癸合化火。
func quickCheckTianGan(x, y int) int {
	if x == y {
		return -1 // 相同不合
	}

	x %= 5
	y %= 5

	if x == y {
		switch x {
		case 0:
			return 4 // 甲 阳木 + 己 阴土 = 合化土
		case 1:
			return 0 // 庚 阳金 + 乙 阴木 = 合化金
		case 2:
			return 2 // 丙 阳火 + 辛 阴金 = 合化水
		case 3:
			return 1 // 壬 阳水 + 丁 阴火 = 合化木
		case 4:
			return 3 // 戊 阳土 + 癸 阴水 = 合化火
		}
	}

	return -1
}

// 计算合
func GetTianGanHe(nGan1, nGan2, n5x int) string {
	if nGan1%2 == 0 {
		return "天干" + GetTianGanFromNumber(nGan1) + GetTianGanFromNumber(nGan2) + "合" + GetWuXingFromNumber(n5x)
	} else {
		return "天干" + GetTianGanFromNumber(nGan2) + GetTianGanFromNumber(nGan1) + "合" + GetWuXingFromNumber(n5x)
	}
}

// 查找天干五合
func CheckTianGanWuHe(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	// 查找天干五合是否存在
	// var WU_XING_STR = [5]string{
	// "金", "木", "水", "火", "土"}

	var nTgYear, nTgMonth, nTgDay, nTgHour int //
	nTgYear = pSiZhu.YearZhu.Gan.Value
	nTgMonth = pSiZhu.MonthZhu.Gan.Value
	nTgDay = pSiZhu.DayZhu.Gan.Value
	nTgHour = pSiZhu.HourZhu.Gan.Value

	var nCount int = 0

	// 检查年月
	if nHe := quickCheckTianGan(nTgYear, nTgMonth); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgMonth, nHe)
		nCount++
	}
	// 检查月日
	if nHe := quickCheckTianGan(nTgMonth, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgMonth, nTgDay, nHe)
		nCount++
	}
	// 检查日时
	if nHe := quickCheckTianGan(nTgDay, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgDay, nTgHour, nHe)
		nCount++
	}

	// 检查年日
	if nHe := quickCheckTianGan(nTgYear, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgDay, nHe)
		nCount++
	} else
	// 检查月时
	if nHe := quickCheckTianGan(nTgMonth, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgMonth, nTgHour, nHe)
		nCount++
	}
	// 检查年时
	if nHe := quickCheckTianGan(nTgYear, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgHour, nHe)
		nCount++
	}
}
