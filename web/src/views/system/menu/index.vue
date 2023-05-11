<template>
  <div class="container">
    <Breadcrumb />
    <a-card class="general-card">
      <a-row style="margin-bottom: 16px; align-items: center">
        <a-col :span="12">
          <a-button
            v-permission="'system:menu:add'"
            type="primary"
            @click="handleAdd(undefined)"
          >
            <template #icon>
              <icon-plus />
            </template>
            新建菜单
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
      <a-table
        v-permission="'system:menu:query'"
        row-key="id"
        :loading="loading"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
      >
        <template #title="{ record }">
          <a-space>
            <icon v-if="record.icon" :type="record.icon" :size="20" />
            {{ record.title }}
          </a-space>
        </template>

        <template #isHidden="{ record }">
          <a-tag v-if="record.is_hidden" color="red">隐藏</a-tag>
          <a-tag v-else color="green">显示</a-tag>
        </template>

        <template #type="{ record }">
          <a-tag v-if="record.type === 'R'" color="arcoblue">根菜单</a-tag>
          <a-tag v-if="record.type === 'M'" color="arcoblue">菜单</a-tag>
          <a-tag v-if="record.type === 'A'" color="orange"
            >接口｜{{ record.method }}</a-tag
          >
          <a-tag v-if="record.type === 'G'" color="green">接口组</a-tag>
        </template>

        <template #createdAt="{ record }">
          {{ $formatTime(record.created_at) }}
        </template>
        <template #updatedAt="{ record }">
          {{ $formatTime(record.updated_at) }}
        </template>

        <template #operations="{ record }">
          <a-space class="cursor-pointer">
            <a-tag
              v-permission="'system:menu:add'"
              color="arcoblue"
              @click="handleAdd(record.id)"
            >
              <template #icon><icon-plus /></template> 新建
            </a-tag>

            <a-tag
              v-if="record.id != 1"
              v-permission="'system:menu:update'"
              color="orangered"
              @click="handleUpdate(record)"
            >
              <template #icon><icon-edit /></template> 修改
            </a-tag>

            <a-popconfirm
              v-if="record.id != 1"
              v-permission="'system:menu:delete'"
              content="您确认删除此菜单"
              type="warning"
              @ok="handleDelete(record.id)"
            >
              <a-tag color="red">
                <template #icon><icon-delete /></template> 删除
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
          label="父菜单"
          :rules="[
            {
              required: true,
              message: '父菜单是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-cascader
            v-model="submitForm.parent_id"
            check-strictly
            :options="renderData"
            :field-names="{ value: 'id', label: 'title' }"
            placeholder="请选择父菜单"
            allow-search
          />
        </a-form-item>

        <a-form-item
          field="title"
          label="标题"
          :rules="[
            {
              required: true,
              message: '标题是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.title"
            placeholder="请输入菜单标题"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          field="weight"
          label="菜单权重"
          :rules="[
            {
              required: true,
              message: '菜单权重是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input-number
            v-model="submitForm.weight"
            placeholder="请输入菜单权重"
            :default-value="0"
            mode="button"
          />
        </a-form-item>

        <a-form-item
          field="type"
          label="菜单类型"
          :rules="[
            {
              required: true,
              message: '菜单类型是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-select
            v-model="submitForm.type"
            :options="menuTypes"
            placeholder="请选择菜单类型"
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="name"
          label="菜单标识"
          :rules="[
            {
              required: true,
              message: '菜单标识是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.name"
            placeholder="请输入菜单标识"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="icon"
          label="菜单图标"
          :rules="[
            {
              required: true,
              message: '菜单图标是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-select
            v-model="submitForm.icon"
            allow-search
            placeholder="请选择菜单图标"
          >
            <template v-for="item in iconOptions" :key="item">
              <a-option>
                <a-space align="center">
                  <icon :type="item" :size="18" />{{ item }}
                </a-space>
              </a-option>
            </template>
          </a-select>
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'A' || submitForm.type === 'BA'"
          field="method"
          label="请求方法"
          :validate-trigger="['change', 'input']"
          :rules="[
            {
              required: true,
              message: '请求方法是必填项',
            },
          ]"
        >
          <a-radio-group v-model="submitForm.method">
            <a-radio value="GET">GET</a-radio>
            <a-radio value="POST">POST</a-radio>
            <a-radio value="PUT">PUT</a-radio>
            <a-radio value="DELETE">DELETE</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          v-if="submitForm.type !== 'M' && submitForm.type !== 'BA'"
          field="permission"
          label="权限指令"
          :rules="[
            {
              required: true,
              message: '权限指令是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.permission"
            placeholder="请输入权限指令"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type !== 'G'"
          field="path"
          :label="
            submitForm.type === 'A' || submitForm.type === 'BA'
              ? '接口路径'
              : '菜单路径'
          "
          :rules="[
            {
              required: true,
              message: '菜单路径是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.path"
            placeholder="请输入菜单路径"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="component"
          label="菜单组件"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.component"
            placeholder="请输入菜单组件"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="redirect"
          label="跳转路由"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.redirect"
            placeholder="请输入跳转路由"
            allow-clear
          />
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="hidden"
          label="是否隐藏"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.is_hidden" :default-value="false">
            <a-radio :value="false">否</a-radio>
            <a-radio :value="true">是</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="cache"
          label="是否缓存"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.is_cache" :default-value="false">
            <a-radio :value="false">否</a-radio>
            <a-radio :value="true">是</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          v-if="submitForm.type === 'M'"
          field="is_home"
          label="设为首页"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.is_home" :default-value="false">
            <a-radio :value="false">否</a-radio>
            <a-radio :value="true">是</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Message, SelectOptionData } from '@arco-design/web-vue';
  import { deepClone } from '@/utils';
  import { getMenus, addMenu, updateMenu, deleteMenu } from '@/api/system/menu';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';

  import icons from '@/utils/icon';
  import icon from './components/icon.vue';

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

  const iconOptions = computed(() => icons);

  const menuTypes = computed<SelectOptionData[]>(() => [
    {
      label: '菜单',
      value: 'M',
    },
    {
      label: '接口',
      value: 'A',
    },
    {
      label: '接口组',
      value: 'G',
    },
    {
      label: '基础接口',
      value: 'BA',
    },
  ]);

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '菜单标题',
      dataIndex: 'locale',
      slotName: 'title',
    },
    {
      title: '菜单路由',
      dataIndex: 'path',
      slotName: 'path',
    },
    {
      title: '菜单类型',
      dataIndex: 'type',
      slotName: 'type',
    },
    {
      title: '是否隐藏',
      dataIndex: 'is_hidden',
      slotName: 'isHidden',
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
      const { data } = await getMenus(params);
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
    submitForm.value.weight = 0;
    isAddVisible.value = true;
    modalVisible.value = true;
  };

  const handleDelete = async (id: number) => {
    await deleteMenu({ id });
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
      await addMenu(data);
      Message.success('创建成功');
    } else {
      await updateMenu(data);
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
