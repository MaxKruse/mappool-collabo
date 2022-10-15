import { axiosClient } from ".";
import { Round, Suggestion } from "../models/Tournament";


export async function addRound(round: Round): Promise<any | null> {
    const response = await axiosClient.post<any>("/rounds/", round).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function addSuggestion(roundId: number, suggestion: Suggestion): Promise<Suggestion | null> {
    const response = await axiosClient.post<Suggestion>(`/rounds/${roundId}/suggest`, suggestion).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function deleteRound(roundId: Round): Promise<any | null> {
    const response = await axiosClient.delete(`/rounds/${roundId}`).catch((error) => {
        console.log(error);
        return null;
    })
    
    return response?.data ?? null;
}
