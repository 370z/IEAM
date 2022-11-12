<template>
  <v-col>
    <v-card class="content-card" elevation="6">
      <v-form ref="form" @submit.prevent="onsubmit">
        <v-row>
          <v-col cols="12">
            <label for="">หัวข้อรายการ</label>
            <v-text-field
              v-model="list.title"
              placeholder="กรอกชื่อรายการ"
              dense
              outlined
              :rules="rules.req"
            ></v-text-field>
          </v-col>
          <v-col>
            <label for="">ประเภท</label>
            <v-select
              :items="type"
              item-value="id"
              item-text="name"
              v-model="list.typeId"
              placeholder="เลือกประเภท"
              dense
              outlined
              :rules="rules.req"
            ></v-select>
          </v-col>
          <v-col>
            <label for="">ยอดรวม</label>
            <v-text-field
              v-model="list.total"
              type="number"
              placeholder="กรอกยอดรวม"
              dense
              outlined
              :rules="rules.req"
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
              :items="list.list"
              :loading="false"
              hide-default-footer
              loading-text="Loading... Please wait"
              no-data-text="ไม่พบข้อมูล"
            >
              <template v-slot:[`item.actions`]="{ item }">
                <div>
                  <v-menu>
                    <template v-slot:activator="{ on, attrs }">
                      <v-icon v-bind="attrs" v-on="on"
                        >mdi-dots-vertical</v-icon
                      >
                    </template>
                    <v-list>
                      <v-list-item link>
                        <v-list-item-title @click="del(item)"
                          ><v-icon>mdi-delete</v-icon> ลบ</v-list-item-title
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
          <v-col style="cursor: pointer">
            <v-btn color="primary" depressed small @click="dialog = true"
              >เพิ่มรายการ</v-btn
            >
          </v-col>
        </v-row>
        <v-divider class="my-5" />
        <v-row>
          <v-col cols="12">
            <label for="">หลักฐาน</label>
          </v-col>
          <v-col cols="4">
            <v-file-input
              accept=".jpeg,.jpg,.png,image/jpeg,image/png"
              show-size
              counter
              label="รูปหลักฐาน"
              @change="upload"
              type="file"
              prepend-icon="mdi-file-image"
              :rules="rules.req"
            />
            <!-- <input
      type="file"
      accept=".jpeg,.jpg,.png,image/jpeg,image/png"
      aria-label="upload image button"
      @change="upload"
    /> -->
          </v-col>
        </v-row>
        <v-row class="space-center">
          <v-btn
            class="mx-2"
            min-width="10%"
            color="primary"
            type="submit"
            rounded
            depressed
            >บันทึกรายการ</v-btn
          >
          <v-btn
            class="mx-2"
            min-width="10%"
            color="error"
            rounded
            depressed
            outlined
            @click="$router.go(-1)"
            >กลับ</v-btn
          >
        </v-row>
      </v-form>
    </v-card>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span>{{ this.mode != "e" ? "เพิ่มรายการ" : "แก้ไขรายการ" }}</span>
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
          <v-form ref="formlist" v-model="isValid" @submit.prevent="onaddlist">
            <v-row>
              <v-col cols="12">
                <span>ขื่อรายการ</span>
                <v-text-field
                  v-model="detaillist.title"
                  outlined
                  dense
                  placeholder="ขื่อรายการ"
                  :rules="rules.req"
                />
              </v-col>
              <v-col cols="6">
                <span>จำนวน</span>
                <v-text-field
                  v-model="detaillist.number"
                  outlined
                  dense
                  placeholder="จำนวน"
                  @change="total()"
                  type="number"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col cols="6">
                <span>หน่วย</span>
                <v-text-field
                  v-model="detaillist.unit"
                  outlined
                  dense
                  placeholder="หน่วย"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col>
                <span>ราคาต่อหน่วย</span>
                <v-text-field
                  v-model="detaillist.unit_price"
                  outlined
                  dense
                  placeholder="ราคาต่อหน่วย"
                  suffix="บาท"
                  type="number"
                  @change="total()"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <span>ยอดรวม</span>
                <v-text-field
                  v-model="detaillist.total"
                  disabled
                  outlined
                  dense
                  placeholder="ยอดรวม"
                  suffix="บาท"
                  type="number"
                  :rules="rules.req"
                  hide-details=""
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="space-center">
              <v-btn class="ml-5" color="primary" type="submit">ยืนยัน</v-btn>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script>
