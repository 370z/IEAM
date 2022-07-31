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
                <v-btn color="#E39257" dark block depressed rounded to="accountIE/addlist">
                  <v-icon  left @click="del(item)">mdi-plus </v-icon
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
          <template v-slot:[`item.actions`]="{  }">
            <div>
              <v-icon @click="dialog = true">mdi-eye</v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
      <v-dialog v-model="dialog" max-width="70%">
        <v-card>
          <v-card-title>
            <span>รายจ่าย</span>
            <v-spacer></v-spacer>
            <v-menu bottom left>
              <template v-slot:activator="{ on, attrs }">
                <v-btn icon v-bind="attrs" v-on="on" @click="dialog = false">
                  <v-icon>mdi-close</v-icon>
                </v-btn>
              </template>
            </v-menu>
          </v-card-title>
          <v-card-text>
            <v-row class="d-flex">
              <v-col>
                <p>หัวเรื่อง : รายจ่ายประจำเดือน มิถุนายน ปีการศึกษา2565</p>
                <p>เลขใบฟอร์ม : EP487539</p>
              </v-col>
              <v-col class="ml-10">
                <p>วันที่ : 2022-04-22</p>
                <p>จำนวนเงิน : 6,821.00 บาท</p>
              </v-col>
            </v-row>
            <v-col>
              <v-simple-table>
                <template v-slot:default>
                  <thead>
                    <tr>
                      <th class="text-left">รายการ</th>
                      <th class="text-left">จำนวน</th>
                      <th class="text-left">หน่วย</th>
                      <th class="text-left">ราคา/หน่วย</th>
                      <th class="text-left">รวม</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="item in form.list" :key="item.name">
                      <td>{{ item.title }}</td>
                      <td>{{ item.number }}</td>
                      <td>{{ item.unit }}</td>
                      <td>{{ item.unittonumber }}</td>
                      <td>{{ item.total }}</td>
                    </tr>
                  </tbody>
                </template>
              </v-simple-table>
            </v-col>
            <v-col>
              <p>เอกสารแนบ :</p>
              <v-img
                lazy-src="https://picsum.photos/id/11/10/6"
                max-height="150"
                max-width="250"
                src="https://picsum.photos/id/11/500/300"
              ></v-img>
            </v-col>
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
      dialog: true,
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
          value: "calories",
          sortable: false,
          width: "10%",
        },
        { text: "รายการ", value: "name", sortable: false, width: "%" },
        { text: "จำนวนเงิน", value: "protein", sortable: false, width: "15%" },
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
          value: "iron",
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
      desserts:[],
      form: {
        type: "ikp0jkp",
        number: "sfd",
        total: 100000,
        date: "2022-01-02",
        img: "wqwdsdfsdfsdfsdfsdf.com",
        list: [
          {
            title: "sdsadasdasd",
            number: 50,
            unit: 20,
            total: 20000,
            unittonumber: 10,
          },
          {
            title: "sdsadasdasd",
            number: 50,
            unit: 20,
            total: 20000,
            unittonumber: 10,
          },
        ],
      },
    };
  },
};
</script>

<style>
</style>