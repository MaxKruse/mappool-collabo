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
            <li class="home"><RouterLink to="/">Home</RouterLink></li>
            <li class="dashboard"><RouterLink to="/">Dashboard</RouterLink></li>
            <li v-if="isUserLoggedIn()" class="logout" @click="logout()">Logout</li>
            <li v-if="isUserLoggedIn()" class="profile-image"><img :src="store.user?.avatar_url" alt=""></li>
        </ul>
    </nav>
</template>

<style scoped>

/*
    The navbar should be floating on top of the page.
    It should have its 3 non-image list items aligned to the left.
    The Login button is going to be a green hue, the logout a red hue.
    All other buttons get the color theme.
    The color theme is going to be purple, violet, blue and pinks.
    The image should be aligned to the right, fully contain its content and be a circle of 64x64.
    Keep in mind a small margin and padding so the buttons are not literally on the edge of the screen.
    Also appropriately size the font and list items for 1080p displays.
    The image should be on the right edge of the screen. Still inline with the buttons, but on the far right edge
*/

nav {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 64px;
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 32px;
    margin: 0;
    z-index: 1;
}

ul {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    list-style: none;
    margin: 0;
    padding: 0;
}

li {
    margin: 0 16px;
    padding: 0;
    font-size: 1.2rem;
}

li.login {
    color: #fff;
    background-color: #00b894;
    border-radius: 4px;
    padding: 8px 16px;
}

li.logout {
    color: #fff;
    background-color: #d63031;
    border-radius: 4px;
    padding: 8px 16px;
}

li.profile-image {
    margin-right: 0;
}

li.profile-image img {
    width: 64px;
    height: 64px;
    border-radius: 50%;
    object-fit: cover;

}


</style>