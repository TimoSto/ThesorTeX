import ErrorSuccessDisplay from "./ErrorSuccessDisplay.vue";

export default {
    title: "VueComponentLibrary/ErrorSuccessDisplay",
    component: ErrorSuccessDisplay,
    render: (args: any, {argTypes}: any) => ({
        props: Object.keys(argTypes),
        components: {ErrorSuccessDisplay},
        template: "<ErrorSuccessDisplay v-bind=\"$props\" />",
    }),
    argTypes: {
        valid: Boolean,
        message: String,
        errorTitle: String,// e.g. "Something went wrong"
        errorSuffix: String,// e.g. "If you think this is a bug, ..."
        close: String,
    },
};

export const Success = {
    args: {
        valid: true,
        message: "Success Message"
    },
};

export const Error = {
    args: {
        valid: false,
        message: "Error Message",
        errorTitle: "Error 1",
        close: "Close it"
    },
};
