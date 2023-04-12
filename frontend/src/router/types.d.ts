import 'vue-router'

declare module 'vue-router' {
  // 自定义路由元信息
  interface RouteMeta {
    title?: string
    icon?: title
    role?: string // 路由权限 admin/editor/*
    hideInMenu?: boolean // 是否在导航菜单中隐藏
    hideChildrenInMenu?: boolean // 是否在导航菜单中隐藏子路由
    activeMenu?: string // 路由活跃时高亮指定路由，如子路由活跃时，同时高亮父路由
  }
}
