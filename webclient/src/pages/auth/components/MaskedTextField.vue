<template>
  <v-text-field
    :label="label"
    v-on:input="maskInput"
    v-model="inputValue"
    ref="tf"
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
  watch: {
    // valueMasked() {
    //   this.value += this.valueMasked.charAt(this.valueMasked.length - 1)
    //   console.log(this.value)
    //   this.valueMasked = this.valueMasked.substring(this.valueMasked.length - 2)
    //   this.valueMasked += '#';
    // }
  },
  data() {
    return {
      value: '',
      valueMasked: '',
      inputValue: ''
    }
  },
  methods: {
    maskInput() {
      const input = this.$refs.tf;
      if( input ) {
        const currentState = {
          CurrentValue: this.value,
          CurrentMask: this.valueMasked,
          InputState: this.inputValue
        }
        const newState = MaskValue(currentState);
        console.log(newState, currentState)
        this.value = newState.CurrentValue;
        this.valueMasked = newState.CurrentMask;
        this.inputValue = newState.CurrentMask;
      }
    }
  }
})
</script>

<style scoped>

</style>