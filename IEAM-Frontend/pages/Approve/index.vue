<template>
  <div class="content-size">
    <layout-title-page
      title="จัดการอนุมัติ"
      details="การอนุมัติเอกสาร รายรับ รายจ่าย เบิก ยืม คืน"
    ></layout-title-page>
    <v-col class="content">
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
                ><h1 class="font-20">
                  รายการทั้งหมด ( {{ account_wbr_approve.length }} )
                </h1></v-col
              >
              <v-spacer />
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
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.number_form`]="{ item }">
            <span
              style="
                cursor: pointer;
                color: #2096f3;
                text-decoration: underline;
              "
              @click="view(item)"
            >
              {{ item.number_form }}
            </span>
          </template>
          <template v-slot:[`item.name`]="{ item }">
            {{ item.user.firstname + " " + item.user.lastname }}
          </template>
          <template v-slot:[`item.total`]="{ item }">
            {{ item.total.toLocaleString("en-US") }}
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
          <template v-slot:[`item.actions`]="{ item }">
            <div v-if="item.isapprove == 0">
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="success"
                    rounded
                    x-small
                    fab
                    v-bind="attrs"
                    v-on="on"
                    elevation="false"
                    @click="
                      (dialog_approve = true), (confirm.id_form = item.id_form)
                    "
                  >
                    <v-icon> mdi-check </v-icon></v-btn
                  >
                </template>
                <span>อนุมัติ</span>
              </v-tooltip>
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    class="ml-2"
                    color="error"
                    rounded
                    x-small
                    fab
                    v-bind="attrs"
                    v-on="on"
                    elevation="false"
                    @click="
                      change_status('cancel', (confirm.id_form = item.id_form))
                    "
                  >
                    <v-icon>mdi-close </v-icon></v-btn
                  >
                </template>
                <span>ไม่อนุมัติ</span>
              </v-tooltip>
            </div>
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
              <!-- <v-icon @click="view(item)"> mdi-eye </v-icon> -->
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
              type="password"
              v-model="confirm.password"
              outlined
              dense
              hide-details
              placeholder="กรอกรห้สผ่าน"
            ></v-text-field>
            <v-btn
              class="ml-5"
              color="primary"
              :disabled="!confirm"
              @click="change_status('confirm')"
              >ยืนยัน</v-btn
            >
          </v-card-text>
        </v-card>
      </v-dialog>
      <v-dialog v-model="dialog" max-width="70%">
        <layout-dialog-transfermomey
          :detail="account_wbr_approve_detail"
          @closedialog="dialog = !dialog"
          :show="dialog"
        />
      </v-dialog>
    </v-col>
  </div>
</template>

<script>
import env from "~/env";
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
      confirm: {},
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
          value: "number_form",
          sortable: false,
          width: "10%",
        },
        { text: "ชื่อ-นามสกุล", value: "name", sortable: false, width: "15%" },
        { text: "จำนวนเงิน", value: "total", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type.nametype",
          sortable: false,
          width: "15%",
        },
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
    this.lv = [2];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
  },
  methods: {
    async gettransfermdetail() {
      await this.$axios
        .get(`${env.BASE_URL}/form/all`)
        .then((response) => {
          this.account_wbr_approve = response.data?.data;
        });
    },
    view(item) {
      this.account_wbr_approve_detail = item;
      this.dialog = !this.dialog;
    },
    async change_status(type) {
      if (type == "confirm") {
        await this.$axios
          .patch(`${env.BASE_URL}/approve/approve`, {
            ...this.confirm,
            isapprove: 1,
          })
          .then((response) => {
            this.dialog_approve = !this.dialog_approve;
            alert("ทำการอนุมัติเรียบร้อย");
            window.location.reload(true);
          })
          .catch((err) => {
            alert(err);
          });
      } else {
        if (confirm("ทำการไม่อนุมัติใช่ไหม")) {
          await this.$axios
            .patch(`${env.BASE_URL}/approve/approve`, {
              ...this.confirm,
              isapprove: 2,
            })
            .then((response) => {
              this.confirm;
              alert("ทำการไม่อนุมัติเรียบร้อย");
              window.location.reload(true);
            })
            .catch((err) => {
              alert(err);
            });
        }
      }
    },
  },
};
</script>

<style>
</style>