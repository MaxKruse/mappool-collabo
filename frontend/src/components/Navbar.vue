<script setup lang="ts">
import { useDefaultStore } from '../store';
import { isUserLoggedIn } from '../utils/userFuncs';

const store = useDefaultStore();

</script>

<template>
    <nav>
        <ul>
            <li>Home</li>
            <li>Admin</li>
            <li v-if="isUserLoggedIn()">Logout</li>
            <li v-else>Login</li>
        </ul>
    </nav>
</template>

<style scoped>


</style>