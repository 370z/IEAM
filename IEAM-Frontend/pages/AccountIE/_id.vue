<template>
  <div class="content-size">
    <layout-title-page
      title="แก้ไขรายการ"
      details="เพิ่มรายการบัญชีรายรับ-รายจ่าย"
    />
    <c-form mode="e"/>
  </div>
</template>

<script>
export default {
  mounted() {
    this.lv =  [3];
    if(!this.lv.includes(this.$auth.state.user.level)){
      this.$router.push("/dashboard");
    }
  },
};
</script>

<style>
</style>