<template>
  <el-card class="box-card">
    <el-form
      ref="artForm"
      :model="productInfo"
      :inline="true"
      label-position="top"
      :rules="productRules"
    >
      <el-row>
        <el-col :span="5">
          <el-form-item label="商品名称：">
            <el-input
              size="big"
              v-model="productInfo.name"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
          <el-col :span="5">
          <el-form-item label="商品价格：">
            <el-input
              size="big"
              v-model="productInfo.price"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
          <el-col :span="5">
          <el-form-item label="商品数量：">
            <el-input
              size="big"
              v-model="productInfo.num"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
                  <el-col :span="5">
          <el-form-item label="商家名称：">
            <el-input
              size="big"
              v-model="productInfo.unit"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="14">
          <el-form-item label="商品描述：">
            <el-input
              type="textarea"
              :rows="8"
              style="width: 400px"
              v-model="productInfo.desc"
            >
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="文章封面"  >
           <input type="file"/>
          </el-form-item>
          <el-image :src="productInfo.pic" v-show="productInfo.pic"></el-image>
        </el-col>
      </el-row>
    </el-form>
     <div class="btn">
        <el-button type="primary" @click="submitForm"
          >完成</el-button
        >
        <el-button type="danger" @click="back"
          >返回</el-button
        >
        </div>
  </el-card>
</template>

<script>
export default {
  name: "AddProduct",
  props:['id'],
  data() {
    return {
      productInfo: {
        id: 0,
        name: '',
        desc: '',
        price:undefined,
        unit:'',
        num:undefined,
        pic: '',
      },
       productRules: {
        name: [{ required: true, message: '请输入商品名称', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入商品描述', trigger: 'change' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'change' },
        ],
        unit: [{ required: true, message: '请输入文章内容', trigger: 'change' }],
      },
    };
  },
  methods: {
    async getInfo(id){
      const {data:res}=await this.$http.get(`product/info`,{
        params:{
          id
        }
      })
      console.log(res);
      if(res.code!=200){
        this.$message.error(res.message)
        return
      }
      this.productInfo=res.data;
    },
    async submitForm(){
          this.$refs.artForm.validate(async (valid) => {
          if(valid){
          const form_data=new FormData();
          form_data.append("id",this.productInfo.id)
          form_data.append("name",this.productInfo.name)
          form_data.append("desc",this.productInfo.desc)
          form_data.append("price",this.productInfo.price)
          form_data.append("unit",this.productInfo.unit)
          form_data.append("num",this.productInfo.num)
        form_data.append("pic",document.querySelector('input[type=file]').files[0])
          if(this.id){
          const { data: res } = await this.$http.put(`product/edit`, form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$router.push('/productList')
          this.$message({
            type:'success',
            message:'更新商品成功！'
          })
        }else{
          const { data: res } = await this.$http.post('product/add',form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$router.push('/productList')
          this.$message({
            type:'success',
            message:'添加文章成功！'
          })
        };
        }else{
          return this.$message.error('输入数据非法，请重新输入')
        }})
    },
    back(){
      this.$router.push('/productList')
    }
  },
  created(){
    if(this.id){
      this.getInfo(this.id)
    }
  }
};
</script>

<style scoped>
</style>
