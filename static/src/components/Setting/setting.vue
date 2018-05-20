<template lang="pug">
div(:class="['setting', {'show':visible}]")
    div.box
        div.close(@click="close")
            i.fa.fa-times
        div.title Settings
        div.unit
            div.header
                div.title Default Board
            ul
                li(v-for="board in boards") {{board}}
        div.unit
            div.header
                div.title Category
                div.new
                    i.fa.fa-plus &nbsp;New
            ul
                li(v-for="category in categories") {{category.title}} [delete]

        div.unit
            div.header
                div.title Label
                div.new
                    i.fa.fa-plus &nbsp;New
            ul
                li(v-for="label in labels") 
                    | {{label.title}} 
                    i.fa.fa-remove
            
        div.unit
            div.header
                div.title Webhook Secret
                div.new
                    i.fa.fa-refresh &nbsp;Regenerate
            form
                input(type="text" readonly="readonly")
</template>

<script lang="ts" src="./setting.ts"></script>
<style lang="scss" scoped src="./setting.scss"></style>