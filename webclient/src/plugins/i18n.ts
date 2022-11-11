import Vue from 'vue';
import VueI18n from 'vue-i18n';

Vue.use(VueI18n)

export default function CreateI18n(german: any, english: any): VueI18n {
    return new VueI18n({
        locale: window.navigator.language.indexOf('de') >= 0 ? 'de' : 'en',
        fallbackLocale: 'en',
        messages: {
            de: german,
            en: english
        }
    });
}
