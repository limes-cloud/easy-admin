<template>
  <div v-if="loading" class="loadingBox">
    <a-spin dot :tip="title" />
  </div>
  <a-config-provider :locale="locale">
    <router-view />
    <global-setting />
  </a-config-provider>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';
  import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
  import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';
  import GlobalSetting from '@/components/global-setting/index.vue';
  import useLocale from '@/hooks/locale';
  import { useAppStore } from '@/store';

  const appStore = useAppStore();

  const loading = computed(() => appStore.isLoading);
  const title = computed(() => appStore.loadingTitle);
  const { currentLocale } = useLocale();
  const locale = computed(() => {
    switch (currentLocale.value) {
      case 'zh-CN':
        return zhCN;
      case 'en-US':
        return enUS;
      default:
        return enUS;
    }
  });
</script>

<style lang="less" scoped>
  .loadingBox {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
    background-color: #fff;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
  }
</style>

<style lang="less">
  .cursor-pointer {
    cursor: pointer;
  }
  .arco-table-td-content {
    display: flex !important;
  }

  .container {
    padding: 0 20px 20px 20px;
    :deep(.arco-list-content) {
      overflow-x: hidden;
    }

    :deep(.arco-card-meta-title) {
      font-size: 14px;
    }
  }

  .model-footer {
    padding: 15px;
    display: flex;
    justify-content: space-between;
  }
</style>
