package pattern

import (
	"math/rand"
	"strings"
	"time"

	"github.com/mattn/go-gimei"
)

/*
## 参考文献
[3] 女子高生「おじさんLINEごっこ」の実例に学ぶキモがられる態度とは | ニュース3面鏡 | ダイヤモンド・オンライン
https://diamond.jp/articles/-/143111?page=2
*/

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
	// 地名
	"{LOCATION}": []string{
		"愛知", "青森", "秋田", "石川", "茨城", "岩手", "愛媛", "大分", "大阪", "岡山", "沖縄", "香川", "鹿児島", "神奈川", "岐阜", "京都", "熊本", "群馬", "高知", "埼玉", "佐賀", "滋賀", "静岡", "島根", "千葉", "東京", "徳島", "栃木", "鳥取", "富山", "長崎", "長野", "奈良", "新潟", "兵庫", "広島", "福井", "福岡", "福島", "北海道", "三重", "宮城", "宮崎", "山形", "山口", "山梨", "和歌山",
	},
	// 食べ物
	"{FOOD}": []string{
		"お寿司🍣",
		"イタリアン🍕🍝",
		"パスタ🍝",
		"バー🍷",
		"ラーメン🍜",
	},
	// 下ネタの後は「ナンチャッテ」「冗談（笑）」を使う[3]
	"{NANCHATTE}": []string{
		"ナンチャッテ{EMOJI_POS}",
		"なんちゃって{EMOJI_POS}",
		"なんてね{EMOJI_POS}",
		"冗談{EMOJI_POS}",
		"", // おじさんはたまに本気
	},
}

// 文章中複数回変更&繰り返されるタグ
var flexTags = map[string][]string{
	// ポジティブな表現の絵文字/顔文字
	"{EMOJI_POS}": []string{
		"❗",
		"☺",
		"😄",
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
		"（笑）",
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
		// 引数として名前が存在した場合にはそれを使う
		uniqTags["{TARGET_NAME}"] = []string{targetName + randomNameSuffix()}
	} else {
		// 指定がない場合には gimei から選定
		uniqTags["{TARGET_NAME}"] = []string{randomFirstName() + randomNameSuffix()}
	}

	for tag, pat := range uniqTags {
		content := pat[rand.Intn(len(pat))]
		message = strings.ReplaceAll(message, tag, content)
	}

	for tag, pat := range flexTags {
		n := strings.Count(message, tag)
		for i := 0; i < n; i++ {
			content := combineMultiplePatterns(pat, rand.Intn(emojiNumber)+1)
			// タグを置換
			message = strings.Replace(message, tag, content, 1)
		}
	}
	return message
}

// combineMultiplePatterns: 複数のパターンをnumber分ランダムにつなげる
func combineMultiplePatterns(patterns []string, number int) string {
	result := ""
	// TODO: 同じパターンは使い回さないようにしたほうが自然か？
	for i := 0; i < number; i++ {
		result += patterns[rand.Intn(len(patterns))]
	}
	return result
}

// gimei から女性の名前を無作為に選定
func randomFirstName() string {
	name := gimei.NewFemale()
	switch rand.Intn(3) {
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
	// "時に「◯◯チャン」とカタカナにしてくるのも、おじさんの常套手段だ。"[3]
	case n < 40:
		return "チャン"
	// 多くの場合「ちゃん」にする
	default:
		return "ちゃん"
	}
}
