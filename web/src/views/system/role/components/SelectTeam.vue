<template>
  <!-- <div>
   
  </div> -->
  <a-modal v-model:visible="visible" @ok="handleOk" @cancel="visible = false">
    <template #title> 选择部门 </template>
    <div class="content">
      <a-tree
        v-model:checked-keys="checkedKeys"
        :checkable="true"
        :check-strictly="true"
        :data="(data as any[])"
        :field-names="{
          key: 'id',
          title: 'name',
          children: 'children',
        }"
      />
    </div>
  </a-modal>
  <a-button type="outline" size="small" @click="visible = !visible"
    >已选择{{ keys.length }}个部门</a-button
  >
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';

  const props = defineProps({
    keys: {
      type: Array,
      default: (): number[] => {
        return [];
      },
    },
    data: {
      type: Array,
      default: () => {
        return [];
      },
    },
  });

  const checkedKeys = ref<number[]>([]);
  const visible = ref(false);

  const emit = defineEmits(['select']);
  const handleOk = () => {
    emit('select', checkedKeys);
  };

  onMounted(() => {
    checkedKeys.value = props.keys as number[];
  });
</script>

<style lang="less" scoped>
  .content {
    height: auto;
    max-height: 250px;
    overflow-y: scroll;
  }
</style>
