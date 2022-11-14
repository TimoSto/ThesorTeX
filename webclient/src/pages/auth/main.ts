import Vue from 'vue'
import App from '@/pages/auth/Auth.vue'
import router from '@/pages/auth/router'
import vuetify from '@/common/plugins/vuetify'
import CreateI18n from "@/common/plugins/i18n";
import {GermanTranslations} from "@/pages/auth/i18n/German";
import {EnglishTranslations} from "@/pages/auth/i18n/English";

Vue.config.productionTip = false

const i18n = CreateI18n(GermanTranslations, EnglishTranslations)

new Vue({
  router,
  vuetify,
  i18n,
  render: h => h(App)
}).$mount('#app')
