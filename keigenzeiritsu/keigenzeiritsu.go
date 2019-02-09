package keigenzeiritsu

func isKeigenItem(str string) bool {
	return isKeigenzeiritsu(name2Category(str))
}

func isKeigenzeiritsu(str string) bool {
	m := map[string]bool{"food": true, "beverage": true}
	result := m[str]
	return result
}

func name2Category(str string) string {
	m := map[string]string{
		"オロナミンC":                     "beverage",
		"リポビタンD":                     "quasi_drug",
		"手巻き直火焼き紅しゃけ":                "food",
		"からあげ棒":                      "food",
		"キリン 生茶 (小型容器） 555ml ペットボトル": "beverage",
		"キリンチューハイ氷結グレープフルーツ350ml缶":   "liquor",
		"新ルルＡ錠ｓ 50錠":                 "drug",
	}
	return m[str]
}

func itemPriceIncludeTax(item string) int {
	var zei float64
	if isKeigenItem(item) {
		zei = 1.08
	} else {
		zei = 1.10
	}

	return int(zei * (float64(itemPriceExcludeTax(item))))
}

func itemPriceExcludeTax(item string) int {
	m := map[string]int{
		"オロナミンC":                     105,
		"リポビタンD":                     146,
		"手巻き直火焼き紅しゃけ":                139,
		"からあげ棒":                      114,
		"キリン 生茶 (小型容器） 555ml ペットボトル": 140,
		"キリンチューハイ氷結グレープフルーツ350ml缶":   141,
		"新ルルＡ錠ｓ 50錠":                 871,
	}
	return m[item]
}
