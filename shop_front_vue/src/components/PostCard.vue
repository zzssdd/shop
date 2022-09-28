<template>
  <router-link :to="`/product/${post.id}`">
    <vs-card class="post-card">
      <template #title>
        <h1>{{ post.name }}</h1>
      </template>
      <template #img>
        <img :src="post.pic" alt />
      </template>
      <template #text>
        <p class="post-txt">商品价格：{{ post.price }}</p>
        <p class="post-txt">商家：{{ post.unit }}</p>
      </template>
      <template #interactions>
      </template>
    </vs-card>
  </router-link>
</template>

<script>

export default {
  name: 'PostCard',
  props: [
    'post',
  ],
  filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		}
}
</script>

<style>
.post-txt {
  overflow: hidden;
  text-overflow: ellipsis;
  display: auto;
}

.post-card .vs-card {
  height: 400px !important;
  width: 600px !important;
  max-width: 1000px !important;
}

.post-card-tag {
  position: absolute;
  bottom: 20px;
}
</style>
