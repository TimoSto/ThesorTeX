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