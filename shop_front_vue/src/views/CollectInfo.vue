<template>
  <div class="center">
    <img :src="post.pic" class="left" />
    <div class="right">
      <div>
        <h1 class="headline red">商品价格：{{post.price}}</h1>
        <h3>商品名称：{{ post.name }}</h3>
        <small>商家：{{post.unit}}</small><br/>
        <small>
          发布日期: {{ post.createTime|formatDate }}
        </small>
        <h4>
          {{post.desc}}
        </h4>
      </div>
      <hr>
      <div class="edit">
      <vs-button size="xl" warn>立即购买</vs-button>
      </div>
      <div class="edit">
      <vs-button size="xl" danger @click="delCollect">取消收藏</vs-button>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'CollectInfo',
  data: function () {
    return {
      postId: this.$route.path.split('collect/')[1],
      post: {},
    }
  },
  methods: {
    async getPost() {
      // because of markdown file rendering, can't use the triditional dynamic route matching
      this.postId = this.$route.path.split('collect/')[1]
      const { data: res } = await this.$http.get(`product/info`,{
        params:{
          id:this.postId
        }
      } )
      this.post=res.data
      console.log(res);
    },
    async delCollect(){
      const form_data=new FormData()
      form_data.append("pid",this.postId)
        const { data: res } = await this.$http.delete(`collect/del`,{
            params:{
                id:this.postId
            }
        })
        this.$router.push('/collect')
        return alert(res.msg)
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
  mounted: function () {
    this.getPost()
  },
  watch: {
    $route (to, from) {
      this.getPost()
    }
  }
}
</script>

<style scoped>
.left {
  float: left;
  width: 60%;
  height: 100%;
}
.right{
  float: left;
  width: 30%;
  height: 100%;
}
.margin {
  margin-left: 20px;
}

.post img {  /* limit image max width to 100vw in a post */
  max-width: 100%;
}
.red{
  color: red;
}
.edit{
  float: left;
  width: 40%;
}
</style>
