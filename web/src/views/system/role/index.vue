<template>
  <div class="container">
    <Breadcrumb />
    <a-card class="general-card">
      <a-row style="margin-bottom: 16px; align-items: center">
        <a-col :span="12">
          <a-button
            v-permission="'system:role:add'"
            type="primary"
            @click="handleAdd()"
          >
            <template #icon>
              <icon-plus />
            </template>
            创建
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
        v-permission="'system:role:query'"
        row-key="id"
        :loading="loading"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
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
            <template #checked> 启用 </template>
            <template #unchecked> 禁用 </template>
          </a-switch>
        </template>

        <template #createdAt="{ record }">
          {{ $formatTime(record.created_at) }}
        </template>
        <template #updatedAt="{ record }">
          {{ $formatTime(record.updated_at) }}
        </template>

        <template v-if="$hasPermission('system:role:menu')" #menu="{ record }">
          <a-tag
            v-if="record.keyword != 'superAdmin'"
            class="cursor-pointer"
            color="arcoblue"
            @click="handleUpdateMenu(record)"
          >
            <template #icon><icon-menu /></template> 菜单管理
          </a-tag>
        </template>

        <template #operations="{ record }">
          <a-space class="cursor-pointer">
            <a-tag
              v-permission="'system:role:add'"
              color="arcoblue"
              @click="handleAdd(record.id)"
            >
              <template #icon><icon-plus /></template> 新建
            </a-tag>

            <a-tag
              v-if="record.id != 1"
              v-permission="'system:role:update'"
              color="arcoblue"
              @click="handleUpdate(record)"
            >
              <template #icon><icon-edit /></template>修改
            </a-tag>
            <a-popconfirm
              v-permission="'system:role:delete'"
              content="您确认删除此角色"
              type="warning"
              @ok="handleDelete(record.id)"
            >
              <a-tag v-if="record.id != 1" color="red">
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
      :title="isAddVisible ? '创建' : '修改'"
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
          label="父角色"
          :rules="[
            {
              required: true,
              message: '父角色是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-cascader
            v-model="submitForm.parent_id"
            check-strictly
            :options="renderData"
            :field-names="{ value: 'id', label: 'name' }"
            placeholder="请选择父角色"
            allow-search
          />
        </a-form-item>
        <a-form-item
          field="name"
          label="角色名称"
          :rules="[
            {
              required: true,
              message: '角色名称是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.name"
            allow-clear
            placeholder="请输入角色名称"
          />
        </a-form-item>
        <a-form-item
          field="keyword"
          :label="'角色标识'"
          :rules="[
            {
              required: true,
              message: '角色标识是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.keyword"
            allow-clear
            placeholder="请输入角色标识"
          />
        </a-form-item>
        <a-form-item
          field="status"
          label="角色状态"
          :rules="[
            {
              required: true,
              message: '角色状态是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.status" :default-value="true">
            <a-radio :value="true">启用</a-radio>
            <a-radio :value="false">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item
          field="data_scope"
          label="数据权限"
          :rules="[
            {
              required: true,
              message: '数据权限是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-select
            v-model="submitForm.data_scope"
            allow-search
            placeholder="数据权限"
          >
            <a-option value="ALLTEAM">所在部门及子部门</a-option>
            <a-option value="CURTEAM">仅所在部门</a-option>
            <a-option value="DOWNTEAM">仅子部门</a-option>
            <a-option value="CUSTOM">自定义部门</a-option>
          </a-select>
        </a-form-item>

        <a-form-item
          v-if="submitForm.data_scope === 'CUSTOM'"
          field="team_ids"
          label="选择部门"
          :rules="[
            {
              required: true,
              message: '请选择部门',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <SelectTeam
            :keys="submitForm.team_ids"
            :data="teams"
            @select="
              (keys) => {
                submitForm.team_ids = keys;
              }
            "
          ></SelectTeam>
        </a-form-item>

        <a-form-item
          field="description"
          label="角色描述"
          :rules="[
            {
              required: true,
              message: '角色描述是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-textarea
            v-model="submitForm.description"
            allow-clear
            placeholder="请输入角色描述"
          />
        </a-form-item>
      </a-form>
    </a-drawer>

    <a-drawer
      v-model:visible="menuModalVisible"
      title="设置角色菜单权限"
      width="380px"
      @cancel="menuModalVisible = false"
      @before-ok="handleSubmitMenu"
    >
      <a-tree
        v-model:checked-keys="roleMenuIds"
        v-model:half-checked-keys="halfRoleMenuIds"
        :checkable="true"
        :data="roleMenus"
        :only-check-leaf="true"
        :field-names="{
          key: 'id',
          icon: 'icon_',
        }"
      />
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Message, Modal } from '@arco-design/web-vue';
  import { deepClone } from '@/utils';
  import {
    getRoles,
    addRole,
    updateRole,
    deleteRole,
    getRoleMenu,
    updateRoleMenu,
    getRoleMenuids,
  } from '@/api/system/role';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';
  import { getTeams } from '@/api/system/team';
  import SelectTeam from './components/SelectTeam.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const menuModalVisible = ref<boolean>(false);
  const modalVisible = ref<boolean>(false);
  const isAddVisible = ref<boolean>(false);
  const submitFormRef = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<any[]>([]);

  const roleMenus = ref<any[]>([]);
  const roleMenuIds = ref([]);
  const halfRoleMenuIds = ref([]);
  const teams = ref<any[]>([]);
  const submitForm = ref<any>({});
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('large');
  const curRoleId = ref();
  const basePagination: Pagination = {
    current: 1,
    pageSize: 20,
  };
  const pagination = reactive({
    ...basePagination,
  });

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '角色名称',
      dataIndex: 'name',
      slotName: 'name',
    },
    {
      title: '角色标识',
      dataIndex: 'keyword',
    },
    {
      title: '角色状态',
      dataIndex: 'status',
      slotName: 'status',
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
      title: '菜单管理',
      slotName: 'menu',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchTeams = async (params?: any) => {
    const { data } = await getTeams(params);
    teams.value = [data];
  };

  const fetchData = async (params?: any) => {
    setLoading(true);
    try {
      const { data } = await getRoles(params);
      renderData.value = [data];
      pagination.current = 1;
      pagination.total = data.length;
    } finally {
      setLoading(false);
    }
  };

  fetchData();
  fetchTeams();

  const handleSubmitMenu = async () => {
    let ids = roleMenuIds.value;
    ids = ids.concat(halfRoleMenuIds.value);
    await updateRoleMenu({ role_id: curRoleId.value, menu_ids: ids });
    Message.success('设置角色菜单成功');
    menuModalVisible.value = false;
    return true;
  };

  const handleUpdateMenu = async (item: any) => {
    curRoleId.value = item.id;
    const { data } = await getRoleMenu({ role_id: item.parent_id });
    roleMenus.value = [data];

    const idsInfo = await getRoleMenuids({ role_id: item.id });
    roleMenuIds.value = idsInfo.data;
    halfRoleMenuIds.value = [];
    menuModalVisible.value = true;
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

  const handleChangeStatus = (record: any) => {
    const status = record.status ? '启用' : '禁用';
    Modal.info({
      title: '状态变更提示',
      content: () => `您确认要 '${status}'此角色？`,
      closable: true,
      hideCancel: false,
      onOk: async () => {
        updateRole(record)
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

  const handleAdd = (id?: number) => {
    submitForm.value = {};
    if (id) {
      submitForm.value.parent_id = id;
    }
    submitForm.value.status = true;
    isAddVisible.value = true;
    modalVisible.value = true;
  };

  const handleDelete = async (id: number) => {
    await deleteRole({ id });
    await fetchData();
    Message.success('删除成功');
  };

  const handleUpdate = (data: any) => {
    submitForm.value = deepClone(data);
    if (data.data_scope === 'CUSTOM') {
      submitForm.value.team_ids = JSON.parse(data.team_ids);
    } else {
      submitForm.value.team_ids = [];
    }
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
    if (data.data_scope === 'CUSTOM') {
      data.team_ids = JSON.stringify(data.team_ids);
    } else {
      data.team_ids = null;
    }

    if (isAddVisible.value) {
      await addRole(data);
      Message.success('创建成功');
    } else {
      await updateRole(data);
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

<script lang="ts">
  export default {
    name: 'Role',
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
