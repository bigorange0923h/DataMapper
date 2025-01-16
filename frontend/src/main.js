import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import * as Icons from '@ant-design/icons-vue';

const app = createApp(App);
app.use(Antd)
// 全局注册图标
Object.keys(Icons).forEach((key) => {
    app.component(key, Icons[key]);
});
app.mount('#app')

