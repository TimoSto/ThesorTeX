<script setup>

import {computed} from "vue";

const headShapePath = computed(() => {

  let path = "M ";

  // move to left temple
  path += "50, 50"

  // forehead
  path += "q100 -12 200 0";

  // ear hair
  path += "c 36 -9 54 -11 86 13";
  path += "q -22 3 -32 20";

  // hair under ear
  path += "q 15 5 23 18"
  path += "c -1 5 -14 2 -16 6";
  path += "q 12 4 15 14";
  path += "q -2 1 -7 -1 ";
  path += "c 7 5 7 13 11 15";
  path += "c -1 3 -6 1 -8 6";

  // beard 1
  path += "q 9 2 23 25";
  path += "c 3 5 -7 4 -12 8 ";

  // beard 2
  path += "q 17.5 5.5 40 35"
  path += "q 2 6 -20 10";

  // beard 3
  path += "q 26 20 23 27";
  path += "q -5 3 -32 8";
  path += "q 33 9 57 41";
  path += "q -42 9 -44 20";

  // chin
  path += "q -6 9 -18 21"
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

  path += ""

  return path;
})

</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" version="1.0" width="300" height="300"
       viewBox="0 0 300 300">
    <metadata>Mascot for ThesorTeX (https://github.com/TimoSto/ThesorTeX)</metadata>

    <g id="head_v1">
      <path id="outer_bottom" stroke-width="2" stroke="black" fill="#E2E0D4"
            d="M94 100 l -15 39 q -10 14 -15 40 l 10 -9 q -3 5 0 23 q 10 19 60 53 q 10 7 25 0 q 50 -40 60 -50 q 7 -5 8 -14 l 5 6 l 18 1 q 2 -5 -7 -20 l -2 -9 q -25 2 -55 -3 q -15 6 -29 9 l -50 0 q -15 -4 -19 -7 l 23 -57 Q 105 98 94 100" />
      <path id="inner_color" stroke-width="1" stroke="none" fill="#8D6436"
            d="M110 100 q -16 10 -25 60 l 5 0 q 15 14 19 7 l 50 0 q 14 -4 29 -9 q 27 5 55 3 l 5 -5 l 15 10 q 10 -10 -23 -37 l 15 10 q 5 -5 -20 -50 l -20 -15 L 110 100" />
      <path id="inner_border" stroke-width="2" stroke="black" fill="none"
            d="M242 162 l 7 -5 l 15 10 q 10 -10 -23 -37 l 15 10 q 5 -5 -20 -50 l -20 -15 L 110 100" />
      <path id="left_ear" stroke-width="2" stroke="black" fill="#8D6436"
            d="M95 90 l -24 10 l -4 21 q 5 3 20 -5 L 95 90" />
      <path id="right_ear" stroke-width="2" stroke="black" fill="#8D6436"
            d="M200 71 l 10 -7 q 20 -5 45 -3 q 2 20 -22 30 q -5 -5 -20 -2 L 200 71" />
      <g id="left_eye">
        <path stroke="black" stroke-width="2" fill="white"
              d="M 100 120 q -20 18 2 31 q 16 7 29 -8 q 7 -8 0 -21 q -9 -17 -31 -2" />
        <path stroke="black" fill="black" d="M 109 153 q 0 -21 21 -33 q 8 15 1 23 q -14 15 -23 10" />
      </g>
      <g id="right_eye">
        <path stroke="black" stroke-width="2" fill="white"
              d="M 161 121 q -20 18 2 31 q 16 7 29 -8 q 7 -8 0 -21 q -9 -17 -31 -2" />
        <path stroke="black" fill="black" d="M 170 153 q 0 -21 21 -33 q 8 15 1 23 q -14 15 -23 10" />
      </g>
      <g id="snout">
        <path id="mouth" stroke="black" fill="#BD3C4C" stroke-width="2"
              d="M 160 183 q -5 23 -20 30 q -3 2 -9 0 q -11 -2 -16 -25 L 160 183" />
        <path id="snout_shaft_right" stroke="black" fill="white" stroke-width="2"
              d="M 140 167 q 40 2 11 25 q -20 9 -25 -5 l 0 -5" />
        <path id="snout_shaft_left" stroke="black" fill="white" stroke-width="2"
              d="M 125 180 l 0 5 q -6 11 -23 -5 q -10 -10 10 -14" />
        <path id="nose" stroke="black" stroke-width="2" fill="black"
              d="M 140 167 q -17 -9 -30 1 q 4 10 15 15 q 11 -5 15 -15" />
      </g>
    </g>

  </svg>

  <svg xmlns="http://www.w3.org/2000/svg" version="1.0" width="400" height="450"
       viewBox="-100 0 500 450">
    <g id="head">
      <path stroke="black" stroke-width="2" fill="none" :d="headShapePath"></path>
    </g>
  </svg>
</template>

<style scoped lang="scss">

</style>