Feature: Display the config dialog
  The config dialog is globally accessible and contains the config for the app

  Scenario: The config dialog is accessible via mouse clicks
    Given the page "http://localhost:8448/app/" was opened
    When I click the tree dot button
    Then I expect to see a menu in the top right corner