Feature: Start page

  Scenario: Initial display
    Given the url "/" was opened
    Then the sidebar is empty
    And the sidebar is disabled
    And the sidebar is closed
    And the title of the app is " ThesorTeX "
    And the title of the main area is "Willkommen bei ThesorTeX!"
    And following projects are displayed
      | project |
      | example |

  Scenario: Navigating to a project
    Given the url "/" was opened
    When the project " example " is opened
    Then the title of the app is " ThesorTeX  - Projektansicht"
  