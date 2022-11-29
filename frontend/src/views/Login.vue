<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getSelf } from '../compositions/useUser';
import { useDefaultStore } from '../store';


const route = useRoute()
const router = useRouter()

const store = useDefaultStore();

onMounted(async () => {
    // check if we got a "token" query param
    const token = route.query.token as string | undefined;
    if (token) {
        // set the token as the auth_token localStorage item
        localStorage.setItem('auth_token', token as string)

        // get the current user based on the auth_token and commit him to the store
        try {
            const user = await getSelf();
            store.setUser(user);
        } catch (e) {
            localStorage.removeItem('auth_token');
            console.error(e);
        }

        // redirect to the home page
        router.push('/')
    }
})

</script>

<template>
    <div>
        <a href="/api/oauth/login">
            <h1>Login</h1>
        </a>
    </div>
</template>