<script setup lang="ts">

import {NavState, useApplicationStateStore} from "../../../stores/ApplicationStateStore/ApplicationStateStore";

const applicationStateStore = useApplicationStateStore();

// methods
function openPage(page: string) {
  applicationStateStore.openPage(page);
}

function goBack() {
  applicationStateStore.goBack(1);
}

</script>

<template>
  <div class="container">
    <div class="page" v-for="(e,n) in applicationStateStore.history"
         :class="`${n === applicationStateStore.history.length-1 && applicationStateStore.navState != NavState.Back || n === applicationStateStore.history.length-2 && applicationStateStore.navState === NavState.Back ? 'opened' : ''}`">
      <slot :name="n" :openPage="openPage" :goBack="goBack" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  height: 100%;
  width: 100%;
  padding: 0;

  & .page {
    flex: 1 0 auto;
    height: 100%;
    width: 0;
    overflow-x: hidden;
    transition: width .75s ease-in-out;

    &.opened {
      width: 100%;
    }
  }
}
</style>