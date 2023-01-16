import * as components from "vuetify/components"
import * as directives from "vuetify/directives"
import {createVuetify} from "vuetify";
import {MountingOptions} from "@vue/test-utils";
import {createI18n} from "vue-i18n";

const vuetify = createVuetify({ components, directives });

export default function CreateVuetifyMountingOptions(options: MountingOptions<any>, locale?: any): MountingOptions<any> {
    if( !options.global ) {
        options.global = {}
    }

    if( !options.global.plugins ) {
        options.global.plugins = []
    }

    options.global.plugins.push(vuetify);

    if( locale ) {
        const i18n = createI18n({
            locale: "en",
            messages: {
                en: locale
            },
            legacy: false
        })
        options.global.plugins.push(i18n);
    }

    return options;
}
