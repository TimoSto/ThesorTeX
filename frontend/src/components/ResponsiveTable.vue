<template>
  <v-table style="border: 1px solid rgba(var(--v-theme-on-background), 0.2)">
    <thead>
      <tr>
        <th
          v-for="(h,i) in headers"
          :key="`header-${i}`"
          :style="`width: ${h.width}; ${h.minWidth !== '' ? `min-width: ${h.minWidth}` : ''}`"
        >
          <span
            v-if="h.icon !== ''"
          >
            <v-btn
              flat
              text
              @click="$emit('btnClicked', h.event)"
            >
              <v-icon>{{ h.icon }}</v-icon>
            </v-btn>
          </span>
          <span
            v-if="h.icon === ''"
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
            v-if="c.icon !== ''"
          >
            <v-btn
              flat
              text
            >
              <v-icon>{{ c.icon }}</v-icon>
            </v-btn>
          </span>
          <span
            v-if="c.icon === '' && !c.slot"
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
  icon: string,
  event: string,
  hideUnder: number
  slot?: boolean
}

export interface ResponsiveTableHeaderCell extends ResponsiveTableCell{
  width: string,
  minWidth: string
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

<style scoped>
.hide-ripple {
  opacity: 1!important;
}
</style>
