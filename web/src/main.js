import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
import './assets/css/icon.css'

//import axios from 'axios';

const app = createApp(App)

// Vue.config.productionTip = false
// Vue.prototype.$axios = Axios
// axios.defaults.baseURL = '/api'
// app.config.globalProperties.$axios = axios



installElementPlus(app)
app
    .use(store)
    .use(router)
    .mount('#app')