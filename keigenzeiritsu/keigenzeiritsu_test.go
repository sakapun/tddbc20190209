package keigenzeiritsu

import "testing"

func Test_単体の価格を計算して返す(t *testing.T) {
	t.Run("オロナミンCは113である", func(t *testing.T) {
		checkEqualInt(t, itemPriceTax("オロナミンC"), 113)
	})
	//t.Run("リポビタンDはquasi_drugである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("リポビタンD"), "quasi_drug")
	//})
	//t.Run("手巻き直火焼き紅しゃけはfoodである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("手巻き直火焼き紅しゃけ"), "food")
	//})
	//t.Run("からあげ棒はfoodである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("からあげ棒"), "food")
	//})
	//t.Run("キリン 生茶 (小型容器） 555ml ペットボトルはbeverageである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("キリン 生茶 (小型容器） 555ml ペットボトル"), "beverage")
	//})
	//t.Run("キリンチューハイ氷結グレープフルーツ350ml缶はliquorである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("キリンチューハイ氷結グレープフルーツ350ml缶"), "liquor")
	//})
	//t.Run("新ルルＡ錠ｓ 50錠はdrugである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("新ルルＡ錠ｓ 50錠"), "drug")
	//})
}

func Test_単体の価格を税抜きで返す(t *testing.T) {
	t.Run("オロナミンCは105である", func(t *testing.T) {
		checkEqualInt(t, itemPriceExcludeTax("オロナミンC"), 105)
	})
	// todo: テストちゃんと書く
	//t.Run("リポビタンDはquasi_drugである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("リポビタンD"), "quasi_drug")
	//})
	//t.Run("手巻き直火焼き紅しゃけはfoodである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("手巻き直火焼き紅しゃけ"), "food")
	//})
	//t.Run("からあげ棒はfoodである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("からあげ棒"), "food")
	//})
	//t.Run("キリン 生茶 (小型容器） 555ml ペットボトルはbeverageである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("キリン 生茶 (小型容器） 555ml ペットボトル"), "beverage")
	//})
	//t.Run("キリンチューハイ氷結グレープフルーツ350ml缶はliquorである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("キリンチューハイ氷結グレープフルーツ350ml缶"), "liquor")
	//})
	//t.Run("新ルルＡ錠ｓ 50錠はdrugである", func(t *testing.T) {
	//	checkEqualStr(t, name2Category("新ルルＡ錠ｓ 50錠"), "drug")
	//})
}

func Test_商品名は軽減税率対象である(t *testing.T) {
	t.Run("オロナミンCは軽減対象", func(t *testing.T) {
		checkEqualBool(t, isKeigenItem("オロナミンC"), true)
	})
	t.Run("リポビタンDは軽減対象ではない", func(t *testing.T) {
		checkEqualBool(t, isKeigenItem("リポビタンD"), false)
	})
}

func Test_カテゴリが軽減税率対象である(t *testing.T) {
	t.Run("foodは軽減税率である", func(t *testing.T) {
		checkEqualBool(t, isKeigenzeiritsu("food"), true)
	})
	t.Run("beverageは軽減税率である", func(t *testing.T) {
		checkEqualBool(t, isKeigenzeiritsu("beverage"), true)
	})
	t.Run("liquorは軽減税率でない", func(t *testing.T) {
		checkEqualBool(t, isKeigenzeiritsu("liquor"), false)
	})
	t.Run("drugは軽減税率でない", func(t *testing.T) {
		checkEqualBool(t, isKeigenzeiritsu("drug"), false)
	})
}

func Test_食べ物のカテゴリー分け(t *testing.T) {
	t.Run("オロナミンCはbeverageである", func(t *testing.T) {
		checkEqualStr(t, name2Category("オロナミンC"), "beverage")
	})
	t.Run("リポビタンDはquasi_drugである", func(t *testing.T) {
		checkEqualStr(t, name2Category("リポビタンD"), "quasi_drug")
	})
	t.Run("手巻き直火焼き紅しゃけはfoodである", func(t *testing.T) {
		checkEqualStr(t, name2Category("手巻き直火焼き紅しゃけ"), "food")
	})
	t.Run("からあげ棒はfoodである", func(t *testing.T) {
		checkEqualStr(t, name2Category("からあげ棒"), "food")
	})
	t.Run("キリン 生茶 (小型容器） 555ml ペットボトルはbeverageである", func(t *testing.T) {
		checkEqualStr(t, name2Category("キリン 生茶 (小型容器） 555ml ペットボトル"), "beverage")
	})
	t.Run("キリンチューハイ氷結グレープフルーツ350ml缶はliquorである", func(t *testing.T) {
		checkEqualStr(t, name2Category("キリンチューハイ氷結グレープフルーツ350ml缶"), "liquor")
	})
	t.Run("新ルルＡ錠ｓ 50錠はdrugである", func(t *testing.T) {
		checkEqualStr(t, name2Category("新ルルＡ錠ｓ 50錠"), "drug")
	})
}

//func カテゴリが軽減税率対象ならtrueをでなければfalseを(t )

func TestKeigenzeiritsu(t *testing.T) {
	checkEqualStr(t, "hoge", "hoge")
}

func checkEqualBool(t *testing.T, hoge bool, fuga bool) {
	if hoge != fuga {
		t.Errorf("adfasdfsaf")
	}
}

func checkEqualStr(t *testing.T, hoge string, fuga string) {
	if hoge != fuga {
		t.Errorf("adfasdfsaf")
	}
}

func checkEqualInt(t *testing.T, hoge int, fuga int) {
	if hoge != fuga {
		t.Errorf("adfasdfsaf")
	}
}
