<template>
  <div
    class="container"
    :class="animationClass + (props.instantSwitch ? ' disable-animations' : '')"
  >
  </div>
</template>
<script setup lang="ts">

// data
const props = defineProps({
  instantSwitch: {
    type: String,
    required: true
  }
})

let animationClass = "";
</script>