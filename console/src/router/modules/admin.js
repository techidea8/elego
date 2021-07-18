import Layout from '@/layout'
const opRoutes =   {
    path: '/admin',
    component: Layout,
    redirect: '/admin/member',
    alwaysShow: true, // will always show the root menu
    name: 'Admin',
    meta: {
      title: '运营中心',
		icon: 'el-icon-aim',
      roles:[3],
    },
    children: [
      {
        path: 'member',
        component: () => import('@/views/user/member'),
        name: 'memberList',
        meta: {
          title: '会员管理',
          roles:[3],
          query:{role:1}
        }
      },
      {
        path: 'staff',
        component: () => import('@/views/user/staff'),
        name: 'staffList',
        meta: {
          title: '店员管理',
          roles:[3],
          query:{role:3}
        }
      }
    ]
  }
  export default opRoutes
