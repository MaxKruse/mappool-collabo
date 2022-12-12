<script setup lang="ts">
import { onBeforeMount, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getSelf } from '../compositions/useUser';
import { useDefaultStore } from '../store';

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