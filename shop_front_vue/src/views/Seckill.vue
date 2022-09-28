<template>
<div>
 <div class="grid-3_xs-1_sm-2_md-2" v-if="List.length">
    <div :key="index" v-for="(post, index) in List">
    <vs-card class="seckill-card">
      <template #title>
        <h1>{{ post.name }}</h1>
      </template>
      <template #img>
        <img :src="post.pic" alt />
      </template>
      <template #text>
        <h2 class="post-txt text-decoration">商品价格：{{ post.price }}</h2>
        <h1 class="post-txt">秒杀价格：<b class="red">{{post.secPrice}}</b></h1>
        <h3>剩余数量：{{post.num}}</h3>
        <h5 class="post-txt">活动时间：{{ post.startTime }}~{{post.endTime}}</h5>
        <p>{{post.secDesc}}</p>
          <vs-button size="xl" danger @click="getSeckill(post.id,post.pid,post.secPrice)">
            立即抢购
          </vs-button>
      </template>
    </vs-card>
    </div>
 </div>
 <Nothing v-else />
 <div class="center">
 </div>
</div>
</template>

<script>
import Nothing from '@/components/Nothing.vue'
export default {
    name:'Seckill',
    components: {
    Nothing
  },
    data(){
        return{
            List:[],
        }
    },
    methods:{
        async getList(){
            const { data: res } = await this.$http.get('seckill/list' ,{
                params:{
                    currentPage:this.page,
                }
            })
            if (res.code!=200) {
                return alert(msg)
            }
            this.List=res.data
            this.total=(parseInt(res.count)+5)/6
            console.log(res);
        },
        async getSeckill(id,pid,price){
           const { data: res } = await this.$http.post('seckill/seckill' ,{
                params:{
                    id,
                }
            })
            if (res.code!=200) {
             return  alert(res.msg) 
            }
            const {data:res2}=await this.$http.get('seckill/result')
            if (res2.code==200) {
            const form_data=new FormData()
            form_data.append("pid",pid)
            form_data.append("price",price)
            const { data: res3 } = await this.$http.post(`payment/add`,form_data)
            return alert(res3.msg)
            }
            return alert(res2.msg)
        }
    },
      filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		},
    mounted(){
        this.getList()
    },
    watch:{
    page(){
       this.getList()
    }
  }
}
</script>

<style>
 .text-decoration{
  text-decoration:line-through
 }
 .red{
  color: red;
 }
.seckill-card .vs-card{
  height: 700px !important;
  max-height: 1000px !important;
  width: 600px !important;
  max-width: 1000px !important;
 }
</style>