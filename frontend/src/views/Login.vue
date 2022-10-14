<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';


const route = useRoute()
const router = useRouter()

const { query } = route

// check if we got a "token" query param
if (query.token) {
    // set the token as the auth_token localStorage item
    localStorage.setItem('auth_token', query.token as string)
    // redirect to the home page
    router.push('/')
}
</script>

<template>
    <div>
        <a href="/api/oauth/login">
            <h1>Login</h1>
        </a>
    </div>
</template>