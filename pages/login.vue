<template>
  <div class="background">
    <div class="login__container">
      <div class="wrapper__login">
        <!-- <img src="~/assets/images/login_logofull.png" width="60%" /> -->
        <div class="my-5">
          <h1 style="color: #192786">LOGIN</h1>
          <p>ระบบบริหารจัดการบัญชีรายรับรายจ่าย</p>
          <p>ด้วยแพลตฟอร์มออนไลน์</p>
        </div>

        <form class="login__form" @submit.prevent="login">
          <v-row justify="center" align="center">
            <v-col justify="center" align-self="center" cols="12">
              <v-text-field
                v-model="formlogin.username"
                placeholder="ชื่อผู้ใช้"
                prepend-inner-icon="mdi-account"
                hide-details
                solo
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="formlogin.password"
                placeholder="รห้สผ่าน"
                prepend-inner-icon="mdi-key"
                :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                :type="show ? 'text' : 'password'"
                @click:append="show = !show"
                hide-details
                solo
              ></v-text-field>
            </v-col>
            <v-col cols="12" align="right" class="mr-4">
              <nuxt-link to="/">ลืมรห้สผ่าน</nuxt-link>
            </v-col>
            <v-col cols="12">
              <c-login-button class="login__button" type="submit"
                >เข้าสู่ระบบ</c-login-button
              >
            </v-col>
          </v-row>
        </form>
      </div>
    </div>
    <v-alert v-if="alert.snackbar" border="left" dense :type="alert.type">{{
      alert.msg
    }}</v-alert>
    <!-- <c-alert :alert="alert" /> -->
    <c-loader :isloader="loader" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      alert: {
        snackbar: false,
        msg: "",
        type: "",
      },
      loader: false,
      show: false,
      formlogin: {
        username: "",
        password: "",
      },
    };
  },
  mounted() {
    if (this.$auth.loggedIn) {
      this.$router.push("/dashboard");
    }
  },
  methods: {
    async login() {
      this.loader = true;
      await this.$auth
        .loginWith("local", { data: this.formlogin })
        .then((response) => {
          this.loader = false;
          alert(response.data.msg);

          this.$router.push({ path: "/dashboard" });
        })
        .catch((err) => {
          console.log(err);
          alert(err.response.data?.msg);
          this.loader = false;
        });
    },
  },
};
</script>

<style lang="scss" >
.background {
  width: 100%;
  height: 100%;
  // background-image: url("https://images.unsplash.com/photo-1595113316349-9fa4eb24f884?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1772&q=100");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  // filter: blur(8px);
  // -webkit-filter: blur(8px);
}

.login__container {
  min-width: 300px;
  max-width: 400px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
  padding: 15px;
  border-radius: 25px;

  p {
    margin: 0;
  }
}
</style>