<script>
import { fetchFilms } from "../api.js";
import FilmItem from "./FilmItem.vue";

export default {
  name: "FilmsList",
  components: {
    FilmItem,
  },
  props: {},
  created() {
    this.fetchData();
  },
  mounted() {},
  data() {
    return {
      films: [],
      loading: false,
    };
  },
  methods: {
    fetchData: async function() {
      this.loading = true;
      try {
        const data = await fetchFilms();
        this.films = data;
        this.loading = false;
      } catch (err) {
        this.loading = false;
        console.error(err);
      }
    },
  },
};
</script>
<template>
  <section id="films">
    <ul v-if="films.length">
      <li v-for="film in films" :key="film.id">
        <film-item :film="film" />
      </li>
    </ul>
    <p v-else-if="loading">{{ "loading..." }}</p>
    <p v-else>No films available</p>
  </section>
</template>
