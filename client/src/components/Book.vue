<template>
  <div class="my-3" v-if="getSelectedBook !== null">
    <h2>Selected Book:</h2>
    <p class="text-dark">Title: {{getSelectedBook.title}}</p>
    <p class="text-dark">Author: {{getSelectedBook.author}}</p>
    <p class="text-dark">Genre: {{getSelectedBook.genre}}</p>
    <p class="text-dark">Issuance Date: {{getSelectedBook.issuanceDate}}</p>
    <router-link to="/home"><button class="btn btn-warning">Back</button></router-link>
  </div>
</template>

<script>
import {mapGetters, mapActions} from "vuex";

export default {
  name: "Book",
  computed: mapGetters(['getSelectedBook']),
  methods: mapActions(['fetchOneBook']),
  created() {
    const id = this.$route.params['id']
    this.fetchOneBook(id).catch(() => {
      this.$router.push("/forbidden")
    })

  }
}
</script>

<style scoped>

</style>
