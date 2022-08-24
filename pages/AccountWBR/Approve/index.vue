<template>
  <div class="content-size">
    <layout-title-page
      title="การจัดการอนุมัติ เบิก/ยืม/คืน"
      details="การอนุมัติในการ เบิก/ยืม/คืน"
    ></layout-title-page>
    <div class="content">
      <v-card class="content-card" elevation="6">
        <v-data-table
          :headers="headers"
          :items="account_wbr_approve"
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
                <v-select
                  v-model="search"
                  :menu-props="{ offsetY: true }"
                  :items="items"
                  item-value="value"
                  item-text="text"
                  dense
                  outlined
                  rounded
                  hide-details
                ></v-select>
              </v-col>
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.isapprove`]="{ item }">
            <span
              :style="`color: ${item.isapprove == 0 ? '#FFC83B' : '#30C03E'}`"
            >
              {{ item.isapprove == 0 ? "รอการอนุมัติ" : "อนุมัติเรียบร้อย" }}
            </span>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <div v-if="item.isapprove == 0">
              <v-btn color="success" rounded x-small fab elevation="false">
                <v-icon @click="dialog_approve = true">
                  mdi-check
                </v-icon></v-btn
              >
              <v-btn
                class="ml-2"
                color="error"
                rounded
                x-small
                fab
                elevation="false"
              >
                <v-icon @click="view(item)">mdi-close </v-icon></v-btn
              >
            </div>
            <div v-if="item.isapprove == 1">
              <v-icon @click="view(item)"> mdi-eye </v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
      <v-dialog v-model="dialog_approve" max-width="500px">
        <v-card>
          <v-card-title>
            <span>ยืนยันรห้สผ่าน</span>
            <v-spacer></v-spacer>
            <v-menu bottom left>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  icon
                  v-bind="attrs"
                  v-on="on"
                  @click="dialog_approve = false"
                >
                  <v-icon>mdi-close</v-icon>
                </v-btn>
              </template>
            </v-menu>
          </v-card-title>
          <v-card-text class="d-flex">
            <v-text-field
              outlined
              dense
              hide-details
              placeholder="กรอกรห้สผ่าน"
            ></v-text-field>
            <v-btn class="ml-5" color="primary">ยืนยัน</v-btn>
          </v-card-text>
        </v-card>
      </v-dialog>
      <v-dialog v-model="dialog" max-width="70%">
        <layout-dialog-transfermomey
          :detail="account_wbr_approve_detail"
          @dialogview-click="dialog = !dialog"
        />
      </v-dialog>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      dialog_approve: false,
      dialog: false,
      search: null,
      items: [
        { value: null, text: "ทั้งหมด" },
        { value: "0", text: "รอการอนุมัติ" },
        { value: "1", text: "อนุมัติเรียบร้อย" },
      ],
      account_wbr_approve: [],
      account_wbr_approve_detail: {},
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "10%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "id_form",
          sortable: false,
          width: "10%",
        },
        { text: "ชื่อ-นามสกุล", value: "name", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type",
          sortable: false,
          width: "15%",
        },
        { text: "จำนวนเงิน", value: "total", sortable: false, width: "15%" },
        {
          text: "สถานะการอนุมัติ",
          align: "center",
          value: "isapprove",
          sortable: false,
          width: "10%",
        },
        {
          text: "อนุมัติ",
          align: "center",
          value: "actions",
          sortable: false,
          width: "15%",
        },
      ],
    };
  },
  mounted() {
    this.gettransfermdetail();
  },
  methods: {
    async gettransfermdetail() {
      await axios
        .get(`${process.env.BASE_URL}/transfermoney/all`)
        .then((response) => {
          console.log(response.data.data);
          this.account_wbr_approve = response.data?.data;
        });
    },
    view(item) {
      this.account_wbr_approve_detail = item;
      this.dialog = !this.dialog;
      console.log(this.account_wbr_approve_detail);
    },
  },
};
</script>

<style>
</style>