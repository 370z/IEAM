<template>
  <div class="content-size">
    <layout-title-page
      title="การจัดการบัญชี โอนเงิน"
      details="สถานะการโอนเงินของแบบฟอร์มแต่ละประเภท"
    ></layout-title-page>
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
            <v-col cols="3"><h1 class="font-20">รายการทั้งหมด (100)</h1></v-col>
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
            {{ item.iron == 0 ? "รอการดำเนินการ" : "ดำเนินการสำเร็จ" }}
          </span>
        </template>
        <template v-slot:[`item.actions`]="{ item }">
          <v-btn color="primary" v-if="item.iron == 0" outlined rounded small>
            <v-icon left @click="del(item)"> mdi-file-image </v-icon
            >แนบรูปภาพ</v-btn
          >
          <v-btn color="primary" v-if="item.iron == 1" outlined rounded small>
            <v-icon left @click="del(item)">mdi-image-search </v-icon
            >ดูหลักฐาน</v-btn
          >
        </template>
      </v-data-table>
    </v-card>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card class="p-5">
        <p>หมายเลขฟอร์ม : 1312312313112</p>
        <v-img
          lazy-src="https://picsum.photos/id/11/10/6"
          max-width="250"
          src="https://picsum.photos/id/11/500/300"
        ></v-img>
        <v-container>
          <v-col cols="12" align="center">
            <v-btn
              depressed
              color="error"
              class="mr-4 rounded-pill"
              @click="close"
            >
              ยกเลิก
            </v-btn>
          </v-col>
        </v-container>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      search: null,
      dialog: true,
      items: [
        { value: null, text: "ทั้งหมด" },
        { value: "0", text: "รอการดำเนินการ" },
        { value: "1", text: "ดำเนินการสำเร็จ" },
      ],
      desserts:[],
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
          value: "calories",
          sortable: false,
          width: "10%",
        },
        { text: "ชื่อ-นามสกุล", value: "name", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "carbs",
          sortable: false,
          width: "15%",
        },
        { text: "จำนวนเงิน", value: "protein", sortable: false, width: "15%" },
        {
          text: "สถานะการโอน",
          align: "center",
          value: "iron",
          sortable: false,
          width: "10%",
        },
        {
          text: "หลักฐาน",
          align: "center",
          value: "actions",
          sortable: false,
          width: "15%",
        },
      ],
    };
  },
};
</script>

<style>
</style>