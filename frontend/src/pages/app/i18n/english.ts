import {I18nKeys} from "@/pages/app/i18n/keys";

export const english: I18nKeys = {
  Overview: {
    Welcome: "Welcome to ThesorTeX!",
    Project: "Project",
    Created: 'Created',
    LastModified: 'Last edited',
    NumberOfEntries: 'Number of entries',
    CreateProject: 'Create new project',
    ProjectName: 'Projectname',
    DeleteProject: 'Delete project',
    SureDeleteProject: 'Do you really want to delete the project {project} permanently',
  },
  Project: {
    Title: 'Projektansicht',
    Entry: {
      Heading: 'Bibliography entries',
      Key: 'Key',
      Entry: 'Entry'
    },
    Category: {
      Heading: 'Bibliography categories',
      Category: 'Category',
      Example: 'Example'
    }
  },
  EntryEditor: {
    Title: 'Edit entry',
  },
  Common: {
    Abort: 'Abort',
    Create: 'Create',
    Delete: 'Delete',
    Close: 'Close'
  },
  Rules: {
    ProjectAlreadyExists: 'A project with this name already exists',
    NoSpaces: 'No spaces allowed'
  },
  Errors: {
    Title: 'An error occurred',
    ErrorHint: 'If you think this error is a bug, fell free to create an issue in {github}.',
    CreateIssue: 'Create issue',
    ErrorCreating: 'The project could not be created',
    ErrorDeleting: 'The project could not be deleted'
  },
  Success: {
    CreateProject: 'Project successfully created',
    DeleteProject: 'Project successfully deleted'
  }
}
