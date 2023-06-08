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

export const Simple = Template.bind({});
Simple.storyName = "Default";
Simple.args = {
    label: "hallo",
    test: "test slot",
};