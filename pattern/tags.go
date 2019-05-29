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
	// おじさんの一人称
	"{FIRST_PERSON}": []string{
		"僕",
		"ボク",
		"俺",
		"オレ",
		"小生",
		"オジサン",
		"おじさん",
	},
	// 曜日
	"{DAY_OF_WEEK}": []string{
		"月", "火", "水", "木", "金", "土", "日",
	},
	// 食べ物
	"{FOOD}": []string{
		"お寿司🍣",
		"イタリアン🍕🍝",
		"パスタ🍝",
		"バー🍷",
		"ラーメン🍜",
	},
}

// 文章中複数回変更&繰り返されるタグ
var flexTags = map[string][]string{
	// ポジティブな表現の絵文字/顔文字
	"{EMOJI_POS}": []string{
		"❗",
		"☺",
		"💕",
		"😍",
		"♬",
		"♫",
		"☀",
		"👊",
		"😘",
		"(^_^)",
		"(^o^)",
		"(^з<)",
	},
	// ネガティヴな表現の絵文字/顔文字
	"{EMOJI_NEG}": []string{
		"💦",
		"💔",
		"😿",
		"🙀",
		"😱",
		"😰",
		"(◎ ＿◎;)",
		"(T_T)",
		"^^;",
		"(^_^;",
		"(・_・;",
		"(￣Д￣；；",
		"(^▽^;)",
		"(-_-;)",
	},
	// ニュートラルな感情を表す絵文字/顔文字
	"{EMOJI_NEUT}": []string{
		"(^^;;",
		"💤",
		"😴",
		"🙂",
		"🤑",
		"✋",
		"😪",
		"🛌",
		"😎",
		"（￣▽￣）",
		"(＃￣З￣)",
	},
	// 疑問を投げかけるときに利用される絵文字/顔文字
	"{EMOJI_ASK}": []string{
		"❓",
		"❗❓",
		"🤔",
		"😜⁉️",
		"（￣ー￣?）",
	},
}

// ConvertTags ; message内にあるタグを置換して結果を返す
func ConvertTags(message, targetName string, emojiNumber int) string {
	rand.Seed(time.Now().UnixNano())
	if targetName != "" {
		uniqTags["{TARGET_NAME}"] = []string{targetName + randomNameSuffix()}
	} else {
		uniqTags["{TARGET_NAME}"] = []string{randomFirstName() + randomNameSuffix()}
	}

	for tag, pat := range uniqTags {
		content := pat[rand.Intn(len(pat))]
		message = strings.ReplaceAll(message, tag, content)
	}

	for tag, pat := range flexTags {
		n := strings.Count(message, tag)
		for i := 0; i < n; i++ {
			content := combineMultiplePatterns(pat, emojiNumber)
			// タグを置換
			message = strings.Replace(message, tag, content, 1)
		}
	}
	return message
}

// combineMultiplePatterns: 複数のパターンをランダムにつなげる
func combineMultiplePatterns(patterns []string, number int) string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < rand.Intn(number+1); i++ {
		result += patterns[rand.Intn(len(patterns))]
	}
	return result
}

func randomFirstName() string {
	rand.Seed(time.Now().UnixNano())
	name := gimei.NewFemale()
	switch rand.Intn(2) {
	case 0:
		return name.First.Kanji()
	case 1:
		return name.First.Katakana()
	}
	return name.First.Hiragana()
}

// 「ちゃん」「チャン」などをランダムに返す
func randomNameSuffix() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	switch {
	// たまに呼び捨て
	case n < 5:
		return ""
	// そこそこ「チャン」にする
	case n < 40:
		return "チャン"
	// 多くの場合「ちゃん」にする
	default:
		return "ちゃん"
	}
}
