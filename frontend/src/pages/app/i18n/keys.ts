
export const i18nKeys = {
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
    },
    Fields: {
      Title: 'EntryEditor.Fields.Title'
    }
  },
  CategoryEditor: {
    Title: 'CategoryEditor.Title'
  },
  Common: {
    Abort: 'Common.Abort',
    Create: 'Common.Create',
    Delete: 'Common.Delete',
    Close : 'Common.Close'
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
    ErrorDeleting: 'Errors.ErrorDeleting'
  },
  Success: {
    CreateProject: 'Success.CreateProject',
    DeleteProject: 'Success.DeleteProject'
  }
}

export type I18nKeys = typeof i18nKeys;
