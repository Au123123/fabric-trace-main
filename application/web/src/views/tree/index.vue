<!--
 * @Author: Au123123 2502894164@qq.com
 * @Date: 2024-12-25 17:35:30
 * @LastEditors: Au123123 2502894164@qq.com
 * @LastEditTime: 2025-02-20 10:21:05
 * @FilePath: \fabric-trace-main\application\web\src\views\tree\index.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="app-container">
    <el-input
      v-model="filterText"
      placeholder="输入关键词进行过滤"
      style="margin-bottom:30px;"
    />

    <el-tree
      ref="tree2"
      :data="newsImageTraceData"
      :props="defaultProps"
      :filter-node-method="filterNode"
      class="filter-tree"
      default-expand-all
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      filterText: "",
      // 基于区块链的新闻图片溯源数据
      newsImageTraceData: [
        {
          id: 1,
          label: "新闻图片：重大活动现场图",
          children: [
            {
              id: 2,
              label: "摄影记者：张三",
              children: [
                {
                  id: 3,
                  label: "拍摄信息",
                  children: [
                    {
                      id: 4,
                      label: "拍摄时间：2025-02-20 10:00",
                    },
                    {
                      id: 5,
                      label: "拍摄地点：活动现场 A 区",
                    },
                    {
                      id: 6,
                      label: "区块链交易 ID：TX001",
                    },
                    {
                      id: 7,
                      label: "区块链记录时间：2025-02-20 10:05",
                    },
                  ],
                },
                {
                  id: 8,
                  label: "媒体机构处理",
                  children: [
                    {
                      id: 9,
                      label: "媒体机构：XX 报社",
                    },
                    {
                      id: 10,
                      label: "编辑处理时间：2025-02-20 11:00",
                    },
                    {
                      id: 11,
                      label: "报道标题：重大活动精彩瞬间",
                    },
                    {
                      id: 12,
                      label: "区块链交易 ID：TX002",
                    },
                    {
                      id: 13,
                      label: "区块链记录时间：2025-02-20 11:05",
                    },
                  ],
                },
                {
                  id: 14,
                  label: "发布平台传播",
                  children: [
                    {
                      id: 15,
                      label: "发布平台：新闻网站 X",
                    },
                    {
                      id: 16,
                      label: "上传时间：2025-02-20 12:00",
                    },
                    {
                      id: 17,
                      label: "热门传播时间：2025-02-20 14:00",
                    },
                    {
                      id: 18,
                      label: "区块链交易 ID：TX003",
                    },
                    {
                      id: 19,
                      label: "区块链记录时间：2025-02-20 12:05",
                    },
                  ],
                },
              ],
            },
          ],
        },
      ],
      defaultProps: {
        children: "children",
        label: "label",
      },
    };
  },
  watch: {
    filterText(val) {
      this.$refs.tree2.filter(val);
    },
  },
  methods: {
    filterNode(value, data) {
      if (!value) return true;
      return data.label.indexOf(value) !== -1;
    },
  },
};
</script>