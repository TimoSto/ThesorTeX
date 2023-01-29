Feature: Category editor

  Scenario: Creating a new category
    Given the url "/" was opened
    And the project " test " is opened
    When a new category is created
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is ""
    And the save button in the editor is disabled
    When "c1" is entered as name
    Then the save button in the editor is enabled
    When a bib field is added
    And "field" is entered as the name of the field at index 0
    And " " is entered as the suffix of the field at index 0
    When a bib field is added
    And "field2" is entered as the name of the field at index 1
    And "(" is entered as the prefix of the field at index 1
    And ")" is entered as the suffix of the field at index 1
    And the save button in the editor is clicked
    And the back button is clicked
    Then the displayed example entry for "c1" is "field (field2)"
