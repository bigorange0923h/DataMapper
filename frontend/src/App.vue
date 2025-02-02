<template>

  <a-layout style="min-height: 100vh">
    <a-layout-sider v-model:collapsed="collapsed" collapsible style="background: #fff"  @collapse="onCollapse" >
      <div class="logo"/>
      <a-menu v-model:selectedKeys="selectedKeys" theme="light" mode="inline">
        <a-menu-item key="1">
          <pie-chart-outlined/>
          <span>Option 1</span>
        </a-menu-item>
        <a-menu-item key="2">
          <desktop-outlined/>
          <span>Option 2</span>
        </a-menu-item>
        <a-sub-menu key="sub1">
          <template #title>
            <span>
              <user-outlined/>
              <span>User</span>
            </span>
          </template>
          <a-menu-item key="3">Tom</a-menu-item>
          <a-menu-item key="4">Bill</a-menu-item>
          <a-menu-item key="5">Alex</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="sub2">
          <template #title>
            <span>
              <team-outlined/>
              <span>Team</span>
            </span>
          </template>
          <a-menu-item key="6">Team 1</a-menu-item>
          <a-menu-item key="8">Team 2</a-menu-item>
        </a-sub-menu>
        <a-menu-item key="9">
          <file-outlined/>
          <span>File</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-row>
        <a-col :span="24">col</a-col>
      </a-row>
      <a-layout-header class="header top"  style="background: #fff">
        <a-tabs v-model:activeKey="activeKey" type="editable-card" @edit="onEdit">
          <a-tab-pane v-for="pane in panes" :key="pane.key" :tab="pane.title" :closable="pane.closable">
            {{ pane.content }}
          </a-tab-pane>
        </a-tabs>
      </a-layout-header>
      <a-layout-content style="margin: 0 16px">
        <!--目录路径-->
        <a-breadcrumb style="margin: 16px 0">
          <a-breadcrumb-item>Home</a-breadcrumb-item>
          <a-breadcrumb-item>List</a-breadcrumb-item>
          <a-breadcrumb-item>App</a-breadcrumb-item>
        </a-breadcrumb>
        <!--API-->
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
        <!--入参-->

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

const RequestComponent = ref();

const collapsed = ref(false);
const selectedKeys = ref(['1']);
const apiType = ref('GET')

// 引用子组件
const importStart = async () => {
  console.log(RequestComponent.value.selectFilePath);

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
  color: #000 !important;              /* 设置文字颜色 */
  border-top: 1px solid #d9d9d9;       /* 添加边框 */
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