import VueI18n from 'vue-i18n';

export default function CreateI18n(german: any, english: any): VueI18n.I18n<{ de: any, en: any }> {
  return VueI18n.createI18n({
    legacy: false, // you must set `false`, to use Composition API
    locale: window.navigator.language.indexOf('de') >= 0 ? 'de' : 'en',
    fallbackLocale: 'en', // set fallback locale
    messages: {
      de: german,
      en: english
    }, // set locale messages
    // If you need to specify other options, you can set other options
    // ...
  })
}
