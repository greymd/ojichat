package generator

import (
	"math/rand"
	"time"

	"github.com/greymd/ojichat/pattern"
)

// Config ... main で受け取られる引数、オプション
type Config struct {
	TargetName string `docopt:"<name>"`
	EmojiNum   int    `docopt:"--emoji"`
}

// Start ... おじさんの文言を生成
func Start(config Config) (string, error) {
	rand.Seed(time.Now().UnixNano())
	selectedMessage := ""
	// アルゴリズム (ONARA) を無作為に選定
	selectedOnara := pattern.Onara[rand.Intn(len(pattern.Onara))]

	// アルゴリズムそれぞれの感情に対応した文言を選定
	for _, s := range selectedOnara {
		selected := pattern.OnaraMessages[s]
		selectedMessage += selected[rand.Intn(len(selected))]
	}

	// タグを変換
	result := pattern.ConvertTags(selectedMessage, config.TargetName, config.EmojiNum)
	return result, nil
}
