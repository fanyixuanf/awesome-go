import vue from '@vitejs/plugin-vue'
import {defineConfig} from "vite";

export default defineConfig({
    base: './',
    plugins: [vue()],
    optimizeDeps: {
        include: ['schart.js']
    },
    server: {
        cors: true, // 默认启用并允许任何源
        open: true, // 在服务器启动时自动在浏览器中打开应用程序
        //反向代理配置，注意rewrite写法，开始没看文档在这里踩了坑
        proxy: {
            '/api': {
                target: 'http://172.26.0.2:8088',   //代理接口
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            }
        },
    build: {
          target: 'es2015',
          minify: 'terser', // 是否进行压缩,boolean | 'terser' | 'esbuild',默认使用terser
          manifest: false, // 是否产出maifest.json
          sourcemap: false, // 是否产出soucemap.json
          outDir: 'dist', // 产出目录
          rollupOptions:{
            external: ['vue','wangEditor','element-plus'],
            output: {
                // 在 UMD 构建模式下为这些外部化的依赖提供一个全局变量
                globals: {
                  vue: 'Vue',
                  'wang-editor': 'wangEditor',
                  'element-plus': 'elementPlus'
                }
              },
          },
        },
    },
})

