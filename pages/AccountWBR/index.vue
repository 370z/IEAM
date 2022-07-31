<template>
  <div class="content-size">
    <layout-title-page
      title="การจัดการบัญชี รายรับ/รายจ่าย"
      details="การจัดการบัญชีรายรับ-รายจ่าย"
    ></layout-title-page>
    <div class="content">
      <v-card class="content-card" elevation="6">
        <v-data-table
          :headers="headers"
          :items="desserts"
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
          <template v-slot:[`item.iron`]="{ item }">
            <span :style="`color: ${item.iron == 0 ? '#FFC83B' : '#30C03E'}`">
              {{ item.iron == 0 ? "รอการดำเนินการ" : "อนุมัติเรียบร้อย" }}
            </span>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <div>
              <v-icon @click="del(item)">mdi-eye</v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
      <v-dialog v-model="dialog" max-width="500px">
        <v-card>
          <v-card-title>
            <span>ยืนยันรห้สผ่าน</span>
            <v-spacer></v-spacer>
            <v-menu bottom left>
              <template v-slot:activator="{ on, attrs }">
                <v-btn icon v-bind="attrs" v-on="on" @click="dialog = false">
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
    </div>
  </div>
</template>

<script>
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
      desserts:[],
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
          value: "calories",
          sortable: false,
          width: "10%",
        },
        { text: "ชื่อ-นามสกุล", value: "name", sortable: false, width: "20%" },
        {
          text: "ประเภท",
          align: "center",
          value: "carbs",
          sortable: false,
          width: "15%",
        },
        { text: "จำนวนเงิน", value: "protein", sortable: false, width: "10%" },
        {
          text: "สถานะการอนุมัติ",
          align: "center",
          value: "iron",
          sortable: false,
        },
        {
          text: "สถานะการโอน",
          align: "center",
          value: "iron",
          sortable: false,
        },
        {
          text: "เพิ่มเติม",
          align: "center",
          value: "actions",
          sortable: false,
          width: "7%",
        },
      ],
    };
  },
};
</script>

<style>
</style>