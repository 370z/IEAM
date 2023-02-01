<template>
  <div class="content-size">
    <layout-title-page
      title="Dashboard"
      details="ดูภาพรวมของตัวระบบ"
    ></layout-title-page>
    <div class="content">
      <v-row>
        <v-col v-if="this.$auth.state.user.level != 1">
          <c-card-dashboard
            :dashbord_data="dashbord_data[0]"
            :currency="items[0].currency"
            :color="items[0].color"
            :icon="items[0].icon"
          ></c-card-dashboard>
        </v-col>
        <v-col v-if="this.$auth.state.user.level != 1">
          <c-card-dashboard
            :dashbord_data="dashbord_data[1]"
            :currency="items[1].currency"
            :color="items[1].color"
            :icon="items[1].icon"
          ></c-card-dashboard>
        </v-col>
        <v-col>
          <c-card-dashboard
            :dashbord_data="dashbord_data[2]"
            :currency="items[2].currency"
            :color="items[2].color"
            :icon="items[2].icon"
          ></c-card-dashboard>
        </v-col>
        <v-col>
          <c-card-dashboard
            :dashbord_data="dashbord_data[3]"
            :currency="items[3].currency"
            :color="items[3].color"
            :icon="items[3].icon"
          ></c-card-dashboard>
        </v-col>
        <v-col>
          <c-card-dashboard
            :dashbord_data="dashbord_data[4]"
            :currency="items[4].currency"
            :color="items[4].color"
            :icon="items[4].icon"
          ></c-card-dashboard>
        </v-col>
      </v-row>
      <v-row v-if="this.$auth.state.user.level == 3">
        <v-col cols="3">
          <c-card-status :data="dashbord_numberform"></c-card-status>
        </v-col>
        <v-col cols="9">
          <c-cardnotreture :datatable="dashbord_notreturn"></c-cardnotreture>
        </v-col>
      </v-row>
      <v-row class="flex-column">
        <v-col>
          <v-card class="content-card" elevation="6">
            <v-row>
              <h2 class="mb-5">ประวัติทำรายการ</h2>
            </v-row>
              <table-history-user
                :datatable="desserts"
                v-if="this.$auth.state.user.level == 1"
              />
              <table-history-admin
                v-if="this.$auth.state.user.level != 1"
                :datatable="desserts"
              />
          </v-card>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
import moment from "moment";
import env from "~/env";
export default {
  data() {
    return {
      dashbord_notreturn: [],
      dashbord_data: {},
      dashbord_numberform:{},
      items: [
        {
          color: "#4caf4f",
          icon: "mdi-bank-transfer-in",
        },
        {
          color: "#ff3b2f",
          icon: "mdi-bank-transfer-out",
        },
        {
          color: "#f9ae00",
          icon: "mdi-clipboard-list",
        },
        {
          color: "#ffe000",
          icon: "mdi-text-box-minus-outline",
        },
        {
          color: "#009fe3",
          icon: "mdi-text-box-plus-outline",
        },
      ],
      desserts: [],
    };
  },
  mounted() {
    this.gethistory();
    this.getdashbord_historynotreturn();
    this.getdashbord_data();
  },
  methods: {
    async getdashbord_historynotreturn() {
      if (this.$auth.state.user.level != 1) {
        await this.$axios
          .get(`${env.BASE_URL}/dashbord/dashbord_historynotreturn`)
          .then((response) => {
            this.dashbord_notreturn = response.data?.data;
          });
      }
    },
    async getdashbord_data() {
      if (this.$auth.state.user.level == 1) {
        await this.$axios
          .get(`${env.BASE_URL}/dashbord/dashbord`, {
            params: { id: this.$auth.state.user.accountId },
          })
          .then((response) => {
            this.dashbord_data = response.data?.data;
          });
      } else {
        await this.$axios
          .get(`${env.BASE_URL}/dashbord/dashbord_all`)
          .then((response) => {
            this.dashbord_data = response.data?.data;
            this.dashbord_numberform = response.data?.data_numberform;
          });
      }
    },
    async gethistory() {
      if (this.$auth.state.user.level == 2) {
        await this.$axios
          .get(`${env.BASE_URL}/history/approve`)
          .then((response) => {
            this.desserts = response.data?.data;
          });
      } else if (this.$auth.state.user.level == 3) {
        await this.$axios
          .get(`${env.BASE_URL}/history/finance`)
          .then((response) => {
            this.desserts = response.data?.data;
          });
      } else {
        await this.$axios
          .get(`${env.BASE_URL}/history/user`, {
            params: { id: this.$auth.state.user.accountId },
          })
          .then((response) => {
            this.desserts = response.data?.data;
          });
      }
    },
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L LT");
      return result;
    },
  },
};
</script>
<style>
</style>