<template>
  <v-text-field
    :label="label"
    v-on:input="maskInput"
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
      valueMasked: ''
    }
  },
  methods: {
    maskInput() {
      const input = this.$refs.tf;
      if( input ) {
        const valueWithAdded = (input).$el.querySelector('input').value;
        const currentState = {
          CurrentValue: this.value,
          CurrentMask: this.valueMasked,
          InputState: valueWithAdded
        }
        const newState = MaskValue(currentState);
        console.log(newState, currentState)
      }
    }
  }
})
</script>

<style scoped>

</style>