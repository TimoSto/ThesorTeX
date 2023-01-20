export const i18nKeys = {
    Common: {
        Abort: "Common.Abort",
        Create: "Common.Create",
        Close: "Common.Close",
        Error: "Common.Error",
        ContactBug: "Common.ContactBug",
        Error404: "Common.Error404",
        Error400: "Common.Error400",
        Error500: "Common.Error500",
        Delete: "Common.Delete"
    },
    Rules: {
        NoSpaces: "Rules.NoSpaces",
        NoEmpty: "Rules.NoEmpty",
        ProjectAlreadyExists: "Rules.ProjectAlreadyExists"
    },
    App: {},
    MainPage: {
        Welcome: "MainPage.Welcome",
        Project: "MainPage.Project",
        ProjectCreated: "MainPage.ProjectCreated",
        ProjectLastModified: "MainPage.ProjectLastModified",
        ProjectNumberOfEntries: "MainPage.ProjectNumberOfEntries",
        CreateProject: "MainPage.CreateProject",
        ProjectName: "MainPage.ProjectName",
        SuccessCreation: "MainPage.SuccessCreation",
        ErrorCreation: "MainPage.ErrorCreation"
    },
    ProjectPage: {
        Title: "ProjectPage.Title",
        Entries: "ProjectPage.Entries",
        EntryKey: "ProjectPage.EntryKey",
        Entry: "ProjectPage.Entry",
        EntryCategory: "ProjectPage.EntryCategory",
        Categories: "ProjectPage.Categories",
        CategoryName: "ProjectPage.CategoryName",
        CategoryExample: "ProjectPage.CategoryExample",
        ErrorReadingData: "ProjectPage.ErrorReadingData",
        DeleteTitle: "ProjectPage.DeleteTitle",
        DeleteMessage: "ProjectPage.DeleteMessage",
        SuccessDelete: "ProjectPage.SuccessDelete",
        ErrorDelete: "ProjectPage.ErrorDelete"
    },
    CategoryEditor: {
        Title: "CategoryEditor.Title"
    }
}

export type I18nKeys = typeof i18nKeys;
