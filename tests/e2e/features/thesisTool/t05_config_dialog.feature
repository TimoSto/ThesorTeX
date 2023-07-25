Feature: The config dialog can be used to configure the application

  Scenario: Opening the config dialog
    Given the url "/" was opened
    When the config button is clicked
    Then the config dialog is shown
    And the save button in the dialog is disabled

# TODO: fix backdrop click
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

  Scenario: Setting a port
    Given the url "/" was opened
    And the config button is clicked
    When "4444" is entered as custom port
    Then the save button in the dialog is enabled
    When the save button in the dialog is clicked
    Then the save button in the dialog is disabled
    When the browser is reloaded
    And the config button is clicked
    Then the value for the custom port is "4444"

  Scenario: Setting a projects dir
    Given the url "/" was opened
    And the config button is clicked
    When "/myprojects" is entered as custom projects path
    Then the save button in the dialog is enabled
    When the save button in the dialog is clicked
    Then the save button in the dialog is disabled
    When the browser is reloaded
    And the config button is clicked
    Then the value for the custom projects path is "/myprojects"

  Scenario: Setting auto open
    Given the url "/" was opened
    And the config button is clicked
    When auto open is activated
    Then the save button in the dialog is enabled
    When the save button in the dialog is clicked
    Then the save button in the dialog is disabled
    When the browser is reloaded
    And the config button is clicked
    Then auto open is enabled
