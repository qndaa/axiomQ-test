import {createBookRequest, fetchOneBookRequest, getApi} from "@/api/api";

const state = {
    books: [],
    selectedBook: null
}

const getters = {
    getBooks: function (state) {
        return state.books;
    }, getSelectedBook: function (state) {
        return state.selectedBook
    }
}

const actions = {
    async fetchAllBooks({commit}) {
        await getApi().get('/book').then((response) => {
            commit('setBooks', response.data);
        })
    },  createBook({commit}, params) {
        return createBookRequest(params).then((response) => {
            commit('addBook', response.data);
        })
    }, async fetchOneBook({commit}, id) {
        await fetchOneBookRequest(id).then(response => {
            commit('setSelectedBook', response.data)
        })
    }
}

const mutations = {
    setBooks: function (state, books) {
        state.books = books;
    }, addBook: function (state, book) {
        state.books = [...state.books, book]
    }, setSelectedBook: function (state, book) {
        state.selectedBook = book;
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}
