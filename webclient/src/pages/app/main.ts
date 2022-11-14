import Vue from 'vue'
import App from '@/pages/app/App.vue'
import router from '@/pages/app/router'
import store from '@/pages/app/store'
import vuetify from '@/common/plugins/vuetify'
import CreateI18n from "@/common/plugins/i18n";
import {GermanTranslations} from "@/pages/app/i18n/German";
import {EnglishTranslations} from "@/pages/app/i18n/English";

Vue.config.productionTip = false

const i18n = CreateI18n(GermanTranslations, EnglishTranslations)

new Vue({
  router,
  store,
  vuetify,
  i18n,
  render: h => h(App)
}).$mount('#app')
