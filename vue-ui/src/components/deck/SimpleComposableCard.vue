<script setup lang="ts">
const width = ref(100)
const height = ref(100)
</script>

<template>
    <svg :viewBox="'0 0 ' + width + ' ' + height">
        <path id="heart" d="M10,30 A20,20,0,0,1,50,30 A20,20,0,0,1,90,30 Q90,60,50,90 Q10,60,10,30 Z" />

    </svg>
</template>