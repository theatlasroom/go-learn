<script>
export default {
  name: "FilmItem",
  props: {
    film: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {};
  },
};
</script>
<template>
  <div class="film-item">
    <header>
      <h1 class="film-title">
        {{ film.title }}
        <span class="film-year">{{ `(` + film.year + `)` }}</span>
      </h1>
    </header>
    <section></section>
  </div>
</template>
