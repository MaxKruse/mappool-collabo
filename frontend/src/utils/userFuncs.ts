import { useDefaultStore } from "../store";

const store = useDefaultStore();

export const isUserLoggedIn = () => {
    return store.user !== null && store.user !== undefined && store.user.avatar_url !== "";
}