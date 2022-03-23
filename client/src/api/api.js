import axios from "axios";

const ROOT_URL = "http://localhost:8080/api"

export const getApi = () => {
    return axios.create({
        baseURL: ROOT_URL
    })
}


export const loginRequest = (params) => {
    return getApi().post('/login', params)
}

export const createBookRequest = (params) => {

    const config = {
        headers: { Authorization: `Bearer ${window.localStorage.getItem('token')}` }
    };

    return getApi().post('/book', params, config)
}

export const fetchOneBookRequest = (id) => {
    const config = {
        headers: { Authorization: `Bearer ${window.localStorage.getItem('token')}` }
    };

    return getApi().get(`/book/${id}`, config)
}
