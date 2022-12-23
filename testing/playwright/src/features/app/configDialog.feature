Feature: Display the config dialog
  The config dialog is globally accessible and contains the config for the app

  Scenario: The three dot menu is accessible via mouse navigation
    Given the page "http://localhost:8448/app/" was opened
    When I click the tree dot button
    Then I expect to see a menu in the top right corner
    Then I expect the three dot menu to contain the following entries "Settings"

  Scenario: The three dot is accessible via keyboard navigation
    Given the page "http://localhost:8448/app/" was opened
    When I press the tab key 1 times
    And I press the enter key
    Then I expect to see a menu in the top right corner
    Then I expect the three dot menu to contain the following entries "Settings"
    When I press the tab key 1 times
    # Then I expect the first menu element to be focussed