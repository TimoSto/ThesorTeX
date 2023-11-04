Feature: Accessibility

  Scenario: Start page - AXE core
    Given the url "/" was opened
    Then the wcag guidelines are met

  Scenario: Going to the documentation (keyboard ctrl)
    Given the url "/" was opened
    When the TAB key is pressed 2 times
    Then the button with the title "To the documentation" is focussed

  Scenario: Config dialog - AXE core
    Given the url "/" was opened
    When a new project is added
    Then the dialog for project creation is shown
    And the wcag guidelines are not yet met

  Scenario: New project dialog - AXE core
    Given the url "/" was opened
    And a new project is added
    Then the wcag guidelines are not yet met
