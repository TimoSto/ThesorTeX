import {I18nKeys} from "@/pages/app/i18n/keys";

export const german: I18nKeys = {
  Overview: {
    Welcome: "Willkommen bei ThesorTeX!",
    Project: "Projekt",
    Created: 'Erstellt',
    LastModified: 'Zuletzt bearbeitet',
    NumberOfEntries: 'Anzahl Einträge',
    CreateProject: 'Neues Projekt erstellen',
    ProjectName: 'Projektname',
    DeleteProject: 'Projekt löschen',
    SureDeleteProject: 'Willst du das Projekt {project} wirklich unwiderruflich löschen?',
  },
  Project: {
    Title: 'Projektansicht',
    Entry: {
      Heading: 'Literatur-Einträge',
      Key: 'Schlüssel',
      Entry: 'Eintrag'
    },
    Category: {
      Heading: 'Literatur-Kategorien',
      Category: 'Kategorie',
      Example: 'Beispiel'
    }
  },
  EntryEditor: {
    Title: 'Eintrag bearbeiten',
    General: {
      Title: 'Allgemein',
    },
    Fields: {
      Title: 'Felder'
    }
  },
  Common: {
    Abort: 'Abbrechen',
    Create: 'Erstellen',
    Delete: 'Löschen',
    Close: 'Schließen'
  },
  Rules: {
    ProjectAlreadyExists: 'Ein Projekt mit diesem Namen existiert bereits',
    NoSpaces: 'Leerzeichen nicht erlaubt'
  },
  Errors: {
    Title: 'Es ist ein Fehler aufgetreten',
    ErrorHint: 'Falls du glaubst, dass dieser Fehler ein Bug ist, leg dazu gern ein Issue in {github} an.',
    CreateIssue: 'Issue anlegen',
    ErrorCreating: 'Das Projekt konnte nicht erstellt werden',
    ErrorDeleting: 'Das Projekt konnte nicht gelöscht werden'
  },
  Success: {
    CreateProject: 'Projekt erfolgreich erstellt',
    DeleteProject: 'Projekt erfolgreich gelöscht'
  }
}
