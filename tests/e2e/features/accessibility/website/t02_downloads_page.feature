Feature: Accessibility

  Scenario: Downloads page
    Given the url "/#/downloads" was opened
    Then the wcag guidelines are not yet met

  Scenario: Navigating through the page using tab (Overview button)
    Given the url "/#/downloads" was opened
    When the "Tab" key is pressed
    Then the button with the text "Overview" is focussed

  Scenario: Navigating through the page using tab (Tutorials button)
    Given the url "/#/downloads" was opened
    When the "Tab" key is pressed 2 times
    Then the button with the text "Tutorials" is focussed

  Scenario: Navigating through area 1 card 1
    Given the url "/#/downloads" was opened
    When the "Tab" key is pressed 3 times
    Then the button with the text "Learn more" in area 1 and card 1 is focussed
    When the "Tab" key is pressed
    Then the button with the text "Download" in area 1 and card 1 is focussed

  Scenario: Jump from area 1 card 1 to tutorials
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 3 times
    When the "Enter" key is pressed
    Then the url is "/#/tutorials?target=thesisTemplate"
    And area 2 is shown in full height

  Scenario: Jump from area 1 card 1 to the downloads table
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 4 times
    When the "Enter" key is pressed
    Then the url is "/#/downloads"
    And area 2 is shown in full height

  Scenario: Navigating through area 1 card 2
    Given the url "/#/downloads" was opened
    When the "Tab" key is pressed 5 times
    Then the button with the text "Learn more" in area 1 and card 2 is focussed
    When the "Tab" key is pressed
    Then the button with the text "Download" in area 1 and card 2 is focussed

  Scenario: Jump from area 1 card 2 to tutorials
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 5 times
    When the "Enter" key is pressed
    Then the url is "/#/tutorials?target=thesisTool"
    And area 3 is shown in full height

  Scenario: Jump from area 1 card 2 to the downloads table
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 6 times
    When the "Enter" key is pressed
    Then the url is "/#/downloads"
    And area 3 is shown in full height

  Scenario: Navigating through area 1 card 3
    Given the url "/#/downloads" was opened
    When the "Tab" key is pressed 7 times
    Then the button with the text "Learn more" in area 1 and card 3 is focussed
    When the "Tab" key is pressed
    Then the button with the text "Download" in area 1 and card 3 is focussed

  Scenario: Jump from area 1 card 1 to tutorials
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 7 times
    When the "Enter" key is pressed
    Then the url is "/#/tutorials?target=cvTemplate"
    And area 4 is shown in full height

  Scenario: Jump from area 1 card 1 to the downloads table
    Given the url "/#/downloads" was opened
    And the "Tab" key is pressed 8 times
    When the "Enter" key is pressed
    Then the url is "/#/downloads"
    And area 4 is shown in full height

    # TODO: test tables
