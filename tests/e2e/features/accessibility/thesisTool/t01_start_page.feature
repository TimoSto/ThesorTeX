Feature: Accessibility

  Scenario: Start page - AXE core
    Given the url "/" was opened
    Then the wcag guidelines are met

  Scenario: Going to the documentation (keyboard ctrl)
    Given the url "/" was opened
    When the TAB key is pressed 1 times
    Then the button with the title "To the documentation" is focussed
    When the enter key is pressed
    Then a second tab was opened with "https://thesortex.com/#/tutorials?target=ThesisTool"

  Scenario: A11y dialog
    Given the url "/" was opened
    When the TAB key is pressed 2 times
    Then the button with the title "Open settings affecting accessibility" is focussed
    When the enter key is pressed
    Then the a11y dialog is shown
    When the TAB key is pressed 1 times
    Then the checkbox with the aria label "Activate focus borders" is focussed
    When the TAB key is pressed 1 times
    Then the button with the text "Close" is focussed
    When the enter key is pressed
    Then no dialog is shown

  Scenario: Config dialog - AXE core
    Given the url "/" was opened
    When a new project is added
    Then the dialog for project creation is shown
    And the wcag guidelines are not yet met

  Scenario: New project dialog - AXE core
    Given the url "/" was opened
    And a new project is added
    Then the wcag guidelines are not yet met

  Scenario: New Project dialog - keyboard control
    Given the url "/" was opened
    When the TAB key is pressed 4 times
    Then the button with the title "Create new project" is focussed
    When the enter key is pressed
    Then the dialog for project creation is shown
    When the TAB key is pressed 1 times
    And "a11yproject1" is typed on the keyboard
    And the TAB key is pressed 1 times
    Then the button with the text "Abort" is focussed
    And the TAB key is pressed 1 times
    Then the button with the text "Create" is focussed
    When the enter key is pressed
    # Then no dialog is shown TODO: why fails this??
