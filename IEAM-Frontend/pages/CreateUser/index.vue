<template>
  <div class="content-size">
    <layout-title-page
      title="เพิ่มผู้ใช้งาน"
      details="จัดการเพิ่ม ลบ แก้ไข ข้อมูลของผู้ใช้งาน"
    />
    <v-col class="content">
      <v-card class="content-card" elevation="6">
        <v-data-table
          :headers="headers"
          :items="userall"
          :search="search"
          :loading="false"
          loading-text="Loading... Please wait"
          no-data-text="ไม่พบข้อมูล"
        >
          <template v-slot:top>
            <v-row
              ><v-col cols="3"
                ><h1 class="font-20">ผู้ใช้งาน ( {{userall.length}} )</h1></v-col
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
                  @click="(dialog = !dialog), (mode = 'c')"
                >
                  <v-icon left>mdi-plus </v-icon>เพิ่มผู้ใช้งาน</v-btn
                >
              </v-col>
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.name`]="{ item }">
            {{ item.firstname + " " + item.lastname }}
          </template>
          <template v-slot:[`item.level`]="{ item }">
            {{
              item.level == 1
                ? "ผู้ใช้งาน"
                : item.level == 2
                ? "ผู้อนุมัติ"
                : "งานการเงิน"
            }}
          </template>
          <template v-slot:[`item.actions`]="{ index }">
            <div>
              <v-icon @click="view(index)" color="#2096f3">mdi-eye </v-icon>
              <v-icon @click="edit(index)" color="#FFC83B">mdi-pencil</v-icon>
              <v-icon @click="del(index)" color="#CD2126">mdi-delete </v-icon>
            </div>
          </template>
        </v-data-table>
      </v-card>
    </v-col>
    <v-dialog v-model="dialog" max-width="50%">
      <v-card>
        <v-card-title>
          <span> เพิ่มผู้ใช้งาน </span>
          <v-spacer></v-spacer>
          <v-menu bottom left>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on" @click="dialog = !dialog">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </template>
          </v-menu>
        </v-card-title>
        <v-card-text>
          <v-form ref="formlist" @submit.prevent="submit">
            <v-row>
              <v-col>
                <label for="">ชื่อผู้ใช้</label>
                <v-text-field
                  v-model="userdetail.username"
                  :disabled="this.mode == 'v' ? true : false"
                  placeholder="กรอกชื่อผู้ใช้งาน"
                  dense
                  outlined
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col>
                <label for="">รห้สผ่าน </label>
                <v-text-field
                  v-model="userdetail.password"
                  :disabled="this.mode == 'v' ? true : false"
                  type="password"
                  placeholder="กรอกรห้สผ่าน"
                  dense
                  outlined
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <label for="">ชื่อ</label>
                <v-text-field
                  v-model="userdetail.firstname"
                  :disabled="this.mode == 'v' ? true : false"
                  placeholder="กรอกชื่อผู้ใช้งาน"
                  dense
                  outlined
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col>
                <label for="">นามสกุล</label>
                <v-text-field
                  v-model="userdetail.lastname"
                  :disabled="this.mode == 'v' ? true : false"
                  placeholder="กรอกรห้สผ่าน"
                  dense
                  outlined
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <label for="">เบอร์โทร</label>
                <v-text-field
                  v-model="userdetail.phonenumber"
                  :disabled="this.mode == 'v' ? true : false"
                  placeholder="กรอกชื่อผู้ใช้งาน"
                  dense
                  outlined
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col>
                <label for="">สถานะ</label>
                <v-select
                  :items="level"
                  item-value="id"
                  item-text="name"
                  v-model="userdetail.level"
                  :disabled="this.mode == 'v' ? true : false"
                  placeholder="เลือกสถานะ"
                  dense
                  outlined
                ></v-select>
              </v-col>
            </v-row>
            <v-row class="space-center">
              <v-btn class="ml-5" color="primary" type="submit">ยืนยัน</v-btn>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import env from "~/env";
export default {
  data() {
    return {
      dialog: false,
      mode: "c",
      userall: [],
      search: null,
      userdetail: {
        username: null,
        password: null,
        firstname: null,
        lastname: null,
        phonenumber: null,
        level: 1,
      },
      level: [
        { id: 1, name: "ผู้ใช้งาน" },
        { id: 2, name: "ผู้อนุมัติ" },
        { id: 3, name: "งานการเงิน" },
      ],
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "5%",
        },

        { text: "ชื่อผู้ใช้", value: "username", sortable: false, width: "%" },
        { text: "ชื่อ", value: "name", sortable: false, width: "%" },
        { text: "เบอร์โทร", value: "phonenumber", sortable: false, width: "%" },
        {
          text: "สถานะ",
          align: "center",
          value: "level",
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
      rules: {
        req: [(v) => !!v || "กรุณาใส่ข้อมูล"],
      },
    };
  },
  watch: {
    dialog(val) {
      if (!val) {
        this.userdetail = {};
        this.$refs.formlist.reset()
      }
    },
  },
  mounted() {
    this.lv = [3];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
    this.getUserall();
  },
  methods: {
    async getUserall() {
      await this.$axios
        .get(`${env.BASE_URL}/user/`)
        .then((response) => {
          this.userall = response.data?.data;
        });
    },
    view(item) {
      this.mode = "v";
      this.userdetail = this.userall[item];
      this.dialog = !this.dialog;
    },
    edit(item) {
      this.mode = "e";
      this.userdetail = JSON.parse(JSON.stringify(this.userall[item]));
      this.dialog = !this.dialog;
    },
    async submit() {
      if (this.$refs.formlist.validate()) {
        if (this.mode == "c") {
          await this.$axios
            .post(`${env.BASE_URL}/user/create`, this.userdetail)
            .then((response) => {
              alert("เพิ่มข้อมูลสำเร็จ");
            });
        }
        if (this.mode == "e") {
          await this.$axios
            .patch(`${env.BASE_URL}/user/edite`, this.userdetail)
            .then((response) => {
              alert("แก้ไขข้อมูลสำเร็จ");
            });
        }
        this.dialog = !this.dialog;
        if (this.mode != "v") {
          this.getUserall();
        }
      }
    },
    async del(item) {
      if (confirm("ต้องการลบข้อมูลนี้ใช่ไหม")) {
        this.$axios
          .delete(`${env.BASE_URL}/user/delete`, {
            params: { id: this.userall[item].accountId },
          })
          .then((response) => {
            this.userall.splice(item, 1);
            alert("ลบข้อมูลสำเร็จ");
          });
      }
    },
  },
};
</script>

<style>
</style>