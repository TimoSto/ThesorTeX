<template>
  <v-img :src="image" @click="opened=true" style="cursor: pointer"
         :style="imageStyles? imageStyles : 'max-height: calc(100vh - 100px)'" />
  <v-dialog v-model="opened" fullscreen scrim="black" class="darker-scrim">
    <v-img :src="image"
           style="max-height: calc(100% - 50px); max-width: calc(100% - 50px); margin-top: 25px; margin-left: 25px;" />
    <v-btn icon position="fixed" size="large" style="right: 25px; top: 25px;" @click="opened=false">
      <v-icon size="large">mdi-close</v-icon>
    </v-btn>
  </v-dialog>
</template>

<script lang="ts" setup>
import {ref} from "vue";

defineProps({
  image: {
    type: String,
    required: true
  },
  imageStyles: String
});

const opened = ref(false);
</script>

<style lang="scss">
.darker-scrim {
  & .v-overlay__scrim {
    opacity: 0.7;
  }
}
</style>