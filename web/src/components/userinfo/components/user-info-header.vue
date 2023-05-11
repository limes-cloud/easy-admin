<template>
  <div class="header">
    <a-space :size="12" direction="vertical" align="center">
      <Upload
        :file="userInfo.avatar"
        :multiple="false"
        :size="100"
        shape="circle"
        dir="avatar"
        @confirm="uploadCallback"
      ></Upload>

      <a-typography-title :heading="6" style="margin: 0">
        {{ userInfo.name }} - {{ userInfo.nickname }}
      </a-typography-title>
      <div class="user-msg">
        <a-space :size="18">
          <div>
            <icon-user-group />
            <a-typography-text>{{ userInfo.team.name }}</a-typography-text>
          </div>
          <div>
            <icon-safe />
            <a-typography-text> {{ userInfo.role.name }} </a-typography-text>
          </div>
          <div>
            <icon-email />
            <a-typography-text> {{ userInfo.email }} </a-typography-text>
          </div>
        </a-space>
      </div>
    </a-space>
  </div>
</template>

<script lang="ts" setup>
  import { updateUserInfo } from '@/api/system/user';
  import Upload from '@/components/upload/index.vue';

  import { useUserStore } from '@/store';
  import { Message } from '@arco-design/web-vue';

  const userInfo = useUserStore();

  const uploadCallback = async (files: any[]) => {
    // todo更新用户信息
    if (files.length) {
      console.log({ id: userInfo.id, avatar: files[0].url });
      await updateUserInfo({ id: userInfo.id, avatar: files[0].url });
      Message.success('更换头像成功');
      // userInfo.info();
    }
  };
</script>

<style scoped lang="less">
  .header {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 204px;
    color: var(--gray-10);
    background: url('@/assets/images/user-bg.png') no-repeat;
    background-size: cover;
    border-radius: 4px;

    :deep(.arco-avatar-trigger-icon-button) {
      color: rgb(var(--arcoblue-6));

      :deep(.arco-icon) {
        vertical-align: -1px;
      }
    }
    .user-msg {
      .arco-icon {
        color: rgb(var(--gray-10));
      }
      .arco-typography {
        margin-left: 6px;
      }
    }
  }
</style>
