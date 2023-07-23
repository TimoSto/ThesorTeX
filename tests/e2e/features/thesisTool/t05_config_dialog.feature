Feature: The config dialog can be used to configure the application

  Scenario: Opening the config dialog
    Given the url "/" was opened
    When the config button is clicked
    Then the config dialog is shown
    And the save button in the dialog is disabled

#  Scenario: Closing the config dialog (backdrop)
#    Given the url "/" was opened
#    And the config button is clicked
#    Then the config dialog is shown
#    When the backdrop of the dialog is clicked
#    Then no dialog is shown

  Scenario: Closing the config dialog (button)
    Given the url "/" was opened
    And the config button is clicked
    Then the config dialog is shown
    When the close button of the dialog is clicked
    Then no dialog is shown