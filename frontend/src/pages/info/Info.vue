<template>
  <v-app>
    <v-main class="test">
      <div class="t2">

      </div>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";

const { t } = useI18n()
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
