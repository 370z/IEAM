<template>
  <v-app>
    <v-navigation-drawer
      v-if="this.$auth.state.loggedIn"
      v-model="mini"
      mini-variant
      mini-variant-width="70"
      app
      style="border-right: 5px solid #eb3c41; drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));"
    >
      <v-list-item class="px-5">
        <v-img :src="require('../assets/images/login_logo.png')" />
      </v-list-item>

      <v-divider class="ma-auto" style="width: 70%"></v-divider>
      <v-list>
        <v-list-item-group class="groupManu">
          <v-row v-for="item in items" :key="item.title">
            <!-- v-if="item.lv.includes($auth.state.user.level)" -->
            <v-list-item
              v-if="item.lv.includes($auth.state.user.level)"
              :to="item.to"
              :class="item.to === $route.path ? 'activeManu' : '' + 'manu'"
              active-class="activeManu"
              exact
            >
              <v-tooltip right>
                <template v-slot:activator="{ on, attrs }">
                  <v-icon v-bind="attrs" v-on="on">
                    {{ item.icon }}
                  </v-icon>
                </template>
                <span>{{ item.title }}</span>
              </v-tooltip>
            </v-list-item>
          </v-row>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      elevation="1"
      v-if="this.$auth.state.loggedIn"
      dense
      app
      style="border-bottom: 2px solid #eb3c41; background: #ffffff"
    >
      <v-toolbar-title>
        <v-icon class="mx-2 my-2" @click.stop="mini = !mini">{{
          !mini ? "mdi-chevron-right" : "mdi-chevron-left"
        }}</v-icon
        ><span style="font-size: 15px">{{ title }}</span>
      </v-toolbar-title>
      <v-spacer />
      <span>{{ name }}</span
      ><v-icon @click="logout">mdi-login-variant</v-icon>
    </v-app-bar>

    <!-- <v-container class="container--fluid" style="margin-top: 54px"> -->
    <v-main>
      <Nuxt />
    </v-main>
    <!-- </v-container> -->
  </v-app>
</template>

<script>
export default {
  name: "DefaultLayout",
  data() {
    return {
      name: "",
      mini: true,
      IS_LOGIN: false,
      items: [
        {
          icon: "mdi-layers",
          title: "Dashboard",
          to: "/dashboard",
          lv: [0, 1, 2],
        },
        {
          icon: "mdi-cash-sync",
          title: "การจัดการบัญชี โอนเงิน",
          to: "/transfermoney",
          lv: [2],
        },
        {
          icon: "mdi-cash-multiple",
          title: "การจัดการบัญชี รายรับ/รายจ่าย",
          to: "/accountIE",
          lv: [2],
        },
        {
          icon: "mdi-account-multiple-plus-outline",
          title: "การจัดการผู้ใช้งาน",
          to: "/createUser",
          lv: [2],
        },
        {
          icon: "mdi-clipboard-check",
          title: "การจัดการอนุมัติ เบิก/ยืม/คืน",
          to: "/approve",
          lv: [1],
        },
        // {
        //   icon: "mdi-clipboard",
        //   title: "การจัดการบัญชี เบิก/ยืม/คืน ",
        //   to: "/history",
        //   lv: [1,2],
        // },
        {
          icon: "mdi-clipboard-plus",
          title: "เพิ่มแบบฟอร์ม เบิก/ยืม/คืน",
          to: "/documents",
          lv: [0, 1, 2],
        },
      ],
      title: "ระบบบริหารจัดการบัญชีรายรับรายจ่าย ด้วยแพลตฟอร์มออนไลน์ ",
    };
  },
  mounted() {
    if (this.$auth.state.loggedIn) {
      this.name = `${this.$auth.state.user.firstname} ${this.$auth.state.user.lastname}`;
    }
  },
  methods: {
    async logout() {
      await this.$auth.logout();
      this.$router.push("/login");
    },
  },
};
</script>

<style lang="scss" scoped>
.groupManu {
  .manu::before {
    width: 0px;
  }

  .manu a {
    text-decoration: none;
  }

  .activeManu::before {
    width: 5px;
    background: $color_blue;
    border-radius: 0px 5px 5px 0px;
    opacity: 1;
    box-shadow: $shadow_layouts;
    transition: ease-in-out 0.3s;
  }
  .activeManu > i {
    color: $color_blue;
  }
}
</style>