<script setup lang="ts">
import {DocumentationPack} from "../api/GetDocumentation";
import {onMounted, ref} from "vue";
import GetDSGVO from "../api/GetDSGVO";
import DocumentationPanel from "./DocumentationPanel.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";

const jsonDSGVO = ref<DocumentationPack | undefined>(undefined);

const {t} = useI18n();

onMounted(async () => {
    jsonDSGVO.value = await GetDSGVO();

    console.log(jsonDSGVO.value);
});
</script>

<template>
    <v-dialog width="950">
        <template v-slot:activator="{props}">
            <v-btn class="mx-2" variant="text" color="white" style="font-weight: bold" v-bind="props">
                {{ jsonDSGVO?.Title }}
            </v-btn>
        </template>
        <v-card theme="light">
            <v-toolbar density="comfortable" color="white">
                <v-toolbar-title>{{ jsonDSGVO?.Title }}</v-toolbar-title>
                <v-spacer/>
                <a href="/dsgvo" download="Datenschutzerklärung.pdf" style="margin-right:12px">
                    <v-btn icon color="black" :title="t(i18nKeys.Common.Download)">
                        <v-icon>mdi-download</v-icon>
                    </v-btn>
                </a>
            </v-toolbar>
            <v-card-text>
                <v-expansion-panels multiple v-if="jsonDSGVO != undefined">
                    <DocumentationPanel v-for="d in (jsonDSGVO as DocumentationPack).Docs" :doc="d"/>
                </v-expansion-panels>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<style scoped lang="scss">

</style>