package parser

import "testing"

func TestGetRelevantNotes(t *testing.T) {
	file := `# v1.0.0

## Updated Dependencies

- Typescript asf
- Vue sdfs

## Changes build mechanics

- Build entirely using bazel

## Added

- Some cool feature

# v0.0.2

## Updated Dependencies

- Typescript 4.9.4 -> 4.9.5
- Vue 3.2.36 -> 3.2.47

## Changes build mechanics

- Build entirely using bazel

# v0.0.1

## Added

- The webapp for the local app was setup.
- Categories can be created, edited and deleted in the gui.
- Entries can be created, edited and deleted in the gui.`

	notes := GetRelevantNotes([]byte(file), "0.0.2")

	expected := `# v0.0.2

## Updated Dependencies

- Typescript 4.9.4 -> 4.9.5
- Vue 3.2.36 -> 3.2.47

## Changes build mechanics

- Build entirely using bazel`

	if notes != expected {
		t.Errorf("expected '%v' but got '%v'", expected, notes)
	}
}

func TestGetRelevantNotes_Last(t *testing.T) {
	file := `# v1.0.0

## Updated Dependencies

- Typescript asf
- Vue sdfs

## Changes build mechanics

- Build entirely using bazel

## Added

- Some cool feature

# v0.0.2

## Updated Dependencies

- Typescript 4.9.4 -> 4.9.5
- Vue 3.2.36 -> 3.2.47

## Changes build mechanics

- Build entirely using bazel`

	notes := GetRelevantNotes([]byte(file), "0.0.2")

	expected := `# v0.0.2

## Updated Dependencies

- Typescript 4.9.4 -> 4.9.5
- Vue 3.2.36 -> 3.2.47

## Changes build mechanics

- Build entirely using bazel`

	if notes != expected {
		t.Errorf("expected '%v' but got '%v'", expected, notes)
	}
}

func TestGetRelevantNotes_First(t *testing.T) {
	file := `# v1.0.0

## Updated Dependencies

- Typescript asf
- Vue sdfs

## Changes build mechanics

- Build entirely using bazel

## Added

- Some cool feature

# v0.0.2

## Updated Dependencies

- Typescript 4.9.4 -> 4.9.5
- Vue 3.2.36 -> 3.2.47

## Changes build mechanics

- Build entirely using bazel`

	notes := GetRelevantNotes([]byte(file), "1.0.0")

	expected := `# v1.0.0

## Updated Dependencies

- Typescript asf
- Vue sdfs

## Changes build mechanics

- Build entirely using bazel

## Added

- Some cool feature`

	if notes != expected {
		t.Errorf("expected '%v' but got '%v'", expected, notes)
	}
}
