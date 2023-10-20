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

    path += "q 6 10 10 14";
    path += "q 2 5 -10 -3";

    // chin
    path += "q -10 5 -20 15";

    path += "q -45 15 -90 0";

    path += "q -10 -10 -20 -15";

    // left beard
    path += "q -8 5 -10 3";
    path += "q 5 -4 10 -13";

    path += "q -4 2 -16 2";
    path += "q 14 -8 12 -16";

    path += "q -4 2 -13 -1";
    path += "q 14 -12 18 -16";

    path += "q 2 2 -10 7";
    path += "q 0 -5 10 -22";

    path += "q -2 5 -10 5";
    path += "q -13 15 9 -12";

    // left side
    path += "q 5 -55 11 -50";

    return path;
});

const leftEar = computed(() => {
    let path = `M ${props.startX + 15} ${props.startY} `;

    path += "q -7 1 -14 -6 q -9 -6 -27 1 q -5 0 -17 18 q -12 35 20 21 q 11 -5 27 -15 l 11 -19";

    return path;
});

const rightEar = computed(() => {
    let path = `M ${props.startX + 85} ${props.startY} `;

    path += "q 7 1 14 -6 q 9 -6 27 1 q 10 3 14 19 q 1 18 -13 12 q -16 -6 -24 -5 l -18 -21";

    return path;
});

</script>

<template>
    <g data-type="head">
        <path data-type="left-ear" stroke="black" stroke-width="2" fill="white" :d="leftEar"/>
        <path data-type="right-ear" stroke="black" stroke-width="2" fill="white" :d="rightEar"/>
        <path data-type="head-shape" stroke="black" stroke-width="2" fill="white" :d="shape"/>
    </g>

</template>

<style scoped lang="scss">

</style>