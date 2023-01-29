Feature: Start page

  Scenario: Initial display
    Given the url "/" was opened
    Then the sidebar is empty
    And the sidebar is disabled
    And the sidebar is closed
    And the title of the app is " ThesorTeX "
    And the title of the main area is "Welcome to ThesorTeX!"
    And following projects are displayed
      | project |
      | example |

  Scenario: Navigating to a project
    Given the url "/" was opened
    When the project " example " is opened
    Then the title of the app is " ThesorTeX  - Projectview"

  Scenario: Create new project
    Given the url "/" was opened
    When a new project is added
    Then the dialog for project creation is shown
    When the name "test" is entered into the projectname field
    And the create project button is clicked
    Then following projects are displayed
      | project |
      | example |
      | test    |
  