module.exports = {
  configureWebpack: {
    devServer: {
      headers: { "Access-Control-Allow-Origin": "*" },
      proxy: {
        "/api": {
          target: "http://localhost:3000",
        },
      },
    },
  },
};
