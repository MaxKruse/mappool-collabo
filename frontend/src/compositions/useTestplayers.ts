import { axiosClient } from ".";
import { StaffDto } from "../models/Staff";


export async function addTestplayer(testplayer: StaffDto): Promise<any | null> {
    const response = await axiosClient.post<any>("/testplayers/", testplayer).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function deleteTestplayer(testplayer: StaffDto): Promise<any | null> {
    const response = await axiosClient.delete(`/testplayers/${testplayer.TournamentID}/${testplayer.UserID}`).catch((error) => {
        console.log(error);
        return null;
    })
    
    return response?.data ?? null;
}