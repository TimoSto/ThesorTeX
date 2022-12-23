const { DEFAULT_THEME } = require('@cucumber/pretty-formatter');

module.exports = {
    default: {
        formatOptions: {
            theme: {
                ...DEFAULT_THEME,
                'step text': 'magenta'
            }
        },
        requireModule: [
            'ts-node/register'
        ],
        require: [
            'src/support/test.setup.ts',
            'src/step_definitions/**/*.ts'
        ],
        paths: [
            'src/features/**/*.feature'
        ]
    }
}
