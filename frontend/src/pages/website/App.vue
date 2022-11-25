<template>
  <v-layout>
    <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

    <v-app-bar
      color="background"
      elevation="0"
    >

      <v-toolbar-title>ThesorTeX</v-toolbar-title>

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
export default {
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
