// vite.config.ts
import { defineConfig } from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/vite@5.4.5_@types+node@22.5.5/node_modules/vite/dist/node/index.js";
import vue from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/@vitejs+plugin-vue@5.1.3_vite@5.4.5_@types+node@22.5.5__vue@3.5.5_typescript@5.6.2_/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import AutoImport from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/unplugin-auto-import@0.18.3_@vueuse+core@11.0.3_vue@3.5.5_typescript@5.6.2___rollup@4.21.3/node_modules/unplugin-auto-import/dist/vite.js";
import Components from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/unplugin-vue-components@0.27.4_@babel+parser@7.25.6_rollup@4.21.3_vue@3.5.5_typescript@5.6.2_/node_modules/unplugin-vue-components/dist/vite.js";
import { ElementPlusResolver } from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/unplugin-vue-components@0.27.4_@babel+parser@7.25.6_rollup@4.21.3_vue@3.5.5_typescript@5.6.2_/node_modules/unplugin-vue-components/dist/resolvers.js";
import Icons from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/unplugin-icons@0.19.3_@vue+compiler-sfc@3.5.5/node_modules/unplugin-icons/dist/vite.js";
import IconsResolver from "file:///D:/CodeProjects/XuanQiong/admin/node_modules/.pnpm/unplugin-icons@0.19.3_@vue+compiler-sfc@3.5.5/node_modules/unplugin-icons/dist/resolver.js";
var vite_config_default = defineConfig({
  build: {
    // 配置静态资源的公共路径
    assetsDir: "static"
  },
  plugins: [
    vue(),
    AutoImport({
      resolvers: [
        // 自动导入函数
        ElementPlusResolver(),
        // 自动导入图标组件
        IconsResolver({
          prefix: "Icon"
        })
      ]
    }),
    Components({
      resolvers: [
        // 自动导入 Element Plus 组件
        ElementPlusResolver(),
        // 自动注册图标组件
        IconsResolver({
          enabledCollections: ["ep"]
        })
      ]
    }),
    Icons({
      autoInstall: true
      // 自动安装图标集
    })
  ],
  resolve: {
    alias: {
      // 设置别名 '@' 指向 '/src' 目录
      //'@': '/src',
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJEOlxcXFxDb2RlUHJvamVjdHNcXFxcWHVhblFpb25nXFxcXGFkbWluXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJEOlxcXFxDb2RlUHJvamVjdHNcXFxcWHVhblFpb25nXFxcXGFkbWluXFxcXHZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9EOi9Db2RlUHJvamVjdHMvWHVhblFpb25nL2FkbWluL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcbmltcG9ydCB2dWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJ1xuaW1wb3J0IEF1dG9JbXBvcnQgZnJvbSAndW5wbHVnaW4tYXV0by1pbXBvcnQvdml0ZSdcbmltcG9ydCBDb21wb25lbnRzIGZyb20gJ3VucGx1Z2luLXZ1ZS1jb21wb25lbnRzL3ZpdGUnXG5pbXBvcnQgeyBFbGVtZW50UGx1c1Jlc29sdmVyIH0gZnJvbSAndW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzJ1xuaW1wb3J0IEljb25zIGZyb20gJ3VucGx1Z2luLWljb25zL3ZpdGUnXG5pbXBvcnQgSWNvbnNSZXNvbHZlciBmcm9tICd1bnBsdWdpbi1pY29ucy9yZXNvbHZlcidcblxuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcbiAgYnVpbGQ6IHtcbiAgICAvLyBcdTkxNERcdTdGNkVcdTk3NTlcdTYwMDFcdThENDRcdTZFOTBcdTc2ODRcdTUxNkNcdTUxNzFcdThERUZcdTVGODRcbiAgICBhc3NldHNEaXI6ICdzdGF0aWMnXG4gIH0sXG4gIHBsdWdpbnM6IFtcbiAgICB2dWUoKSxcbiAgICBBdXRvSW1wb3J0KHtcbiAgICAgIHJlc29sdmVyczogW1xuICAgICAgICAvLyBcdTgxRUFcdTUyQThcdTVCRkNcdTUxNjVcdTUxRkRcdTY1NzBcbiAgICAgICAgRWxlbWVudFBsdXNSZXNvbHZlcigpLFxuICAgICAgICAvLyBcdTgxRUFcdTUyQThcdTVCRkNcdTUxNjVcdTU2RkVcdTY4MDdcdTdFQzRcdTRFRjZcbiAgICAgICAgSWNvbnNSZXNvbHZlcih7XG4gICAgICAgICAgcHJlZml4OiAnSWNvbicsXG4gICAgICAgIH0pLFxuICAgICAgXSxcbiAgICB9KSxcbiAgICBDb21wb25lbnRzKHtcbiAgICAgIHJlc29sdmVyczogW1xuICAgICAgICAvLyBcdTgxRUFcdTUyQThcdTVCRkNcdTUxNjUgRWxlbWVudCBQbHVzIFx1N0VDNFx1NEVGNlxuICAgICAgICBFbGVtZW50UGx1c1Jlc29sdmVyKCksXG4gICAgICAgIC8vIFx1ODFFQVx1NTJBOFx1NkNFOFx1NTE4Q1x1NTZGRVx1NjgwN1x1N0VDNFx1NEVGNlxuICAgICAgICBJY29uc1Jlc29sdmVyKHtcbiAgICAgICAgICBlbmFibGVkQ29sbGVjdGlvbnM6IFsnZXAnXSxcbiAgICAgICAgfSksXG4gICAgICBdLFxuICAgIH0pLFxuICAgIEljb25zKHtcbiAgICAgIGF1dG9JbnN0YWxsOiB0cnVlLCAvLyBcdTgxRUFcdTUyQThcdTVCODlcdTg4QzVcdTU2RkVcdTY4MDdcdTk2QzZcbiAgICB9KSxcbiAgXSxcbiAgcmVzb2x2ZToge1xuICAgIGFsaWFzOiB7XG4gICAgICAvLyBcdThCQkVcdTdGNkVcdTUyMkJcdTU0MEQgJ0AnIFx1NjMwN1x1NTQxMSAnL3NyYycgXHU3NkVFXHU1RjU1XG4gICAgICAvLydAJzogJy9zcmMnLFxuICAgIH0sXG4gIH0sXG59KVxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUF1UixTQUFTLG9CQUFvQjtBQUNwVCxPQUFPLFNBQVM7QUFDaEIsT0FBTyxnQkFBZ0I7QUFDdkIsT0FBTyxnQkFBZ0I7QUFDdkIsU0FBUywyQkFBMkI7QUFDcEMsT0FBTyxXQUFXO0FBQ2xCLE9BQU8sbUJBQW1CO0FBRTFCLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLE9BQU87QUFBQTtBQUFBLElBRUwsV0FBVztBQUFBLEVBQ2I7QUFBQSxFQUNBLFNBQVM7QUFBQSxJQUNQLElBQUk7QUFBQSxJQUNKLFdBQVc7QUFBQSxNQUNULFdBQVc7QUFBQTtBQUFBLFFBRVQsb0JBQW9CO0FBQUE7QUFBQSxRQUVwQixjQUFjO0FBQUEsVUFDWixRQUFRO0FBQUEsUUFDVixDQUFDO0FBQUEsTUFDSDtBQUFBLElBQ0YsQ0FBQztBQUFBLElBQ0QsV0FBVztBQUFBLE1BQ1QsV0FBVztBQUFBO0FBQUEsUUFFVCxvQkFBb0I7QUFBQTtBQUFBLFFBRXBCLGNBQWM7QUFBQSxVQUNaLG9CQUFvQixDQUFDLElBQUk7QUFBQSxRQUMzQixDQUFDO0FBQUEsTUFDSDtBQUFBLElBQ0YsQ0FBQztBQUFBLElBQ0QsTUFBTTtBQUFBLE1BQ0osYUFBYTtBQUFBO0FBQUEsSUFDZixDQUFDO0FBQUEsRUFDSDtBQUFBLEVBQ0EsU0FBUztBQUFBLElBQ1AsT0FBTztBQUFBO0FBQUE7QUFBQSxJQUdQO0FBQUEsRUFDRjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
