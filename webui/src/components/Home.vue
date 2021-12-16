<template>
  <div class="mx-auto" max-width="100%">
    <v-img src="@/assets/food.jpg" height="500px"></v-img>
    <Banner />
    <Search />
  </div>
</template>

<script>
import Search from "./Search";
import Banner from "./Banner";

export default {
  components: {
    Search,
    Banner,
  },
  data: () => ({}),
};
</script>
