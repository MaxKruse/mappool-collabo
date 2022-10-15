import { axiosClient } from ".";
import { Vote } from "../models/Tournament";


export async function addVote(suggestionId: number, vote: Vote): Promise<any | null> {
    const response = await axiosClient.post<any>(`/votes/${suggestionId}`, vote).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function deleteVote(voteId: number): Promise<any | null> {
    const response = await axiosClient.delete(`/votes/${voteId}`).catch((error) => {
        console.log(error);
        return null;
    })
    
    return response?.data ?? null;
}
