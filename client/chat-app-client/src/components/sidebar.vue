<template>
    <div id="sidebar">
        {{ username }}
        {{ id }}
    </div>
</template>

<script>

    export default {
        name: "sidebar", 
        props: [
            "username",
            "id"
        ]
    }
</script>

<style scoped>
    #sidebar {
        width: 400px;
        height: 100%;
        background-color: pink;
    }
</style>