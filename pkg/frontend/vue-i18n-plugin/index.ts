import {createI18n, I18n} from "vue-i18n";
import {useI18n} from "vue-i18n";

export default function CreateI18n(german: any, english: any): I18n<{ de: any, en: any }> {
    return createI18n({
        legacy: false, // you must set `false`, to use Composition API
        locale: window.navigator.language.indexOf('de') >= 0 ? 'de' : 'en',
        fallbackLocale: 'en', // set fallback locale
        messages: {
            de: german,
            en: english
        },
    })
}

export { useI18n }
