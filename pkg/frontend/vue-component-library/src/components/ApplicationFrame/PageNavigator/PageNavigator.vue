<script setup lang="ts">
import {NavState, useApplicationStateStore} from "../../../stores/ApplicationStateStore/ApplicationStateStore";

const applicationStateStore = useApplicationStateStore();

const props = defineProps({
  instantNav: Boolean
});

// methods
function openPage(page: string) {
  applicationStateStore.openPage(page);
}

function goBack() {
  applicationStateStore.goBack(1);
}

</script>

<template>
  <div class="container" :class="applicationStateStore.instantNav ? 'disable-animations' : ''">
    <div class="page" :id="`page-${n+1}`" v-for="(e,n) in applicationStateStore.history"
         :class="`${n === applicationStateStore.history.length-1 && applicationStateStore.navState != NavState.Back || n === applicationStateStore.history.length-2 && applicationStateStore.navState === NavState.Back ? 'opened' : ''}`">
      <h2>nav {{ applicationStateStore.instantNav }}</h2>
      <slot :name="n" :openPage="openPage" :goBack="goBack" />
    </div>
  </div>

  <v-dialog>

  </v-dialog>
</template>

<style scoped lang="scss">
.disable-animations {
  & * {
    transition: none !important;
  }
}

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