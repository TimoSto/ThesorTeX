import {I18nKeys} from "./keys";

export const german: I18nKeys = {
    Common: {
        Abort: "Abbrechen",
        Create: "Erstellen",
        Close: "Schließen",
        Error: "Etwas ist schief gelaufen",
        ContactBug: "Wenn du denkst, das ist ein Bug, lege gern ein Issue unter {link} an",
        Error404: "Die angefragte Ressource wurde nicht gefunden.",
        Error400: "Die Aktion ist unzulässig.",
        Error500: "Bei der Verarbeitung ist ein Fehler aufgetreten.",
        Delete: "Löschen",
        Attribute: "Attribut",
        Value: "Wert",
    },
    Rules: {
        NoSpaces: "Keine Leerzeichen erlaubt",
        NoEmpty: "Kein leerer Wert erlaubt",
        ProjectAlreadyExists: "Projektname bereits vergeben",
        NameAlreadyExists: "Name bereits vergeben"
    },
    App: {},
    MainPage: {
        Welcome: "Willkommen bei ThesorTeX!",
        Project: "Projekt",
        ProjectCreated: "Erstellt",
        ProjectLastModified: "Zuletzt bearbeitet",
        ProjectNumberOfEntries: "Einträge",
        CreateProject: "Neues Projekt anlegen",
        ProjectName: "Name des Projektes",
        SuccessCreation: "Das Projekt NAME wurde erfolgreich erstellt",
        ErrorCreation: "Das Projekt konnte nicht erstellt werden."
    },
    ProjectPage: {
        Title: " - Projektansicht",
        Entries: "Einträge im Literaturverzeichnis",
        EntryKey: "Schlüssel",
        Entry: "Eintrag",
        EntryCategory: "Kategorie",
        Categories: "Kategorien für Einträge",
        CategoryName: "Name",
        CategoryExample: "Beispiel für Eintrag",
        ErrorReadingData: "Beim Abrufen der Projektdaten ist ein Fehler aufgetreten",
        DeleteTitle: "Projekt löschen",
        DeleteMessage: "Willst du das Projekt {project} wirklich löschen?",
        SuccessDelete: "Das Projekt PROJECTNAME wurde erfolgreich gelöscht",
        ErrorDelete: "Das Projekt konnte nicht gelöscht werden"
    },
    CategoryEditor: {
        Title: " - Kategorie bearbeiten",
        General: "Allgemein",
        Name: "Name",
        CitaviCategory: "Citavi-Kategorie",
        CitaviFilter: "Citavi-Filter",
        BibEntry: "Literatureintrag",
        Cites: "Zitate",
        Prefix: "Prefix",
        Suffix: "Suffix",
        Italic: "Kursiv",
        Preformatted: "Vorformatiert",
        CitaviMapping: "Citavi-Attribute",
        SuccessSave: "Kategorie erfolgreich gespeichert",
        ErrorSave: "Die Kategorie konnte nicht gespeichert werden."
    }
};