import axios from "axios";
export default {
  props: {
    mode: {
      type: String,
      default: "c",
    },
  },
  data() {
    return {
      type: [
        { id: 1, name: "รายรับ" },
        { id: 2, name: "รายจ่าย" },
        { id: 3, name: "การเบิก" },
        { id: 4, name: "การยืม" },
        { id: 5, name: "การคืน" },
      ],
      isValid: null,
      dialog: false,
      isSelecting: false,
      selectedFile: null,
      list: {
        number_form: null,
        title: null,
        typeId: null,
        total: null,
        list: [],
        img_url: "",
        user: {
          accountId: 3,
          username: "dsfds",
        },
      },
      detaillistbackup: {},
      detaillist: {
        tirftal: "",
        number: 0,
        unit: "",
        unit_price: 0,
        total: 0,
      },
      headers: [
        {
          text: "รายการ",
          align: "center",
          value: "title",
          sortable: false,
        },
        {
          text: "จำนวน",
          value: "number",
          sortable: false,
        },
        { text: "หน่วย", value: "unit", sortable: false },
        { text: "ราคาต่อหน่วย", value: "unit_price", sortable: false },
        {
          text: "ยอดรวม",
          align: "center",
          value: "total",
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
      rules: {
        req: [(v) => !!v || "กรุณาใส่ข้อมูล"],
      },
    };
  },

  watch: {
    dialog() {
      if (!this.dialog) {
        this.detaillist = {
          tirftal: "",
          number: 1,
          unit: "",
          unit_price: 0,
          total: 0,
        };
      }
    },
  },
  mounted() {
    if (this.mode == "e") {
      this.$axios
        .get(`${process.env.BASE_URL}/form/detail`, {
          params: this.$route.params,
        })
        .then((response) => {
          this.list = response.data?.data[0];
        });
    }
  },
  methods: {
    onaddlist() {
      if (this.$refs.formlist.validate()) {
        this.list.list.push(this.detaillist);
        this.dialog = !this.dialog;
      }
    },
    upload(e) {
      this.selectedFile = e;
    },

    total() {
      this.detaillist.total =
        parseInt(this.detaillist.unit_price) * parseInt(this.detaillist.number);
    },

    del(item) {
      this.list.list.splice(this.list.list.indexOf(item), 1);
    },
    async onsubmit() {
      if (this.$refs.form.validate()) {
        if (this.mode != "e") {
          await this.$axios
            .post(`${process.env.BASE_URL}/form/add`, {
              ...this.list,
              total: + this.list.total,
            })
            .then((response) => {
              alert("เพิ่มข้อมูลสำเร็จ(รอการอนุมัติ)");
              this.$router.go(-1)
            });
        } else {
          await this.$axios
            .put(`${process.env.BASE_URL}/form/update`, {
              ...this.list,
              total: +this.list.total,
            })
            .then((response) => {
              alert("แก้ไขข้อมูลสำเร็จ(รอการอนุมัติ)");
              this.$router.go(-1)
            });
        }
      }
    },
    // edit(item) {
    //   this.mode = "e";
    //   var id = this.list.detaillist.indexOf(item);
    //   this.detaillist = JSON.parse(
    //     JSON.stringify(this.list.detaillist[id])
    //   );
    //   this.detaillistbackup = JSON.parse(
    //     JSON.stringify(this.list.detaillist[id])
    //   );
    //   // this.detaillist = this.list.detaillist[id];
    //   this.dialog = !this.dialog;
    // },
  },
};
</script>

<style>
</style>