<template>
<div>
<div class="center" v-if="List.length>0">
      <vs-table
        v-model="selected"
        >
        <template #thead>
          <vs-tr>
            <vs-th>
              <vs-checkbox
                :indeterminate="selected.length == List.length"
                @change="selected = $vs.checkAll(selected, List)"
              />
            </vs-th>
            <vs-th>
              
            </vs-th>
            <vs-th>
              商品名称
            </vs-th>
            <vs-th>
              商品介绍
            </vs-th>
            <vs-th>
              添加时间
            </vs-th>
            <vs-th>
              价格
            </vs-th>
          </vs-tr>
        </template>
        <template #tbody>
          <vs-tr
            :key="i"
            v-for="(tr, i) in List"
            :data="tr"
            :is-selected="!!selected.includes(tr)"
          >
            <vs-td checkbox>
              <vs-checkbox :val="tr" v-model="selected" />
            </vs-td>
            <vs-td class="td-head" @click="getDetail(tr.id)">
               <img :src="tr.pic" alt="" class="pic">
            </vs-td>
            <vs-td>
            {{ tr.name }}
            </vs-td>
            <vs-td>
            {{ tr.desc }}
            </vs-td>
            <vs-td>
            {{ tr.addTime|formatDate }}
            </vs-td>
              <vs-td>
            {{ tr.price }}
            </vs-td>
          </vs-tr>
        </template>
      </vs-table>
      <span class="data">
        <b>
       总价格为:{{selected|ComputePrice }}
        </b>
      </span>
      <vs-button size="xl" class="btn-margin" @click="buy()">
        立即支付
      </vs-button>
    </div>
      <Nothing v-else />
      <div class="center">
        <vs-pagination v-model="page" :length="total" />
       </div>
    </div>
</template>

<script>
import Nothing from '@/components/Nothing.vue'
export default {
    name:'Collect',
    components: {
    Nothing
    },
    data(){
        return{
            List:[],
            page:1,
            total:0,
            selected:[],
            totalPrice:0,
        }
    },
    components: {
    Nothing
  },
    methods:{
        async getList(){
            const { data: res } = await this.$http.get('collect/list' ,{
                params:{
                    currentPage:this.page,
                }
            })
            if (res.code!=200) {
                return alert(res.msg)
            }
            this.List=res.data
            this.total=(parseInt(res.count)+5)/6
            console.log(res);
        },
        getDetail(id){
          this.$router.push(`/collect/${id}`)
        },
        async buy(){
          for (let i = 0; i < this.selected.length; i++) {
            const form_data=new FormData()
            form_data.append("pid",this.selected[i].pid)
            form_data.append("price",this.selected[i].price)
            const { data: res } = await this.$http.post(`payment/add`,form_data)
            if (res.code!=200) {
              return alert(res.msg)
            }
            alert("下单成功！")
          }
        }
    },
      filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			},
      ComputePrice:function(value){
        let sum=0
        for (let i = 0; i < value.length; i++) {
          sum+= value[i].price;
        }
        return sum
      }
		},
    mounted(){
        this.getList()
    },
    watch:{
    page(){
       this.getList()
    },
  }
}
</script>

<style>
  .data{
    display: block;
  }
  .pic{
    width:400px !important;
    height:200px !important;
  }
  .td-head{
    width: 450px;
  }
  .btn-margin{
    margin-left:40px !important ;
  }
</style>