<script setup>
import {LoadMappingConfig, SelectExcelFile} from '../../wailsjs/go/main/App'
import { DeleteOutlined } from '@ant-design/icons-vue';
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
    title: '操作',
    dataIndex: 'operate',
    editable: true,
    inputType: "operate"

  },
];
let dataSource = ref([]);

const selectFilePath = ref("");
// 当前正在编辑的单元格
const editingKey = ref(null);
const editingValue = ref("");

async function selectFile() {

  selectFilePath.value = await SelectExcelFile();
  // todo 基于文件路径提前获取对应文件头,并且加载配置表格
  if (selectFilePath.value) {
    console.log("dataSource:" + dataSource);
    // todo 没有实时根据 返回加载数据
    LoadMappingConfig("csv", selectFilePath.value).then(result => {
      dataSource.value = result
      console.log("dataSource:" + dataSource);
    })
  }

}

const count = computed(() => dataSource.value.length + 1);
// 新增行
const handleAdd = () => {
  const newData = {
    id: `${count.value}`,
    sourceName: null,
    sourceType: "string",
    targetName: null,
    targetType: "string",
  };
  dataSource.value.push(newData);
};

// 判断当前单元格是否是正在编辑的状态
const isEditing = (id, dataIndex) => {
  return editingKey.value === `${id}-${dataIndex}`;
};

// 设置当前单元格为编辑状态，并保存初始值
const edit = (id, dataIndex, value) => {
  editingKey.value = `${id}-${dataIndex}`;
  editingValue.value = value;
};
// 保存修改后的值并退出编辑状态
const save = (id, dataIndex) => {
  const record = dataSource.value.find(item => item.id === id);
  if (record) {
    record[dataIndex] = editingValue.value;
  }
  editingKey.value = null; // 清空编辑状态
};

// 删除对应行
const onDelete = id => {
  dataSource.value = dataSource.value.filter(item => item.id !== id);
};

defineExpose({
  selectFilePath
});
</script>

<template>
  <a-space style="margin-top: 10px">
    <p v-if="selectFilePath" style="color: #141414">选择的文件路径: {{ selectFilePath }}</p>
    <a-button class="editable-add-btn" style="margin-bottom: 8px" @click="handleAdd">Add</a-button>
    <a-button @click="selectFile">选择文件</a-button>
  </a-space>

  <a-divider/>
  <a-table :columns="columns" :data-source="dataSource" :pagination="false" :scroll="{  y: 300 }" rowKey="id">
    <template #bodyCell="{ title, record, column, index }">
      <!-- 判断是否是正在编辑的单元格 -->
      <div v-if="isEditing(record.id, column.dataIndex)">
        <template v-if="column.inputType === 'input'">
          <a-input
              v-model:value="editingValue"
              @blur="save(record.id, column.dataIndex)"
              @pressEnter="save(record.id, column.dataIndex)"
              style="width: 100%"
              autofocus
          />
        </template>
        <template v-else-if="column.inputType === 'select'">
          <a-select v-model:value="editingValue" @blur="save(record.id, column.dataIndex)"
                    @change="save(record.id, column.dataIndex)"
                    style="width: 100%" autofocus>
            <a-select-option value="string">string</a-select-option>
            <a-select-option value="int">int</a-select-option>
          </a-select>
        </template>
      </div>
      <div v-else-if="column.inputType === 'operate'">
        <a-popconfirm
            v-if="dataSource.length"
            title="确认删除?"
            @confirm="onDelete(record.id)">
          <DeleteOutlined/>
        </a-popconfirm>
      </div>
      <div
          v-else
          @click="edit(record.id, column.dataIndex, record[column.dataIndex])"
          style="cursor: pointer;">
        {{ record[column.dataIndex] || '-' }}
      </div>
    </template>
  </a-table>
</template>

<style scoped>
</style>