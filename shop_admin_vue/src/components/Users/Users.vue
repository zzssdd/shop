<template>
  <div>
    <el-table :data="UserList" border style="width: 100%">
      <el-table-column label="ID" align="center"  prop="id" width="250"></el-table-column>
      <el-table-column label="注册时间" align="center"  width="700">
        <template slot-scope="scope" align="center" >
          <span>{{ scope.row.create_time | DateForm }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="email" align="center" label="邮箱" width="700">
      </el-table-column>
    </el-table>
      <div class="block">
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 40]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
  </div>
  </div>
  
</template>

<script>
import day from "dayjs";

export default {
  name: "Users",
  data() {
    return {
      UserList: [],
      total:0,
      currentPage:1,
      pageSize:10,
    };
  },
  filters: {
    DateForm: function (date) {
      return date ? day(date).format("YYYY年MM月DD日 HH:mm") : "暂无";
    },
  },
  methods: {
    async getUsersList() {
      const { data: res } = await this.$http.get("user/list",{
        params:{
          currentPage:this.currentPage,
          pageSize:this.pageSize,
        }
      });
      if (res.code !== 200) {
          // window.sessionStorage.clear();
          // this.$router.push("/login");
        this.$message.error(res.$message);
      }
      this.UserList = res.users;
      this.total=res.total
    },
    handleSizeChange(val){
      this.pageSize=val
      this.getUsersList()
    },
    handleCurrentChange(val){
      this.currentPage=val
      this.getUsersList()
    }
  },
  created() {
    this.getUsersList();
  },
};
</script>

<style scoped>
.el-image {
  width: 200px;
  height: 100px;
}
</style>
