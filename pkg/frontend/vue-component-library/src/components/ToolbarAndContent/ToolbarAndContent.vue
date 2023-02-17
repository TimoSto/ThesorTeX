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
      :class="props.hideBar ? 'hidden-bar' : ''"
    >
      <slot name="content" />
    </v-sheet>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from "vue";

const emit = defineEmits(["scroll", "noScroll"]);

const props = defineProps({
  hideBar: {
    type: Boolean,
    default: false
  },
  barColor: {
    type: String,
    default: "background"
  },
  page: String
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
  });
});

watch(elevation, () => {
  if (elevation.value === 0) {
    emit("noScroll");
  } else {
    emit("scroll");
  }
});

watch(() => props.page, () => {
  root.value!.querySelector(".content")!.scrollTo(0, 0);
});
</script>

<style scoped lang="scss">
.content {
  padding: 0;
  margin: 0;
  width: 100%;
  overflow-y: auto;
  height: calc(100vh - 112px);

  &.hidden-bar {
    height: calc(100vh - 64px);
  }
}
</style>