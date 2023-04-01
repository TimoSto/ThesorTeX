package documentations

import "testing"

func TestGetDocumentation(t *testing.T) {
	t.Run("german doc for thesis template", func(t *testing.T) {
		doc, err := GetDocumentation("thesisTemplate", "JSON", "de")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(doc) == 0 {
			t.Errorf("unexpected empty doc")
		}
	})

	//t.Run("english doc for thesis template", func(t *testing.T) {
	//	doc, err := GetDocumentation("thesisTemplate", "JSON", "en")
	//	if err != nil {
	//		t.Errorf("unexpected error: %v", err)
	//		return
	//	}
	//	if len(doc) == 0 {
	//		t.Errorf("unexpected empty doc")
	//	}
	//})
}
