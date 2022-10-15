import { axiosClient } from ".";
import { Map } from "../models/Tournament";


export async function getMap(id: number): Promise<Map | null> {
    const response = await axiosClient.get<Map>(`/maps/${id}`).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

// Upload a replay file (FormFile) for a specified map (id)
export async function uploadReplay(id: number, file: File): Promise<any | null> {
    const formData = new FormData();
    formData.append("replay", file);

    const response = await axiosClient.post<any>(`/maps/${id}/replay`, formData).catch((error) => {
        console.log(error);
        return null;
    })

    return response?.data ?? null;
}

// Download a file from the server by specifying the id
export async function downloadReplay(id: number): Promise<void> {
    const response = await axiosClient.get<any>(`/maps/${id}/replay`, { responseType: "blob" }).catch((error) => {
        console.log(error);
        return;
    })

    const fileBlob = new Blob([response?.data], { type: "application/octet-stream" });
    const url = window.URL.createObjectURL(fileBlob);
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `replay-${id}.osr`);
    document.body.appendChild(link);
    link.click();
    link.remove();
    return;
}