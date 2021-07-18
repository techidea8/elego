const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
 device: state => state.app.device,  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  roles: state => state.account.roles,
  permission_routes: state => state.permission.routes,
  errorLogs: state => state.errorLog.logs,
  token:state=>state.account.token,
  avatar:state=>state.account.avatar,
}
export default getters
