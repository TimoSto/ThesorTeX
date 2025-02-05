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
        General: "Allgemein",
        Continue: "Fortfahren",
        Save: "Speichern",
        Back: "Zurück"
    },
    Rules: {
        NoSpaces: "Keine Leerzeichen erlaubt",
        NoEmpty: "Kein leerer Wert erlaubt",
        ProjectAlreadyExists: "Projektname bereits vergeben",
        NameAlreadyExists: "Name bereits vergeben",
        KeyAlreadyExists: "Schlüssel bereits vergeben"
    },
    App: {
        OpenSidebar: "Erweitern",
        CloseSidebar: "Minimieren",
        ProjectsList: "Liste der Projekte",
        ProjectsListShort: "Projekte",
        GoToDocs: "Zur Dokumentation"
    },
    UpdateDialog: {
        Title: "Eine neue Version is verfügbar: {version}",
        Content: "Gehe zu {link} und lade die aktuellste Version herunter. Beende die Applikation (STG + C oder Terminalfenster schließen), entpacke die ZIP-Datei und ersetze die Programm-Datei. Nun kannst du die Applikation durch einen Doppelklick wieder starten.",
        ToDownloads: "Zu den Downloads",
        Ignore: "Ignorieren"
    },
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
        ErrorDelete: "Das Projekt konnte nicht gelöscht werden",
        UnknownCategory: "Unbekannte Kategorie",
        DragNDrop: "Citavi-Export hochladen",
        WrongFileExt: "Die Dateiendung muss .bib sein.",
        UploadTitle: "Einträge hinzufügen",
        UploadEntries: "Folgende Einträge werden hinzugefügt und werden bestehende Einsträge mit demselben Schlüssel überschreiben:",
        UploadUnknowns: "Folgende Einträge konnten keiner Kategorie zugeordnet werden:",
        Add: "Hinzufügen",
        UploadNoEntriesFound: "Es wurden keine Einträge gefunden.",
        UploadSuccess: "Die Einträge wurden erfolgreich hochgeladen",
        UploadError: "Die Einträge konnten nicht gespeichert werden."
    },
    CategoryEditor: {
        Title: " - Kategorie bearbeiten",
        NewCategory: "Neue Kategorie erstellen",
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
        ErrorSave: "Die Kategorie konnte nicht gespeichert werden.",
        DeleteTitle: "Kategorie löschen",
        DeleteMessage: "Möchtest du die Kategorie {name} wirklich löschen?",
        SuccessDelete: "Kategorie erfolgreich gelöscht",
        ErrorDelete: "Die Kategorie konnte nicht gelöscht werden.",
        StillUsed: "Die Kategorie wird noch für folgenden Einträgen verwendet:",
        StillUsedSuffix: "Wenn du die Kategorie löscht, werden diese Einträge nicht mehr korrekt dargestellt. Du kannst dann aber andere Kategorien für sie auswählen."
    },
    EntryEditor: {
        Title: " - Eintrag bearbeiten",
        NewEntry: "Neuen Eintrag erstellen",
        Key: "Schlüssel",
        Category: "Kategorie",
        Fields: "Felder",
        Field: "Feld",
        SuccessSave: "Eintrag erfolgreich gespeichert",
        ErrorSave: "Der Eintrag konnte nicht gespeichert werden.",
        DeleteTitle: "Eintrag löschen",
        DeleteMessage: "Möchtest du den Eintrag {key} wirklich löschen?",
        SuccessDelete: "Eintrag erfolgreich gelöscht",
        ErrorDelete: "Der Eintrag konnte nicht gelöscht werden"
    },
    UnsavedChanges: {
        Title: "Es liegen ungespeicherte Änderungen vor",
        Message: "Wenn du fortfährst, werden diese Änderungen verworfen."
    },
    Config: {
        Title: "Konfiguration",
        OpenConfig: "Konfiguration öffnen",
        Port: "Port",
        Dir: "Verzeichnis für Projekte",
        Open: "Beim Starten Browserfenster öffnen",
        PortNumInvalid: "Die Portnummer is invalid",
        PortReserved: "Dieser Port ist reserviert (1024 - 49151 verfügbar)",
        LogLevel: "Log-Level"
    }
};
