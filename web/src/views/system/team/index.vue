<template>
  <div class="container">
    <Breadcrumb />
    <a-card class="general-card">
      <a-row style="margin-bottom: 16px; align-items: center">
        <a-col :span="12">
          <a-button
            v-permission="'system:team:add'"
            type="primary"
            @click="handleAdd(undefined)"
          >
            <template #icon>
              <icon-plus />
            </template>
            新建部门
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
            <div class="action-icon" @click="fetchData"
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
          <a-tooltip content="列设置">
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
      <a-table
        row-key="id"
        :loading="loading"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
      >
        <template #createdAt="{ record }">
          {{ $formatTime(record.created_at) }}
        </template>
        <template #updatedAt="{ record }">
          {{ $formatTime(record.updated_at) }}
        </template>

        <template #operations="{ record }">
          <a-space class="cursor-pointer">
            <a-tag
              v-permission="'system:team:add'"
              color="arcoblue"
              @click="handleAdd(record.id)"
            >
              <template #icon><icon-edit /></template>新建
            </a-tag>

            <a-tag
              v-if="record.id != 1"
              v-permission="'system:team:update'"
              color="orangered"
              @click="handleUpdate(record)"
            >
              <template #icon><icon-edit /></template>修改
            </a-tag>

            <a-popconfirm
              v-if="record.id != 1"
              v-permission="'system:team:delete'"
              content="您确认要删除此部门？"
              type="warning"
              @ok="handleDelete(record.id)"
            >
              <a-tag color="red">
                <template #icon><icon-delete /></template>删除
              </a-tag>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>

    <!--修改和新建弹窗表单-->
    <a-drawer
      v-model:visible="modalVisible"
      :title="isAddVisible ? '新建' : '修改'"
      width="380px"
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
          field="parent_id"
          label="上级部门"
          :rules="[
            {
              required: true,
              message: '上级部门是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-cascader
            v-model="submitForm.parent_id"
            check-strictly
            :field-names="{ value: 'id', label: 'name' }"
            :options="renderData"
            placeholder="请选择上级部门"
            allow-search
          />
        </a-form-item>

        <a-form-item
          field="name"
          label="部门名称"
          :rules="[
            {
              required: true,
              message: '部门名称是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.name"
            placeholder="请输入部门名称"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          field="description"
          label="部门描述"
          :rules="[
            {
              required: true,
              message: '部门描述是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-textarea
            v-model="submitForm.description"
            placeholder="请输入部门描述"
            allow-clear
          />
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Message } from '@arco-design/web-vue';
  import { deepClone } from '@/utils';
  import { getTeams, addTeam, updateTeam, deleteTeam } from '@/api/system/team';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const modalVisible = ref<boolean>(false);
  const isAddVisible = ref<boolean>(false);
  const submitFormRef = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<any[]>([]);
  const submitForm = ref<any>({});
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('large');

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '部门名称',
      dataIndex: 'name',
      slotName: 'name',
    },
    {
      title: '部门描述',
      dataIndex: 'description',
      slotName: 'description',
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      slotName: 'createdAt',
    },
    {
      title: '更新时间',
      dataIndex: 'updated_at',
      slotName: 'updatedAt',
    },
    {
      title: '操作人员',
      dataIndex: 'operator',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async (params?: any) => {
    setLoading(true);
    try {
      const { data } = await getTeams(params);
      renderData.value = [data];
    } finally {
      setLoading(false);
    }
  };

  fetchData();

  const handleSelectDensity = (
    val: string | number | Record<string, any> | undefined
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

  const handleAdd = (id?: number) => {
    submitForm.value = {};
    if (id) {
      submitForm.value.parent_id = id;
    }
    isAddVisible.value = true;
    modalVisible.value = true;
  };

  const handleDelete = async (id: number) => {
    await deleteTeam({ id });
    await fetchData();
    Message.success('删除成功');
  };

  const handleUpdate = (data: any) => {
    submitForm.value = deepClone(data);
    isAddVisible.value = false;
    modalVisible.value = true;
  };

  // 进行新增/修改请求
  const handleSubmitForm = async () => {
    const isError = await submitFormRef.value.validate();
    if (isError) {
      return false;
    }

    const data: any = deepClone(submitForm.value);
    if (isAddVisible.value) {
      await addTeam(data);
      Message.success('创建成功');
    } else {
      await updateTeam(data);
      Message.success('更新成功');
    }
    await fetchData();
    return true;
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
