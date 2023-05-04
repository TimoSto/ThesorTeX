Feature: Documentation and tutorials

  Scenario: displaying the first area
    Given the url "/#/tutorials" was opened
    Then the url is "/#/tutorials"
    And area 1 is shown in full height
    And the title of area 1 is "How can I use the templates and tools?"

  Scenario: switching to the start page
    Given the url "/#/tutorials" was opened
    When the button with the text "Overview" is clicked
    Then the url is "/#/"
    And area 1 is shown in full height
    And the title of area 1 is "Use LaTeX more comfortably"

  Scenario: switching to downloads
    Given the url "/#/tutorials" was opened
    When the button with the text "Downloads" is clicked
    Then the url is "/#/downloads"
    And area 1 is shown in full height
    And the title of area 1 is not existent