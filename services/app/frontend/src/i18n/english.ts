import {I18nKeys} from "./keys";

export const english: I18nKeys = {
    Common: {
        Abort: "Abort",
        Create: "Create",
        Close: "Close",
        Error: "Something went wrong",
        ContactBug: "If you think, this is a bug, feel free to create an issue under {link}",
        Error404: "The requested resource was not found.",
        Error400: "The requested action is not allowed.",
        Error500: "An error occurred during the processing.",
        Delete: "Delete",
        Attribute: "Attribute",
        Value: "Value",
    },
    Rules: {
        NoSpaces: "No spaces allowed",
        NoEmpty: "No empty value allowed",
        ProjectAlreadyExists: "Projectname already exists"
    },
    App: {},
    MainPage: {
        Welcome: "Welcome to ThesorTeX!",
        Project: "Project",
        ProjectCreated: "Created",
        ProjectLastModified: "Last modified",
        ProjectNumberOfEntries: "Entries",
        CreateProject: "Create new project",
        ProjectName: "Name of the project",
        SuccessCreation: "The project NAME was successfully created",
        ErrorCreation: "The project could not be created."
    },
    ProjectPage: {
        Title: " - Projectview",
        Entries: "Entries in bibliography",
        EntryKey: "Key",
        Entry: "Entry",
        EntryCategory: "Category",
        Categories: "Categories for entries",
        CategoryName: "Name",
        CategoryExample: "Example entry",
        ErrorReadingData: "An error occurred reading the project data",
        DeleteTitle: "Delete project",
        DeleteMessage: "Do you really want to delete the project {project}?",
        SuccessDelete: "The project PROJECTNAME was successfully deleted",
        ErrorDelete: "The project could not be deleted"
    },
    CategoryEditor: {
        Title: " - Category editor",
        General: "General",
        Name: "Name",
        BibEntry: "Bibliography entry",
        Cites: "Cites",
        Prefix: "Prefix",
        Suffix: "Suffix",
        Italic: "Italic",
        Preformatted: "Preformatted"
    }
}
