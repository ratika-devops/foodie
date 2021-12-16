<template>
  <div style="paddint-right: 30px">
    <v-card
      class="d-flex justify-center mb-6"
      style="border-radius: 15px !important"
      elevation="15"
      tile
    >
      <v-radio-group
        v-model="searchBy"
        row
        style="padding-right: 50px !important; padding-left: 120px !important"
      >
        <h3>Search by:</h3>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <v-radio label="FoodTruck" value="truck"></v-radio>
        <v-radio label="Cousine" value="cousine"></v-radio>
        <v-radio label="Location" value="address"></v-radio>
      </v-radio-group>
      <Settings />
    </v-card>
    <v-card
      class="d-flex justify-center mb-6"
      style="border-radius: 15px !important"
      flat
      tile
      elevation="15"
    >
      <v-icon large style="padding-left: 20px; padding-right: 10px; top: -5px">
        mdi-magnify
      </v-icon>

      <div class="search-style">
        <v-text-field
          v-model="searchWhat"
          solo
          label='Type something and click on "Search" button'
        ></v-text-field>
      </div>
      <v-btn
        class="white--text"
        color="red"
        style="top: 27px"
        large
        @click="onSearch"
        type="submit"
        elevation="15"
      >
        Search</v-btn
      >&nbsp;&nbsp;&nbsp;
    </v-card>
  </div>
</template>

<script>
import Settings from "./Settings";

export default {
  components: {
    Settings,
  },
  data: () => ({
    descriptionLimit: 60,
    entries: [],
    isLoading: false,
    searchWhat: null,
    searchBy: "truck",
  }),

  methods: {
    onSearch() {
      this.$emit("onSearch", this.searchBy, this.searchWhat);
    },
  },
};
</script>

<style scoped>
.search-style {
  width: 600px;
  padding-top: 25px;
  padding-right: 10px;
  margin: 0;
}
</style>
