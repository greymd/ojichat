package pattern

// OjisanEmotion ... おじさんの感情列挙体
type OjisanEmotion int

const (
	// GREEDING ... 挨拶
	GREEDING OjisanEmotion = iota
	// QUESTION ... 疑問
	QUESTION
	// REPORTING ... 報告
	REPORTING
	// CHEERING ... 応援
	CHEERING
	// INVITATION ... 誘い
	INVITATION
	// PRAISING ... 褒める
	PRAISING
	// ADMIRATION ... 自分が参った表現(感服)
	ADMIRATION
	// SYMPATHY ... 気遣い/慰め/同情
	SYMPATHY
)

// Onara ... Ojisan NArikiri Randomized Algorhytm: おじさんなりきり乱択アルゴリズム
// おじさんの感情表現の順番を表す。
// 既に近年の研究によりおじさんなりきるために効果的と考えられるアルゴリズムが特定されている。
var Onara = [][]OjisanEmotion{
	// GQS パターン
	[]OjisanEmotion{GREEDING, QUESTION, SYMPATHY},
	// GR パターン
	[]OjisanEmotion{GREEDING, REPORTING},
	// GC パターン
	[]OjisanEmotion{GREEDING, CHEERING},
	// GQI パターン
	[]OjisanEmotion{GREEDING, QUESTION, INVITATION},
	// PA パターン
	[]OjisanEmotion{PRAISING, ADMIRATION},
	// S パターン (短いので SS にする)
	[]OjisanEmotion{SYMPATHY, SYMPATHY},
}

// OnaraMessages .. メッセージのテンプレート
var OnaraMessages = [][]string{
	GREEDING: []string{
		"{TARGET_NAME}チャン{EMOJI_POS}",
		"{TARGET_NAME}ちゃん{EMOJI_POS}",
		"{TARGET_NAME}ちゃん、オハヨー{EMOJI_POS}",
		"{TARGET_NAME}ちゃん、ヤッホー{EMOJI_POS}何してる{EMOJI_ASK}",
	},
	QUESTION: []string{
		"今週の{DAY_OF_WEEK}曜日、仕事が早く終わりそうなんだけど、ご飯でもどう{EMOJI_ASK}",
		"今日はどんな一日だった{EMOJI_ASK}",
	},
	REPORTING: []string{
		"{FIRST_PERSON}は、近所に新しくできたラーメン屋さんに行ってきたよ。味はまぁまぁだったカナ{EMOJI_POS}",
	},
	CHEERING: []string{
		"今日も頑張ってね{EMOJI_POS}",
		"{TARGET_NAME}ちゃんにとって素敵な1日になりますように{EMOJI_POS}",
	},
	INVITATION: []string{
		"突然だけど、{TARGET_NAME}ちゃんは{FOOD}好きカナ{EMOJI_ASK}{DAY_OF_WEEK}曜日ご飯行こうよ",
	},
	PRAISING: []string{
		"可愛すぎ{EMOJI_POS}",
	},
	ADMIRATION: []string{
		"今から寝ようと思ってたのに、目が覚めちゃったよ{EMOJI_POS}どうしてくれるんだ{EMOJI_POS}",
	},
	SYMPATHY: []string{
		"今日も大変だったんだね{EMOJI_NEG}",
		"{FIRST_PERSON}は{TARGET_NAME}ちゃんの味方だからね{EMOJI_POS}",
		"今日はよく休んでね{EMOJI_NEUT}",
		"くれぐれも体調に気をつけて{EMOJI_NEUT}",
	},
}
