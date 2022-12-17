import {I18nKeys} from "@/pages/app/i18n/keys";

export const german: I18nKeys = {
  Settings: {
    Port: 'Port',
    PortHint: 'Die Anwendung wird auf diesem Port deines Computers erreichbar sein. Der Standard ist http://localhost:8081.',
    ProjectFolder: 'Speicherort für Projekte',
    ProjectFolderHint: 'In diesem Ordner werden deine Projekte gespeichert.',
  },
  Sidebar: {
    Projects: 'Projekte'
  },
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
      Key: 'Schlüssel',
      Category: 'Kategorie'
    },
    Fields: {
      Title: 'Felder'
    }
  },
  CategoryEditor: {
    Title: 'Kategorie bearbeiten',
    General: {
      Title: 'Allgemein',
      Name: 'Name',
      CitaviCategory: 'Citavi-Kategorie',
      CitaviFilter: 'Citavi-Filter'
    },
    Bib: {
      Title: 'Literaturverzeichnis',
      Field: 'Feld',
      Italic: 'kursiv',
      Prefix: 'Prefix',
      Suffix: 'Suffix',
      Formatted: 'Formatiert',
      CitaviAttributes: 'Citavi-Attribute'
    },
    Cite: {
      Title: 'Zitate'
    }
  },
  Common: {
    Abort: 'Abbrechen',
    Create: 'Erstellen',
    Delete: 'Löschen',
    Close: 'Schließen',
    Attribute: 'Attribut',
    Value: 'Wert',
    Field: 'Feld',
    Settings: 'Einstellungen',
    Save: 'Speichern'
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
    ErrorDeleting: 'Das Projekt konnte nicht gelöscht werden',
    ErrorSaving: 'Beim Speichern ist ein Fehler aufgetreten.'
  },
  Success: {
    CreateProject: 'Projekt erfolgreich erstellt',
    DeleteProject: 'Projekt erfolgreich gelöscht',
    SaveCategory: 'Kategorie erfolgreich gespeichert',
    SaveEntry: 'Eintrag erfolgreich gespeichert',
    SaveConfig: 'Einstellungen erfolgreich gespeichert',
  }
}
