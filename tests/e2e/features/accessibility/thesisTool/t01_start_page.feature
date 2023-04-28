Feature: Accessibility

  Scenario: Start page
    Given the url "/" was opened
    Then the wcag guidelines are met

  Scenario: Config dialog
    Given the url "/" was opened
    When a new project is added
    Then the dialog for project creation is shown
    And the wcag guidelines are not yet met

  Scenario: New project dialog
    Given the url "/" was opened
    And a new project is added
    Then the wcag guidelines are not yet met
