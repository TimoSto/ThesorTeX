<script setup>
import {computed} from "vue";

const props = defineProps({
    startX: Number,
    startY: Number,
});

const shape = computed(() => {
    let path = `M ${props.startX} ${props.startY} `;

    // forehead
    path += "q 50 -12 90 0";

    // right side
    path += "q 15 5 20 55";

    // right beard
    path += "q 13 15 13 20";
    path += "q 2 5 -10 -5";

    path += "q 14 16 15 20";
    path += "q 2 5 -9 -2";

    path += "q 14 12 18 16";
    path += "q -4 2 -18 -1";

    return path;
})
</script>

<template>
    <g data-type="head">
        <path data-type="head-shape" stroke="black" stroke-width="2" fill="none" :d="shape"/>
    </g>

</template>

<style scoped lang="scss">

</style>