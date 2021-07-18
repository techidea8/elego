import Layout from '@/layout'
const systemRoutes =   {
    path: '/system',
    component: Layout,
    redirect: '/system/weixin',
    alwaysShow: true, // will always show the root menu
    name: 'System',
    meta: {
      title: '系统配置',
    icon: 'el-icon-aim',
      roles:[4],
    },
    children: [
      
      {
        path: 'weixin',
        component: () => import('@/views/weixin/list'),
        name: 'weixinList',
        meta: {
          title: '微信配置',
          roles:[4],
          query:{role:3}
        }
      }
      ,{
        path: 'dict',
        component: () => import('@/views/dict/list'),
        name: 'dictList',
        meta: {
          title: '字典管理',
          roles:[3]
         
        }
      }
      
    ]
  }
  export default systemRoutes
