<template>
  <div>

    <v-dialog
      width="450"
      v-model="errorMessage"
      >
      <v-card>
        <v-card-title>An Error Occurred</v-card-title>
        <v-card-text v-html="$t(errorMessage)"></v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn text>
            To Issues
          </v-btn>
          <v-btn text>
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-snackbar :value="successMessage !== ''">
    <span v-html="successMessage">
    </span>

      <v-progress-linear
          absolute
          bottom
          :value="Math.floor(100 * (currentTime / 5000))"
      />
    </v-snackbar>

  </div>
</template>

<script lang="ts">
//TODO: use this at others and therefore make state implement a global interface?
import Vue from "vue";

export default Vue.extend({
  name: "ResponseHandler",

  computed: {
    errorMessage: {
      get(): string {
        const v = this.$store.state.actionResult.error;
        if(v !== '') {
          this.syncPbar();
        }
        return v;
      },
      set(v: string) {
        if(v !== '') {
          this.syncPbar();
        }
      }
    },
    successMessage(): string {
      return this.$store.state.actionResult.success;
    }
  },

  data() {
    return {
      currentTime: 0,
    };
  },
  methods: {
    syncPbar() {
      //Create a timeout every 100 miliseconds
      setTimeout(() => {
        //Increment the time counter by 100
        this.currentTime += 100;

        //If our current time is larger than the timeout
        if (5000 <= this.currentTime) {

          //Wait 500 miliseconds for the dom to catch up, then reset the snackbar
          setTimeout(() => {
            this.currentTime = 0; // reset the current time
            this.$emit('timeoutReached')
          }, 500);
        } else {
          //Recursivly update the progress bar
          this.syncPbar();
        }
      }, 100);
    }
  },
});
</script>

<style scoped>

</style>