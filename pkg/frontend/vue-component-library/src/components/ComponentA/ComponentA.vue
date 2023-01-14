<script setup lang="ts">
interface Props {
  msg: string
}

const props = defineProps<Props>()
</script>

<template>
  <div>
    Hello {{ msg }}! <v-icon>mdi-delete</v-icon>
    <v-btn color="primary">test</v-btn>
  </div>
</template>