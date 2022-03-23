import {loginRequest} from "@/api/api";

const state = {
    token: window.localStorage.getItem('token')
};

const getters = {
    isLoggedIn: state => !!state.token
};

const actions = {
    login({ commit }, params) {
        commit('setToken', null)
        return loginRequest(params);
    }, setToken({commit}, token) {
        commit('setToken', token)
        window.localStorage.setItem('token', token)
    }, logout({commit}) {
        commit('setToken', null);
        window.localStorage.removeItem('token')
    }
}

const mutations = {
    setToken: (state, token) => {
        state.token = token;
    }}

export default {
    state,
    getters,
    actions,
    mutations
}
