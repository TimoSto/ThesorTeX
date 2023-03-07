Feature: Entry editor

  Scenario: Creating a new entry
    Given the url "/" was opened
    And the project " test " is opened
    When a new entry is created
    Then the title of the app is " ThesorTeX  - Edit entry"
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
    And the displayed entry for "e2" is "hallo ():  In:  , "

  Scenario: Edit entry
    Given the url "/" was opened
    And the project " test " is opened
    And the entry "e2" is opened
    When "feld2" is entered into the input at index 1
    And the save button in the editor is clicked
    And the back button is clicked
    Then following entries are displayed
      | key       |
      | e2        |
      | testEntry |
    And the displayed entry for "e2" is "hallo (feld2):  In:  , "

  Scenario: Unsafe close abort
    Given the url "/" was opened
    And the project " test " is opened
    And the entry "e2" is opened
    When "feld3" is entered into the input at index 2
    And the back button is clicked
    Then the user is prompted that there are unsaved changes
    When the close is aborted
    Then the title of the app is " ThesorTeX  - Edit entry"
    And the title of the main area is "e2"

  Scenario: Unsafe close confirm
    Given the url "/" was opened
    And the project " test " is opened
    And the entry "e2" is opened
    When "feld3" is entered into the input at index 2
    And the back button is clicked
    Then the user is prompted that there are unsaved changes
    When the close is confirmed
    Then the title of the app is " ThesorTeX  - Projectview"
    And the title of the main area is "test"
    When the back button is clicked
    Then the title of the app is " ThesorTeX "
    And the title of the main area is "Welcome to ThesorTeX!"

  Scenario: Delete entry
    Given the url "/" was opened
    And the project " test " is opened
    And the entry "e2" is opened
    And the delete button in the editor is clicked
    Then the user is asked to confirm the deletion of the entry
    When the deletion is confirmed
    Then the editor-page is closed
    And the title of the app is " ThesorTeX  - Projectview"
    And the title of the main area is "test"
    And following entries are displayed
      | key       |
      | testEntry |