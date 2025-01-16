<template>

  <a-layout style="min-height: 100vh">
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <div class="logo"/>
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
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
      <a-layout-header class="header top">
        <a-menu
            v-model:selectedKeys="selectedKeys"
            theme="dark"
            mode="horizontal"
            :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="1">nav 1</a-menu-item>
          <a-menu-item key="2">nav 2</a-menu-item>
          <a-menu-item key="3">nav 3</a-menu-item>
        </a-menu>
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
            <a-button type="primary" @click="loadMappingConfig">开始</a-button>
          </a-space>
          <a-space>
            <RequestParamMapper/>
          </a-space>
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

const collapsed = ref(false);
const selectedKeys = ref(['1']);
const apiType = ref('GET')
const loadMappingConfig = ()=>{

}
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

[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}

</style>