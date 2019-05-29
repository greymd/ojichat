package generator

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/greymd/ojichat/pattern"
	"github.com/ikawaha/kagome/tokenizer"
	"github.com/miiton/kanaconv"
)

// PunctuationConfig ... 句読点挿入の設定
type PunctuationConfig struct {
	TargetHinshis []string
	Rate          int
}

var pconfigs = []PunctuationConfig{
	{
		TargetHinshis: []string{},
		Rate:          0,
	},
	{
		TargetHinshis: []string{"助動詞"},
		Rate:          30,
	},
	{
		TargetHinshis: []string{"助動詞", "助詞"},
		Rate:          60,
	},
	{
		TargetHinshis: []string{"助動詞", "助詞"},
		Rate:          100,
	},
}

// Config ... main で受け取られる引数、オプション
type Config struct {
	TargetName        string `docopt:"<name>"`
	EmojiNum          int    `docopt:"-e"`
	PunctiuationLebel int    `docopt:"-p"`
}

// Start ... おじさんの文言を生成
func Start(config Config) (string, error) {
	rand.Seed(time.Now().UnixNano())
	selectedMessage := ""
	// アルゴリズム (ONARA) を無作為に選定
	selectedOnara := pattern.Onara[rand.Intn(len(pattern.Onara))]

	// アルゴリズム内で表現されたそれぞれの感情に対応した文言を選定
	for _, s := range selectedOnara {
		selected := pattern.OnaraMessages[s]
		selectedMessage += selected[rand.Intn(len(selected))]
		// 挨拶以外の感情に関しては語尾を最大2文字までカタカナに変換するおじさんカタカナ活用を適用する
		if s != pattern.GREEDING {
			selectedMessage = katakanaKatsuyou(selectedMessage, rand.Intn(3))
		}
	}

	// タグを変換
	selectedMessage = pattern.ConvertTags(selectedMessage, config.TargetName, config.EmojiNum)

	level := config.PunctiuationLebel
	if level < 0 || level > 3 {
		return "", fmt.Errorf("句読点挿入頻度レベルが不正です: %v", level)
	}
	// 句読点レベルに応じて、おじさんがよくやる句読点を適宜挿入する
	result := insertPunctuations(selectedMessage, pconfigs[level])

	return result, nil
}

// カタカナ活用を適用する
func katakanaKatsuyou(message string, number int) string {
	var reg *regexp.Regexp
	if number < 1 {
		return message
	}
	reg = regexp.MustCompile(`^(.+)(\p{Hiragana}{` + strconv.Itoa(number) + `})([^\p{Hiragana}]*)$`)
	hiraganas := reg.FindSubmatch([]byte(message))
	if len(hiraganas) != 4 {
		return message
	}
	return string(hiraganas[1]) + kanaconv.HiraganaToKatakana(string(hiraganas[2])) + string(hiraganas[3])
}

// 句読点レベルに応じ、助詞、助動詞の後に句読点を挿入する
func insertPunctuations(message string, config PunctuationConfig) string {
	if config.Rate == 0 {
		return message
	}
	rand.Seed(time.Now().UnixNano())
	result := ""
	t := tokenizer.New()
	tokens := t.Tokenize(message)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		features := token.Features()
		hinshiFlag := false
		for _, hinshi := range config.TargetHinshis {
			if hinshi == features[0] {
				hinshiFlag = true
				break
			}
		}
		if hinshiFlag && rand.Intn(100) <= config.Rate {
			result += token.Surface + "、"
		} else {
			result += token.Surface
		}
	}
	return result
}
