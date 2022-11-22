<template>
  <div>

    <v-dialog
      width="450"
      v-model="errorOpen"
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

    <v-snackbar :value="successOpen">
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
import {MutationTypes} from "../store/mutation-types";

export default Vue.extend({
  name: "ResponseHandler",
  computed: {
    errorMessage(): string {
      return this.$store.state.actionResult.error;
    },
    errorOpen(): boolean {
      return this.errorMessage !== '';
    },
    successMessage(): string {
      return this.$store.state.actionResult.success;
    },
    successOpen(): boolean {
      return this.successMessage !== '';
    },
  },

  data() {
    return {
      currentTime: 0,
      counter: 0,
    };
  },

  watch: {
    successMessage(v: string) {
      console.log(v);
      if( v !== "" ) {
        this.currentTime = 0;
        this.counter++;
        this.syncSnackbarTime(this.counter);
      }
    }
  },

  methods: {
    syncSnackbarTime(count: number) {
      this.currentTime += 100;
      console.log(count)
      if( count === this.counter ) {
        if( this.currentTime < 5000 ) {
          setTimeout(() => {
            this.syncSnackbarTime(count);
          }, 100)
        } else {
          this.$nextTick(() => {
            this.currentTime = 0;
            console.log('kill', count)
            this.$store.commit(MutationTypes.SET_SUCCESS, "");
          })
        }
      }
    }
  },
});
</script>

<style scoped>

</style>