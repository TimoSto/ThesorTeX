Feature: The start page of the website contains general information and links

  Scenario: displaying the first area
    Given the url "/" was opened
    Then the url is "/#/"
    And area 1 is shown in full height
    And the title of area 1 is "Use LaTeX more comfortably"

  Scenario: switching to downloads
    Given the url "/" was opened
    When the button with the text "Downloads" is clicked
    Then the url is "/#/downloads"
    And area 1 is shown in full height
    And the title of area 1 is not existent

  Scenario: switching to tutorials
    Given the url "/" was opened
    When the button with the text "Tutorials" is clicked
    Then the url is "/#/tutorials"
    And area 1 is shown in full height
    And the title of area 1 is "How can I use the templates and tools?"

  Scenario: Jumping to area 2
    Given the url "/" was opened
    When the list item with the text "Template for academic papers" is clicked
    Then area 2 is shown in full height
    And the title of area 2 is "Template for academic papers"