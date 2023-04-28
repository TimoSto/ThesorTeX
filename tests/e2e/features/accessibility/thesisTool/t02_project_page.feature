Feature: Deleting project and overview over entries and categories

  Scenario: Start page
    Given the url "/" was opened
    When the project " example " is opened
    Then the wcag guidelines are not yet met