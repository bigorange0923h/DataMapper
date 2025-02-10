<template>

  <a-layout style="min-height: 100vh">
    <a-layout-sider v-model:collapsed="collapsed" collapsible style="background: #fff" @collapse="onCollapse">
      <div class="logo"/>
      <a-menu v-model:selectedKeys="selectedKeys" theme="light" mode="inline">
        <template v-for="item in menuItems" :key="item.key">
          <template v-if="item.children">
            <a-sub-menu :key="item.key">
              <template #title>
                <span>{{ item.title }}</span>
              </template>
              <a-menu-item v-for="child in item.children" :key="child.key">
                {{ child.title }}
              </a-menu-item>
            </a-sub-menu>
          </template>
          <template v-else>
            <a-menu-item :key="item.key">
              {{ item.title }}
            </a-menu-item>
          </template>
        </template>
        <a-menu-item key="add">
          <PlusOutlined />
          <span>新建</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-row>
        <a-col :span="24">col</a-col>
      </a-row>
      <a-layout-header class="header top" style="background: #fff">

      </a-layout-header>
      <a-layout-content style="margin: 0 16px">
        <a-tabs v-model:activeKey="activeKey" type="editable-card" @edit="onEdit">
          <a-tab-pane v-for="pane in panes" :key="pane.key" :tab="pane.title" :closable="pane.closable">
            <div :style="{ background: '#fff', padding: '24px', minHeight: '280px' }">
              <a-space>
                <a-input placeholder="请输入接口地址" style="min-width: 110vh;">
                  <template #addonBefore>
                    <a-select v-model:value="apiType" style="width: 90px">
                      <a-select-option value="POST">POST</a-select-option>
                      <a-select-option value="GET">GET</a-select-option>
                    </a-select>
                  </template>
                </a-input>
                <a-button type="primary" @click="importStart">开始导入</a-button>
              </a-space>
              <a-flex :vertical="value === 'vertical'">
                <RequestParamMapper ref="RequestComponent"/>
                <ImportResponse/>
              </a-flex>
            </div>
          </a-tab-pane>
        </a-tabs>


      </a-layout-content>
      <a-layout-footer style="text-align: center">
        footer
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script setup>
import {ref} from 'vue';
import RequestParamMapper from "./components/RequestParamMapper.vue";
import ImportResponse from "./components/ImportResponse.vue";
import {GetMenusData} from "../wailsjs/go/main/App.js";

const RequestComponent = ref();

const collapsed = ref(false);
const selectedKeys = ref(['1']);
const apiType = ref('GET')
const menuItems = ref([
  {
    key: 'sub1',
    title: '文件夹1',
    children: [
      {
        title: '页面 1',
        key: 'setting:1',
      },
      {
        title: '页面 2',
        key: 'setting:2',
      }
    ],
  }, {
    key: 'sub1',
    title: '页面3'
  }
]);


// 引用子组件
const importStart = async () => {
  console.log("开始导入");
  console.log(GetMenusData());
/*  GetMenusData().then(res => {
    console.log(res)
  })*/

}
const onCollapse = (collapseState) => {
  collapsed.value = collapseState;
};

const panes = ref([
  {
    title: 'Tab 1',
    content: 'Content of Tab 1',
    key: '1',
  },
  {
    title: 'Tab 2',
    content: 'Content of Tab 2',
    key: '2',
  },
  {
    title: 'Tab 3',
    content: 'Content of Tab 3',
    key: '3',
    closable: false,
  },
]);
const activeKey = ref(panes.value[0].key);
const newTabIndex = ref(0);
const add = () => {
  activeKey.value = `newTab${++newTabIndex.value}`;
  panes.value.push({
    title: 'New Tab',
    content: 'Content of new Tab',
    key: activeKey.value,
  });
};
const remove = targetKey => {
  let lastIndex = 0;
  panes.value.forEach((pane, i) => {
    if (pane.key === targetKey) {
      lastIndex = i - 1;
    }
  });
  panes.value = panes.value.filter(pane => pane.key !== targetKey);
  if (panes.value.length && activeKey.value === targetKey) {
    if (lastIndex >= 0) {
      activeKey.value = panes.value[lastIndex].key;
    } else {
      activeKey.value = panes.value[0].key;
    }
  }
};
const onEdit = (targetKey, action) => {
  if (action === 'add') {
    add();
  } else {
    remove(targetKey);
  }
};
</script>
<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}

.ant-layout-sider-trigger {
  background: #fff;
}

.a-layout-sider-trigger {
  background-color: #f0f0f0 !important; /* 与 light 主题保持一致 */
  color: #000 !important; /* 设置文字颜色 */
  border-top: 1px solid #d9d9d9; /* 添加边框 */
}


.ant-layout .ant-layout-sider-trigger {
  position: fixed;
  bottom: 0;
  z-index: 1;
  height: 48px;
  color: pink !important;
  line-height: 48px;
  text-align: center;
  background: #fff !important;
  cursor: pointer;
  transition: all 0.2s;
}
</style>