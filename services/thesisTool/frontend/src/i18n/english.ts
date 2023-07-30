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
        General: "General",
        Continue: "Continue",
        Save: "Save",
        Back: "Back"
    },
    Rules: {
        NoSpaces: "No spaces allowed",
        NoEmpty: "No empty value allowed",
        ProjectAlreadyExists: "Projectname already exists",
        NameAlreadyExists: "Name is already used",
        KeyAlreadyExists: "This key already exists"
    },
    App: {
        OpenSidebar: "Expand",
        CloseSidebar: "Minimize",
        ProjectsList: "List of projects",
        ProjectsListShort: "Projects",
        GoToDocs: "To the documentation",
    },
    UpdateDialog: {
        Title: "A new version is available: {version}",
        Content: "Go to {link} and download the latest version. Close the application (STG + C or close terminal window), unpack the ZIP file and replace the programme file. Now you can start the application again with a double click.",
        ToDownloads: "To the downloads",
        Ignore: "Ignore"
    },
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
        ErrorDelete: "The project could not be deleted",
        UnknownCategory: "Unknown category",
        DragNDrop: "Upload citavi-export",
        WrongFileExt: "The file ending has to be .bib.",
        UploadTitle: "Add entries",
        UploadEntries: "Following entries will be added to your project and override existing entries with the same key:",
        UploadUnknowns: "Following entries could not be assigned to a category:",
        Add: "Add",
        UploadNoEntriesFound: "Es wurden keine Einträge gefunden",
        UploadSuccess: "The entries were successfully uploaded",
        UploadError: "The entries could not be saved."
    },
    CategoryEditor: {
        Title: " - Category editor",
        NewCategory: "Create new category",
        Name: "Name",
        CitaviCategory: "Citavi category",
        CitaviFilter: "Citavi filter",
        BibEntry: "Bibliography entry",
        Cites: "Cites",
        Prefix: "Prefix",
        Suffix: "Suffix",
        Italic: "Italic",
        Preformatted: "Preformatted",
        CitaviMapping: "Citavi attributes",
        SuccessSave: "Category successfully saved",
        ErrorSave: "The category could not be saved.",
        DeleteTitle: "Delete category",
        DeleteMessage: "Do you really want to delete the category {name}?",
        SuccessDelete: "Category successfully deleted",
        ErrorDelete: "The category could not be deleted.",
        StillUsed: "The category is still being used for following entries:",
        StillUsedSuffix: "If you delete it, these entries are not displayed correctly anymore. But you can select different categories for them."
    },
    EntryEditor: {
        Title: " - Edit entry",
        NewEntry: "Create new entry",
        Key: "Key",
        Category: "Category",
        Fields: "Fields",
        Field: "Field",
        SuccessSave: "Entry successfully saved",
        ErrorSave: "The entry could not be saved.",
        DeleteTitle: "Delete entry",
        DeleteMessage: "Do you really want to delete the entry {key}?",
        SuccessDelete: "Entry successfully deleted",
        ErrorDelete: "The entry could not be deleted.",
    },
    UnsavedChanges: {
        Title: "There are unsaved changes",
        Message: "If you continue, these changes will be discarded."
    },
    Config: {
        Title: "Configuration",
        OpenConfig: "Open configuration",
        Port: "Port",
        Dir: "Directory for projects",
        Open: "Open browser on startup",
        PortNumInvalid: "The portnumber is invalid",
        PortReserved: "This port is reserved (1024 - 49151 available)"
    }
};
