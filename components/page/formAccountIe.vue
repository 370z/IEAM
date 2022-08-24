<template>
  <div>
    <v-card class="content-card" elevation="6">
      <v-row>
        <v-col cols="12">
          <label for="">หัวข้อรายการ</label>
          <v-text-field
            v-model="list.title"
            placeholder="กรอกชื่อรายการ"
            dense
            outlined
          ></v-text-field>
        </v-col>
        <v-col>
          <label for="">ประเภท</label>
          <v-text-field
            v-model="list.type"
            placeholder="กรอกชื่อรายการ"
            dense
            outlined
          ></v-text-field>
        </v-col>
        <v-col>
          <label for="">ยอดรวม</label>
          <v-text-field
            v-model="list.total"
            type="number"
            placeholder="กรอกยอดรวม"
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
            :items="list.detaillist"
            :loading="false"
            hide-default-footer
            loading-text="Loading... Please wait"
            no-data-text="ไม่พบข้อมูล"
          >
            <template v-slot:[`item.actions`]="{ item }">
              <div>
                <v-menu>
                  <template v-slot:activator="{ on, attrs }">
                    <v-icon v-bind="attrs" v-on="on">mdi-dots-vertical</v-icon>
                  </template>
                  <v-list>
                    <!-- <v-list-item link>
                      <v-list-item-title @click="edit(item)"
                        ><v-icon>mdi-pencil</v-icon> แก้ไข</v-list-item-title
                      >
                    </v-list-item> -->
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
        <v-col style="cursor: pointer">
          <v-btn
            color="primary"
            depressed
            small
            :loading="isSelecting"
            @click="handleFileImport"
            >แนบไฟล์</v-btn
          >
          <div class="mt-2">
            {{ this.selectedFile ? this.selectedFile.name : "" }}
          </div>
        </v-col>
        <input
          ref="uploader"
          class="d-none"
          type="file"
          @change="onFileChanged"
        />
      </v-row>
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
          <v-row>
            <v-col cols="12">
              <span>ขื่อรายการ</span>
              <v-text-field
                v-model="detaillist.title"
                outlined
                dense
                hide-details
                placeholder="ขื่อรายการ"
              />
            </v-col>
            <v-col cols="6">
              <span>จำนวน</span>
              <v-text-field
                v-model="detaillist.number"
                outlined
                dense
                hide-details
                placeholder="จำนวน"
                @change="total()"
                type="number"
              ></v-text-field>
            </v-col>
            <v-col cols="6">
              <span>หน่วย</span>
              <v-text-field
                v-model="detaillist.unit"
                outlined
                dense
                hide-details
                placeholder="หน่วย"
              ></v-text-field>
            </v-col>
            <v-col>
              <span>ราคาต่อหน่วย</span>
              <v-text-field
                v-model="detaillist.unit_price"
                outlined
                dense
                hide-details
                placeholder="ราคาต่อหน่วย"
                suffix="บาท"
                type="number"
                @change="total()"
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <span>ยอดรวม</span>
              <v-text-field
                v-model="detaillist.total"
                disabled
                outlined
                dense
                hide-details
                placeholder="ยอดรวม"
                suffix="บาท"
                type="number"
              ></v-text-field>
            </v-col>
          </v-row>
          {{ detaillist }}
          <v-row class="space-center">
            <v-btn class="ml-5" color="primary" @click="addlist()"
              >ยืนยัน</v-btn
            >
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
      mode: "",
      dialog: false,
      isSelecting: false,
      selectedFile: null,
      list: {
        tital: null,
        type: null,
        total: null,
        detaillist: [],
        img_url: "",
      },
      detaillistbackup: {},
      detaillist: {
        tirftal: "",
        number: 0,
        unit: "",
        unit_price: 0,
        total: 0,
      },
      items: [
        { text: "แก้ไข", icon: "mdi-pencil", func: "" },
        { text: "ลบ", icon: "mdi-delete-forever", func: "del(item.index)" },
      ],
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
  methods: {
    addlist() {
      if (this.mode != "e") {
        this.list.detaillist.push(this.detaillist);
      }
      // else {
      //     console.log("Fsdf");
      //   var id = this.detaillist.indexOf(this.detaillist);
      //   console.log(id);
      //   this.list.detaillist[id] = this.detaillist
      //   this.mode = "";
      // }
      this.dialog = !this.dialog;
    },
    handleFileImport() {
      this.isSelecting = true;
      window.addEventListener(
        "focus",
        () => {
          this.isSelecting = false;
        },
        { once: true }
      );
      this.$refs.uploader.click();
    },
    onFileChanged(e) {
      this.selectedFile = e.target.files[0];
    },

    total() {
      this.detaillist.total =
        parseInt(this.detaillist.unit_price) * parseInt(this.detaillist.number);
    },

    del(item) {
      this.list.detaillist.splice(this.list.detaillist.indexOf(item), 1);
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