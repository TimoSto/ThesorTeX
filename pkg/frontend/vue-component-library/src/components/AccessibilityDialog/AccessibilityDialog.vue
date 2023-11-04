<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {i18nKeys} from "./i18n/keys";

const props = defineProps({
  keydownTarget: Document,
  i18n: {
    type: Function,
    default: (k: string) => k
  }
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

function openDialog() {
  opened.value = true;
}

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
  <slot :openDialog="openDialog"></slot>
  <v-dialog v-model="opened" width="500" height="500" theme="light">
    <v-card>
      <v-card-title>{{ i18n(i18nKeys.AccessibilityDialog.Title) }}</v-card-title>
      <v-card-text>
        {{ i18n(i18nKeys.AccessibilityDialog.Text) }}
        <v-list>
          <v-list-item>
            <v-list-item-title>
              {{ i18n(i18nKeys.AccessibilityDialog.FocusBorders) }}
            </v-list-item-title>
            <template v-slot:prepend>
              <v-icon icon="mdi-image-filter-center-focus-strong"></v-icon>
            </template>
            <template v-slot:append>
              <v-checkbox-btn v-model="focusBordersActivated"
                              :aria-label="i18n(i18nKeys.AccessibilityDialog.FocusBorders)" />
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
        <v-btn color="primary" @click="opened = false">
          {{ i18n(i18nKeys.AccessibilityDialog.Close) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped lang="scss">

</style>