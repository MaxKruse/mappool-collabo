import { axiosClient } from ".";
import { Tournament } from "../models/Tournament";


export async function getTournaments(): Promise<Tournament[] | null> {
    const response = await axiosClient.get<Tournament[]>("/tournaments/").catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function getTournament(id: number): Promise<Tournament | null> {
    const response = await axiosClient.get<Tournament>(`/tournaments/${id}`).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function createTournament(tournament: Tournament): Promise<Tournament | null> {
    const response = await axiosClient.post<Tournament>("/tournaments/", tournament).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function updateTournament(tournament: Tournament): Promise<Tournament | null> {
    const response = await axiosClient.put<Tournament>(`/tournaments/${tournament.id}`, tournament).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function deleteTournament(id: number): Promise<boolean> {
    await axiosClient.delete(`/tournaments/${id}`).catch((error) => {
        console.log(error);
        return false;
    })
    
    return true;
}