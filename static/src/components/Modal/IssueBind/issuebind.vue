<template lang="pug">
div
    div(:class="['modal', {'show':formVisible}]")
        | Bind Issue
        hr
        form
            div.field
                div
                    select-component(v-model="selectedIssue" :options="issue_options")
            div.field.submit
                button(type="button" @click="submit") 決定
                
    div(:class="['overlay', {'show':formVisible}]" @click="closeModal")
</template>

<script lang="ts" src="./issuebind.ts"></script>