<template lang="pug">
div
    div(:class="['modal', {'show':formVisible}]")
        | New Card
        hr
        form
            div.field
                div
                    select#pipes-new(v-model="form.pipe_id")
                        option(v-for="pipe in pipes" :value="pipe.id") {{pipe.title}}

            div.field
                div
                    select#types-new(v-model="form.category_id")
                        option(v-for="category in categories" :value="category.id") {{category.title}}

            div.field
                div
                    input(v-model="form.title" type="text" placeholder="ストーリー名")

            div.field
                div
                    datepicker-component(v-model="form.due" placeholder="期限" format="yyyy-MM-dd")

            div.field
                div.checkboxes
                    label(v-for="label in labels" :style="{background:label.bgcolor, color:label.color}")
                        input(type="checkbox" :value="label.id" v-model="form.label_ids")
                        | {{label.title}}  

            div.field
                div
                    textarea(v-model="form.memo" placeholder="メモ")

            div.field.submit
                button(type="button" @click="submit") 決定
                
    div(:class="['overlay', {'show':formVisible}]" @click="closeModal")
</template>

<script lang="ts" src="./ticketcreate.ts"></script>