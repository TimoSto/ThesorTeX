<script setup lang="ts">
</script>

<template>
  <v-app>
    <div class="outer">
      <div class="inner">
        <v-btn>test</v-btn>
      </div>
    </div>
  </v-app>
</template>

<style scoped lang="scss">
.outer{
  background-color: red;
  padding: 10px;
  & .inner {
    background-color: green;
    padding: 100px;
  }
}
</style>
