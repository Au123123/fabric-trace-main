<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column
        align="center"
        label="ID"
        width="95"
      >
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="新闻事件主题">
        <template slot-scope="scope">
          {{ scope.row.newsEventTheme }}
        </template>
      </el-table-column>
      <el-table-column
        label="摄影记者"
        width="110"
        align="center"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.photographer }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="拍摄时间"
        width="180"
        align="center"
      >
        <template slot-scope="scope">
          {{ scope.row.shootingTime }}
        </template>
      </el-table-column>
      <el-table-column
        label="媒体机构"
        width="180"
        align="center"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.mediaOrganization }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="原始发布时间"
        width="180"
        align="center"
      >
        <template slot-scope="scope">
          {{ scope.row.originalReleaseTime }}
        </template>
      </el-table-column>
      <el-table-column
        label="传播路径"
        width="300"
        align="center"
      >
        <template slot-scope="scope">
          <span
            v-for="(step, index) in scope.row.propagationPath"
            :key="index"
          >
            {{ step }}
            <span v-if="index < scope.row.propagationPath.length - 1"> -> </span>
          </span>
        </template>
      </el-table-column>
      <el-table-column
        label="发布平台"
        width="180"
        align="center"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.publishingPlatform }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="当前状态"
        width="110"
        align="center"
      >
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
// 假设这里是新的获取新闻图片溯源数据的 API 函数
import { getNewsImageTraceList } from "@/api/newsImageTrace";

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        valid: "success",
        suspected: "warning",
        invalid: "danger",
      };
      return statusMap[status];
    },
  },
  data() {
    return {
      list: null,
      listLoading: true,
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      getNewsImageTraceList().then((response) => {
        this.list = response.data.items;
        this.listLoading = false;
      });
    },
  },
};
</script>