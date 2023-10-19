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

                    const userToken = response.data.token;
                    localStorage.setItem("userToken", userToken);
                    this.$router.push("/home");

                } catch (error) {
                    console.error(error);
                }
           }
        }
    }
</script>

<style scoped>

</style>