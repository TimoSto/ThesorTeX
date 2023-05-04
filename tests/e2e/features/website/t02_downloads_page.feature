Feature: Lists of downloads

  Scenario: displaying the first area
    Given the url "/#/downloads" was opened
    Then the url is "/#/downloads"
    And area 1 is shown in full height
    And the title of area 1 is not existent
    And three cards are shown for the products