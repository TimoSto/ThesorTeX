Feature: Entry editor

  Scenario: Creating a new entry
    Given the url "/" was opened
    And the project " test " is opened
    When a new entry is created
    Then the title of the app is " ThesorTeX "
    And the title of the main area is ""
    And the save button in the editor is disabled
    When "e2" is entered as key
    And the first category is selected
    Then the save button in the editor is enabled
    And the fields have a length greater than 0
    When "hallo" is entered into the input at index 0
    And the save button in the editor is clicked
    And the back button is clicked
    Then following entries are displayed
      | key       |
      | e2        |
      | testEntry |
