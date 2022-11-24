<template>
  <v-app>
    <v-main class="test">
      <div class="t2">

      </div>
      <HelloWorld />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
  import HelloWorld from '../../components/HelloWorld.vue'
</script>

<style lang="scss" scoped>
.test {
  background-color: red;
  & .t2 {
    width: 50px;
    height: 50px;
    background-color: darkblue;
  }
}
</style>
