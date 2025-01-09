<script setup>
import {LoadMappingConfig, SelectExcelFile} from '../../wailsjs/go/main/App'

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
let data = [];
let filePath ="";

async function selectFile() {
  filePath = await SelectExcelFile();
  // todo 基于文件路径提前获取对应文件头,并且加载配置表格
  if(filePath) {
    LoadMappingConfig(filePath).then(result=>{
      console.log(result)
      data = result
    })
  }

}
</script>

<template>
  <a-space style="margin-top: 10px">
    <p v-if="filePath" style="color: #141414">选择的文件路径: {{ filePath }}</p>
    <a-button @click="selectFile">选择文件</a-button>
  </a-space>

  <a-divider />
  <a-table :columns="columns" :data-source="data" :pagination="false" :scroll="{  y: 300 }"/>
</template>

<style scoped>

</style>