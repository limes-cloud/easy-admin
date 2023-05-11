<template>
  <div class="container">
    <Breadcrumb />
    <a-card class="general-card">
      <a-row>
        <a-col :flex="1">
          <a-form
            :model="searchForm"
            :label-col-props="{ span: 6 }"
            :wrapper-col-props="{ span: 18 }"
            label-align="left"
            layout="horizontal"
            auto-label-width
          >
            <a-row :gutter="16">
              <a-col :span="6">
                <a-form-item field="phone" label="手机号">
                  <a-input
                    v-model="searchForm.phone"
                    placeholder="请输入手机号"
                    allow-clear
                  />
                </a-form-item>
              </a-col>
              <a-col :span="6">
                <a-form-item field="status" label="登陆状态">
                  <a-select
                    v-model="searchForm.status"
                    allow-search
                    allow-clear
                    placeholder="请选择登陆状态"
                  >
                    <a-option value="true">成功</a-option>
                    <a-option value="false">失败</a-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :span="6">
                <a-form-item field="time" label="登陆时间">
                  <a-range-picker
                    v-model="searchForm.time"
                    allow-clear
                    style="width: 100%"
                  />
                </a-form-item>
              </a-col>
              <a-col :span="6">
                <a-space>
                  <a-button type="primary" @click="search">
                    <template #icon>
                      <icon-search />
                    </template>
                    搜索
                  </a-button>
                  <a-button class="ml-15" @click="reset">
                    <template #icon>
                      <icon-refresh />
                    </template>
                    重置
                  </a-button>
                </a-space>
              </a-col>
            </a-row>
          </a-form>
        </a-col>
      </a-row>
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px; align-items: center">
        <a-col
          :span="12"
          style="
            display: flex;
            flex: 1;
            align-items: center;
            justify-content: end;
          "
        >
          <a-tooltip content="刷新">
            <div class="action-icon" @click="search"
              ><icon-refresh size="18"
            /></div>
          </a-tooltip>
          <a-dropdown @select="handleSelectDensity">
            <a-tooltip content="字体大小">
              <div class="action-icon"><icon-line-height size="18" /></div>
            </a-tooltip>
            <template #content>
              <a-doption
                v-for="item in $densityList"
                :key="item.value"
                :value="item.value"
                :class="{ active: item.value === size }"
              >
                <span>{{ item.name }}</span>
              </a-doption>
            </template>
          </a-dropdown>
          <a-tooltip content="行设置">
            <a-popover
              trigger="click"
              position="bl"
              @popup-visible-change="popupVisibleChange"
            >
              <div class="action-icon"><icon-settings size="18" /></div>
              <template #content>
                <div id="tableSetting">
                  <div
                    v-for="(item, index) in showColumns"
                    :key="item.dataIndex"
                    class="setting"
                  >
                    <div style="margin-right: 4px; cursor: move">
                      <icon-drag-arrow />
                    </div>
                    <div>
                      <a-checkbox
                        v-model="item.checked"
                        @change="
                          handleChange($event, item as TableColumnData, index)
                        "
                      >
                      </a-checkbox>
                    </div>
                    <div class="title">
                      {{ item.title === '#' ? 'index' : item.title }}
                    </div>
                  </div>
                </div>
              </template>
            </a-popover>
          </a-tooltip>
        </a-col>
      </a-row>
      <a-space direction="vertical" fill>
        <a-table
          v-permission="'system:login:log:query'"
          row-key="id"
          :loading="loading"
          :pagination="false"
          :columns="(cloneColumns as TableColumnData[])"
          :data="renderData"
          :bordered="false"
          :size="size"
        >
          <template #index="{ rowIndex }">
            {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
          </template>

          <template #status="{ record }">
            <a-tag v-if="record.status" color="green">登陆成功</a-tag>
            <a-tag v-else color="red">登陆失败</a-tag>
          </template>

          <template #createdAt="{ record }">
            {{ $formatTime(record.created_at) }}
          </template>
        </a-table>
        <a-pagination
          :total="pagination.total"
          :current="pagination.current"
          :page-size="pagination.pageSize"
          show-total
          show-jumper
          show-page-size
          @change="onPageChange"
          @page-size-change="onPageSizeChange"
        />
      </a-space>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { getLoginLog } from '@/api/system/login_log';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const newSearchForm = () => {
    return {
      phone: undefined,
      time: [],
      status: '',
    };
  };

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<any[]>([]);
  const searchForm = ref(newSearchForm());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('large');

  const pagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
  });

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '序号',
      dataIndex: 'index',
      slotName: 'index',
    },
    {
      title: '手机号',
      dataIndex: 'phone',
    },
    {
      title: 'ip地址',
      dataIndex: 'ip',
    },
    {
      title: '登陆地址',
      dataIndex: 'address',
    },
    {
      title: '浏览器',
      dataIndex: 'browser',
    },
    {
      title: '设备类型',
      dataIndex: 'device',
    },
    {
      title: '登陆状态',
      dataIndex: 'status',
      slotName: 'status',
    },
    {
      title: '登陆备注',
      dataIndex: 'description',
    },
    {
      title: '登陆时间',
      dataIndex: 'created_at',
      slotName: 'createdAt',
    },
  ]);

  const fetchData = async () => {
    setLoading(true);

    try {
      // 组装请求参数
      const req: any = { ...searchForm.value };
      req.page = pagination.current;
      req.page_size = pagination.pageSize;

      req.status = undefined;
      if (searchForm.value.status) {
        req.status = searchForm.value.status === 'true';
      }
      if (searchForm.value.time && searchForm.value.time.length === 1) {
        req.start = Math.floor(
          new Date(searchForm.value.time[0]).getTime() / 1000
        );
      }
      if (searchForm.value.time && searchForm.value.time.length === 2) {
        req.start = Math.floor(
          new Date(searchForm.value.time[0]).getTime() / 1000
        );
        req.end = Math.floor(
          new Date(searchForm.value.time[1]).getTime() / 1000
        );
      }

      const { data, total } = await getLoginLog(req);
      renderData.value = data;
      pagination.total = total;
    } finally {
      setLoading(false);
    }
  };

  fetchData();

  const search = () => {
    pagination.current = 1;
    pagination.pageSize = 10;
    fetchData();
  };

  const onPageChange = (current: number) => {
    pagination.current = current;
    fetchData();
  };

  const onPageSizeChange = (size: number) => {
    pagination.pageSize = size;
    fetchData();
  };

  const reset = () => {
    searchForm.value = newSearchForm();
  };

  const handleSelectDensity = (
    val: string | number | Record<string, any> | undefined,
    e: Event
  ) => {
    size.value = val as SizeProps;
  };

  const handleChange = (
    checked: boolean | (string | boolean | number)[],
    column: Column,
    index: number
  ) => {
    if (!checked) {
      cloneColumns.value = showColumns.value.filter(
        (item) => item.dataIndex !== column.dataIndex
      );
    } else {
      cloneColumns.value.splice(index, 0, column);
    }
  };

  const exchangeArray = <T extends Array<any>>(
    array: T,
    beforeIdx: number,
    newIdx: number,
    isDeep = false
  ): T => {
    const newArray = isDeep ? cloneDeep(array) : array;
    if (beforeIdx > -1 && newIdx > -1) {
      // 先替换后面的，然后拿到替换的结果替换前面的
      newArray.splice(
        beforeIdx,
        1,
        newArray.splice(newIdx, 1, newArray[beforeIdx]).pop()
      );
    }
    return newArray;
  };

  const popupVisibleChange = (val: boolean) => {
    if (val) {
      nextTick(() => {
        const el = document.getElementById('tableSetting') as HTMLElement;
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        const sortable = new Sortable(el, {
          onEnd(e: any) {
            const { oldIndex, newIndex } = e;
            exchangeArray(cloneColumns.value, oldIndex, newIndex);
            exchangeArray(showColumns.value, oldIndex, newIndex);
          },
        });
      });
    }
  };

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      cloneColumns.value.forEach((item, index) => {
        item.checked = true;
      });
      showColumns.value = cloneDeep(cloneColumns.value);
    },
    { deep: true, immediate: true }
  );
</script>

<script lang="ts">
  export default {
    name: 'Locale',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
</style>
