import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    status: '',
    token: localStorage.getItem('token') || '',
    role: '',
    username: '',
    equipments: [{ invId: 1, description: 'Test', image: 'https://i.unishop.by/models/smartbuy/olt_240_gb_sb240gb-jlt-25sat3_0.jpg', name: 'An Item', price: 999 },
      { invId: 2, description: 'Test', image: 'https://ae01.alicdn.com/kf/HTB1nt4PPXXXXXc2XpXXq6xXFXXXb/NT-128-80-22mm-Hot-sale-NGFF-M-2-128GB-SSD-hard-drive-for-Ultrabook-Notebook.jpg', name: 'Thing', price: 1499 },
      { invId: 3, description: 'Test', image: 'https://pbs.twimg.com/media/ECqyMX9U4AAwxZw.jpg:large', name: 'Doo-dad', price: 499 },
      { invId: 4, description: 'Test', image: 'https://a.allegroimg.com/s1440/064146/f7abfff14d91b79cec28a0e8e0c9', name: 'Процессор Intel Core i9-9900K', price: 39300 }],
    inCart: [],
  },
  mutations: {
    auth_request(state) {
      state.status = 'loading';
    },
    auth_success(state, token) {
      state.status = 'success';
      state.token = token;
    },
    auth_error(state) {
      state.status = 'error';
    },
    logout(state) {
      state.status = '';
      state.token = '';
    },
    info_request(state, username, role) {
      state.status = 'info';
      state.username = username;
      state.role = role;
    },
    add_to_cart(state, payload) {
      const id = state.inCart.indexOf(state.inCart.find(item => item.id === payload.id));
      console.log(id);
      if (id !== -1) {
        state.inCart[id].count = payload.count;
      } else {
        state.inCart.push({ id: payload.id, count: payload.count });
      }
    },
    remove_from_cart(state, index) {
      state.inCart.splice(state.inCart.find(item => item === index), 1);
    },
    clean(state) {
      state.inCart = [];
    },
  },
  actions: {
    login({ commit }, user) {
      return new Promise((resolve, reject) => {
        commit('auth_request');
        axios({ url: 'http://localhost:1984/login', data: user, method: 'POST' })
          .then((resp) => {
            const token = resp.data.token;
            localStorage.setItem('token', token);
            axios({ url: 'http://localhost:1984/whoami',
              method: 'GET',
              headers: { Authorization: `Bearer ${token}` },
            })
              .then((res) => {
                resolve(res);
                commit('info_request', res.data.username, res.data.role);
              })
              .catch((err) => {
                reject(err);
              });
            commit('auth_success', token);
            resolve(resp);
          })
          .catch((err) => {
            commit('auth_error');
            localStorage.removeItem('token');
            reject(err);
          });
      });
    },
    register({ commit }, user) {
      return new Promise((resolve, reject) => {
        commit('auth_request');
        axios({ url: 'http://157.245.201.175:1984/register', data: user, method: 'POST' })
          .then((resp) => {
            console.log(resp);
            resolve(resp);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    logout({ commit }) {
      return new Promise((resolve) => {
        commit('logout');
        localStorage.removeItem('token');
        delete axios.defaults.headers.common.Authorization;
        resolve();
      });
    },
    addToCart({ commit }, payload) {
      commit('add_to_cart', payload);
    },
    removeFromCart({ commit }, index) {
      commit('remove_from_cart', index);
    },
    clean({ commit }) {
      commit('clean');
    },
  },
  getters: {
    isLoggedIn: state => !!state.token,
    authStatus: state => state.status,
    getUsername: state => state.username,
    getRole: state => state.role,
    getToken: state => state.token,
    getEquipments: state => state.equipments,
    getCart: state => state.inCart,
  },
});
