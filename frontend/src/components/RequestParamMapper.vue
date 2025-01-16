<script setup>
import {LoadMappingConfig, SelectExcelFile} from '../../wailsjs/go/main/App'
import {computed, ref} from "vue";


const columns = [
  {
    title: '原始入参名称',
    dataIndex: 'sourceName',
    key: 'sourceName',
    editable: true,
    inputType: "input"
  },
  {
    title: '原始入参类型',
    dataIndex: 'sourceType',
    key: 'sourceType',
    editable: true,
    inputType: "select"
  },
  {
    title: '目标参数名称',
    dataIndex: 'targetName',
    key: 'targetName',
    editable: true,
    inputType: "input"
  },
  {
    title: '目标参数类型',
    dataIndex: 'targetType',
    key: 'targetType',
    editable: true,
    inputType: "select"
  },
  {
    title: 'Action',
    dataIndex: 'action',
    key: 'action',
    editable: true
  },
];
let dataSource = ref([]);

let filePath = "";

async function selectFile() {

  filePath = await SelectExcelFile();
  // todo 基于文件路径提前获取对应文件头,并且加载配置表格
  if (filePath) {
    console.log("dataSource:" + dataSource);
    // todo 没有实时根据 返回加载数据
    LoadMappingConfig("csv", filePath).then(result => {
      dataSource.value = result
      console.log("dataSource:" + dataSource);
    })
  }

}

const count = computed(() => dataSource.value.length + 1);

const handleAdd = () => {
  const newData = {
    key: `${count.value}`,
    sourceName: "name",
    sourceType: "string",
    targetName: null,
    targetType: "string",
  };
  dataSource.value.push(newData);
};

// 当前正在编辑的单元格
const editingCell = ref({key: null, dataIndex: null});

// 鼠标悬浮时设置为编辑模式
const handleMouseOver = (key, dataIndex) => {
  editingCell.value = {key, dataIndex};
};

// 鼠标离开时清除编辑状态
const handleMouseLeave = (key, dataIndex) => {
  if (editingCell.value.key === key && editingCell.value.dataIndex === dataIndex) {
    editingCell.value = {key: null, dataIndex: null};
  }
};

// 失去焦点或回车后退出编辑
const handleBlur = (key, dataIndex) => {
  console.log("进入编辑")
  if (editingCell.value.key === key && editingCell.value.dataIndex === dataIndex) {
    editingCell.value = {key: null, dataIndex: null};
  }
};
// 选中某个选项后退出编辑状态
const handleChange = (key, dataIndex) => {
  editingCell.value = {key: null, dataIndex: null};
};
// 判断当前单元格是否正在编辑
const isEditing = (key, dataIndex) =>
    editingCell.value.key === key && editingCell.value.dataIndex === dataIndex;
</script>

<template>
  <a-space style="margin-top: 10px">
    <p v-if="filePath" style="color: #141414">选择的文件路径: {{ filePath }}</p>
    <a-button class="editable-add-btn" style="margin-bottom: 8px" @click="handleAdd">Add</a-button>
    <a-button @click="selectFile">选择文件</a-button>
  </a-space>

  <a-divider/>
  <a-table :columns="columns" :data-source="dataSource" :pagination="false" :scroll="{  y: 300 }">
    <template #bodyCell="{ column, record, index }">
      <div
          v-if="column.editable"
          @mouseover="handleMouseOver(record.key, column.dataIndex)"
          @mouseleave="handleMouseLeave(record.key, column.dataIndex)"
      >
        <template v-if="isEditing(record.key, column.dataIndex)">
          <template v-if="column.inputType === 'select'">
            <a-select
                :options="column.options"
                v-model:value="record[column.dataIndex]"
                @change="handleChange(record.key, column.dataIndex)"
                @blur="handleBlur(record.key, column.dataIndex)"
                style="width: 100%;"
            />
          </template>
          <template v-else>
            <a-input
                v-model:value="record[column.dataIndex]"
                @blur="handleBlur(record.key, column.dataIndex)"
                @pressEnter="handleBlur(record.key, column.dataIndex)"
            />
          </template>
        </template>
        <template v-else>
          {{ record[column.dataIndex] }}
        </template>
      </div>
      <template v-else>
        {{ record[column.dataIndex] }}
      </template>
    </template>
  </a-table>
</template>

<style scoped>
</style>