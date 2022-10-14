import axios, { AxiosHeaders, AxiosRequestConfig } from "axios"

import { User } from "../models/User";

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

export async function getSelf(): Promise<User | null> {
    const response = await axiosClient.get<User>("/users/self").catch((error) => {
        
        console.log(error);
        return null;
    })

    return await response?.data ?? null;
}