<script setup lang="ts">
import type { PropType } from 'vue';

const props = defineProps<{
    cards?: string[],
}>()

// const props = defineProps({
//     cards: Array as PropType<string[]>
// })
</script>

<template>
    <section class="first-deck">
        <div v-for="card in props.cards" class="first-deck--single" :key="card">
            | {{ card }} |
        </div>
    </section>
</template>

<style scoped>
.first-deck {
    position: relative;
}

.first-deck--single {
    height: 175px;
    width: 100px;
    border: 2px solid black;

    filter: drop-shadow(black 5px 5px 2px);
}
</style>