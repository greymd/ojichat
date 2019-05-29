package pattern

import (
	"strings"
	"testing"
)

func TestConvertTags(t *testing.T) {
	expected := "ウンコ"
	actual := ConvertTags("{TARGET_NAME}チャン！", "ウンコ", 1)
	t.Log(actual)
	if strings.Count(actual, expected) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRandomFirstName1(t *testing.T) {
	anexpected := ""
	actual := randomFirstName()
	t.Log(actual)
	if actual == anexpected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, anexpected)
	}
}
