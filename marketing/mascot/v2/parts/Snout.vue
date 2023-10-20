<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
    startX: Number,
    startY: Number,
    right: Boolean,
});

const shape = computed(() => {
    let path = `M ${props.startX} ${props.startY}`;

    // upper shaft
    path += "q -11 -4 -40 3";

    // part covered by nose
    path += "l -11 7 ";

    // lower shaft
    path += "q 10 40 70 42";

    // mouth
    path += "q 4 0 15 -21";

    // smile wrinkle
    path += "m -7 1 q 9 -2 13 0";

    return path;
});

const nose = computed(() => {
    let path = `M ${props.startX - 40} ${props.startY + 3}`;

    // upper
    path += "q -11 -2 -15 0";

    // left
    path += "q -5 7 7 15";

    // right
    path += "q 10 -6 7 -15";

    return path;
});

// TODO: maybe add whiskers

</script>

<template>
    <g data-type="snout">
        <path data-type="snout-shaft" stroke="black" stroke-width="2" fill="white" :d="shape"/>
        <path data-type="nose" stroke="black" fill="#58676A" stroke-width="2" :d="nose"/>
    </g>
</template>

<style scoped lang="scss">

</style>