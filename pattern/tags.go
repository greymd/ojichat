package pattern

import (
	"math/rand"
	"strings"
	"time"

	"github.com/mattn/go-gimei"
)

// 文章中一種類に統一されるタグ
var uniqTags = map[string][]string{
	// 対象の名前
	"{TARGET_NAME}": []string{},
	// お姫様の一人称
	"{FIRST_PERSON}": []string{
		"私",
		"わたし",
		"我",
		"うち",
		"小生",
	},
	// 曜日
	"{DAY_OF_WEEK}": []string{
		"月", "火", "水", "木", "金", "土", "日",
	},
	// 地名
	"{LOCATION}": []string{
		"愛知", "青森", "秋田", "石川", "茨城", "岩手", "愛媛", "大分", "大阪", "岡山", "沖縄", "香川", "鹿児島", "神奈川", "岐阜", "京都", "熊本", "群馬", "高知", "埼玉", "佐賀", "滋賀", "静岡", "島根", "千葉", "東京", "徳島", "栃木", "鳥取", "富山", "長崎", "長野", "奈良", "新潟", "兵庫", "広島", "福井", "福岡", "福島", "北海道", "三重", "宮城", "宮崎", "山形", "山口", "山梨", "和歌山",
	},
	// 外食
	"{RESTAURANT}": []string{
		"お寿司🍣",
		"イタリアン🍝",
		"バー🍷",
		"ラーメン屋さん🍜",
		"中華🍜",
		"フレンチ🍽",
	},
	// 食べ物
	"{FOOD}": []string{
		"お惣菜🍖",
		"サラダ🥗",
		"ピザ🍕",
		"パスタ🍝",
		"スイーツ🍮",
		"ケーキ🎂",
		"ビール🍺",
		"タピオカ🥟",
		"ぎょうざ🥟",
	},
	// 天気
	"{WEATHER}": []string{
		"曇り☁️",
		"晴れ🌞",
		"快晴☀️",
		"大雨☔️",
		"雨☂️",
		"雪❄️",
		"台風🌀",
	},
	// TODO: 「かな？」と「かい？」の語尾の違いも組み込んでも良いかもしれない
	// おじさんの欲望の地、ホテル
	"{TAPIOKA}": []string{
		"タピオカ",
		"ゴンチャ",
		"THE ALLEY",
	},
	// デートの種類
	"{DATE}": []string{
		"デート❤",
		"カラオケ🎤",
		"ドライブ🚗",
	},
	// おじさんは比喩で相手を持ち上げる (川柳)
	"{METAPHOR}": []string{
		"イケメン",
		"かっこいい",
		"天才",
		"ゴリラ",
		"優男",
	},
}

// 文章中複数回変更&繰り返されるタグ
var flexTags = map[string][]string{
	// ポジティブな表現の絵文字/顔文字
	"{EMOJI_POS}": []string{
		"🍣🍣🍣",
		"🍒🍒🍒",
		"😃",
		"‼️‼️",
		"😄",
		"😆",
		"😚",
		"😘",
		"😚",
		"💕",
		"💗",
		"😍",
		"😁",
		"😉",
		"☹️",
		"😂",
		"😊",
		"☺️",
		"^_^",
		"（笑）",
		"❕❕",
		"(T^T)",
		"💯",
		"💮",
	},
	// ネガティヴな表現の絵文字/顔文字
	"{EMOJI_NEG}": []string{
		"💦💦",
		"💔",
		"😱",
		"😰",
		"😭😭",
		"😓",
		"😣",
		"😖",
		"😥",
		"😢",
		"😮",
		"😬",
		"🤧",
		"🥺",
		"😠",
		"☹️",
		"😑",
		"💢💢",
		"(💢･ᴗ･)",
		":((;´☋`;)):",
		"(　˙-˙　)",
		"(^_^;",
	},
	// ニュートラルな感情を表す絵文字/顔文字
	"{EMOJI_NEUT}": []string{
		"💤",
		"😴",
		"🙂",
		"🤑",
		"✋",
		"😪",
		"😎",
		"😒",
		"😙",
		"😏",
		"😳",
		"😌",
		"🤭",
		"😑",
		"(^^;;",
		"💚💜❤️💛",
		"📛",
		"🍣🍣🍣",
		"🍒🍒🍒",
		"👏👏",
	},
	// 疑問を投げかけるときに利用される絵文字/顔文字
	"{EMOJI_ASK}": []string{
		"‼️‼️",
		"❓",
		"❔",
		"🤔🤔🤔",
		"😧😧😧",
		"✋(⌒-⌒; )",
	},
}

// ConvertTags ; message内にあるタグを置換して結果を返す
func ConvertTags(message, targetName string, emojiNumber int) string {
	rand.Seed(time.Now().UnixNano())
	if targetName != "" {
		// 引数として名前が存在した場合にはそれを使う
		uniqTags["{TARGET_NAME}"] = []string{targetName + randomNameSuffix()}
	} else {
		// 指定がない場合には gimei から選定
		uniqTags["{TARGET_NAME}"] = []string{randomFirstName() + randomNameSuffix()}
	}

	for tag, pat := range uniqTags {
		content := pat[rand.Intn(len(pat))]
		message = strings.Replace(message, tag, content, -1)
	}

	for tag, pat := range flexTags {
		n := strings.Count(message, tag)
		for i := 0; i < n; i++ {
			content := ""
			if emojiNumber > 0 {
				content = combineMultiplePatterns(pat, rand.Intn(emojiNumber)+1)
			} else {
				// Ohimesama could be seriously
				content = "。"
			}
			// タグを置換
			message = strings.Replace(message, tag, content, 1)
		}
	}
	return message
}

// combineMultiplePatterns: 複数のパターンをnumber分ランダムにつなげる
func combineMultiplePatterns(patterns []string, number int) string {
	result := ""
	if number <= len(patterns) {
		for i := 0; i < number; i++ {
			index := rand.Intn(len(patterns) - i)
			result += patterns[index]
			patterns[index], patterns[len(patterns)-1-i] = patterns[len(patterns)-1-i], patterns[index]
		}
	} else {
		for i := 0; i < number; i++ {
			result += patterns[rand.Intn(len(patterns))]
		}
	}
	return result
}

// gimeiから下の名前だけ取得
func randomFirstName() string {
	name := gimei.NewName()
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(3) {
	case 0:
		return name.First.Kanji()
	case 1:
		return name.First.Katakana()
	}
	return name.First.Hiragana()
}

// 「ちゃん」「先輩」などをランダムに返す
func randomNameSuffix() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	switch {
	// たまに呼び捨てにする
	case n < 10:
		return ""
	// あざとい
	case n < 30:
		return "先輩"
	// うざい
	case n < 40:
		return "（笑）"
	// 多くの場合「ちゃん」にする
	default:
		return "ちゃん"
	}
}
