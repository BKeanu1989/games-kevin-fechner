<script setup lang="ts">
import { ref } from 'vue';

const username = ref("")
const password = ref("")

const submitNewUser = () => {
    console.log("submit new user", username.value, password.value)
}
</script>

<template>
    <div class="new-user">
        <h1 class="text-3xl">Neuer Nutzer</h1>
        <form @submit.prevent="submitNewUser" action="">
            <div class="group">
                <label for="username">Benutzername:</label>
                <input type="text" id="username" v-model="username">
            </div>
            <div class="group">
                <label for="password">Benutzername:</label>
                <input type="text" v-model="password">
            </div>
            <button type="submit">Senden</button>
        </form>
    </div>
</template>