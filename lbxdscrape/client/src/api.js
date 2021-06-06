export const fetchFilms = async () => {
  const resp = await fetch("/api/films");
  return await resp.json();
};
