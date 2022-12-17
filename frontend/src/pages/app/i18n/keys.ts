
export const i18nKeys = {
  Settings: {
    Port: 'Settings.Port',
    PortHint: 'Settings.PortHint',
    ProjectFolder: 'Settings.ProjectFolder',
    ProjectFolderHint: 'Settings.ProjectFolderHint',
  },
  Sidebar: {
    Projects: 'Sidebar.Projects'
  },
  Overview: {
    Welcome: 'Overview.Welcome',
    Project: "Overview.Project",
    Created: 'Overview.Created',
    LastModified: 'Overview.LastModified',
    NumberOfEntries: 'Overview.NumberOfEntries',
    CreateProject: 'Overview.CreateProject',
    ProjectName: 'Overview.ProjectName',
    DeleteProject: 'Overview.DeleteProject',
    SureDeleteProject: 'Overview.SureDeleteProject',
  },
  Project: {
    Title: 'Project.Title',
    Entry: {
      Heading: 'Project.Entry.Heading',
      Key: 'Project.Entry.Key',
      Entry: 'Project.Entry.Entry'
    },
    Category: {
      Heading: 'Project.Category.Heading',
      Category: 'Project.Category.Category',
      Example: 'Project.Category.Example'
    }
  },
  EntryEditor: {
    Title: 'EntryEditor.Title',
    General: {
      Title: 'EntryEditor.General.Title',
      Key: 'EntryEditor.General.Key',
      Category: 'EntryEditor.General.Category',
    },
    Fields: {
      Title: 'EntryEditor.Fields.Title'
    }
  },
  CategoryEditor: {
    Title: 'CategoryEditor.Title',
    General: {
      Title: 'CategoryEditor.General.Title',
      Name: 'CategoryEditor.General.Name',
      CitaviCategory: 'CategoryEditor.General.CitaviCategory',
      CitaviFilter: 'CategoryEditor.General.CitaviFilter',
    },
    Bib: {
      Title: 'CategoryEditor.Bib.Title',
      Field: 'CategoryEditor.Bib.Field',
      Italic: 'CategoryEditor.Bib.Italic',
      Prefix: 'CategoryEditor.Bib.Prefix',
      Suffix: 'CategoryEditor.Bib.Suffix',
      Formatted: 'CategoryEditor.Bib.Formatted',
      CitaviAttributes: 'CategoryEditor.Bib.CitaviAttributes',
    },
    Cite: {
      Title: 'CategoryEditor.Cite.Title'
    }
  },
  Common: {
    Abort: 'Common.Abort',
    Create: 'Common.Create',
    Delete: 'Common.Delete',
    Close : 'Common.Close',
    Attribute: 'Common.Attribute',
    Value: 'Common.Value',
    Field: 'Common.Field',
    Settings: 'Common.Settings',
    Save: 'Common.Save'
  },
  Rules: {
    ProjectAlreadyExists: 'Rules.ProjectAlreadyExists',
    NoSpaces: 'Rules.NoSpaces'
  },
  Errors: {
    Title: 'Errors.Title',
    ErrorHint: 'Errors.ErrorHint',
    CreateIssue: 'Errors.CreateIssue',
    ErrorCreating: 'Errors.ErrorCreating',
    ErrorDeleting: 'Errors.ErrorDeleting',
    ErrorSaving: 'Errors.ErrorSaving'
  },
  Success: {
    CreateProject: 'Success.CreateProject',
    DeleteProject: 'Success.DeleteProject',
    SaveCategory: 'Success.SaveCategory',
    SaveEntry: 'Success.SaveEntry',
    SaveConfig: 'Success.SaveConfig'
  }
}

export type I18nKeys = typeof i18nKeys;
