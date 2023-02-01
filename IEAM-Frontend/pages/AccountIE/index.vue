<template>
  <div class="content-size">
    <layout-title-page
      title="จัดการ รายรับ/รายจ่าย"
      details="การจัดการบัญชีรายรับ-รายจ่าย"
    />
    <v-col class="content">
      <v-card class="content-card" elevation="6">
        <v-data-table
          :headers="headers"
          :items="form"
          :search="search"
          :loading="false"
          loading-text="Loading... Please wait"
          no-data-text="ไม่พบข้อมูล"
        >
          <template v-slot:top>
            <v-row>
              <v-col cols="3"
                ><h1 class="font-20">
                  รายการทั้งหมด ( {{ form.length }} )
                </h1></v-col
              ></v-row
            >
            <v-row>
              <v-col cols="3">
                <v-text-field
                  v-model="search"
                  append-icon="mdi-magnify"
                  placeholder="Search"
                  dense
                  outlined
                  rounded
                  hide-details
                ></v-text-field>
              </v-col>
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
                  <v-icon left>mdi-plus </v-icon>เพิ่มรายการ</v-btn
                >
              </v-col>
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.total`]="{ item }">
            {{ item.total.toLocaleString("en-US") }}
          </template>
          <template v-slot:[`item.type`]="{ item }">
            <span
              :style="`color: ${item.type.typeId == 2 ? '#EB3C41' : '#34ACFF'}`"
            >
              {{ item.type.typeId == 2 ? "รายจ่าย" : "รายรับ" }}
            </span>
          </template>
          <template v-slot:[`item.CreatedAt`]="{ item }">
            <div>
              {{ relativeTime(item.CreatedAt) }}
            </div>
          </template>
          <template v-slot:[`item.isapprove`]="{ item }">
            <span
              :style="`color: ${
                item.isapprove == 0
                  ? '#FFC83B'
                  : item.isapprove == 1
                  ? '#48A451'
                  : '#CD2126'
              }`"
            >
              {{
                item.isapprove == 0
                  ? "รอการอนุมัติ"
                  : item.isapprove == 1
                  ? "อนุมัติเรียบร้อย"
                  : "ไม่อนุมัติ"
              }}
            </span>
          </template>
          <template v-slot:[`item.actions`]="{ item, index }">
            <div>
              <div v-if="item.isapprove != 0">
                <v-btn
                  class="ml-2"
                  color="blue"
                  style="color: #ffffff"
                  small
                  rounded
                  elevation="false"
                  @click="view(item)"
                >
                  <v-icon>mdi-eye</v-icon> ดูรายละเอียด</v-btn
                >
              </div>
              <v-icon
                v-if="item.isapprove == 0"
                @click="view(item)"
                color="#2096f3"
                >mdi-eye
              </v-icon>
              <v-icon
                v-if="item.isapprove == 0"
                @click="edit(item.id_form)"
                color="#FFC83B"
                >mdi-pencil</v-icon
              >
              <v-icon
                v-if="item.isapprove == 0"
                @click="del(item.id_form, index)"
                color="#CD2126"
                >mdi-delete
              </v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
    </v-col>
    <v-dialog v-model="dialog" max-width="70%">
      <layout-dialog-transfermomey
        :detail="formdetail"
        :show="dialog"
        @closedialog="dialog = !dialog"
      />
    </v-dialog>
  </div>
</template>

<script>
import moment from "moment";
import env from "~/env";
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
          value: "number_form",
          sortable: false,
          width: "10%",
          align: "center",
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
          text: "สถานะการอนุมัติ",
          align: "center",
          value: "isapprove",
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
      form: [],
      formdetail: {},
    };
  },
  mounted() {
    this.gettransfermdetail();
    this.lv = [3];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
  },
  methods: {
    async gettransfermdetail() {
      await this.$axios
        .get(`${env.BASE_URL}/ie/all`)
        .then((response) => {
          this.form = response.data?.data;
        });
    },
    view(item) {
      this.formdetail = item;
      this.dialog = !this.dialog;
    },
    edit(item) {
      this.$router.push(`AccountIE/${item}`);
    },
    async del(item, index) {
      if (confirm("ต้องการลบข้อมูลนี้ใช่ไหม")) {
        this.$axios
          .delete(`${env.BASE_URL}/form/delete`, {
            params: { id: item },
          })
          .then((response) => {
            alert("ลบข้อมูลสำเร็จ");
            this.form.splice(index, 1);
          });
      }
    },
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L");
      return result;
    },
    dialogviewclick() {
      this.dialog = item;
    },
  },
};
</script>

<style>
</style>