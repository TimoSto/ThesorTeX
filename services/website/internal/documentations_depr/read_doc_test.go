package documentations

import "testing"

func TestReadDoc(t *testing.T) {
	result, err := readDoc("thesis_template_usage/abbreviations", "de")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := Doc{
		Title:   "Wie kann ich das Abkürzungsverzeichnis bearbeiten?",
		Content: "Die Abkürzungen werden aus der Datei *abkuerzungen.csv* ausgelesen. Dort kannst du sie in der Form ***Abkürzung***;\n***Bedeutung*** eintragen.\n\nWenn du den Befehl ***\\printabbreviations*** in deinem Dokument aufrufst, wird das Abkürzungsverzeichnis ausegeben.",
	}

	if result.Title != expected.Title {
		t.Errorf("expected title %v, got %v", expected.Title, result.Title)
	}

	if result.Content != expected.Content {
		t.Errorf("expected content %v, got %v", expected.Content, result.Content)
	}
}
