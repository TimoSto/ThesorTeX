Feature: Entry editor

  Scenario: Creating a new entry
    Given the url "/" was opened
    And the project " test " is opened
    When a new entry is created
    Then the title of the app is " ThesorTeX "
    And the title of the main area is ""
    And the save button in the editor is disabled
