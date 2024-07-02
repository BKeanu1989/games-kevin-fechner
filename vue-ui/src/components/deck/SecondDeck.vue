<script setup lang="ts">
import type { PropType } from 'vue';
import type { Deck } from "@/types/index"
const props = defineProps<{
    cards?: Deck[],
}>()

// const props = defineProps({
//     cards: Array as PropType<string[]>
// })
</script>

<template>
    <section class="second-deck">
        <div v-for="card in props.cards" class="second-deck--single" :key="card.name">
            | {{ card }} |
        </div>
    </section>
</template>

<style scoped>
.second-deck {
    position: relative;
    display: flex;
    flex-direction: column;
    perspective: 800px;
}

.second-deck--single {
    height: 25px;
    width: 100px;
    border: 2px solid black;
    transform: rotateX(65deg);
    filter: drop-shadow(black 5px 5px 2px);
}
</style>