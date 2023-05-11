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
            auto-label-width
          >
            <a-row :gutter="16">
              <a-col :span="6">
                <a-form-item field="title" label="通知标题">
                  <a-input
                    v-model="searchForm.title"
                    allow-clear
                    placeholder="请输入通知标题"
                  />
                </a-form-item>
              </a-col>

              <a-col :span="6">
                <a-form-item field="status" label="通知状态">
                  <a-select
                    v-model="searchForm.status"
                    allow-search
                    allow-clear
                    placeholder="通知状态"
                  >
                    <a-option value="true">启用</a-option>
                    <a-option value="false">禁用</a-option>
                  </a-select>
                </a-form-item>
              </a-col>

              <a-col :span="6">
                <a-form-item field="createdTime" label="创建时间">
                  <a-range-picker
                    v-model="searchForm.time"
                    style="width: 100%"
                    allow-clear
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
        <a-col :span="12">
          <a-button
            v-permission="'system:notice:add'"
            type="primary"
            @click="handleAdd"
          >
            <template #icon>
              <icon-plus />
            </template>
            新建
          </a-button>
        </a-col>
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
          row-key="id"
          :loading="loading"
          :pagination="pagination"
          :columns="(cloneColumns as TableColumnData[])"
          :data="renderData"
          :bordered="false"
          :size="size"
          @page-change="onPageChange"
        >
          <template #index="{ rowIndex }">
            {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
          </template>

          <template #status="{ record }">
            <a-switch
              v-model="record.status"
              type="round"
              @change="handleChangeStatus(record)"
            >
              <template #checked> 发布 </template>
              <template #unchecked> 私密 </template>
            </a-switch>
          </template>

          <template #createdAt="{ record }">
            {{ $formatTime(record.created_at) }}
          </template>
          <template #updatedAt="{ record }">
            {{ $formatTime(record.updated_at) }}
          </template>

          <template #operations="{ record }">
            <a-space class="cursor-pointer">
              <a-tag color="arcoblue" @click="handlePreview(record)">
                <template #icon><icon-eye /></template>查看
              </a-tag>
              <a-tag
                v-permission="'system:notice:update'"
                color="orangered"
                @click="handleUpdate(record)"
              >
                <template #icon><icon-edit /></template>修改
              </a-tag>
              <a-popconfirm
                content="您确认要删除此用户么？"
                type="warning"
                @ok="handleDelete(record.id)"
              >
                <a-tag v-permission="'system:notice:delete'" color="red">
                  <template #icon><icon-delete /></template>删除
                </a-tag>
              </a-popconfirm>
            </a-space>
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

    <!--修改和新建弹窗表单-->
    <a-drawer
      v-model:visible="modalVisible"
      :title="isAddVisible ? '新建通知' : '修改通知'"
      width="680px"
      @cancel="modalVisible = false"
      @before-ok="handleSubmitForm"
    >
      <a-form
        ref="submitFormRef"
        :model="submitForm"
        label-align="left"
        layout="horizontal"
        auto-label-width
      >
        <a-form-item
          field="title"
          label="通知标题"
          :rules="[
            {
              required: true,
              message: '通知标题是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.title"
            allow-clear
            placeholder="请输入通知标题"
          />
        </a-form-item>

        <a-form-item
          field="type"
          label="通知类型"
          :rules="[
            {
              required: true,
              message: '通知类型是必填项',
            },
          ]"
        >
          <a-select
            v-model="submitForm.type"
            allow-search
            allow-clear
            placeholder="请选择通知类型"
          >
            <a-option
              v-for="(item, index) in $noticeList"
              :key="index"
              :value="item.key"
              >{{ item.name }}</a-option
            >
          </a-select>
        </a-form-item>

        <a-form-item
          field="content"
          label="通知标题"
          :rules="[
            {
              required: true,
              message: '通知标题是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <tinymce
            v-model="submitForm.content"
            allow-clear
            placeholder="请输入通知标题"
          />
        </a-form-item>
      </a-form>
    </a-drawer>

    <a-drawer
      v-model:visible="perviewVisible"
      title="查看通知"
      width="600px"
      @cancel="perviewVisible = false"
      @before-ok="perviewVisible = false"
    >
      <MessageInfo
        v-if="perviewVisible && perviewId != 0"
        :id="perviewId"
      ></MessageInfo>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Message, Modal } from '@arco-design/web-vue';
  import { deepClone } from '@/utils';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';
  import tinymce from '@/components/tinymce/index.vue';
  import {
    getNotices,
    getNotice,
    addNotice,
    updateNotice,
    deleteNotice,
  } from '@/api/system/notice';
  import MessageInfo from '@/components/message-info/index.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const newSearchForm = () => {
    return {
      title: undefined,
      time: [],
      status: undefined,
      is_mine: undefined,
    };
  };

  const perviewVisible = ref<boolean>(false);
  const modalVisible = ref<boolean>(false);
  const isAddVisible = ref<boolean>(false);
  const submitFormRef = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<any[]>([]);
  const searchForm = ref(newSearchForm());
  const submitForm = ref<any>({});
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('large');
  const perviewId = ref<number>(0);
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
      width: 50,
    },
    {
      title: '通知标题',
      dataIndex: 'title',
    },

    {
      title: '通知类型',
      dataIndex: 'type',
    },
    {
      title: '通知状态',
      dataIndex: 'status',
      slotName: 'status',
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      slotName: 'createdAt',
      width: 170,
    },
    {
      title: '更新时间',
      dataIndex: 'updated_at',
      slotName: 'updatedAt',
      width: 170,
    },
    {
      title: '操作人员',
      dataIndex: 'operator',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 200,
    },
  ]);

  // fetchData 请求后端接口查询数据
  const fetchData = async () => {
    setLoading(true);
    try {
      const req: any = { ...searchForm.value };
      req.page = pagination.current;
      req.page_size = pagination.pageSize;

      req.status = undefined;
      if (searchForm.value.status) {
        req.status = searchForm.value.status === 'true';
      }

      req.is_mine = undefined;
      if (searchForm.value.is_mine) {
        req.status = searchForm.value.is_mine === 'true';
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

      const { data, total } = await getNotices(req);
      renderData.value = data;
      pagination.total = total;
    } finally {
      setLoading(false);
    }
  };

  // 进入则请求
  fetchData();

  // search 数据搜索
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

  const handleAdd = () => {
    submitForm.value = {};
    isAddVisible.value = true;
    modalVisible.value = true;
  };

  const handleDelete = async (id: number) => {
    await deleteNotice({ id });
    await fetchData();
    Message.success('删除成功');
  };

  const handleUpdate = async (req: any) => {
    // submitForm.value = deepClone(data);
    const { data } = await getNotice({ id: req.id });
    submitForm.value = data;
    isAddVisible.value = false;
    modalVisible.value = true;
  };

  const handleChangeStatus = (record: any) => {
    const status = record.status ? '发布' : '私密';
    Modal.info({
      title: '状态变更提示',
      content: () => `您确认要设置此公告为 '${status}'状态？`,
      closable: true,
      hideCancel: false,
      onOk: async () => {
        updateNotice(record)
          .then(() => {
            Message.success('更新成功');
          })
          .catch(() => {
            record.status = !record.status;
          });
      },
      onCancel: () => {
        record.status = !record.status;
      },
    });
  };

  // 进行新增/修改请求
  const handleSubmitForm = async () => {
    const isError = await submitFormRef.value.validate();
    if (isError) {
      return false;
    }

    const data: any = deepClone(submitForm.value);
    if (isAddVisible.value) {
      await addNotice(data);
      Message.success('新建成功');
    } else {
      await updateNotice(data);
      Message.success('更新成功');
    }
    await fetchData();
    return true;
  };

  const handlePreview = (data: any) => {
    perviewId.value = data.id;
    perviewVisible.value = true;
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
    name: 'Notice',
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
