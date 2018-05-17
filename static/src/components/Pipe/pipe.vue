
<template lang="pug">
div.pipe
    div.title
        | {{pipe.title}}
        div.right
            i.pointer.fa.fa-plus(@click="newTicket")

    div.ticket(v-for="ticket in pipe.tickets" :key="ticket.id" :ticket="ticket" :selected_board="selected_board")
            
    draggable-component.tickets(v-model="pipe.tickets" :options="{group:'ticket'}" @start="drag=true" @end="drag=false" @change="movedTicket")
        ticket-component(v-for="ticket in pipe.tickets" :key="ticket.id" :ticket="ticket" :selected_board="selected_board")
    
    
</template>

<script lang="ts" src="./pipe.ts"></script>
<style lang="scss" scoped src="./pipe.scss"></style>