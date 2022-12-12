<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useDefaultStore } from '../store';

const store = useDefaultStore();
const router = useRouter();

const isUserLoggedIn = () => store.user !== null;

const logout = async () => {
    store.setUser(null);
    localStorage.removeItem('auth_token');
    await router.push("/login");
}

</script>

<template>
    <nav>
        <ul>
            <li>Home</li>
            <li>Admin</li>
            <li v-if="isUserLoggedIn()" @click="logout()">Logout</li>
            <li v-else><RouterLink to="/login">Login</RouterLink></li>
        </ul>
    </nav>
</template>

<style scoped>


</style>