<template>
  <v-dialog
      v-model="errorOpened"
      width="450"
  >
    <v-card>
      <v-card-title>{{ errorTitle }}</v-card-title>
      <v-card-text>
        <p class="error-message">
          {{ message }}
        </p>
        <p class="suffix">
          {{ errorSuffix }}
        </p>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary"
          @click="emit('close')"
        >
          {{ close }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="successOpened">
    <span>
      {{ message }}
    </span>
    <v-progress-linear
        absolute
        bottom
        color="primary"
        :model-value="Math.floor(100 * (successTimeout / 5000))"
        style="margin-top: 4px"
    />
  </v-snackbar>
</template>

<script lang="ts" setup>
import {computed, ref, watch} from "vue";

// globals
const emit = defineEmits(["close"]);

// props
const props = defineProps({
  valid: Boolean,
  message: String,
  errorTitle: String,// e.g. "Something went wrong"
  errorSuffix: String,// e.g. "If you think this is a bug, ..."
  close: String,// e.g. "Close" or "Understood"
});

// data
const successTimeout = ref(2500);

let successTimeoutInterval: NodeJS.Timer;

// computed
const errorOpened = computed({
  get(): boolean {
    return !props.valid && props.message !== "";
  },
  set(v: boolean) {
    if( !v ) {
      emit("close");
    }
  }
});

const successOpened = computed({
  get(): boolean {
    return props.valid && props.message !== "";
  },
  set(v: boolean) {
    if( !v ) {
      // emit
    }
  }
});

// watchers
watch(() => props.message, () => {
  if( props.message === "" ) {
    successTimeout.value = 5000;
    clearInterval(successTimeoutInterval);
  } else if( props.valid ) {
    startTimeout();
  }
});

// methods
function startTimeout() {
  successTimeout.value = 5000;
  if( successTimeoutInterval ) {
    clearInterval(successTimeoutInterval);
  }
  successTimeoutInterval = setInterval(() => {
    successTimeout.value -= 5;
    if( successTimeout.value === 0 ) {
      emit("close");
    }
  }, 5);
}
</script>

<style scoped lang="scss">
.error-message {
  font-style: italic;
}
</style>