<script>
import FilmsList from "./components/FilmsList.vue";
import PageTitle from "./components/PageTitle.vue";
import SummaryStats from "./components/SummaryStats.vue";

export default {
  name: "App",
  components: {
    PageTitle,
    SummaryStats,
    FilmsList,
  },
};
</script>
<template>
  <div id="app">
    <header>
      <page-title />
      <summary-stats />
    </header>
    <main>
      <!-- inifinite scroll, probably using an intersection observer to auto fetch the next page on scroll -->
      <films-list />
    </main>
  </div>
</template>
