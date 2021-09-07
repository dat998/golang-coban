<template>
   <div class="calculator">
        <div class="tieu-de">
            <input id="result" v-model="result" >
        </div>
        <div class="number">
            <table class="tab-number">
                <tr>
                    <button name="sign" class="math" value="+" @click="clickNumber('+')">+</button>
                    <button name="sign" class="math" value="-" @click="clickNumber('-')">-</button>
                    <button name="sign" class="math" value="*" @click="clickNumber('*')">*</button>
                    <button name="sign" class="math" value="/" @click="clickNumber('/')">/</button>
                </tr>
                <tr>
                    <button class="value" value="1" @click="clickNumber(1)">1</button>
                    <button class="value" value="2" @click="clickNumber(2)">2</button>
                    <button class="value" value="3" @click="clickNumber(3)">3</button>
                    <button class="value" value="4" @click="clickNumber(4)">4</button>
                </tr>
                <tr>
                    <button class="value" value="5" @click="clickNumber(5)">5</button>
                    <button class="value" value="6" @click="clickNumber(6)">6</button>
                    <button class="value" value="7" @click="clickNumber(7)">7</button>
                    <button class="value" value="8" @click="clickNumber(8)">8</button>
                </tr>
                <tr>
                    <button class="value" value="9" @click="clickNumber(9)">9</button>
                    <button class="value" value="0" @click="clickNumber(0)">0</button>
                    <button @click="ac()">ac</button>
                    <button class="value" value="=" @click="clickNumber('=')">=</button>
                </tr>
            </table>
        </div>
    </div>
</template>

<script>
export default {
  name: 'Calculator',
  props: {
    msg: String
  },
  data(){
    return{
      result1:"",
      result2:"",
      numberSubOne:"",
      result:"",
      sign:""
    }
  },
  methods:{
    clickNumber(x){
      if(x=='='){
        //alert("equal")
        this.result=eval(this.result1+this.sign+this.result2)
        this.result1=""
        this.result2=""
        this.sign=""
      }else{
        //xu ly kiem tra ky tu dau tien va co phai dau hay khong
        //xu ly xem co 2 dau duoc nhap ke tiep nhau hay k
        if ((this.result=="" ||(this.numberSubOne=="+"||this.numberSubOne=="-"||this.numberSubOne=="*"||this.numberSubOne=="/"))&&(x=="+"||x=="-"||x=="*"||x=="/")){
          console.log("Khong duoc nhap")
        }else{
          
          
          if(this.sign!=""){
            this.result2+=x
            this.result = this.result2
          }
          else{
            if(x=="+"||x=="-"||x=="*"||x=="/"){
              this.sign=x
            }else{
              this.result1+=x
              this.result = this.result1
              this.numberSubOne = x
            }
            
          }
        
        // console.log(x)
        console.log(this.result1)
        console.log(this.result2)
        console.log(this.sign)
      }
     
      }
    },
    ac(){
      this.result = ""
      this.result1=""
      this.result2=""
      this.sign=""
    }
  }
}
</script>

<style>
  .calculator{
      margin: 0;
      padding: 0;
      height: 600px;
      width: 300px;
      background-color: rgb(94, 105, 105);
  }
  .tieu-de{
      display: block;
      height: 20%;
      background-color: rgb(138, 178, 180);
      width: 100%;
  }
  #result{
    font-size: 30px;
    padding: 0;
    width: 90%;
    height:80%;
    margin: 5%;
    border:2px solid rgb(46, 51, 46) ;
    border-radius:5% ;
  }
  .number{
      margin: 0;
      padding: 0;
      width: 100%;
      height: 80%;
  }
  .tab-number{
      margin: 0;
      padding: 0;
      height: 100%;
      width: 100%;
  }
  .tab-number tr{
      width: 100%;
      height: 25%;
  }
  tr, button{
      height:100%;
      width: 25%;
      background-color: cadetblue;
      border-radius: 10%;
      text-align: center;
  }
</style>