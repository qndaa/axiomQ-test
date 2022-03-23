<template>
  <div class="d-flex justify-content-center h-100 mt-5">
    <div class="w-50 d-flex justify-content-center mt-5">
      <form class="w-75 mt-5" v-on:submit.prevent="submit">
        <h2 class="d-flex justify-content-center my-4">Login!</h2>
        <label class="text-dark">Username:</label>
        <input class="form-control" type="text" v-model="loginParams.username"/>
        <label class="text-dark">Password:</label>
        <input class="form-control" type="password" v-model="loginParams.password"/>
        <button class="btn btn-dark w-100 my-3" type="submit">
          Prijava
        </button>
      </form>
    </div>


  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Login",
  data() {
    return {
      loginParams: {
        username: '',
        password: ''
      }
    }
  }, methods: {
    ...mapActions(['login', 'setToken']),
    submit(event) {
      event.preventDefault()
      this.login(this.loginParams).then((response) => {
        this.setToken(response.data);
        alert("Success login!");
        this.$router.push("/home")
      }).catch(() => {
        alert("Invalid credentials!");
      })
    }
  }
}
</script>

<style scoped>

</style>
