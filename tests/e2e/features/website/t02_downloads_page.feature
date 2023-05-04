Feature: Lists of downloads

  Scenario: displaying the first area
    Given the url "/#/downloads" was opened
    Then the url is "/#/downloads"
    And area 1 is shown in full height
    And the title of area 1 is not existent
    And three cards are shown for the products

  Scenario: switching to the start page
    Given the url "/#/downloads" was opened
    When the button with the text "Overview" is clicked
    Then the url is "/#/"
    And area 1 is shown in full height
    And the title of area 1 is "Use LaTeX more comfortably"

  Scenario: switching to tutorials
    Given the url "/#/downloads" was opened
    When the button with the text "Tutorials" is clicked
    Then the url is "/#/tutorials"
    And area 1 is shown in full height
    And the title of area 1 is "How can I use the templates and tools?"

  Scenario: displaying the downloads for the thesis template
    Given the url "/#/downloads" was opened
    When the download button in card 1 is clicked
    Then area 2 is shown in full height
    And the title of area 2 is "Template for academic papers"
    And the following versions are shown in area 2
      | name   | date       |
      | v1.0.0 | 01-01-2022 |

  Scenario: displaying the downloads for the thesis template
    Given the url "/#/downloads" was opened
    When the download button in card 2 is clicked
    Then area 3 is shown in full height
    And the title of area 3 is "Tool for bibliography management"
    And the following versions are shown in area 3
      | name   | date       |
      | v1.1.0 | 01-01-2022 |
      | v1.0.0 | 01-01-2022 |

  Scenario: displaying the downloads for the thesis template
    Given the url "/#/downloads" was opened
    When the download button in card 3 is clicked
    Then area 4 is shown in full height
    And the title of area 4 is "Template for a curriculum vitae"
    And the following versions are shown in area 4
      | name   | date       |
      | v1.0.0 | 01-01-2022 |

  Scenario: Jump to tutorials: thesis Template
    Given the url "/#/downloads" was opened
    And the download button in card 1 is clicked
    When the list item with the text "Learn more" in area 2 is clicked
    Then the url is "/#/tutorials?target=thesisTemplate"
    And area 2 is shown in full height
    And the title of area 2 is "Usage of the template for academic papers"

  Scenario: Jump to tutorials (from card): thesis Template
    Given the url "/#/downloads" was opened
    And the more info button in card 1 is clicked
    Then the url is "/#/tutorials?target=thesisTemplate"
    And area 2 is shown in full height
    And the title of area 2 is "Usage of the template for academic papers"

  Scenario: Jump to tutorials: thesis Tool
    Given the url "/#/downloads" was opened
    And the download button in card 2 is clicked
    When the list item with the text "Learn more" in area 3 is clicked
    Then the url is "/#/tutorials?target=thesisTool"
    And area 3 is shown in full height
    And the title of area 3 is "Usage of the tool for bibliography management"

  Scenario: Jump to tutorials (from card): thesis Tool
    Given the url "/#/downloads" was opened
    And the more info button in card 2 is clicked
    Then the url is "/#/tutorials?target=thesisTool"
    And area 3 is shown in full height
    And the title of area 3 is "Usage of the tool for bibliography management"

  Scenario: Jump to tutorials: cv Template
    Given the url "/#/downloads" was opened
    And the download button in card 3 is clicked
    When the list item with the text "Learn more" in area 4 is clicked
    Then the url is "/#/tutorials?target=cvTemplate"
    And area 4 is shown in full height
    And the title of area 4 is "Usage of the template for a curriculum vitae"

  Scenario: Jump to tutorials (from card): cv Template
    Given the url "/#/downloads" was opened
    And the more info button in card 3 is clicked
    Then the url is "/#/tutorials?target=cvTemplate"
    And area 4 is shown in full height
    And the title of area 4 is "Usage of the template for a curriculum vitae"