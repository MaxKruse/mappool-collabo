import { getAxiosClient } from ".";

import { User } from "../models/User";

export async function getSelf(): Promise<User | null> {
    const response = await getAxiosClient().get<User>("/users/self")

    return response?.data ?? null;
}

export async function getUser(id: number): Promise<User | null> {
    const response = await getAxiosClient().get<User>(`/users/${id}`)

    return response?.data ?? null;
}

export async function getUsers(): Promise<User[] | null> {
    const response = await getAxiosClient().get<User[]>("/users")

    return response?.data ?? null;
}