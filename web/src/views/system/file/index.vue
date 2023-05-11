<template>
  <div class="container">
    <Breadcrumb />
    <a-card class="general-card">
      <File></File>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import File from '@/components/file/index.vue';
</script>

<script lang="ts">
  export default {
    name: 'File',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
</style>
