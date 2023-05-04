Feature: Accessibility

  Scenario: Start page
    Given the url "/" was opened
    Then the wcag guidelines are not yet met

  Scenario: Downloads page
    Given the url "/#/downloads" was opened
    Then the wcag guidelines are not yet met

  Scenario: Start page
    Given the url "/#/tutorials" was opened
    Then the wcag guidelines are not yet met