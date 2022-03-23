<template>
  <header class="bg-dark text-white">
    <div class="container-fluid py-2">
      <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
        <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
          <li>
            <router-link to="/home" class="nav-link px-2 text-secondary">Home</router-link>
          </li>
          <li v-if="isLoggedIn">
            <router-link to="/new-book" class="nav-link px-2 text-secondary">New Book</router-link>
          </li>
        </ul>

        <div class="text-end">
          <router-link to="/login" v-if="!isLoggedIn">
            <button class="btn btn-warning">
              Login
            </button>
          </router-link>
          <button class="btn btn-warning" v-if="isLoggedIn" v-on:click="onClickLogOut()">Logout</button>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import{mapActions, mapGetters} from "vuex";

export default {
  name: "Header",
  computed: {
    ...mapGetters(['isLoggedIn'])
  }, methods: {
    ...mapActions(['logout']),
    onClickLogOut() {
      this.logout();
      alert("Success logout!")
      this.$router.push("/home")
    }
  }
}
</script>

<style scoped>

</style>
