<template>
  <v-table style="border: 1px solid rgba(var(--v-theme-on-background), 0.2)">
    <thead>
      <tr>
        <th
          v-for="(h,i) in headers"
          :key="`header-${i}`"
          :class="h.size"
        >
          <span
            v-if="h.slot"
          >
            <slot
              :name="`h-${i}`"
            />
          </span>
          <span
            v-if="!h.slot"
          >
            {{ h.content }}
          </span>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(r,i) in rows"
        :key="`row-${i}`"
        v-ripple="{
          class: disableRipple ? 'text-background' : ''
        }"
        :style="disableRipple ? '' : 'cursor: pointer'"
        @click="$emit('rowClicked', i)"
      >
        <td
          v-for="(c,n) in r"
          :key="`cell${i}${n}`"
        >
          <span
            v-if="!c.slot"
            v-html="c.content"
          />
          <span
            v-if="c.slot"
          >
            <slot
              :name="`${i}-${n}`"
            />
          </span>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>

<script lang="ts">

import {PropType} from "vue";

export interface ResponsiveTableCell {
  content: string,
  slot?: boolean
}

export interface ResponsiveTableHeaderCell extends ResponsiveTableCell{
  size: string
}

export enum SizeClasses {
  // eslint-disable-next-line no-unused-vars
  IconBtn = "size--icon-btn",
  // eslint-disable-next-line no-unused-vars
  TenPercent = "size--15-percent",
  // eslint-disable-next-line no-unused-vars
  TwentyFivePercent = "size--25-percent",
  // eslint-disable-next-line no-unused-vars
  ThirtyPercent = "size--30-percent"

}

export default {
  props: {
    headers: {
      type: Object as PropType<ResponsiveTableHeaderCell[]>,
      required: true
    },
    rows: {
      type: Object as PropType<ResponsiveTableCell[][]>,
      required: true
    },
    disableRipple: {
      type: Boolean
    }
  },
  emits: ['rowClicked', 'btnClicked'],
}
</script>

<style scoped lang="scss">
.hide-ripple {
  opacity: 1!important;
}
.size{
  &--icon-btn {
    width: 72px;
    max-width: 72px;
    padding: 0 4px!important;
    text-align: right!important;
  }
  &--15-percent {
    width: 15%;
    max-width: 15%;
  }
  &--25-percent {
    width: 25%;
    max-width: 25%;
  }
  &--30-percent {
    width: 30%;
    max-width: 30%;
  }
}
</style>
