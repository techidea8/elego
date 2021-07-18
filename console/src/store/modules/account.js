import { login, register,getInfo,auth } from '@/api/account'
import { getToken, setToken, removeToken } from '@/utils/auth'
import router, { resetRouter } from '@/router'

const state = {
  token: getToken(),
  name: '',
  avatar: '',
  roles: [],
  id:{},
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_ID: (state, roles) => {
    state.id = roles
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
  
    return new Promise((resolve, reject) => {
      login(userInfo).then(response => {
        const { data } = response
        commit('SET_TOKEN', data.token)
        commit('SET_ID', data.id)
        commit('SET_NAME', data.name)
        commit('SET_AVATAR',data.avatar)
       // commit('SET_ROLES', [data.role]) 这不要,让他通过getinfo处理一下
        setToken(data.token)
        console.log('data',data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  register({ commit }, arg) {
    return new Promise((resolve, reject) => {
      register(arg).then(response => {
        const { data } = response
        commit('SET_TOKEN', data.token)
        commit('SET_ID', data.id)
        commit('SET_NAME', data.name)
        commit('SET_AVATAR',data.avatar)
        commit('SET_ROLES', [data.role])
        setToken(data.token)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // get user info
  getInfo({ commit, state,dispatch }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        const { data } = response
        
        console.log("getInfo",response)
        if (!data) {
         return  reject(new Error('网络繁忙,请重试'))
        }

        const {id, role, name, avatar ,token} = data

        // roles must be a non-empty array
        if (!role) {
          reject(new Error('该账户未授权,请重新登陆'))
        }
        
        
        commit('SET_ID', id)
        commit('SET_ROLES', [role])
        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        commit('SET_TOKEN', token)
        setToken(token)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
// get user info
auth({ commit, state,dispatch },token) {
  return new Promise((resolve, reject) => {
    auth({token:token}).then(response => {
      const { data } = response
      
      console.log("getInfo",response)
      if (!data) {
       return  reject(new Error('网络繁忙,请重试'))
      }

      const {id, role, name, avatar ,token} = data

      // roles must be a non-empty array
      if (!role) {
        reject(new Error('该账户未授权,请重新登陆'))
      }
      commit('SET_ID', id)
      commit('SET_ROLES', [role])
      commit('SET_NAME', name)
      commit('SET_AVATAR', avatar)
      commit('SET_TOKEN', token)
      setToken(token)
      resolve(data)
    }).catch(error => {
      reject(error)
    })
  })
},
  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
         dispatch('tagsView/delAllViews', null, { root: true })
        removeToken()
        resetRouter()
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        resolve()

    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },

  // dynamically modify permissions
  changeRoles({ commit, dispatch }, role) {
    return new Promise(async resolve => {
      const token = role + '-token'

      commit('SET_TOKEN', token)
      setToken(token)

      const { role } = await dispatch('getInfo')

      resetRouter()

      // generate accessible routes map based on roles
      const accessRoutes = await dispatch('permission/generateRoutes', [role], { root: true })

      // dynamically add accessible routes
      router.addRoutes(accessRoutes)

      // reset visited views and cached views
      dispatch('tagsView/delAllViews', null, { root: true })

      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
