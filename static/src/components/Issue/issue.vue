<template lang="pug">
li
    div(:class="['state', issue.state]")
        svg(v-if="issue.state == 'open'" height="16" version="1.1" viewbox="0 0 16 16" width="16")
            path(fill-rule="evenodd" d="M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z")
        svg(v-if="issue.state == 'closed'" height="16" version="1.1" viewbox="0 0 14 16" width="14")
            path(fill-rule="evenodd" d="M7 10h2v2H7v-2zm2-6H7v5h2V4zm1.5 1.5l-1 1L12 9l4-4.5-1-1L12 7l-1.5-1.5zM8 13.7A5.71 5.71 0 0 1 2.3 8c0-3.14 2.56-5.7 5.7-5.7 1.83 0 3.45.88 4.5 2.2l.92-.92A6.947 6.947 0 0 0 8 1C4.14 1 1 4.14 1 8s3.14 7 7 7 7-3.14 7-7l-1.52 1.52c-.66 2.41-2.86 4.19-5.48 4.19v-.01z")
            
    div.path
        span.repo {{issue.repo}}
        br
        a(target="_blank" :href="issue.url") {{issue.title}}
        
    div.remove
        i.fa.fa-remove(@click="unbindIssue")
</template>

<script lang="ts" src="./issue.ts"></script>
<style lang="scss" scoped src="./issue.scss"></style>