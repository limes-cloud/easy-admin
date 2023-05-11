<template>
  <a-drawer
    :width="500"
    title="个人信息"
    unmount-on-close
    :visible="visible"
    :hide-cancel="true"
    ok-text="关闭"
    @cancel="cancel"
    @ok="cancel"
  >
    <div>
      <!-- <Breadcrumb :items="['个人信息']" /> -->
      <UserInfoHeader />
      <div class="content">
        <a-tabs default-active-key="1" type="rounded">
          <a-tab-pane key="1" title="基础设置">
            <BasicInformation />
          </a-tab-pane>
          <a-tab-pane key="2" title="安全设置">
            <SecuritySettings />
          </a-tab-pane>
        </a-tabs>
      </div>
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref, watch } from 'vue';
  import UserInfoHeader from './components/user-info-header.vue';
  import BasicInformation from './components/basic-information.vue';
  import SecuritySettings from './components/security-settings.vue';

  const props = defineProps({
    show: Boolean,
  });

  watch(
    () => props.show,
    (show: any) => {
      visible.value = show;
    }
  );

  const visible = ref(false);
  const emit = defineEmits(['cancel']);
  const cancel = () => {
    emit('cancel', false);
  };
</script>

<script lang="ts">
  export default {
    name: 'Info',
  };
</script>

<style scoped lang="less">
  .container {
    // padding: 0 20px 20px 20px;
  }

  .content {
    margin-top: 15px;
    // padding: 20px 0 0 20px;
    min-height: 380px;
    background-color: var(--color-bg-2);
    border-radius: 4px;
  }
</style>
