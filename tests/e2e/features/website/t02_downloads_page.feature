Feature: Lists of downloads

  Scenario: displaying the first area
    Given the url "/#/downloads" was opened
    Then the url is "/#/downloads"
    And area 1 is shown in full height
    And the title of area 1 is not existent
    And three cards are shown for the products

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