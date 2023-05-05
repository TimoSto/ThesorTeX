Feature: Accessibility

  Scenario: Tutorials page
    Given the url "/#/tutorials" was opened
    Then the wcag guidelines are not yet met

  Scenario: Navigating through the page using tab (Overview button)
    Given the url "/#/tutorials" was opened
    When the "Tab" key is pressed
    Then the button with the text "Overview" is focussed

  Scenario: Navigating through the page using tab (Downloads button)
    Given the url "/#/tutorials" was opened
    When the "Tab" key is pressed 2 times
    Then the button with the text "Downloads" is focussed

  Scenario: Navigating through the links in the first area
    Given the url "/#/tutorials" was opened
    When the "Tab" key is pressed 3 times
    Then the first link with the text "PDF" is focussed
    When the "Tab" key is pressed
    Then the first link with the text "example project" is focussed

  # TODO: check revealjs presentation

  Scenario: Navigating over tutorial entries
    Given the url "/#/tutorials" was opened
    When the "Tab" key is pressed 6 times
    Then the expansion panel with the title "How is the template structured?" is focussed
    When the "Tab" key is pressed
    Then the expansion panel with the title "What is the best way to start?" is focussed
    When the "Enter" key is pressed
    Then the expansion panel with the title "What is the best way to start?" is focussed
    And the expansion panel with the title "What is the best way to start?" is expanded
    When the "Tab" key is pressed 5 times
    Then the expansion panel with the title "What can I use this tool for?" is focussed
