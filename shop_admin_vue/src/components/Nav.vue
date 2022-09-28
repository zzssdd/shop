<template>
    <el-row class="tac">
    <el-menu
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router
      >
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-s-fold" ref="Icon"></i>
          <span>活动管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="1-1" route="addActive">
         <i class="el-icon-edit"></i>
          <span>增加活动</span>
            </el-menu-item>
          <el-menu-item index="1-2" route="activeList">
            <i class="el-icon-notebook-1"></i>
            <span>活动列表</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>

       <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-s-fold" ref="Icon"></i>
          <span>商品管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="2-1" route="addProduct">
            <i class="el-icon-edit-outline"></i>
            <span>增加商品</span>
          </el-menu-item>
            <el-menu-item index="2-2" route="productList">
            <i class="el-icon-notebook-2"></i>
            <span>商品列表</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>

      <el-menu-item index="3" route="users">
        <i class="el-icon-chat-dot-square"></i>
        <span slot="title">用户管理</span>
      </el-menu-item>


    </el-menu>
</el-row>
</template>

<script>
export default {
    name:'Nav',
    methods:{
    handleOpen(){
        this.$refs.Icon.className='el-icon-s-unfold'
     },
     handleClose(){
         this.$refs.Icon.className='el-icon-s-fold'
     }
    }
}
</script>

<style>

</style>