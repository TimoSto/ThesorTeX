<template>
  <div ref="root">
    <v-toolbar
      v-if="!props.hideBar"
      :color="props.barColor"
      density="compact"
      :elevation="elevation"
      style="z-index: 1000"
    >
      <slot name="bar" />
    </v-toolbar>
    <v-sheet
      class="content"
    >
      <slot name="content" />
    </v-sheet>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";

const props = defineProps({
  hideBar: {
    type: Boolean,
    default: false
  },
  barColor: {
    type: String,
    default: "background"
  }
});

// data
const elevation = ref(0);
const root = ref<HTMLElement | null>(null);
onMounted(() => {
  root.value!.querySelector(".content")!.addEventListener("scroll", (e) => {
    if ((e.target! as HTMLElement).scrollTop > 0) {
      elevation.value = 1;
    } else {
      elevation.value = 0;
    }
  })
});
</script>

<style scoped>
.content {
  padding: 0;
  margin: 0;
  width: 100%;
  overflow-y: auto;
  height: calc(100vh - 112px)
}
</style>