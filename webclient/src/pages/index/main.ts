import Vue from 'vue'
import App from './Index.vue'
import vuetify from '../../common/plugins/vuetify'
import CreateI18n from "@/common/plugins/i18n";
import {GermanTranslations} from "@/pages/index/i18n/German";
import {EnglishTranslations} from "@/pages/index/i18n/English";

Vue.config.productionTip = false

const i18n = CreateI18n(GermanTranslations, EnglishTranslations)

new Vue({
  vuetify,
  i18n,
  render: h => h(App)
}).$mount('#app')
