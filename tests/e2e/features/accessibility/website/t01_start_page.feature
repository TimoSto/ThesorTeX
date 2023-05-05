Feature: Accessibility

  Scenario: Start page
    Given the url "/" was opened
    Then the wcag guidelines are not yet met

  Scenario: Navigating through the page using tab (Download button)
    Given the url "/" was opened
    When the "Tab" key is pressed
    Then the button with the text "Downloads" is focussed

  Scenario: Navigating through the page using tab (Tutorials button)
    Given the url "/" was opened
    When the "Tab" key is pressed 2 times
    Then the button with the text "Tutorials" is focussed

  Scenario: Navigating through the page using tab (list elements in first area)
    Given the url "/" was opened
    When the "Tab" key is pressed 3 times
    Then the list item with the text "Template for academic papers" is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Tool for bibliography management" is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Template for a curriculum vitae" is focussed

  Scenario: Navigating through the page using tab (list elements in second area)
    Given the url "/" was opened
    When the "Tab" key is pressed 6 times
    Then area 2 is shown in full height
    Then the list item with the text "To the downloads" in area 2 is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Download an example" is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Tutorial" in area 2 is focussed

  Scenario: Navigating through the page using tab (list elements in third area)
    Given the url "/" was opened
    When the "Tab" key is pressed 9 times
    Then area 3 is shown in full height slightly scrolled off
    Then the list item with the text "To the downloads" in area 3 is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Tutorial" in area 3 is focussed

  Scenario: Navigating through the page using tab (list elements in forth area)
    Given the url "/" was opened
    When the "Tab" key is pressed 11 times
    Then area 4 is shown in full height slightly scrolled off
    Then the list item with the text "To the downloads" in area 4 is focussed
    When the "Tab" key is pressed
    Then the list item with the text "Tutorial" in area 4 is focussed

  Scenario: Navigating through the page using tab (list elements in fifth area)
    Given the url "/" was opened
    When the "Tab" key is pressed 13 times
    Then area 5 is shown in full height slightly scrolled off
    Then the first link with the text "here" is focussed
    When the "Tab" key is pressed
    Then the second link with the text "here" is focussed