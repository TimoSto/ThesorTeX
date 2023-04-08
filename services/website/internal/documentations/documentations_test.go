package documentations

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetJsonDocs(t *testing.T) {
	t.Run("german doc", func(t *testing.T) {
		doc, err := GetJsonDocs("de")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		var docObj ThesorTexDocs
		err = json.Unmarshal(doc, &docObj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if diff := cmp.Diff(string(docObj.ThesisTemplate), string(templateDocJSON_DE)); diff != "" {
			t.Errorf("%s", diff)
		}
	})

	t.Run("english doc for thesis template", func(t *testing.T) {
		doc, err := GetJsonDocs("en")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		var docObj ThesorTexDocs
		err = json.Unmarshal(doc, &docObj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if diff := cmp.Diff(string(docObj.ThesisTemplate), string(templateDocJSON_EN)); diff != "" {
			t.Errorf("%s", diff)
		}
	})

	t.Run("doc for thesis template in different language", func(t *testing.T) {
		doc, err := GetJsonDocs("pl")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		var docObj ThesorTexDocs
		err = json.Unmarshal(doc, &docObj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if diff := cmp.Diff(string(docObj.ThesisTemplate), string(templateDocJSON_EN)); diff != "" {
			t.Errorf("%s", diff)
		}
	})
}
