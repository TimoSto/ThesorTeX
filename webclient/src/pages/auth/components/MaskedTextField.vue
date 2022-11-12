<template>
  <v-text-field
    :label="label"
    v-on:input="maskInput"
    v-model="inputValue"
    :append-icon="isMasked ? 'mdi-eye' : 'mdi-eye-off'"
    @click:append="isMasked = !isMasked"
    :rules="appliedRules"
    style="margin-bottom: 8px"
  />
</template>

<script lang="ts">
import Vue from "vue";
import MaskValue, {maskChar} from "./MaskValue";

export interface MaskedTextFieldInterface {
  value: string
}

export default Vue.extend({
  name: "MaskedTextField",
  props: [
      'label',
      'rules'
  ],
  data() {
    return {
      value: '',
      valueMasked: '',
      inputValue: '',
      isMasked: true
    }
  },
  watch: {
    isMasked() {
      if( !this.isMasked ) {
        this.inputValue = this.value;
      } else {
        this.valueMasked = maskChar.repeat(this.inputValue.length);
        this.maskInput()
      }
    }
  },
  methods: {
    maskInput() {
      if( !this.isMasked ) {
        this.value = this.inputValue;
        return;
      }
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
  },
  computed: {
    appliedRules(): ((v: string) => boolean|string)[] {

      if( !this.rules ) return [];

      return [
        (v: string) => {

          let result: boolean|string = true;

          this.rules.forEach((r: (v: string) => boolean|string) => {
            const res = r(this.value)

            if ( res !== true ) {
              return result = res;
            }
          })

          if( result === true ) {
            return true
          }

          return this.$t(result) as string;
        }
      ]
    }
  }
})
</script>

<style scoped>

</style>