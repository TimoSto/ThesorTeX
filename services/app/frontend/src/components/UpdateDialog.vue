<template>
  <v-card position="fixed" elevation="3" max-width="500" style="z-index: 1010; right: 25px; top: 50px;" v-if="display">
    <v-card-title>
      <i18n-t :keypath="i18nKeys.UpdateDialog.Title">
        <template #version>
          {{ version }}
        </template>
      </i18n-t>
    </v-card-title>
    <v-card-text class="text-body-1">
      <i18n-t :keypath="i18nKeys.UpdateDialog.Content">
        <template #link>
          <a href="https://thesortex.com/#/downloads/" target="_blank" style="color: rgb(var(--v-theme-primary))">thesortex.com</a>
        </template>
      </i18n-t>
    </v-card-text>
    <v-card-actions>
      <a href="https://thesortex.com/#/downloads" target="_blank">
        <v-btn color="primary">{{ t(i18nKeys.UpdateDialog.ToDownloads) }}</v-btn>
      </a>
      <v-btn color="primary" @click="ignore">{{ t(i18nKeys.UpdateDialog.Ignore) }}</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {computed, onMounted, ref} from "vue";

const {t} = useI18n();

const ignored = ref([] as string[]);

const display = computed(() => {
  return props.version && props.version.length > 0 && ignored.value.indexOf(props.version) === -1;
});

const props = defineProps({
  version: String
});

const emit = defineEmits(["ignore"]);

function ignore() {
  ignored.value.push(props.version!);
  localStorage.setItem("thesortex.ignoreUpdates", JSON.stringify(ignored.value));
}

onMounted(() => {
  ignored.value = localStorage.getItem("thesortex.ignoreUpdates") ? JSON.parse(localStorage.getItem("thesortex.ignoreUpdates")!) : [];
});

</script>

<style lang="scss" scoped>
@import "../styles/common.scss";
</style>