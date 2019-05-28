package pattern

import "testing"

func TestConvertTags(t *testing.T) {
	expected := "ウンコチャン！"
	actual := ConvertTags("{TARGET_NAME}チャン！", "ウンコ", 1)
	t.Log(actual)
	if actual == expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
