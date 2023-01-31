<template>
      <div class="page-loader" v-if="isloader">
      <!-- <div class="logo">
        <img src="~/assets/images/login_logo.png" alt="">
      </div> -->
      <div class="boxloader">
        <div class="cube"></div>
        <div class="cube"></div>
        <div class="cube"></div>
        <div class="cube"></div>
        <div class="cube"></div>
      </div>
    </div>
</template>

<script>
export default {
    props:{
        'isloader':{
            type: Boolean,
            default: false
        }
    }
    
}
</script>

<style lang="scss" scoped>
$colors: #EB3C41, #3951ef, #ffffff,#EB3C41, #3951ef, #ffffff ;
   .page-loader {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: #1f1f1fcb;
    z-index: 99;
  }
  // -----------------------------------------------------
  .logo{
    margin-bottom: 15px;
    
    img{
      animation: img 7.5s infinite;
      width: 100px;
      height: 100px;
    }
  }
  .boxloader{
    width: 10%;
    display: flex;
    justify-content: space-between;
    .cube{
      width: 15px;
      height: 15px;
      border-radius: 100%;
  
       @for $i from 1 through length($colors) {
        &:nth-child(#{$i}) {
          background-color: nth($colors, $i);
        }
      }
   
      &:first-child {
        animation: left 1.5s infinite;
      }
      &:nth-last-child(4) {
        animation: left 1.5s infinite;
        animation-delay: 0.3s;
      }
      &:nth-last-child(3) {
        animation: left 1.5s infinite;
        animation-delay: 0.5s;
      }
      &:nth-last-child(2) {
        animation: left 1.5s infinite;
        animation-delay: 0.7s;
      }
      &:last-child {
        animation: left 1.5s infinite;
        animation-delay: 0.9s;
      }
    }
  }
  // -----------------------------------------------------
  @keyframes left {
    40% {
      transform: translateX(-10px);
    }
    50% {
      transform: translateX(0);     
    }
    75%{
      transform: translateY(-5px);
    }
    100%{
      transform: translateY(-0px);
    }
  }
  @keyframes img {
    0% {
      
    }
    50% {
      opacity: 0;
    }
  }
</style>