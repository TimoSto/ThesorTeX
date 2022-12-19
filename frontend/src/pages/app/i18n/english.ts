import {I18nKeys} from "@/pages/app/i18n/keys";

export const english: I18nKeys = {
  Settings: {
    Port: 'Port',
    PortHint: 'The application will be available on this port of your computer. The deafult is http://localhost:8081.',
    ProjectFolder: 'Storage location for projects',
    ProjectFolderHint: 'Your projects will be saved in this directory.',
  },
  Sidebar: {
    Projects: 'Projects'
  },
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
    Title: 'Project view',
    Rename: 'Rename project',
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
    General: {
      Title: 'General',
      Key: 'Key',
      Category: 'Category',
    },
    Fields: {
      Title: 'Fields'
    },
    Delete: {
      Title: 'Delete entry',
      Message: 'Do you really want to delete the entry with the key {key}?'
    }
  },
  CategoryEditor: {
    Title: 'Edit category',
    General: {
      Title: 'General',
      Name: 'Name',
      CitaviCategory: 'Citavi category',
      CitaviFilter: 'Citavi filter'
    },
    Bib: {
      Title: 'Bibliography',
      Field: 'Field',
      Italic: 'kursiv',
      Prefix: 'Prefix',
      Suffix: 'Suffix',
      Formatted: 'Formatted',
      CitaviAttributes: 'Citavi attributes'
    },
    Cite: {
      Title: 'Cites'
    },
    Delete: {
      Title: 'Delete category',
      Message: 'Do you really want to delete the category {name}'
    }
  },
  Common: {
    Abort: 'Abort',
    Create: 'Create',
    Delete: 'Delete',
    Close: 'Close',
    Attribute: 'Attribute',
    Value: 'Value',
    Field: 'Field',
    Settings: 'Settings',
    Save: 'Save',
    Rename: 'Rename',
    Continue: 'Continue'
  },
  Rules: {
    ProjectAlreadyExists: 'A project with this name already exists',
    NoSpaces: 'No spaces allowed',
    NotEmpty: 'An empty value is not valid',
    KeyAlreadyExists: 'This key is already in usage',
    CategoryNameAlreadyExists: 'This name is already in usage',
    FieldAlreadyExists: 'A field with this name already exists',
    NoSpecialChars: 'No special chars permitted',
  },
  Errors: {
    Title: 'An error occurred',
    ErrorHint: 'If you think this error is a bug, fell free to create an issue in {github}.',
    CreateIssue: 'Create issue',
    ErrorCreating: 'The project could not be created',
    ErrorDeleting: 'The project could not be deleted',
    ErrorSaving: 'An error occurred during saving',
  },
  Success: {
    CreateProject: 'Project successfully created',
    DeleteProject: 'Project successfully deleted',
    SaveCategory: 'Category successfully saved',
    DeleteCategory: 'Category successfully deleted',
    SaveEntry: 'Entry successfully saved',
    DeleteEntry: 'Entry successfully deleted',
    SaveConfig: 'Settings successfully saved',
  }
}
