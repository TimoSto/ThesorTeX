<template>
  <v-layout>
    <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

    <v-app-bar
      color="background"
      elevation="1"
      fixed
    >

      <LogoSVG
        fill="rgba(0, 136, 52, 0.8)"
        style="height: 35px; margin-left: 20px;"
        :show-title="true"
        font-color="rgba(0,0,0,0.7)"
      />

      <v-spacer></v-spacer>

      <v-btn
        variant="text"
        color="primary"
      >Download</v-btn>
    </v-app-bar>

    <v-main>

    </v-main>
  </v-layout>
</template>

<script>
import LogoSVG from "@/components/LogoSVG";
export default {
  components: {LogoSVG},
  data: () => ({
    drawer: false,
    group: null,
    items: [
      {
        title: 'Foo',
        value: 'foo',
      },
      {
        title: 'Bar',
        value: 'bar',
      },
      {
        title: 'Fizz',
        value: 'fizz',
      },
      {
        title: 'Buzz',
        value: 'buzz',
      },
    ],
  }),

  watch: {
    group () {
      this.drawer = false
    },
  },
}
</script>
