package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/greymd/ojichat/pattern"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	selectedOnara := pattern.Onara[rand.Intn(len(pattern.Onara))]
	for _, s := range selectedOnara {
		selected := pattern.OnaraMessages[s]
		msg := selected[rand.Intn(len(selected))]
		fmt.Printf("%s\n", msg)
	}
}
