import {defineStore} from "pinia";
import {ref} from "vue";

export const useErrorSuccessStore = defineStore('error-success', () => {
  const errorMessage = ref('');

  const successMessage = ref('');

  function handleResponse(success: boolean, successMsg: string, errorMsg: string) {
    if( !success ) {
      errorMessage.value = errorMsg;
    } else {
      successMessage.value = '';
      setTimeout(() => {
        successMessage.value = successMsg;
      })
    }
  }

  function clean() {
    errorMessage.value = '';
    successMessage.value = '';
  }

  return {errorMessage, successMessage, handleResponse, clean}
})
