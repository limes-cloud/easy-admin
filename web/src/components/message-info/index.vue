<template>
  <div class="container">
    <a-spin :loading="loading" tip="内容加载中...">
      <div class="title">这是一个公告{{ renderData.title }}</div>
      <div class="sub-title">
        <a-space>
          <template v-for="(it, ind) in $noticeList" :key="ind">
            <a-tag
              v-if="it.key === renderData.type"
              size="small"
              class="tag"
              :color="it.color"
            >
              {{ it.name }}
            </a-tag>
          </template>
          <span class="time"
            >{{ renderData.operator }}于{{
              $formatTime(renderData.created_at)
            }}发布</span
          >

          <span
            v-if="renderData.created_at != renderData.updated_at"
            class="time"
            >于{{ $formatTime(renderData.created_at) }}修改</span
          >
        </a-space>
      </div>
      <div class="msg-box" v-html="renderData.content"> </div>
    </a-spin>
  </div>
</template>

<script lang="ts" setup>
  import { getNotice } from '@/api/system/notice';
  import useLoading from '@/hooks/loading';
  import { ref } from 'vue';

  const { loading, setLoading } = useLoading(true);

  const props = defineProps({
    id: Number,
  });

  const renderData = ref({
    title: undefined,
    operator: undefined,
    created_at: undefined,
    updated_at: undefined,
    content: undefined,
    type: undefined,
  });
  // fetchData 请求后端接口查询数据
  const fetchData = async () => {
    setLoading(true);
    try {
      const { data } = await getNotice({ id: props.id });
      renderData.value = data;
    } finally {
      setLoading(false);
    }
  };

  // 进入则请求
  fetchData();
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
    .title {
      display: block;
      cursor: pointer;
      color: #000;
      font-size: 16px;
      font-weight: 700;
      margin-bottom: 10px;
    }
    .sub-title {
      .operator,
      .time {
        color: #777;
      }
      margin-bottom: 15px;
      font-size: 13px;
    }

    .msg-box {
    }
  }
</style>
