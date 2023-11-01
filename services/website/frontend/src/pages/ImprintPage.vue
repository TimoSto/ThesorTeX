<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";

const {t} = useI18n();

defineProps({
    smallDisplay: Boolean
});
</script>

<template>
    <div style="display: flex; align-items: start; padding-top: 12px; min-height: calc(100vh - 120px);">

        <v-container style="padding-top: 100px">
            <v-row>
                <v-col :cols="12">
                    <h2 class="text-h4 font-weight-bold text-white pa-4 pt-0 text-center" ref="container0">{{
                            t(i18nKeys.ImprintPage.Title)
                        }}</h2>
                    <v-divider style="margin-bottom: 8px"></v-divider>
                    <p>Timo Stovermann</p>
                    <p>Bree 8</p>
                    <p>46354 Südlohn</p>
                    <p>thesortex.contact@gmail.com</p>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<style scoped lang="scss">

</style>