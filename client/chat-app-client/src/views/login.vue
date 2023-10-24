<template>
    <div class="login">
        <h2>Login</h2>
        <form @submit.prevent="login">
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" v-model="username" required>
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" v-model="password" required> 
            </div>
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                username: "",
                password: "",
            }
        },
        methods: {
           async login() {
                try {
                    const response = await this.$axios.post("http://localhost:2000/user/login",{
                        username: this.username,
                        password: this.password
                    });

                    const userToken = response.data;
                    this.setCookie("userToken", userToken, 30);
                    this.$router.push("/home");

                } catch (error) {
                    console.error(error);
                }
           },
                setCookie(name, value, days) {
                    const date = new Date();
                    date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
                    const expires = "expires=" + date.toUTCString();
                    document.cookie = name + "=" + value + ";" + expires + ";path=/";
           }
        }
    }
</script>

<style scoped>

</style>