<script setup lang="ts">
import { ref } from 'vue';
import FirstDeck from "@/components/deck/FirstDeck.vue";
import SecondDeck from "@/components/deck/SecondDeck.vue";
import type { Deck } from '@/types';

const deck = ref(["a", "2", "3",])
const deck2 = ref<Deck[]>([
    {
        name: "3",
        type: "",
    },
    {
        name: "2",
        type: ""
    },
    {
        name: "a",
        type: "",
    }
    // "a",
    //  "2", "3",
])
// const 

</script>

<template>
    <div>
        <h1 class="container-v2">Card Game</h1>

        <div class="cards">
            {{ deck }}
        </div>

        ---
        <h1>First deck</h1>
        <SecondDeck :cards="deck2"></SecondDeck>
        <FirstDeck :cards="deck"></FirstDeck>
    </div>
</template>