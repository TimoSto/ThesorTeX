<script setup lang="ts">
import {computed} from "vue";
import Eye from "./parts/Eye.vue";
import Snout from "./parts/Snout.vue";
import Head from "./parts/Head.vue";

const headShapePath = computed(() => {

    let path = "M ";

    // move to left temple
    path += "50, 50";

    // forehead
    path += "q100 -12 200 0";

    // ear hair
    path += "c 36 -9 54 -11 86 13";
    path += "q -22 3 -32 20";

    // hair under ear
    path += "q 15 5 23 18";
    path += "c -1 5 -14 2 -16 6";
    path += "q 12 4 15 14";
    path += "q -2 1 -7 -1 ";
    path += "c 7 5 7 13 11 15";
    path += "c -1 3 -6 1 -8 6";

    // beard 1
    path += "q 9 2 23 25";
    path += "c 3 5 -7 4 -12 8 ";

    // beard 2
    path += "q 17.5 5.5 40 35";
    path += "q 2 6 -20 10";

    // beard 3
    path += "q 26 20 23 27";
    path += "q -5 3 -32 8";
    path += "q 33 9 57 41";
    path += "q -42 9 -44 20";

    // chin
    path += "q -6 9 -18 21";
    path += "q -29 -5 -54 15";
    path += "q -124 29 -251 -15";
    path += "q -25 -18 -54 -19";
    path += "q -5 -15 -18 -21";

    // beard 4
    path += "q -5 -20 -44 -20";
    path += "q 17 -43 41 -57";

    // beard 5
    path += "q -8 -3 -15 -2";
    path += "q -2 -17 29 -31";

    // beard 6
    path += "q -5 -9 -16 -6";
    path += "q 6 -33 48 -56";

    // hair under left ear
    path += "q -14 -2 3 -12";
    path += "q -2 -4 -7 -6";
    path += "q 3 -10 17 -13";
    path += "c 5 -3 -2 -5 4 -9";
    path += "q 1 -3 -4 -4";
    path += "q 1 -6 7 -10";

    // ear hair left
    path += "q -17 -14 -29 -15";
    path += "q 22 -19 54 -5";

    path += "";

    return path;
});


</script>

<template>
    <svg xmlns="http://www.w3.org/2000/svg" width="300" height="300"
         viewBox="0 0 300 300">
        <g id="head">
            <Head :start-x="95" :start-y="50"/>
            <Eye :start-x="100" :start-y="90"/>
            <Eye :start-x="150" :start-y="90" right/>
            <Snout :start-x="125" :start-y="110"/>
        </g>
    </svg>
</template>

<style scoped lang="scss">

</style>