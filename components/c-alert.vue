<template>
  <v-snackbar :v-model="alert.snackbar"  :color="alert.color" timeout="1500">
    {{ alert.msg }}
  </v-snackbar>
</template>

<script>
export default {
    
  props: {
    'alert':{
    type: Object,
    default: () => ({}),
  }
  },
};
</script>

<style>
</style>