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
              <a-col :span="8">
                <a-form-item field="team_id" label="用户角色">
                  <a-cascader
                    v-model="searchForm.role_id"
                    check-strictly
                    allow-clear
                    :options="roles"
                    :field-names="{ value: 'id', label: 'name' }"
                    placeholder="请选择用户角色"
                    allow-search
                  />
                </a-form-item>
              </a-col>

              <a-col :span="8">
                <a-form-item field="team_id" label="用户部门">
                  <a-cascader
                    v-model="searchForm.team_id"
                    allow-clear
                    check-strictly
                    :options="teams"
                    :field-names="{ value: 'id', label: 'name' }"
                    placeholder="请选择用户部门"
                    allow-search
                  />
                </a-form-item>
              </a-col>

              <a-col :span="8">
                <a-form-item field="name" label="用户姓名">
                  <a-input
                    v-model="searchForm.name"
                    allow-clear
                    placeholder="请输入用户姓名"
                  />
                </a-form-item>
              </a-col>

              <a-col :span="8">
                <a-form-item field="phone" label="用户电话">
                  <a-input
                    v-model="searchForm.phone"
                    allow-clear
                    placeholder="请输入用户电话"
                  />
                </a-form-item>
              </a-col>

              <a-col :span="8">
                <a-form-item field="status" label="用户状态">
                  <a-select
                    v-model="searchForm.status"
                    allow-search
                    allow-clear
                    placeholder="用户状态"
                  >
                    <a-option value="true">启用</a-option>
                    <a-option value="false">禁用</a-option>
                  </a-select>
                </a-form-item>
              </a-col>

              <a-col :span="8">
                <a-form-item field="createdTime" label="创建时间">
                  <a-range-picker
                    v-model="searchForm.time"
                    style="width: 100%"
                    allow-clear
                  />
                </a-form-item>
              </a-col>
            </a-row>
          </a-form>
        </a-col>
        <a-divider style="height: 84px" direction="vertical" />
        <a-col :flex="'86px'" style="text-align: right">
          <a-space direction="vertical" :size="18">
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              搜索
            </a-button>
            <a-button @click="reset">
              <template #icon>
                <icon-refresh />
              </template>
              重置
            </a-button>
          </a-space>
        </a-col>
      </a-row>
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px; align-items: center">
        <a-col :span="12">
          <a-button
            v-permission="'system:user:add'"
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
          v-permission="'system:user:query'"
          row-key="id"
          :loading="loading"
          :pagination="false"
          :columns="(cloneColumns as TableColumnData[])"
          :data="renderData"
          :bordered="false"
          :size="size"
          :scroll="{
            x: 2000,
            y: 200,
          }"
        >
          <template #index="{ rowIndex }">
            {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
          </template>

          <template #role="{ record }">
            {{ record.role.name }}
          </template>

          <template #team="{ record }">
            {{ record.team.name }}
          </template>

          <template #avatar="{ record }">
            <a-avatar>
              <img
                v-if="record.avatar == ''"
                alt="avatar"
                :src="$logo ? $logo : logo"
              />
              <img alt="avatar" :src="$staticUrl + record.avatar" />
            </a-avatar>
          </template>

          <template #sex="{ record }">
            <a-tag v-if="record.sex" color="arcoblue">
              <template #icon><icon-man /></template>男
            </a-tag>
            <a-tag v-if="!record.sex" color="red">
              <template #icon><icon-woman /></template>女
            </a-tag>
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

          <template #lastLogin="{ record }">
            {{ $formatTime(record.last_login) }}
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
                v-permission="'system:user:update'"
                color="arcoblue"
                @click="handleUpdate(record)"
              >
                <template #icon><icon-edit /></template>修改
              </a-tag>
              <a-popconfirm
                content="您确认要删除此用户么？"
                type="warning"
                @ok="handleDelete(record.id)"
              >
                <a-tag v-permission="'system:user:delete'" color="red">
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
          field="name"
          label="用户姓名"
          :rules="[
            {
              required: true,
              message: '用户姓名是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.name"
            allow-clear
            placeholder="请输入用户姓名"
          />
        </a-form-item>

        <a-form-item
          field="phone"
          label="用户电话"
          :rules="[
            {
              required: true,
              message: '用户电话是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.phone"
            allow-clear
            placeholder="请输入用户电话"
          />
        </a-form-item>

        <a-form-item
          field="email"
          label="用户邮箱"
          :rules="[
            {
              required: true,
              message: '用户邮箱是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.email"
            allow-clear
            placeholder="请输入用户邮箱"
          />
        </a-form-item>

        <a-form-item
          field="sex"
          label="用户性别"
          :rules="[
            {
              required: true,
              message: '用户性别是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.sex">
            <a-radio :value="true">男</a-radio>
            <a-radio :value="false">女</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          field="role_id"
          label="用户角色"
          :rules="[
            {
              required: true,
              message: '用户角色是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-cascader
            v-model="submitForm.role_id"
            check-strictly
            :options="roles"
            placeholder="请选择用户角色"
            :field-names="{ value: 'id', label: 'name' }"
            allow-search
          />
        </a-form-item>

        <a-form-item
          field="team_id"
          label="用户部门"
          :rules="[
            {
              required: true,
              message: '用户部门是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-cascader
            v-model="submitForm.team_id"
            check-strictly
            :options="teams"
            :field-names="{ value: 'id', label: 'name' }"
            placeholder="请选择用户部门"
            allow-search
          />
        </a-form-item>

        <a-form-item
          field="status"
          label="用户状态"
          :rules="[
            {
              required: true,
              message: '用户状态是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-radio-group v-model="submitForm.status">
            <a-radio :value="true">启用</a-radio>
            <a-radio :value="false">禁用</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          v-if="isAddVisible"
          field="password"
          label="用户密码"
          :rules="[
            {
              required: true,
              message: '用户密码是必填项',
            },
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="submitForm.password"
            allow-clear
            placeholder="请输入用户密码"
          />
        </a-form-item>
        <a-form-item v-else field="password" label="用户密码">
          <a-input
            v-model="submitForm.password"
            allow-clear
            placeholder="请输入用户密码"
          />
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import {
    CascaderOption,
    Message,
    Modal,
    SelectOptionData,
  } from '@arco-design/web-vue';
  import { deepClone } from '@/utils';
  import { getUsers, addUser, updateUser, deleteUser } from '@/api/system/user';
  import { getTeams } from '@/api/system/team';
  import { getRoles } from '@/api/system/role';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';
  import logo from '@/assets/logo.png';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const newSearchForm = () => {
    interface SearchForm {
      role_id?: number;
      team_id?: number;
      name?: string;
      phone?: string;
      status?: string;
      time: string[];
    }
    return {} as SearchForm;
  };

  const modalVisible = ref<boolean>(false);
  const isAddVisible = ref<boolean>(false);
  const submitFormRef = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<any[]>([]);
  const roles = ref<SelectOptionData[]>([]);
  const teams = ref<CascaderOption[]>([]);
  const searchForm = ref(newSearchForm());
  const submitForm = ref<any>({});
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('medium');

  const pagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
  });

  const fetchRoles = async () => {
    const { data } = await getRoles();
    roles.value = [data];
  };

  const fetchTeams = async () => {
    const { data } = await getTeams();
    teams.value = [data];
  };

  fetchRoles();
  fetchTeams();

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '序号',
      dataIndex: 'index',
      slotName: 'index',
      width: 70,
    },
    {
      title: '用户角色',
      dataIndex: 'role',
      slotName: 'role',
    },

    {
      title: '用户部门',
      dataIndex: 'team',
      slotName: 'team',
    },
    {
      title: '用户昵称',
      dataIndex: 'nickname',
    },
    {
      title: '用户姓名',
      dataIndex: 'name',
    },
    {
      title: '用户电话',
      dataIndex: 'phone',
    },
    {
      title: '用户头像',
      dataIndex: 'avatar',
      slotName: 'avatar',
    },
    {
      title: '用户邮箱',
      dataIndex: 'email',
    },
    {
      title: '用户性别',
      dataIndex: 'sex',
      slotName: 'sex',
    },
    {
      title: '用户状态',
      dataIndex: 'status',
      slotName: 'status',
    },
    {
      title: '最后登录时间',
      dataIndex: 'last_login',
      slotName: 'lastLogin',
      width: 170,
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

      const { data, total } = await getUsers(req);
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

  const handleAdd = () => {
    submitForm.value = {};
    isAddVisible.value = true;
    modalVisible.value = true;
  };

  const handleDelete = async (id: number) => {
    await deleteUser({ id });
    await fetchData();
    Message.success('删除成功');
  };

  const handleUpdate = (data: any) => {
    submitForm.value = deepClone(data);
    isAddVisible.value = false;
    modalVisible.value = true;
  };

  const handleChangeStatus = (record: any) => {
    const status = record.status ? '启用' : '禁用';
    Modal.info({
      title: '状态变更提示',
      content: () => `您确认要 '${status}'此角色？`,
      closable: true,
      hideCancel: false,
      onOk: async () => {
        updateUser(record)
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
      await addUser(data);
      Message.success('新建成功');
    } else {
      await updateUser(data);
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
    name: 'User',
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
