import {defineStore} from "pinia";
import {ref} from "vue";

export const useUnsaveCloseStore = defineStore('unsave-close', () => {
  const tried = ref(false);

  const triggered = ref(false);

  const fromPage = ref(-1);

  const promiseResolve = ref((value: unknown) => {});

  function trigger(p: number) {
    fromPage.value = p;
    tried.value = true;
    return new Promise((resolve) => {
      promiseResolve.value = resolve;
    })
  }

  return {tried, triggered, fromPage,  promiseResolve, trigger}
})
