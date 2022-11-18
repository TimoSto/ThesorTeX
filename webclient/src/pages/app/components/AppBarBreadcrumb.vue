<template>
  <v-breadcrumbs :items="displayItems">
    <template v-slot:divider>
      <v-icon style="font-size: 22px">mdi-chevron-right</v-icon>
    </template>
    <template v-slot:item="{ item }">
      <v-breadcrumbs-item
          style="font-size: 20px; padding: 0 8px;"
          :style="item.disabled ? '' : 'cursor: pointer'"
          v-ripple="!item.disabled"
          @click="$emit('clicked', item)"
      >
        {{item.text}}
      </v-breadcrumbs-item>
    </template>
  </v-breadcrumbs>
</template>

<script lang="ts">
import Vue from "vue";

interface Item {
  text: string
  disabled: boolean
}

export default Vue.extend({
  name: "AppBarBreadcrumb",
  props: [
      'items'
  ],
  data() {
    return {
      displayItems: [] as Item[]
    }
  },

  mounted() {
    this.updateItems();
  },

  watch: {
    items() {
      this.updateItems();
    }
  },

  methods: {
    updateItems() {
      let items: Item[] = [];

      this.items.forEach((i: string) => {
        items.push({
          text: i,
          disabled: false
        })
      });

      if( items.length > 0 ) {
        items[items.length - 1].disabled = true;
      }

      this.displayItems = items;
    }
  }
})
</script>

<style scoped lang="scss">
@import '@/common/styles/theme_fixes.scss';
</style>

<style>
.v-breadcrumbs__divider {
  padding: 0 !important;
}
</style>