<script>
import FilmsList from "./components/films/List.vue";

export default {
  name: "App",
  components: {
    FilmsList,
  },
};
</script>
<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <films-list></films-list>
  </div>
</template>
