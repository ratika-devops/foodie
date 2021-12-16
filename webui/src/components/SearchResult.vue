<template>
  <div>
    <v-card color="#BB7760" style="padding: 10px" v-if="data && data.length > 0" elevation="15">
      <v-card-title class="white--text">Search results</v-card-title>
      <v-data-table
        :headers="headers"
        :items="data"
        item-key="name"
        class="elevation-5"
        color="green"
      ></v-data-table>
    </v-card>
  </div>
</template>

<script>
export default {
  props: ["data"],
  data: () => ({
    headers: [
      { text: "Truck Name", value: "applicant" },
      { text: "Address", value: "address" },
      { text: "Cousine", value: "fooditems" },
    ],
  }),
};
</script>
