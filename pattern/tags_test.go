package pattern

import (
	"strings"
	"testing"
)

func TestConvertTags1(t *testing.T) {
	expected := "ウンコ"
	actual := ConvertTags("{TARGET_NAME}チャン！", "ウンコ", 1)
	t.Log(actual)
	if strings.Count(actual, expected) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestConvertTags2(t *testing.T) {
	expected := "ウンコ"
	actual := ConvertTags("{TARGET_NAME}チャン！", "", 1)
	t.Log(actual)
	if strings.Count(actual, expected) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestConvertTags3(t *testing.T) {
	expected := "。"
	actual := ConvertTags("{TARGET_NAME}いえい！{EMOJI_POS}", "", 0)
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

func TestCombineMultiplePatterns1(t *testing.T) {
	result := combineMultiplePatterns([]string{"A", "B", "C"}, 3)
	t.Log(result)
	if strings.Count(result, "A") != 1 {
		t.Errorf("handler returned unexpected body: got %v", result)
	}
	if strings.Count(result, "B") != 1 {
		t.Errorf("handler returned unexpected body: got %v", result)
	}
	if strings.Count(result, "C") != 1 {
		t.Errorf("handler returned unexpected body: got %v", result)
	}
}

func TestCombineMultiplePatterns2(t *testing.T) {
	result := combineMultiplePatterns([]string{"A", "B", "C"}, 4)
	t.Log(result)
	if len(result) != 4 {
		t.Errorf("handler returned unexpected body: got %v", result)
	}
}
