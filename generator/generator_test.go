package generator

import "testing"

func TestKatakanaKatsuyou(t *testing.T) {
	expected := "なんちゃッテ"
	actual := katakanaKatsuyou("なんちゃって", 2)
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
