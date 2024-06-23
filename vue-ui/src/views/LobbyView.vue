<script setup lang="ts">
import { onUnmounted, ref } from 'vue';

// @ts-nocheck
const ws = ref<any>(null);

// Attempt to establish a WebSocket connection
try {
    ws.value = new WebSocket('ws://localhost:8080/websocket');
    // Event listener for successful connection
    ws.value.onopen = () => console.log('WebSocket connection established');
} catch (error) {
    // Log error if connection fails
    console.error('Failed to establish WebSocket connection:', error);
}

// ws.value.on("connect_error", (err: any) => {
//     // the reason of the error, for example "xhr poll error"
//     console.log(err.message);

//     // some additional description, for example the status code of the initial HTTP response
//     console.log(err.description);

//     // some additional context, for example the XMLHttpRequest object
//     console.log(err.context);
// });

// Proper WebSocket closure within the onUnmounted lifecycle hook
onUnmounted(() => {
    if (ws.value) {
        ws.value.close();
        console.log('WebSocket connection closed');
    }
});

// Define sendMessage: verifies the connection is open before sending
const sendMessage = (message: any) => {

    if (ws.value && ws.value.readyState === WebSocket.OPEN) {
        ws.value.send(JSON.stringify(message)); // Messages must be stringified
    } else {
        // Error handling if the WebSocket connection is not open
        console.error('WebSocket connection is not open.');
    }
};
</script>

<template>
    <main>
        <h1>Lobby</h1>
        <p>
            Choose your games and lobby size here
            <br />
            or enter a room
        </p>
    </main>
</template>

<style scoped></style>