<script setup lang="ts">
import {onMounted, ref, watch} from "vue";

const props = defineProps({
  title: String,
  keydownTarget: Document
});

const opened = ref(false);
const focusBordersActivated = ref(false);

watch(focusBordersActivated, () => {
  setFocusBorders();
});

function setFocusBorders() {
  if (focusBordersActivated.value) {
    const elem = props.keydownTarget!.createElement("style");
    elem.id = "focus-borders-set";
    props.keydownTarget!.head.append(elem);
    elem.appendChild(props.keydownTarget!.createTextNode("*:focus{outline: 2px solid red!important    ; outline-offset: -2px;}"));
  } else {
    props.keydownTarget!.getElementById("focus-borders-set")?.remove();
  }

  localStorage.setItem("thesortexFocusBorders", String(focusBordersActivated.value));
}

let ctrlPressed = false;

onMounted(() => {
  focusBordersActivated.value = localStorage.getItem("thesortexFocusBorders") === "true";

  props.keydownTarget!.addEventListener("keyup", (e: KeyboardEvent) => {
    if (e.key === "Control") {
      if (!ctrlPressed) {
        ctrlPressed = true;
        setTimeout(() => {
          ctrlPressed = false;
        }, 300);
      } else {
        opened.value = true;
      }
    }
  });
});
</script>

<template>
  <v-dialog v-model="opened" width="500" height="500" theme="light">
    <v-card>
      <v-card-title>{{ title }}</v-card-title>
      <v-card-text>
        Hier kannst du verschiedene Eisntellungen für die Barrierefreiheit aktivieren.
        <v-list>
          <v-list-item title="Fokusrahmen">
            <template v-slot:prepend>
              <v-icon icon="mdi-image-filter-center-focus-strong"></v-icon>
            </template>
            <template v-slot:append>
              <v-checkbox-btn v-model="focusBordersActivated" />
            </template>
          </v-list-item>
          <!--          <v-list-item title="Erhöhte Kontraste">-->
          <!--            <template v-slot:prepend>-->
          <!--              <v-icon icon="mdi-invert-colors"></v-icon>-->
          <!--            </template>-->
          <!--            <template v-slot:append>-->
          <!--              <v-checkbox-btn />-->
          <!--            </template>-->
          <!--          </v-list-item>-->
        </v-list>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary">
          Schließen
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped lang="scss">

</style>