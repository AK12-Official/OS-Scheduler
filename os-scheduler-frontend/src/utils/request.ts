import axios from "axios";
import { ElMessage } from "element-plus";

const request = axios.create({
    //基础路径
    baseURL: import.meta.env.VITE_APP_API_URL,//基础路径会携带/api
    timeout: 5000 // 超时时间的设置
});

request.interceptors.request.use((config) => {
    //config配置对象，有headers属性 请求头，给服务器端携带公共的参数
    //返回配置对象
    return config
});

request.interceptors.response.use((response) => {
    //成功回调

    //简化数据
    return response.data
}, (error) => {
    //失败回调:处理http网络错误

    //定义一个变量：存储网络错误信息
    let message = '';

    //http状态码
    const status = error.response.status;

    switch (status) {
        case 400:
            message = 'Request Param Wrong'; break;
        case 401:
            message = 'Token Expired'; break;
        case 403:
            message = 'Permission Denied'; break;
        case 404:
            message = 'RequestUrl Error'; break;
        case 500:
            message = 'Server Error'; break;
        default:
            message = 'NetWork Error'; break;
    }

    ElMessage({
        type: 'error',
        message //key: value简写
    });

    return Promise.reject(error)
});

export default request