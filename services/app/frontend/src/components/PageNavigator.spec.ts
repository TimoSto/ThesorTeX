import {mount} from "@vue/test-utils";
import GuessAge from "./PageNavigator.vue";
import { expect, test } from "vitest";
test("mount component",async()=>{
    expect(GuessAge).toBeTruthy();

    const wrapper = mount(GuessAge,{
        props:{
            title:"Guess User Age App",
        },
    });
    expect(wrapper.text()).toContain("Guess User Age App");
})
