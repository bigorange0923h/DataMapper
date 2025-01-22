<script setup>
import { onMounted, ref, nextTick } from 'vue';
const count = 3;
const fakeDataUrl = `https://randomuser.me/api/?results=${count}&inc=name,gender,email,nat,picture&noinfo`;
const initLoading = ref(true);
const data = ref([]);
const list = ref([]);
const pagination = {
  onChange: page => {
    console.log(page);
  },
  pageSize: 3,
};
onMounted(() => {
  initLoading.value = false;
  fetch(fakeDataUrl)
      .then(res => res.json())
      .then(res => {

        data.value = res.results;
        list.value = res.results;
      });
});
</script>

<template>
  <a-list
      class="demo-loadmore-list"
      item-layout="horizontal"
      :data-source="list"
      :pagination="pagination"
  >
    <template #renderItem="{ item }">
      <a-list-item>
        <template #actions>
          <a key="list-loadmore-more">重新导入</a>
        </template>
        <a-skeleton   :loading="!!item.loading" active>
          <div>行号</div>
          <a-list-item-meta
              description="返回提示:xxxx">
          </a-list-item-meta>
        </a-skeleton>
      </a-list-item>
    </template>
  </a-list>
</template>

<style scoped>

</style>