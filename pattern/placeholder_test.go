package pattern

import "testing"

func TestConvertTags(t *testing.T) {
	input := "{TARGET_NAME}チャン"
	expected := "優子チャン"
	actual := ConvertTags(input)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
