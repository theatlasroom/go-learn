<script>
import FilmsList from "./components/FilmList.vue";

export default {
  name: "App",
  components: {
    FilmsList,
  },
};
</script>
<template>
  <div id="app">
    <header>
      <section class="title"></section>
      <section class="summary-stats">
        <!-- TODO: stats like % watched, total items in watched list, total watched -->
      </section>
    </header>
    <main>
      <!-- inifinite scroll, probably using an intersection observer to auto fetch the next page on scroll -->
      <films-list />
    </main>
  </div>
</template>
