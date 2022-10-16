import axios from "axios";

export const axiosClient = axios.create({
    baseURL: "/api",
    timeout: 10000,
    validateStatus: (status: number) => {
        return status >= 200 && status < 300
    },
    headers: {
        "Authorization": `Bearer ${localStorage.getItem("auth_token")}`,
    }
});