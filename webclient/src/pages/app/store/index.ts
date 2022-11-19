import Vue, {PluginObject} from 'vue'
import Vuex from 'vuex'
import {appState} from "@/pages/app/store/state";
import {mutations} from "@/pages/app/store/mutations";
import {actions} from "@/pages/app/store/actions";

Vue.use(Vuex)

export default new Vuex.Store({
  state: appState,
  getters: {
  },
  mutations: mutations,
  actions: actions,
  modules: {
  }
})

export const TypedStorePlugin: PluginObject<void> = {
  install(VueInstance: typeof Vue) {
    Object.defineProperty(VueInstance.prototype, '$texStore', {
      get() {
        return this.$store;
      }
    });
  }
};
