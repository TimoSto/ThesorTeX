<template>
  <div>
    <slot v-for="n in pages" :name="n"></slot>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export interface NavAreaMethods extends Vue {
  addPage(name: string): void
}

export default Vue.extend({
  name: "NavArea",
  data() {
    return {
      pages: [] as string[]
    }
  },
  methods: {
    addPage(name: string) {
      this.pages.push(name)
    }
  }
})
</script>

<style scoped>

</style>