<template>
  <v-layout>
    <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

    <AppbarContent
      bar-color="background"
      :fixed="true"
      :level="1"
    >
      <template v-slot:bar>
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
      </template>
      <template v-slot:content>
        <div style="width: 20px; height: 200px; background-color: red; margin-bottom: 300px"></div>
      </template>
    </AppbarContent>
  </v-layout>
</template>

<script>
import LogoSVG from "@/components/LogoSVG";
import AppbarContent from "@/components/AppbarContent";
export default {
  components: {AppbarContent, LogoSVG},
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
