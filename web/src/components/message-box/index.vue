<template>
  <a-spin style="display: block" :loading="loading">
    <a-tabs
      v-model:activeKey="type"
      class="tabs-start"
      type="capsule"
      :justify="true"
      destroy-on-hide
      @change="fetchData()"
    >
      <a-tab-pane v-for="item in tabList" :key="item.key">
        <template #title>
          <span> {{ item.title }}</span>
        </template>
        <a-result v-if="!list.length" status="404">
          <template #subtitle> 暂无消息</template>
        </a-result>
        <div v-for="(ite, index) in list" :key="index" class="message">
          <div class="title" @click="handlePreview(ite)">{{ ite.title }}</div>
          <a-space>
            <template v-for="(it, ind) in $noticeList" :key="ind">
              <a-tag
                v-if="it.key === ite.type"
                size="small"
                class="tag"
                :color="it.color"
              >
                {{ it.name }}
              </a-tag>
            </template>
            <span class="operator">{{ ite.operator }}</span>
            <span class="time">{{ $formatTime(ite.created_at) }}</span>
          </a-space>
        </div>
      </a-tab-pane>
      <!-- <template #extra>
        <a-button
          style="margin-right: 10px"
          type="primary"
          @click="handleAddNotice"
          >发送通知</a-button
        >
      </template> -->
    </a-tabs>
  </a-spin>
  <a-drawer
    v-model:visible="perviewVisible"
    title="查看通知"
    width="600px"
    :hide-cancel="true"
    ok-text="关闭"
    @cancel="perviewVisible = false"
    @before-ok="perviewVisible = false"
  >
    <MessageInfo v-if="perviewId != 0" :id="perviewId"></MessageInfo>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import useLoading from '@/hooks/loading';
  import { getNotices } from '@/api/system/notice';
  import MessageInfo from '@/components/message-info/index.vue';

  interface TabItem {
    key: string;
    title: string;
  }
  const { loading, setLoading } = useLoading(true);
  const type = ref('unread');
  const perviewVisible = ref<boolean>(false);
  const perviewId = ref<number>(0);

  const tabList: TabItem[] = [
    {
      key: 'unread',
      title: '未读消息',
    },
    {
      key: 'read',
      title: '已读消息',
    },
  ];

  const list = ref<any>([]);

  const query = ref({
    page: 1,
    page_size: 10,
    is_read: false,
    status: true,
  });

  async function fetchData() {
    setLoading(true);
    if (query.value.is_read !== (type.value === 'unread')) {
      query.value.page = 1;
      query.value.page_size = 10;
    }
    query.value.is_read = type.value !== 'unread';
    try {
      const { data } = await getNotices(query.value);
      list.value = data;
    } finally {
      setLoading(false);
    }
  }

  fetchData();

  const handlePreview = (data: any) => {
    perviewId.value = data.id;
    perviewVisible.value = true;
  };
</script>

<style scoped lang="less">
  .message {
    padding: 10px 20px;
    margin-bottom: 5px;
    .title {
      display: block;
      cursor: pointer;
      color: #000;
      font-weight: 700;
      &:hover {
        color: rgb(var(--blue-6));
      }
    }

    .time,
    .operator {
      font-size: 13px;
    }
  }

  .tabs-start {
    .arco-tabs-nav-type-capsule {
      justify-content: flex-start !important;
    }
  }
  :deep(.arco-popover-popup-content) {
    padding: 0;
  }

  :deep(.arco-list-item-meta) {
    align-items: flex-start;
  }
  :deep(.arco-tabs-nav) {
    padding: 14px 0 12px 16px;
    border-bottom: 1px solid var(--color-neutral-3);
  }
  :deep(.arco-tabs-content) {
    padding-top: 0;
    .arco-result-subtitle {
      color: rgb(var(--gray-6));
    }
  }
</style>
