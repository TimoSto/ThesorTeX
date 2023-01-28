Feature: Start page

  Scenario: Initial display
    Given the url "/" was opened
    Then the sidebar is empty
    And the sidebar is disabled
    And the sidebar is closed
  