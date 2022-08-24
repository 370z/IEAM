<template>
  <div class="content-size">
    <layout-title-page
      title="การจัดการบัญชี รายรับ/รายจ่าย"
      details="การจัดการบัญชีรายรับ-รายจ่าย"
    />
    <div class="content">
      <v-card class="content-card" elevation="6">
        <v-data-table
          :headers="headers"
          :items="account_ie"
          :search="search"
          :loading="false"
          loading-text="Loading... Please wait"
          no-data-text="ไม่พบข้อมูล"
        >
          <template v-slot:top>
            <v-row>
              <v-col cols="3"
                ><h1 class="font-20">รายการทั้งหมด (100)</h1></v-col
              >
              <v-spacer />
              <v-col cols="2">
                <v-btn
                  color="#E39257"
                  dark
                  block
                  depressed
                  rounded
                  to="accountIE/addlist"
                >
                  <v-icon left >mdi-plus </v-icon
                  >เพิ่มรายการ</v-btn
                >
              </v-col>
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.type`]="{ item }">
            <span :style="`color: ${item.iron == 0 ? '#EB3C41' : '#34ACFF'}`">
              {{ item.iron == 0 ? "รายจ่าย" : "รายรับ" }}
            </span>
          </template>
          <template v-slot:[`item.CreatedAt`]="{ item }">
            <div>
              {{ relativeTime(item.CreatedAt) }}
            </div>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <div>
              <v-icon @click="view(item)">mdi-eye </v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
      <v-dialog v-model="dialog" max-width="70%">
        <layout-dialog-transfermomey :detail="account_iedetail" />
      </v-dialog>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
export default {
  data() {
    return {
      dialog: false,
      search: null,
      items: [
        { value: null, text: "ทั้งหมด" },
        { value: "0", text: "รอการอนุมัติ" },
        { value: "1", text: "อนุมัติเรียบร้อย" },
      ],
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "5%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "id_form",
          sortable: false,
          width: "10%",
        },
        { text: "รายการ", value: "title", sortable: false, width: "%" },
        { text: "จำนวนเงิน", value: "total", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type",
          sortable: false,
          width: "10%",
        },
        {
          text: "วัน",
          align: "center",
          value: "CreatedAt",
          sortable: false,
          width: "10%",
        },
        {
          text: "เพิ่มเติม",
          align: "center",
          value: "actions",
          sortable: false,
          width: "10%",
        },
      ],
      account_ie: [],
      account_iedetail: {},
    };
  },
  mounted() {
    this.gettransfermdetail();
  },
  methods: {
    async gettransfermdetail() {
      await axios
        .get(`${process.env.BASE_URL}/income-expenses/all`)
        .then((response) => {
          console.log(response.data.data);
          this.account_ie = response.data?.data;
        });
    },
    view(item) {
      this.account_iedetail = item;
      this.dialog = !this.dialog;
    },

    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format('L');
      return result;
    },
  },
};
</script>

<style>
</style>