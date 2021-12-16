module.exports = {
  publicPath: "/foodies/ui/",
  transpileDependencies: ["vuetify"],
  devServer: {
    proxy: {
      "^/foodies/api": {
        target: "http://localhost:8084",
        changeOrigin: true,
      },
    },
  },
};
