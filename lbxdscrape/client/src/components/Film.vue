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
  methods: {
    altText(film) {
      return `${film.title} (${film.year})`;
    },
    imageUrl(film) {
      // TODO: send these to our own BE endpoint
      // We should resolve the posters either by:
      // - caching from lbxd 😬
      // - resolve from another provider + cache ie. free movie db or similar
      return `http://letterboxd.com${film.posterUrl}`;
    },
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
    <section>
      <img :src="imageUrl(film)" :alt="altText(film)" />
    </section>
  </div>
</template>
