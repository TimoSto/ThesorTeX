Feature: Deleting project and overview over entries and categories

  Scenario: Displaying entries and categories
    Given the url "/" was opened
    When the project " example " is opened
    Then the title of the app is " ThesorTeX  - Projektansicht"
    And the title of the main area is "example"
    Then following entries are displayed
      | key       |
      | testEntry |
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

  Scenario: Deleting project
    Given the url "/" was opened
    And the project " example " is opened
    When the project is deleted
    Then the user is asked to confirm the deletion of the project
    When the deletion is confirmed
    Then the project-page is closed
    And the title of the app is " ThesorTeX "
    And the title of the main area is "Willkommen bei ThesorTeX!"
