<template>
  <v-card class="d-flex justify-center mb-6" flat tile>
    <div class="neonText">Welcome Foodies !!!!</div>
  </v-card>
</template>

<script>
import Search from "./Search";
export default {
  components: {
    Search,
  },
  data: () => ({}),
};
</script>
<style scoped>
.neonText {
  text-shadow: 0 0 22px #fff, 0 0 10px #fff, 0 0 21px #fff,
    0 0 42px rgb(255, 0, 0), 0 0 82px rgb(255, 0, 0), 0 0 92px rgb(255, 38, 0),
    0 0 102px rgb(255, 115, 0), 0 0 51px rgb(255, 81, 0);
  position: absolute;
  z-index: 100;
  top: -450px;
  font-family: Corinthia;
  font-size: 120px;
  font-weight: bold;
  color: rgba(80, 45, 45, 0.904);
}
</style>
