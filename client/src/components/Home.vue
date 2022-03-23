<template>
  <div>
    <h2 class="my-4 text-dark">Books:</h2>
    <table class="table table-bordered">
      <thead class="thead-dark">
      <tr>
        <th scope="col">#</th>
        <th scope="col">Title</th>
        <th scope="col">Author Name</th>
        <th scope="col">Issuance Date</th>
        <th scope="col">Genre</th>
        <th scope="col" v-if="isLoggedIn"></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="book in getBooks" v-bind:key="book.id">
        <th scope="row">{{ book.ID }}</th>
        <td>{{ book.title }}</td>
        <td>{{ book.authorName }}</td>
        <td>{{ parseDate(book.issuanceDate) }}</td>
        <td>{{ book.genre }}</td>
        <td v-if="isLoggedIn"><router-link :to="'/book/' + book.ID"><button class="btn btn-warning ">-></button></router-link></td>
      </tr>

      </tbody>
    </table>
  </div>
</template>

<script>


import {mapActions, mapGetters} from "vuex";

export default {
  name: "Home",
  computed: mapGetters(['getBooks', 'isLoggedIn'])
  ,
  methods: {
    ...mapActions(['fetchAllBooks']),
    parseDate(date) {
      if (date.indexOf('0001-01-01') !== -1) {
        return '';
      }
      const tokens = date.split('T')
      return tokens[0]
    }
  },
  created() {
    this.fetchAllBooks()
  }
}
</script>

<style scoped>

</style>
