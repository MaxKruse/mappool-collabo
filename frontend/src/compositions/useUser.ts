import { axiosClient } from ".";

import { User } from "../models/User";

export async function getSelf(): Promise<User | null> {
    const response = await axiosClient.get<User>("/users/self").catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function getUser(id: number): Promise<User | null> {
    const response = await axiosClient.get<User>(`/users/${id}`).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function getUsers(): Promise<User[] | null> {
    const response = await axiosClient.get<User[]>("/users").catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}