<template lang="pug">
div
    header
        div.boardlist
            label
                select(v-model="selected_board" v-on:change="getBoard")
                    option(v-for="board in boards" :value="board") {{board}}
        div.logo
            | lssue
        div.options
            //- i.fa.fa-cog

    div.container
        pipe-component(v-for="pipe in pipes" :key="pipe.id" :pipe="pipe" :selected_board="selected_board")

    ticket-create-component(:selected_board="selected_board" :pipes="pipes" :categories="categories" :labels="labels" :issues="issues")
    ticket-update-component(:selected_board="selected_board" :pipes="pipes" :categories="categories" :labels="labels" :issues="issues")
    ticket-delete-component
    issue-bind-component(:selected_board="selected_board" :pipes="pipes" :categories="categories" :labels="labels" :issues="issues")
    issue-unbind-component
</template>

<script lang="ts" src="./board.ts"></script>
<style lang="scss" scoped src="./board.scss"></style>