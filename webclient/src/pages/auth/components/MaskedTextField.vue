<template>
  <v-text-field
    :label="label"
    v-on:input="maskInput"
    v-model="inputValue"
  />
</template>

<script lang="ts">
import Vue from "vue";
import MaskValue from "./MaskValue";

export default Vue.extend({
  name: "MaskedTextField",
  props: [
      'label'
  ],
  data() {
    return {
      value: '',
      valueMasked: '',
      inputValue: ''
    }
  },
  methods: {
    maskInput() {
      const currentState = {
        CurrentValue: this.value,
        CurrentMask: this.valueMasked,
        InputState: this.inputValue
      }
      const newState = MaskValue(currentState);
      this.value = newState.CurrentValue;
      this.valueMasked = newState.CurrentMask;
      this.inputValue = newState.CurrentMask;
    }
  }
})
</script>

<style scoped>

</style>