import Vuex from 'vuex';
import auth from "@/store/modules/auth";
import book from "@/store/modules/book";

export default new Vuex.Store({
    modules: {
        auth,
        book
    }
})
