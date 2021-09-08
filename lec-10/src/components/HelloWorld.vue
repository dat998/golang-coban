<template>
  <div class="main">
    <div class="element">
      <div class="header">
        <img src="../assets/background.svg" />
      </div>
      <div class="body">
        <div class="text">
          <h3>Order Summary</h3>
          <p>You can now listen to milions of songs, audiobooks, and podcasts on any device anywhere you like!</p>
        </div>
        <div class="price">
          <img class="img-music" src="../assets/music.png" alt="">
          <div v-show="isShow" class="price-item">
            <h4>Annual Plan</h4>
            <p>$59.99/year</p>
          </div>
          <div v-show="!isShow" class="price-item">
            <h4>Pro Plan</h4>
            <p>$99.99/year</p>
          </div>
          <div @click="changePlan" value="Change" type="button"><u>Change</u></div>
        </div>
        <div class="footer">
          <button class="btn-pay" @click="showData=true">Proceed to Payment</button>
          <div class="cancel">Cancel Order</div>
        </div>
        
      </div>
    </div>
    
    <modal v-if="showData" @close="showData = false">
      <p>success</p>
      <button @click="showData=false">OK</button>
    </modal>
  </div>
  
</template>

<script>
export default {
  name: 'HelloWorld',
  data(){
    return{
      isShow: true,
      showData: false,
    }
  },
  methods:{
     changePlan() {
      this.isShow = !this.isShow;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.main{
  margin: 0px;
  padding: 0px 10px 0px 10px ;
  height: auto;
  width: 100%;
}
.element{
  height: auto;
  width: 100%;
  max-width: 500px;
  margin: 0px;
  padding: 0px;
  box-shadow: -10px 10px 10px rgb(187, 185, 185);
}  
.header{
  display: flex;
  width: auto;
  height: 30%;
  /* background-image:url("../assets/background.svg") ; */
}
 img{
  display: table;
  width: 100%;
  height: 100%;
  border-radius: 10px 10px 0px 0px;
}
.body{
  width: auto;
  height: 70%;
}
h3{
  text-align: center;
  width: auto;
  padding: 20px;
  margin: 0;
}
p{
  margin: 0;
  text-align: center;
}
.text{
  height: auto;
  width: auto;
}
.price {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgb(205, 211, 224);
  border-radius: 10px;
  width: 94%;
  margin: 3%;
}
.price-item {
  margin-right: 40px;
}
h4 {
  margin-bottom: 0;
}
.btn-pay {
  width: 90%;
  height: 45px;
  border-radius: 10px;
  background-color: rgb(80, 112, 255);
  margin: 10px 0px 10px 0px;
  border: none;
}
.cancel {
  margin: 1rem 0 1rem 0;
}
.img-music{
  width: 100px;
  height: 100px;
  border-radius:50% ;
  margin-right:20px ;
}
.footer{
  padding: 2rem;
}
</style>
