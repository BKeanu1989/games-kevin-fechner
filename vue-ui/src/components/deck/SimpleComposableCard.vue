<script lang="ts">

function makeDraggable(evt: MouseEvent) {
    console.log(evt)
    var selectedElement, offset;

    var svg = evt.target;
    if (!svg) return;
    svg.addEventListener('mousedown', startDrag);
    svg.addEventListener('mousemove', drag);
    svg.addEventListener('mouseup', endDrag);
    svg.addEventListener('mouseleave', endDrag);
    function startDrag(evt: MouseEvent) {
        if (evt.target!.classList.contains('draggable')) {
            selectedElement = evt.target;
            offset = getMousePosition(evt);
            // TODO: check error, being lazy now
            // @ts-ignore
            offset.x -= parseFloat(selectedElement.getAttributeNS(null, "x"));
            // @ts-ignore
            offset.y -= parseFloat(selectedElement.getAttributeNS(null, "y"));
        }
    }
    function drag(evt: MouseEvent) {
        if (selectedElement) {
            evt.preventDefault();
            var coord = getMousePosition(evt);
            selectedElement.setAttributeNS(null, "x", coord.x);
            selectedElement.setAttributeNS(null, "y", coord.y);
        }
    }
    function endDrag(evt: MouseEvent) {
        if (selectedElement) {
            evt.preventDefault();
            var dragX = evt.clientX;
            var dragY = evt.clientY;
            selectedElement.setAttributeNS(null, "x", dragX);
            selectedElement.setAttributeNS(null, "y", dragY);
        }
    }
    function getMousePosition(evt: MouseEvent) {
        var CTM = svg.getScreenCTM();
        return {
            x: (evt.clientX - CTM.e) / CTM.a,
            y: (evt.clientY - CTM.f) / CTM.d
        };
    }
}
</script>
<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue';
// source: https://www.petercollingridge.co.uk/tutorials/svg/interactive/dragging/

const width = ref(100)
const height = ref(300)
const svgElement = ref<HTMLElement | null>(null)


</script>

<template>
    <svg ref="svgElement" :viewBox="'0 0 ' + width + ' ' + height" onload="makeDraggable(evt)"
        xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
        <!-- <path id="heart" d="M10,30 A20,20,0,0,1,50,30 A20,20,0,0,1,90,30 Q90,60,50,90 Q10,60,10,30 Z" /> -->
        <path d="M 10 60 C 20 80, 40 80, 50 60" stroke="black" fill="transparent" />
        <circle cx="20" cy="80" r="2" fill="gold" />
        <circle cx="40" cy="80" r="2" fill="gold" />
        <circle cx="50" cy="60" r="2" fill="green" />
        <circle cx="10" cy="60" r="2" fill="green" />
        <!-- <rect ></rect> -->
        <rect class="draggable" x="20" y="20" width="40" height="80" stroke="black" fill="transparent"
            stroke-width="2" />
        <text x="30" y="50" class="draggable">
            test
        </text>
    </svg>
</template>