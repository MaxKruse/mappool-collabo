import { defineStore } from "pinia";
import { ref } from "vue";
import { User } from "../models/User";

export const useDefaultStore = defineStore("main", () => {
    const user = ref<User | null>(null);

    function setUser(newUser: any) {
        user.value = newUser;
    }

    return {
        user, setUser
    }
});