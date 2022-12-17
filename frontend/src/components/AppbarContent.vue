<template>
  <v-app-bar
    v-if="props.level === 1"
    :color="props.barColor"
    density="default"
  >
    <slot
      name="bar"
    />
  </v-app-bar>
  <v-toolbar
    v-if="props.level > 1 && !props.hideBar"
    :color="props.barColor"
    density="compact"
  >
    <slot name="bar" />
  </v-toolbar>
  <v-sheet
    class="content"
    :class="props.level === 1 ? 'top-level ' : 'inner-level '"
  >
    <slot name="content" />
  </v-sheet>
</template>

<script setup lang="ts">
const props = defineProps({
  barColor: {
    type: String,
    default: ''
  },
  level: {
    type: Number,
    default: -1,
  },
  hideBar: {
    type: Boolean,
    default: false
  }
})
//TODO: elevate on scroll
</script>

<style scoped lang="scss">
.content {
  padding: 0;
  margin: 0;
  //display: flex;
  //flex-direction: column;
  width: 100%;
  overflow-y: auto;
  &.top-level {
    height: calc(100vh - 64px);
    margin-top: 64px;
    //@media only screen and (max-width: 960px) {
    //  height: calc(100vh - 56px);
    //  margin-top: 56px;
    //}
  }
}
</style>
