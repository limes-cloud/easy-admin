<template>
  <svg
    aria-hidden="true"
    class="svg-icon-spin"
    :style="{ width: size + 'px', height: size + 'px' }"
  >
    <use :xlink:href="symbolId" :fill="color" />
  </svg>
</template>

<script lang="ts">
  import { defineComponent, computed } from 'vue';

  export default defineComponent({
    name: 'SvgIcon',
    props: {
      prefix: {
        type: String,
        default: 'icon',
      },
      name: {
        type: String,
        required: true,
      },
      color: {
        type: String,
        default: '#333',
      },
      size: {
        type: Number,
        default: 32,
      },
    },
    setup(props) {
      const symbolId = computed(() => `#${props.prefix}-${props.name}`);
      return { symbolId };
    },
  });
</script>

<style lang="less" scoped>
  .svg-icon-spin {
    fill: currentColor;
    vertical-align: middle;
  }
</style>
