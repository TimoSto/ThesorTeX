import {I18nKeys} from "./keys";

export const german: I18nKeys = {
  Settings: {
    Port: 'Port',
    PortHint: 'Die Anwendung wird auf diesem Port deines Computers erreichbar sein. Der Standard ist http://localhost:8448.',
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
    Rename: 'Projekt umbenennen',
    Import: 'Ziehe eine .bib-Datei hier hinein oder klicke hier, um sie auszuwählen',
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
    },
    Delete: {
      Title: 'Eintrag löschen',
      Message: 'Möchstest du den Eintrag mit dem Schlüssel {key} wirklich löschen?'
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
    },
    Delete: {
      Title: 'Kategorie löschen',
      Message: 'Willst du die Kategorie {name} wirklich löschen?'
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
    Save: 'Speichern',
    Rename: 'Umbenennen',
    Continue: 'Fortfahren'
  },
  Rules: {
    ProjectAlreadyExists: 'Ein Projekt mit diesem Namen existiert bereits',
    NoSpaces: 'Leerzeichen nicht erlaubt',
    NotEmpty: 'Kein leerer Wert erlaubt',
    KeyAlreadyExists: 'Dieser Schlüssel ist bereits vergeben',
    CategoryNameAlreadyExists: 'Dieser Name ist bereits vergeben',
    FieldAlreadyExists: 'Ein Feld mit diesem Namen existiert bereits',
    NoSpecialChars: 'Keine Sonderzeichen erlaubt',
    InvalidPortNumber: 'Der Port ist ungültig',
  },
  Errors: {
    Title: 'Es ist ein Fehler aufgetreten',
    ErrorHint: 'Falls du glaubst, dass dieser Fehler ein Bug ist, leg dazu gern ein Issue in {github} an.',
    CreateIssue: 'Issue anlegen',
    ErrorCreating: 'Das Projekt konnte nicht erstellt werden',
    ErrorDeleting: 'Das Projekt konnte nicht gelöscht werden',
    ErrorSaving: 'Beim Speichern ist ein Fehler aufgetreten.',
    ErrorUpload: 'Beim Hinzufügen der Einträge ist ein Fehler aufgetreten'
  },
  Success: {
    CreateProject: 'Projekt erfolgreich erstellt',
    DeleteProject: 'Projekt erfolgreich gelöscht',
    SaveCategory: 'Kategorie erfolgreich gespeichert',
    DeleteCategory: 'Kategorie erfolgreich gelöscht',
    SaveEntry: 'Eintrag erfolgreich gespeichert',
    DeleteEntry: 'Eintrag erfolgreich gelöscht',
    SaveConfig: 'Einstellungen erfolgreich gespeichert',
    UploadEntries: 'Einträge erfolgreich hinzugefügt'
  },
  UnsafeClose: {
    Title: 'Warnung',
    Message: 'Es liegen ungespeicherte Änderungen vor. Wenn du fortfährst, werden diese verloren gehen.'
  }
}
