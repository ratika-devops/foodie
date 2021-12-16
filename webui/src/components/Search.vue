<template>
  <v-card class="d-flex justify-center mb-6" flat tile>
    <SearchBox class="search" @onSearch="onSearch" />
    <SearchResult :data="searchResult" />
    <v-card
      color="red"
      style="padding: 10px"
      class="white--text"
      v-if="errorMessage"
    >
      <h3>{{ errorMessage }}</h3>
    </v-card>
  </v-card>
</template>
<script>
import SearchBox from "./SearchBox";
import SearchResult from "./SearchResult";
import axios from "axios";

export default {
  components: {
    SearchBox,
    SearchResult,
  },
  data: () => ({
    searchResult: null,
    errorMessage: null,
  }),
  methods: {
    onSearch(searchBy, searchWhat) {
      let self = this;
      if (!searchWhat || searchWhat.length === 0) {
        this.errorMessage =
          "Please type something in search box before clicking search button";
        setTimeout(() => {
          self.errorMessage = null;
        }, 4000);
        return;
      }
      const uri = "/foodies/api/v1/" + "searchby/" + searchBy.trim() + "/" + searchWhat.trim();
      axios.get(uri).then((resp) => {
        self.searchResult = resp.data;
        if (!resp.data || resp.data.length === 0) {
          this.errorMessage =
            "Sorry, no match found for catagory '" +
            searchBy +
            "' when searched for '" +
            searchWhat +
            "'";
          setTimeout(() => {
            self.errorMessage = null;
          }, 6000);
        } else {
          self.errorMessage = null;
        }
      });
    },
  },
};
</script>

<style scoped>
.search {
  position: absolute;
  z-index: 100;
  top: -250px;
}
</style>
