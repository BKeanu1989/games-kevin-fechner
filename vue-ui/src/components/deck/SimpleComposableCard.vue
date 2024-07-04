<script lang="ts">
</script>
<script setup lang="ts">
// @ts-nocheck
import { onMounted, ref } from 'vue';
// source: https://www.petercollingridge.co.uk/tutorials/svg/interactive/dragging/

const width = ref(100)
const height = ref(300)
const svgElement = ref<HTMLElement | null>(null)

onMounted(() => {
    const svg = document.getElementById('svgCanvas');

    // Make rectangle draggable
    const rect = document.getElementById('draggableRect');
    makeElementDraggable(rect);

    // Make text draggable
    const text = document.getElementById('draggableText');
    makeElementDraggable(text);

    function makeElementDraggable(element) {
        let selectedElement = null;
        let offset = { x: 0, y: 0 };

        element.addEventListener('mousedown', (event) => {
            selectedElement = event.target;
            offset.x = event.clientX - selectedElement.getBoundingClientRect().x;
            offset.y = event.clientY - selectedElement.getBoundingClientRect().y;
        });

        svg.addEventListener('mousemove', (event) => {
            if (selectedElement) {
                const x = event.clientX - offset.x;
                const y = event.clientY - offset.y;
                const coords = getMousePosition(event)
                console.log("🚀 ~ svg.addEventListener ~ coords:", coords)
                if (selectedElement.tagName === 'rect') {
                    selectedElement.setAttribute('x', coords.x);
                    selectedElement.setAttribute('y', coords.y);
                } else if (selectedElement.tagName === 'text') {
                    selectedElement.setAttribute('x', coords.x);
                    selectedElement.setAttribute('y', coords.y);
                }
            }
        });

        svg.addEventListener('mouseup', () => {
            selectedElement = null;
        });

        svg.addEventListener('mouseleave', () => {
            selectedElement = null;
        });
    }

    function getMousePosition(evt) {
        var CTM = svg.getScreenCTM();
        console.log("🚀 ~ getMousePosition ~ svg:", svg)
        return {
            x: (evt.clientX - CTM.e) / CTM.a,
            y: (evt.clientY - CTM.f) / CTM.d
        };
    }

})

</script>

<template>
    <svg ref="svgElement" id="svgCanvas" :viewBox="'0 0 ' + width + ' ' + height" onload=""
        xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
        <!-- <path id="heart" d="M10,30 A20,20,0,0,1,50,30 A20,20,0,0,1,90,30 Q90,60,50,90 Q10,60,10,30 Z" /> -->
        <path d="M 10 60 C 20 80, 40 80, 50 60" stroke="black" fill="transparent" />
        <circle cx="20" cy="80" r="2" fill="gold" />
        <circle cx="40" cy="80" r="2" fill="gold" />
        <circle cx="50" cy="60" r="2" fill="green" />
        <circle cx="10" cy="60" r="2" fill="green" />
        <!-- <rect ></rect> -->
        <rect id="draggableRect" x="20" y="20" width="40" height="80" stroke="black" fill="transparent"
            stroke-width="2" />
        <text x="30" y="50" id="draggableText">
            test
        </text>
    </svg>
</template>

<style>
svg {
    border: 1px solid black;
}
</style>