<template>
  <div class="text-center setting-style">
    <v-menu offset-y>
      <template v-slot:activator="{ on, attrs }">
        <v-btn :color="color" dark v-bind="attrs" v-on="on" fab x-small>
          <v-icon dark> {{ icon }} </v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item>
          <v-list-item-title>
            <v-btn v-bind="attrs" v-on="on" block x-small @click="refreshCache">
              <v-icon small> mdi-cached </v-icon> &nbsp;&nbsp;&nbsp;Refresh
              Cache
            </v-btn>
          </v-list-item-title>
        </v-list-item>
        <v-list-item>
          <v-list-item-title>
            <v-btn
              v-bind="attrs"
              v-on="on"
              block
              x-small
              @click="openSwaggerDocs"
            >
              <v-icon small> mdi-open-in-new</v-icon> &nbsp;&nbsp;&nbsp;Swagger
              Docs
            </v-btn>
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data: () => ({
    icon: "mdi-cog-outline",
    color: "primary",
  }),
  methods: {
    refreshCache() {
      let self = this;
      self.icon = "mdi-spin mdi-cog-outline";
      self.color = "orange";
      const uri = "/foodies/api/v1/cache/refresh";
      axios
        .get(uri)
        .then((resp) => {
          console.log("Cache refresh successful: ", resp.data);
          self.color = "green";
          self.icon = "mdi-cog-outline";
        })
        .catch((error) => {
          console.log("Cache refresh error: ", error);
          self.color = "red";
          self.icon = "mdi-cog-outline";
        })
        .finally((error) => {
          setTimeout(() => {
            self.icon = "mdi-cog-outline";
            self.color = "primary"
          }, 3000);
        });
    },
    openSwaggerDocs() {
      window.open("/foodies/swagger/index.html");
    },
  },
};
</script>

<style scoped>
.setting-style {
  padding-left: 50px;
  padding-top: 15px;
  left: 100px !important;
  margin: 0;
}
</style>
