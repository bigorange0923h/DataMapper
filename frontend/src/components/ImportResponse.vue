<script setup>
const columns = [
  {
    title: '原始入参名称',
    dataIndex: 'sourceName',
    key: 'sourceName',
  },
  {
    title: '原始入参类型',
    dataIndex: 'sourceType',
    key: 'sourceType',
  },
  {
    title: '目标参数名称',
    dataIndex: 'targetName',
    key: 'targetName',
  },
  {
    title: '目标参数类型',
    dataIndex: 'targetType',
    key: 'targetType',
  },
  {
    title: 'Action',
    key: 'action',
  },
];
const data = [
  {
    sourceName: 'name',
    sourceType: 'string',
    targetName: 'name',
    targetType: 'string',
  },
  {
    sourceName: 'age',
    sourceType: 'integer',
    targetName: 'age',
    targetType: 'integer',
  },
  {
    sourceName: 'sex',
    sourceType: 'integer',
    targetName: 'sex',
    targetType: 'integer',
  },
  {
    sourceName: 'nickName',
    sourceType: 'string',
    targetName: 'nickName',
    targetType: 'string',
  },
  {
    sourceName: 'mobile',
    sourceType: 'string',
    targetName: 'mobile',
    targetType: 'string',
  } ,
  {
    sourceName: 'phone1',
    sourceType: 'string',
    targetName: 'mobile1',
    targetType: 'phone1',
  }

]
</script>

<template>
  <a-table :columns="columns" :data-source="data" :pagination="false" :scroll="{  y: 300 }"/>
</template>

<style scoped>

</style>