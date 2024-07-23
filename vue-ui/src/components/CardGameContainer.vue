<script setup lang="ts">
import { onMounted } from 'vue';


onMounted(() => {
    initializeDeck()
})
</script>


<script lang="ts">
const deck: HTMLElement | null = document.getElementById('deck');
const hand: HTMLElement | null = document.getElementById('hand');

const cards: string[] = [
    'A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K'
];

function createCardElement(card: string) {
    const cardElement = document.createElement('div');
    cardElement.classList.add('card');
    cardElement.textContent = card;
    cardElement.addEventListener('click', function () {
        moveCardToHand(cardElement);
    });
    return cardElement;
}

function moveCardToHand(cardElement: HTMLElement) {
    hand!.appendChild(cardElement);
}

function shuffle(array: string[]) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
}

function initializeDeck() {
    shuffle(cards);
    deck!.innerHTML = '';
    cards.forEach(card => {
        deck!.appendChild(createCardElement(card));
    });
}

// initializeDeck();
</script>

<template>
    <div>
        <h1 class="text-xl">Card Game Container</h1>
        <div id="game-container">
            <div id="deck"></div>
            <div id="hand"></div>
        </div>
    </div>

</template>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #2c3e50;
    color: #ecf0f1;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    margin: 0;
}

#game-container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

#deck,
#hand {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    margin: 20px;
}

.card {
    width: 100px;
    height: 150px;
    margin: 10px;
    background-color: white;
    border: 1px solid #bdc3c7;
    border-radius: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 24px;
    cursor: pointer;
}
</style>