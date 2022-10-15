import { AxiosResponse } from "axios";
import { axiosClient } from ".";
import { StaffDto } from "../models/Staff";

export async function addMappooler(mappooler: StaffDto): Promise<any | null> {
    const response = await axiosClient.post<any>("/mappoolers/", mappooler).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

export async function deleteMappooler(mappooler: StaffDto): Promise<AxiosResponse<any, any> | null> {
    const resp = await axiosClient.delete(`/mappoolers/${mappooler.TournamentID}/${mappooler.UserID}`).catch((error) => {
        console.log(error);
        return null;
    })
    
    return resp;
}