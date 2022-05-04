import axios from 'axios';

const service = axios.create({
    // process.env.NODE_ENV === 'development' 来判断是否开发环境
    baseURL: '/api',
    timeout: 5000
});

service.interceptors.request.use(
    config => {
        if (sessionStorage.getItem('Authorization')) {
            config.headers["x-token"] = sessionStorage.getItem('Authorization')
        }
        return config;
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

service.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return response.data;
        } else {
            Promise.reject();
        }
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

export default service;
