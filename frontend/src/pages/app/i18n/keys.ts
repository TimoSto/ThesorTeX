
export const i18nKeys = {
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
