Feature: Category editor

  Scenario: Creating a new category
    Given the url "/" was opened
    And the project " test " is opened
    When a new category is created
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is "Create new category"
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
    And following categories are displayed
      | name                   |
      | CitaviArticle          |
      | CitaviArticleDoi       |
      | CitaviBook             |
      | CitaviBookDoi          |
      | CitaviBooklet          |
      | CitaviInBook           |
      | CitaviInBookDoi        |
      | CitaviInCollection     |
      | CitaviInCollectionDoi  |
      | CitaviInProceedings    |
      | CitaviInProceedingsDoi |
      | aufsatz                |
      | c1                     |
    Then the displayed example entry for "c1" is "field (field2)"

  Scenario: Edit category
    Given the url "/" was opened
    And the project " test " is opened
    And the category "c1" is opened
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is "c1"
    When a bib field is added
    And "field3" is entered as the name of the field at index 2
    And ", " is entered as the prefix of the field at index 2
    And the save button in the editor is clicked
    And the back button is clicked
    Then the displayed example entry for "c1" is "field (field2), field3"

  Scenario: unsafe close abort
    Given the url "/" was opened
    And the project " test " is opened
    And the category "c1" is opened
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is "c1"
    When a bib field is added
    And "field3" is entered as the name of the field at index 2
    And ", " is entered as the prefix of the field at index 2
    And the back button is clicked
    Then the user is prompted that there are unsaved changes
    When the close is aborted
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is "c1"

  Scenario: unsafe close confirm
    Given the url "/" was opened
    And the project " test " is opened
    And the category "c1" is opened
    Then the title of the app is " ThesorTeX  - Category editor"
    And the title of the main area is "c1"
    When a bib field is added
    And "field3" is entered as the name of the field at index 2
    And ", " is entered as the prefix of the field at index 2
    And the back button is clicked
    Then the user is prompted that there are unsaved changes
    When the close is confirmed
    Then the title of the app is " ThesorTeX  - Projectview"
    And the title of the main area is "test"
    When the back button is clicked
    Then the title of the app is " ThesorTeX "
    And the title of the main area is "Welcome to ThesorTeX!"

  Scenario: Delete category
    Given the url "/" was opened
    And the project " test " is opened
    And the category "c1" is opened
    And the delete button in the editor is clicked
    Then the user is asked to confirm the deletion of the category
    When the deletion is confirmed
    Then the editor-page is closed
    And the title of the app is " ThesorTeX  - Projectview"
    And the title of the main area is "test"
    And following categories are displayed
      | name                   |
      | CitaviArticle          |
      | CitaviArticleDoi       |
      | CitaviBook             |
      | CitaviBookDoi          |
      | CitaviBooklet          |
      | CitaviInBook           |
      | CitaviInBookDoi        |
      | CitaviInCollection     |
      | CitaviInCollectionDoi  |
      | CitaviInProceedings    |
      | CitaviInProceedingsDoi |
      | aufsatz                |
