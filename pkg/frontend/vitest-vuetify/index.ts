import * as components from "vuetify/components"
import * as directives from "vuetify/directives"
import {createVuetify} from "vuetify";
import {MountingOptions} from "@vue/test-utils";

const vuetify = createVuetify({ components, directives });

export default function CreateVuetifyMountingOptions(options: MountingOptions<any>): MountingOptions<any> {
    if( !options.global ) {
        options.global = {}
    }

    if( !options.global.plugins ) {
        options.global.plugins = []
    }

    options.global.plugins.push(vuetify);

    return options;
}
