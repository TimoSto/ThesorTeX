import Button from "./Button.vue";

export default {
    title: "Examples/Button",
    component: Button,
};

const Template = (args, {argTypes}) => ({
    components: {Button},
    props: Object.keys(argTypes),
    template: `
      <Button v-bind="$props">
      <template v-if="${"test" in args}" #test>${args.test}</template>
      </Button>
    `,
});
/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/vue/api/csf
 * to learn how to use render functions.
 */
export const Simple = Template.bind({});
Simple.storyName = "Default";
Simple.args = {
    label: "hallo",
    test: "test slot",
};