package pattern

var tags = map[string][]string{
	// 対象の名前
	"{TARGET_NAME}": []string{
		"優子",
	},
	// おじさんの一人称
	"{FIRST_PERSON}": []string{
		"僕",
		"ボク",
		"俺",
		"オレ",
		"小生",
	},
	// ポジティブな表現の絵文字/顔文字
	"{EMOJI_POS}": []string{
		"❗",
		"(^_^)",
	},
	// ネガティヴな表現の絵文字/顔文字
	"{EMOJI_NEG}": []string{
		"(T_T)",
		"💦",
	},
	// ニュートラルな感情を表す絵文字/顔文字
	"{EMOJI_NEUT}": []string{
		"(^^;;",
		"💤",
		"😴",
	},
	// 疑問を投げかけるときに利用される絵文字/顔文字
	"{EMOJI_ASK}": []string{
		"❓",
	},
	// 食べ物
	"{FOOD}": []string{
		"お寿司🍣",
		"イタリアン🍝",
	},
	"{DAY_OF_WEEK}": []string{
		"月", "火", "水", "木", "金", "土", "日",
	},
}
