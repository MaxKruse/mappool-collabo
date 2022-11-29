<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getSelf } from '../compositions/useUser';
import { useDefaultStore } from '../store';

const store = useDefaultStore();

const route = useRoute()
const router = useRouter()

onMounted(async () => {
    // check if we got a "token" query param
    const token = route.query.token as string | undefined;
    console.log(token);
    if (!token) {
        return
    }
    // set the token as the auth_token localStorage item
    await localStorage.setItem('auth_token', token as string)

    // get the current user based on the auth_token and commit him to the store
    try {
        const user = await getSelf();
        store.setUser(user);

        // redirect to the home page
        await router.push('/')
    } catch (e) {
        await localStorage.removeItem('auth_token');
        console.error(e);
    }
})

</script>

<template>
    <div>
        <a href="/api/oauth/login">
            <h1>
                Login
            </h1>
        </a>
    </div>
</template>