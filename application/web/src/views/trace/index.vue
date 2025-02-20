<template>
  <div class="trace-container">
    <el-input
      v-model="input"
      placeholder="请输入溯源码查询"
      style="width: 300px;margin-right: 15px;"
    />
    <el-button
      type="primary"
      plain
      @click="NewsImageInfo"
    > 查询 </el-button>
    <el-button
      type="success"
      plain
      @click="AllNewsImageInfo"
    > 获取所有新闻图片信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form
            label-position="left"
            inline
            class="demo-table-expand"
          >
            <div><span
                class="trace-text"
                style="color: #67C23A;"
              >新闻图片基础信息</span></div>
            <el-form-item label="新闻事件主题：">
              <span>{{ props.row.reporter_input.re_newsTheme }}</span>
            </el-form-item>
            <el-form-item label="拍摄地点：">
              <span>{{ props.row.reporter_input.re_shootingLocation }}</span>
            </el-form-item>
            <el-form-item label="拍摄时间：">
              <span>{{ props.row.reporter_input.re_shootingTime }}</span>
            </el-form-item>
            <el-form-item label="摄影记者姓名：">
              <span>{{ props.row.reporter_input.re_photographerName }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.reporter_input.re_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.reporter_input.re_timestamp }}</span>
            </el-form-item>
            <div><span
                class="trace-text"
                style="color: #409EFF;"
              >媒体机构信息</span></div>
            <el-form-item label="报道标题：">
              <span>{{ props.row.media_input.me_reportTitle }}</span>
            </el-form-item>
            <el-form-item label="发布批次：">
              <span>{{ props.row.media_input.me_releaseBatch }}</span>
            </el-form-item>
            <el-form-item label="编辑处理时间：">
              <span>{{ props.row.media_input.me_editingTime }}</span>
            </el-form-item>
            <el-form-item label="媒体机构名称与地址：">
              <span>{{ props.row.media_input.me_mediaName }}</span>
            </el-form-item>
            <el-form-item label="媒体机构电话：">
              <span>{{ props.row.media_input.me_contactNumber }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.media_input.me_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.media_input.me_timestamp }}</span>
            </el-form-item>
            <div><span
                class="trace-text"
                style="color: #E6A23C;"
              >传播轨迹信息</span></div>
            <el-form-item label="转发者姓名：">
              <span>{{ props.row.spreader_input.sp_name }}</span>
            </el-form-item>
            <el-form-item label="转发者联系电话：">
              <span>{{ props.row.spreader_input.sp_phone }}</span>
            </el-form-item>
            <el-form-item label="转发记录：">
              <span>{{ props.row.spreader_input.sp_spreadRecord }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.spreader_input.sp_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.spreader_input.sp_timestamp }}</span>
            </el-form-item>
            <div><span
                class="trace-text"
                style="color: #909399;"
              >发布平台信息</span></div>
            <el-form-item label="图片上传时间：">
              <span>{{ props.row.platform_input.pl_uploadTime }}</span>
            </el-form-item>
            <el-form-item label="热门传播时间：">
              <span>{{ props.row.platform_input.pl_popularTime }}</span>
            </el-form-item>
            <el-form-item label="发布平台名称：">
              <span>{{ props.row.platform_input.pl_platformName }}</span>
            </el-form-item>
            <el-form-item label="发布平台位置：">
              <span>{{ props.row.platform_input.pl_platformLocation }}</span>
            </el-form-item>
            <el-form-item label="发布平台电话：">
              <span>{{ props.row.platform_input.pl_platformPhone }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.platform_input.pl_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.platform_input.pl_timestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="溯源码"
        prop="traceability_code"
      />
      <el-table-column
        label="新闻事件主题"
        prop="reporter_input.re_newsTheme"
      />
      <el-table-column
        label="拍摄时间"
        prop="reporter_input.re_shootingTime"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
// 假设新的 API 函数
import {
  getNewsImageInfo,
  getNewsImageList,
  getAllNewsImageInfo,
  getNewsImageHistory,
} from "@/api/newsImageTrace";

export default {
  name: "Trace",
  data() {
    return {
      tracedata: [],
      loading: false,
      input: "",
    };
  },
  computed: {
    ...mapGetters(["name", "userType"]),
  },
  created() {
    getNewsImageList().then((res) => {
      this.tracedata = JSON.parse(res.data);
    });
  },
  methods: {
    AllNewsImageInfo() {
      getAllNewsImageInfo().then((res) => {
        this.tracedata = JSON.parse(res.data);
      });
    },
    NewsImageHistory() {
      getNewsImageHistory().then((res) => {
        // console.log(res)
      });
    },
    NewsImageInfo() {
      var formData = new FormData();
      formData.append("traceability_code", this.input);
      getNewsImageInfo(formData).then((res) => {
        if (res.code === 200) {
          // console.log(res)
          this.tracedata = [];
          this.tracedata[0] = JSON.parse(res.data);
          return;
        } else {
          this.$message.error(res.message);
        }
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>