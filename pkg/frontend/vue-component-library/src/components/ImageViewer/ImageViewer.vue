<template>
  <v-img :src="image" @click="opened=true" style="cursor: pointer"
         :style="imageStyles? imageStyles : 'max-height: calc(100vh - 100px)'" />
  <v-dialog v-model="opened" scrim="black" class="darker-scrim" :width="adjustedDimensions.width"
            :height="adjustedDimensions.height">
    <v-card theme="light" color="transparent" elevation="0">
      <img :src="image" />
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from "vue";

const props = defineProps({
  image: {
    type: String,
    required: true
  },
  imageStyles: String
});

const opened = ref(false);

const fullWidth = ref(0);
const fullHeight = ref(0);

const imgWidth = ref(0);
const imgHeight = ref(0);

// computed
// TODO: test this
const adjustedDimensions = computed(() => {
  const ratio = imgWidth.value / imgHeight.value;

  const maxDimension = ratio > 1 ? fullWidth.value - 50 : fullHeight.value - 50;
  const scaleFactor = ratio > 1 ? maxDimension / imgWidth.value : maxDimension / imgHeight.value;
  const secondDimensionScaled = ratio > 1 ? imgHeight.value * scaleFactor : imgWidth.value * scaleFactor;

  return {
    width: ratio > 1 ? maxDimension : secondDimensionScaled,
    height: ratio > 1 ? secondDimensionScaled : maxDimension,
  };
});

// watchers
watch(() => props.image, () => {
  checkImageDimensions();
});

onMounted(() => {
  checkImageDimensions();
  fullWidth.value = window.innerWidth;
  fullHeight.value = window.innerHeight;
});

// methods
function checkImageDimensions() {
  const img = new Image();
  img.onload = function () {
    imgWidth.value = img.width;
    imgHeight.value = img.height;
  };
  img.src = props.image;
}
</script>

<style lang="scss">
.darker-scrim {
  & .v-overlay__scrim {
    opacity: 0.7;
  }
}
</style>