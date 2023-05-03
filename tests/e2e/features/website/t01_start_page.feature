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

  Scenario: Jumping to area 3
    Given the url "/" was opened
    When the list item with the text "Tool for bibliography management" is clicked
    Then area 3 is shown in full height
    And the title of area 3 is "Tool for bibliography management"

  Scenario: Jumping to area 4
    Given the url "/" was opened
    When the list item with the text "Template for a curriculum vitae" is clicked
    Then area 4 is shown in full height
    And the title of area 4 is "Template for a curriculum vitae"