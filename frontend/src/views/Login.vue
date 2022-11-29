<script setup lang="ts">
import { onBeforeMount, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getSelf } from '../compositions/useUser';
import { useDefaultStore } from '../store';

const store = useDefaultStore();

const route = useRoute()
const router = useRouter()

onBeforeMount(async () => {
    console.log("before mount login")
    // check if we got a "token" query param
    const token = route.query.token as string | undefined;
    if (!token) {
        return
    }
    // set the token as the auth_token localStorage item
    sessionStorage.setItem('auth_token', token as string)
    await router.push('/')
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