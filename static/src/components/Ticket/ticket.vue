<template lang="pug">

div.ticket(:data-type="ticket.category.title")
    div.right
        i.pointer.fa.fa-pencil(@click="updateTicket") ｜
        i.pointer.fa.fa-plus(@click="bindIssue")
        
    div(v-if="ticket.due" :class="['due', {'delay': overdue }]")
        i.fa.fa-clock-o
        |   {{duedate}}

    div.name {{ticket.title}}

    div.labels
        span(v-for="label in ticket.labels" :key="label.id" :style="{background:label.bgcolor, color:label.color}") {{label.title}}

    ul.refs.show
        issue-component(v-for="issue in ticket.issues" :key="issue.id" :selected_board="selected_board" :ticket="ticket" :issue="issue")
</template>

<script lang="ts" src="./ticket.ts"></script>
<style lang="scss" scoped src="./ticket.scss"></style>