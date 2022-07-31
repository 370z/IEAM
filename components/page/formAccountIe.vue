<template>
  <div>
    <v-card class="content-card" elevation="6">
      <v-row>
        <v-col>
          <label for="">หัวข้อรายการ</label>
          <v-text-field
            placeholder="กรอกชื่อรายการ"
            dense
            outlined
          ></v-text-field>
        </v-col>
        <v-col>
          <label for="">ประเภท</label>
          <v-text-field
            placeholder="กรอกชื่อรายการ"
            dense
            outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <label for="">รายการ</label>
        </v-col>
        <v-col cols="12">
          <v-data-table
            :headers="headers"
            :items="desserts"
            :loading="false"
            hide-default-footer
            loading-text="Loading... Please wait"
            no-data-text="ไม่พบข้อมูล"
          >
            <template v-slot:[`item.actions`]="{}">
              <div>
                <v-menu>
                  <template v-slot:activator="{ on, attrs }">
                    <v-icon v-bind="attrs" v-on="on">mdi-dots-vertical</v-icon>
                  </template>
                  <v-list>
                    <v-list-item
                      link
                      v-for="(item, index) in items"
                      :key="index"
                    >
                      <v-list-item-title
                        ><v-icon>{{ item.icon }}</v-icon>
                        {{ item.text }}</v-list-item-title
                      >
                    </v-list-item>
                  </v-list>
                </v-menu>
              </div>
            </template>
          </v-data-table>
        </v-col>
      </v-row>
      <v-row>
        <v-col style="cursor: pointer" @click="dialog = true">
          <v-icon>mdi-plus </v-icon> เพิ่มรายการ
        </v-col>
      </v-row>
      <v-col class="flex-column">
        <label for="">หลักฐาน</label>
        <v-btn color="primary" small depressed>แนบไฟล์</v-btn>
        <span>sdadasda.png</span>
      </v-col>
      <v-row class="space-center">
        <v-btn class="mx-2" min-width="10%" color="primary" rounded depressed
          >บันทึกรายการ</v-btn
        >
        <v-btn
          class="mx-2"
          min-width="10%"
          color="error"
          rounded
          depressed
          outlined
          >ยกเลิก</v-btn
        >
      </v-row>
    </v-card>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span>เพิ่มรายการ</span>
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
          <v-row>
            <v-col cols="12">
              <v-text-field
                outlined
                dense
                hide-details
                placeholder="กรอกรห้สผ่าน"
              />
            </v-col>
            <v-col>
              <v-text-field
                outlined
                dense
                hide-details
                placeholder="กรอกรห้สผ่าน"
              ></v-text-field>
              <v-text-field
                outlined
                dense
                hide-details
                placeholder="กรอกรห้สผ่าน"
              ></v-text-field>
            </v-col>
            <v-col>
              <v-text-field
                outlined
                dense
                hide-details
                placeholder="กรอกรห้สผ่าน"
              ></v-text-field>
              <v-text-field
                outlined
                dense
                hide-details
                placeholder="กรอกรห้สผ่าน"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row class="space-center">
            <v-btn class="ml-5" color="primary">ยืนยัน</v-btn>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialog: true,
      list: [],
      items: [
        { text: "แก้ไข", icon: "mdi-pencil" },
        { text: "ลบ", icon: "mdi-delete-forever" },
      ],
      headers: [
        {
          text: "รายการ",
          align: "center",
          value: "id",
          sortable: false,
        },
        {
          text: "จำนวน",
          value: "calories",
          sortable: false,
        },
        { text: "หน่วย", value: "name", sortable: false },
        { text: "ราคาต่อหน่วย", value: "protein", sortable: false },
        {
          text: "ยอดรวม",
          align: "center",
          value: "type",
          sortable: false,
          width: "10%",
        },
        {
          text: " ",
          align: "center",
          value: "actions",
          sortable: false,
          width: "5%",
        },
      ],
      desserts: [{ name: "sfsdfsdf" }],
    };
  },
};
</script>

<style>
</style>