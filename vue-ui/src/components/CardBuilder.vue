<script setup lang="ts">
import { ref } from 'vue';

const textInput1 = ref("")
const textInput2 = ref("")
const isImageToUpload = ref(false)
</script>

<template>
    <div class="card-builder">
        <form action="">
            <template v-if="isImageToUpload">
                <p>Todo</p>
                <button type="button" @click="isImageToUpload = false">Switch to input</button>
            </template>
            <template v-else>
                <button type="button" @click="isImageToUpload = true" disabled>Switch to Image
                    Upload</button>
                <i>in the near future</i>
                <div class="group">
                    <label for="text-input">Text</label>
                    <input class="" id="text-input" v-model="textInput1" />
                </div>
                <div class="group">
                    <label for="text-input-2">Text</label>
                    <input class="" id="text-input-2" v-model="textInput2" />
                </div>

            </template>
        </form>
        <div class="result relative">


            <div class="res-inp-1 absolute top-2 left-2 text-2xl">
                {{ textInput1 }}

            </div>
            <div class="res-inp-2 absolute top-2 right-2 text-2xl">
                {{ textInput2 }}

            </div>
        </div>
    </div>
</template>

<style scoped></style>