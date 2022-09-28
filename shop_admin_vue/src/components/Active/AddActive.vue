<template>
  <el-card class="box-card">
    <el-form
      ref="activeForm"
      :model="activeInfo"
      :inline="true"
      label-position="top"
      :rules="activeInfoRules"
    >
      <el-row>
        <el-col :span="5">
          <el-form-item label="活动商品：">
            <el-select v-model="activeInfo.pid" placeholder="请选择商品">
              <el-option
                v-for="item in productList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="秒杀价格：">
            <el-input size="big" v-model="activeInfo.secPrice" style="width: 150px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item label="秒杀数量：">
            <el-input size="big" v-model="activeInfo.num" style="width: 100px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="活动开始时间：">
      <el-date-picker
      v-model="activeInfo.startTime"
      type="datetime"
      format="yyyy-MM-dd HH:mm:ss"
      placeholder="选择日期时间"
      value-format="timestamp">
    </el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="活动结束时间：">
         <el-date-picker
      v-model="activeInfo.endTime"
      type="datetime"
       format="yyyy-MM-dd HH:mm:ss"
      placeholder="选择日期时间"
      value-format="timestamp">
    </el-date-picker>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="14">
          <el-form-item label="活动描述：">
            <el-input type="textarea" :rows="8" style="width: 400px" v-model="activeInfo.secDesc"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <div class="btn">
      <el-button type="primary" @click="submitForm">完成</el-button>
      <el-button type="danger" @click="back">返回</el-button>
    </div>
  </el-card>
</template>

<script>
import day from 'dayjs'
export default {
  name: "AddActive",
  props: ["id"],
  data() {
    return {
      productList: [],
      activeInfo: {
        secPrice: undefined,
        num: undefined,
        pid: undefined,
        startTime: "",
        endTime: "",
        secDesc: "",
      },
      activeInfoRules: {
        startTime: [
          { required: true, message: "请输入文章标题", trigger: "change" }
        ],
        endTime: [
          { required: true, message: "请选择文章分类", trigger: "change" }
        ],
        secDesc: [
          { required: true, message: "请输入文章描述", trigger: "change" },
          { max: 120, message: "描述最多可写120个字符", trigger: "change" }
        ]
      }
    };
  },

  methods: {
    async getProductList() {
      const { data: res } = await this.$http.get("product/list", {
        params: {
          pageSize: 1000000,
          currentPage: 1
        }
      });
      if (res.code !== 200) return this.$message.error(res.message);
      this.productList = res.data;
    },
    async getInfo(id){
      const {data:res}=await this.$http.get(`seckill/info`,{
        params:{
          id
        }
      })
      console.log(res);
      if(res.code!=200){
        this.$message.error(res.message)
        return
      }
      this.activeInfo=res.seckill;
    },
    DateForm(date){
        return date?day(date).format('YYYY-MM-DD HH:mm:ss'):'暂无'
      },
    async submitForm() {
      this.$refs.activeForm.validate(async valid => {
        if (valid) {
          const form_data=new FormData();
          form_data.append("secPrice",this.activeInfo.secPrice)
          form_data.append("num",this.activeInfo.num)
          form_data.append("pid",this.activeInfo.pid)
          form_data.append("startTime",this.DateForm(this.activeInfo.startTime))
          form_data.append("endTime",this.DateForm(this.activeInfo.endTime))
          form_data.append("secDesc",this.activeInfo.secDesc)
          if(this.id){
          const { data: res } = await this.$http.put(`seckill/edit`, {
            params:{
              id:this.id
            }
          },form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$message({
            type:'success',
            message:'更新商品成功！'
          })       
          }else{
          const { data: res } = await this.$http.post(
            "seckill/add",
            form_data
          );
          console.log(res);
          if (res.code !== 200) return this.$message.error(res.message);
          }
          this.$router.push("/activeList");
          this.$message({
            type: "success",
            message: "添加活动成功！"
          });
        } else {
          return this.$message.error("输入数据非法，请重新输入");
        }
      });
    },
    back() {
      this.$router.push("/activeList");
    }
  },
  created() {
    this.getProductList();
      if(this.id){
      this.getInfo(this.id)
    }
  }
};
</script>

<style scoped>
</style>
