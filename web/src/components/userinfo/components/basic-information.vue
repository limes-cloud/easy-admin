<template>
  <a-form ref="formRef" :model="formData" class="form" auto-label-width>
    <a-form-item
      field="nickname"
      label="昵称"
      :rules="[
        {
          required: true,
          message: '昵称是必填项',
        },
      ]"
    >
      <a-input
        v-model="formData.nickname"
        :max-length="15"
        placeholder="请输入昵称"
      />
    </a-form-item>

    <a-form-item
      field="sex"
      label="性别"
      :rules="[
        {
          required: true,
          message: '性别是必填项',
        },
      ]"
    >
      <a-switch
        v-model="formData.sex"
        checked-color="#2196f3"
        unchecked-color="#ff5722"
        type="round"
        :checked-value="true"
        :unchecked-value="false"
      >
        <template #checked> <icon-man />男</template>
        <template #unchecked><icon-woman /> 女 </template>
      </a-switch>
    </a-form-item>

    <a-form-item>
      <a-space>
        <a-button type="primary" @click="validate"> 保存 </a-button>
        <!-- <a-button type="secondary" @click="reset"> 重置 </a-button> -->
      </a-space>
    </a-form-item>
  </a-form>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { updateUserInfo } from '@/api/system/user';
  import { useUserStore } from '@/store';
  import Message from '@arco-design/web-vue/es/message';

  const userInfo = useUserStore();

  const formRef = ref<FormInstance>();
  const formData = ref<any>({
    nickname: userInfo.nickname,
    sex: userInfo.sex,
  });
  const validate = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      await updateUserInfo({ id: userInfo.id, ...formData.value });
      Message.success('基础信息更新成功');
      userInfo.info();
    }
  };
  const reset = async () => {
    await formRef.value?.resetFields();
  };
</script>

<style scoped lang="less">
  .form {
    max-width: 540px;
  }
</style>
