<template>
  <div class="d-flex justify-content-center h-100 mt-5">
    <div class="w-50 d-flex justify-content-center">
      <form class="w-75 mt-5" v-on:submit.prevent="submit">
        <h2 class="d-flex justify-content-center my-4">Add new book:</h2>
        <label class="text-dark mt-2">Title:</label>
        <input class="form-control" type="text" v-model="book.title"/>
        <label class="text-dark mt-2">Author Name:</label>
        <input class="form-control" type="text" v-model="book.authorName"/>
        <label class="text-dark mt-2">Issuance Date:</label>
        <input class="form-control" type="date" v-model="book.date_dto"/>
        <label class="text-dark mt-2">Genre:</label>
        <input class="form-control" type="text" v-model="book.genre"/>
        <button class="btn btn-dark w-100 my-3" type="submit">
          Add
        </button>
      </form>
    </div>


  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "NewBook",
  data() {
    return {
      book: {
        title: '',
        authorName: '',
        genre: '',
        date_dto: '0001-01-01'

      }
    }
  },
  methods: {
    ...mapActions(['createBook']),
    submit() {
      this.createBook(this.book).then(() => {
        this.$router.push("/home")
      })
    }
  }
}
</script>

<style scoped>

</style>
