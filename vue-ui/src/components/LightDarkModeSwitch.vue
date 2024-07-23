<template>
    <div>
        <!-- <button type="button" aria-label="Toggle Dark Mode" :aria-checked="isDark" @click="toggleLocalDark"> -->
        <button type="button" aria-label="Toggle Dark Mode" class="bg-white dark:bg-slate-200 px-5 py-2"
            :aria-checked="localDark" @click="localDark = !localDark">
            <span v-if="localDark">
                <IconMoon width="20" height="20" />
            </span>
            <span v-else>
                <IconSun height="20" width="20" />
            </span>
        </button>
    </div>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
// import { usePreferredDark } from '@vueuse/core'
import IconMoon from './icons/IconMoon.vue';
import IconSun from './icons/IconSun.vue';
// const isDark = usePreferredDark()

const localDark = ref(false)
// const toggleLocalDark = () => {
//     console.log("run toggle")
//     localDark.value = !localDark.value
//     console.log("🚀 ~ toggleLocalDark ~ localDark.value:", localDark.value)
// }

watch(localDark, (newVal, _o) => {
    console.log("watch localDark")
    if (newVal) {
        document.documentElement.classList.add('dark')
    } else {
        document.documentElement.classList.remove('dark')
    }
})


</script>