import '@/style/reset.scss' //引入全局样式
import 'element-plus/dist/index.css'        //引入ElementPlus及样式
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn' //引入ElementPlus中文语言包

import App from './App.vue'


const app = createApp(App)


app.use(createPinia())

// 安装ElementPlus插件
app.use(ElementPlus, {
    locale: zhCn,
})

app.mount('#app')
