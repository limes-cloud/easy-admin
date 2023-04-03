<template>
  <div class="container">
    <div class="logo">
      <img
        alt="logo"
        src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
      />
      <div class="logo-text">{{ appStore.title }}</div>
    </div>
    <div class="content">
      <div class="content-inner">
        <div class="login-form-wrapper">
          <div class="login-form-title">{{ appStore.title }}</div>
          <div class="login-form-sub-title">{{ appStore.desc }}</div>
          <div class="login-form-error-msg">{{ errorMessage }}</div>
          <a-form
            ref="loginForm"
            :model="userInfo"
            class="login-form"
            layout="vertical"
            @submit="handleSubmit"
          >
            <a-form-item
              field="username"
              :rules="[{ required: true, message: '账户不能为空' }]"
              :validate-trigger="['change', 'blur']"
              hide-label
            >
              <a-input
                v-model="userInfo.username"
                size="large"
                placeholder="请输入账户"
              >
                <template #prefix><icon-user /></template>
              </a-input>
            </a-form-item>
            <a-form-item
              field="password"
              :rules="[{ required: true, message: '密码不能为空' }]"
              :validate-trigger="['change', 'blur']"
              hide-label
            >
              <a-input-password
                v-model="userInfo.password"
                size="large"
                placeholder="请输入密码"
                allow-clear
                autocomplete
              >
                <template #prefix>
                  <icon-lock />
                </template>
              </a-input-password>
            </a-form-item>
            <a-form-item
              field="captcha"
              :rules="[{ required: true, message: '验证码不能为空' }]"
              :validate-trigger="['change', 'blur']"
              hide-label
            >
              <a-input
                v-model="userInfo.captcha"
                size="large"
                placeholder="请输入验证码"
                allow-clear
                autocomplete
              >
                <template #prefix>
                  <icon-lock />
                </template>
                <template #append>
                  <img
                    v-if="captchaData.base64"
                    width="100"
                    height="40"
                    :src="captchaData.base64"
                    @click="fetchCaptcha()"
                  />
                </template>
              </a-input>
            </a-form-item>
            <a-space :size="16" direction="vertical">
              <div class="login-form-password-actions">
                <a-checkbox
                  checked="rememberPassword"
                  :model-value="loginConfig.rememberPassword"
                  @change="setRememberPassword as any"
                >
                  记住密码
                </a-checkbox>
              </div>
              <a-button
                size="large"
                type="primary"
                html-type="submit"
                long
                :loading="loading"
              >
                确认登陆
              </a-button>
            </a-space>
          </a-form>
        </div>
      </div>

      <div class="footer">
        <Footer />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import Footer from '@/components/footer/index.vue';
  import { ref, reactive, onUnmounted } from 'vue';
  import { useRouter } from 'vue-router';
  import { Message } from '@arco-design/web-vue';
  import { ValidatedError } from '@arco-design/web-vue/es/form/interface';
  import { useStorage } from '@vueuse/core';
  import { useUserStore, useAppStore } from '@/store';
  import useLoading from '@/hooks/loading';
  import { captcha } from '@/api/system/user';

  const timeInter: any = ref(null);
  const router = useRouter();
  const errorMessage = ref('');
  const captchaData = ref({
    base64: '',
    id: '',
    expire: 30,
  });
  const { loading, setLoading } = useLoading();
  const userStore = useUserStore();
  const appStore = useAppStore();
  const loginConfig = useStorage('login-config', {
    rememberPassword: true,
    username: '',
    password: '',
  });

  onUnmounted(() => {
    clearInterval(timeInter.value);
    timeInter.value = null;
  });

  const userInfo = reactive({
    username: loginConfig.value.username,
    password: loginConfig.value.password,
    captcha: '',
    captcha_id: '',
  });

  const fetchCaptcha = async () => {
    const data = await captcha();
    captchaData.value = data.data;
    userInfo.captcha_id = captchaData.value.id;

    clearInterval(timeInter.value);
    // eslint-disable-next-line no-use-before-define
    autoFetchCaptcha();
  };

  const autoFetchCaptcha = async () => {
    timeInter.value = setInterval(() => {
      fetchCaptcha();
    }, captchaData.value.expire * 1000);
  };

  fetchCaptcha();

  const handleSubmit = async ({
    errors,
    values,
  }: {
    errors: Record<string, ValidatedError> | undefined;
    values: Record<string, any>;
  }) => {
    if (loading.value) return;
    if (!errors) {
      setLoading(true);
      try {
        await userStore.login(values);
        // 不引用导致不回跳转
        // eslint-disable-next-line @typescript-eslint/no-unused-vars

        const { redirect, ...othersQuery } = router.currentRoute.value.query;
        if (redirect && router.hasRoute(redirect as string)) {
          router.push({ name: redirect as string, query: { ...othersQuery } });
        } else {
          router.push({ path: '/', query: { ...othersQuery } });
        }

        Message.success('登陆成功');
        const { rememberPassword } = loginConfig.value;
        const { username, password } = values;
        loginConfig.value.username = rememberPassword ? username : '';
        loginConfig.value.password = rememberPassword ? password : '';
      } catch (err) {
        fetchCaptcha();
        errorMessage.value = (err as Error).message;
      } finally {
        setLoading(false);
      }
    }
  };
  const setRememberPassword = (value: boolean) => {
    loginConfig.value.rememberPassword = value;
  };
</script>

<style lang="less" scoped>
  .login-form {
    &-wrapper {
      width: 300px;
    }

    &-title {
      color: var(--color-text-1);
      font-weight: 500;
      font-size: 24px;
      line-height: 32px;
    }

    &-sub-title {
      color: var(--color-text-3);
      font-size: 16px;
      line-height: 24px;
    }

    &-error-msg {
      height: 32px;
      color: rgb(var(--red-6));
      line-height: 32px;
    }

    &-password-actions {
      display: flex;
      justify-content: space-between;
    }

    &-register-btn {
      color: var(--color-text-3) !important;
    }
  }
</style>

<style lang="less" scoped>
  .container {
    display: flex;
    height: 100vh;
    background: url(@/assets/images/login.jpg) center center fixed no-repeat;
    background-size: cover;

    .content {
      position: relative;
      display: flex;
      flex: 1;
      align-items: center;
      justify-content: flex-end;
      padding-bottom: 40px;
      margin-right: 100px;
    }

    .footer {
      position: absolute;
      right: 0;
      bottom: 0;
      width: 100%;
    }
  }

  .logo {
    position: fixed;
    top: 24px;
    left: 22px;
    z-index: 1;
    display: inline-flex;
    align-items: center;

    &-text {
      margin-right: 4px;
      margin-left: 4px;
      color: var(--color-fill-1);
      font-size: 20px;
    }
  }
</style>

<style lang="less" scoped>
  // responsive
  @media (max-width: @screen-lg) {
    .container {
      .content {
        justify-content: center;
        margin-right: 0px;
      }
    }
  }
</style>
