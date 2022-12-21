Feature: Random
  A random feature using some Playwright stuff
  Scenario: Govuk accessibility statement link
    Given I view 'www.gov.uk'
    When I click 'Accessibility statement'
    Then I expect to be on the accessibility page
