<template>
  <a-modal
    v-model:visible="visible"
    :width="350"
    draggable
    :loading="loading"
    :footer="false"
    @cancel="onClose"
  >
    <template #title>修改{{ info.name }}</template>
    <a-form ref="formRef" :model="formData" auto-label-width autocomplete="off">
      <a-form-item
        field="captcha"
        label="验证码"
        :rules="[{ required: true, message: '验证码不能为空' }]"
        :validate-trigger="['change', 'blur']"
      >
        <a-input
          v-model="formData.captcha"
          placeholder="请输入验证码"
          allow-clear
          autocomplete="off"
        >
        </a-input>
        <a-button
          :style="{ marginLeft: '10px', width: '90px' }"
          type="primary"
          :disabled="!isSend"
          @click="onSendCode"
        >
          发送验证码
        </a-button>
      </a-form-item>

      <a-form-item
        v-if="visible"
        :label="'新' + info.name"
        :field="info.key"
        :rules="[{ required: true, message: '新' + info.name + '不能为空' }]"
        :validate-trigger="['change', 'blur']"
      >
        <a-input-password
          v-if="info.key == 'password'"
          v-model="formData.password"
          placeholder="请输入新密码"
          allow-clear
        />
        <a-input
          v-if="info.key == 'phone'"
          v-model="formData.phone"
          placeholder="请输入新手机号"
          allow-clear
          autocomplete="off"
        />

        <a-input
          v-if="info.key == 'email'"
          v-model="formData.email"
          placeholder="请输入新邮箱"
          allow-clear
          autocomplete="off"
        />
      </a-form-item>
      <a-form-item>
        <a-space :size="15" class="mb-4">
          <a-button type="primary" @click="onSave">提交</a-button>
          <a-button @click="onClose">关闭</a-button>
        </a-space>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
  import { ref, watch, onMounted, reactive } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { sendEmailCaptcha, updateUserByVerify } from '@/api/system/user';
  import { Message } from '@arco-design/web-vue';
  import useUserStore from '@/store/modules/user';

  interface EditProps {
    info: any;
    editVisible: boolean;
    setEditVisible: any;
  }

  const isSend = ref(true);
  const props = withDefaults(defineProps<EditProps>(), {
    editVisible: false,
    Info: {
      key: '',
      name: '',
    },
    setEditVisible: false,
    columns: [],
  });
  const formRef = ref<FormInstance>();
  const { loading, setLoading } = useLoading(true);
  const visible = ref(false);

  const formData = reactive({
    captcha_id: 0,
    password: '',
    captcha: '',
    email: '',
    phone: '',
  });
  watch(
    () => props.editVisible,
    (show: any) => {
      visible.value = show;
    }
  );
  // 发布
  const onSave = () => {
    updateUserByVerify(formData).then(() => {
      Message.success('密码修改成功，正在跳转中...');
      setTimeout(async () => {
        const userStore = useUserStore();
        await userStore.logout();
        window.location.reload();
      }, 2000);
    });
  };
  // 关闭
  const onClose = () => {
    props.setEditVisible();
  };
  // 发送验证码
  const onSendCode = async () => {
    const { data } = await sendEmailCaptcha({ name: 'user' });
    formData.captcha_id = data.id;
    Message.success('验证码发送成功');
    isSend.value = false;
    const timer = setTimeout(() => {
      isSend.value = true;
      clearTimeout(timer);
    }, data.expire * 1000);
  };
  onMounted(() => {
    setLoading(false);
  });
</script>
