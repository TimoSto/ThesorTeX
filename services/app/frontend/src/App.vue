<template>
  <v-app>
    <v-main>
      <v-app-bar
        color="primary"
        elevation="0"
      >
        <v-app-bar-nav-icon
          @click.stop="sidebar = !sidebar"
        />

        <v-app-bar-title>
          ThesorTeX
        </v-app-bar-title>

        <v-spacer />
      </v-app-bar>

      <v-navigation-drawer
        permanent
        :rail="!sidebar"
        :rail-width="68"
      >
        <!--Sidebar content-->
      </v-navigation-drawer>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: "App",
  data() {
    return {
      sidebar: false,
    }
  }
})
</script>

<style scoped>

</style>