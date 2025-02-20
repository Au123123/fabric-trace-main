<template>
  <div>
    <h1 style="color: #1f2f3d; text-align: center;">5分钟构建任意溯源系统</h1>
    <p style="color: #5e6d82; text-align: center;">请对比字段填写，生成个性化的溯源系统。激活码仅可使用一次，</p>
    <p style="color: #5e6d82; text-align: center;">提交前请认真对比，生成后请尽快下载并备份，源码在服务器保留一周后删除。</p>
    <div style="text-align: center; margin-bottom: 20px;">
      <a
        href="https://www.bilibili.com/video/BV1Ar421H7TK"
        style="color: #409EFF; text-decoration: underline;"
      >
        B站：使用教程
      </a>
    </div>
    <el-button
      type="text"
      @click="dialog2Visible = true"
    >购买激活码</el-button>
    <el-dialog
      title="删除此页面"
      :visible.sync="dialog1Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>1. 在fabric-trace/application/web目录下，运行./rmbuildsyspage.sh</span>
      <br>
      <span style="display: block;margin-top: 20px;">2.重新启动前端，npm run dev</span>
      <span
        slot="footer"
        class="dialog-footer"
      >
        <el-button
          type="primary"
          @click="dialog1Visible = false"
        >确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="开发不易，感谢支持！"
      :visible.sync="dialog2Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>激活码售价：399元</span>

      <br>
      <span style="display: block; margin-top: 20px;">请加QQ群776873343联系群主购买</span>
      <br>
      <span style="display: block; margin-top: 20px;">感谢曾付费支持本项目的伙伴！如果您在2024年12月9日前购买过此项目的课程/专栏，</span>
      <br>
      <span>在2024.12月前可75折购买激活码一个</span>
      <span
        slot="footer"
        class="dialog-footer"
      >
        <el-button
          type="primary"
          @click="dialog2Visible = false"
        >确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="构建成功！"
      :visible.sync="dialog3Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>下载地址：</span>
      <br>
      <span style="display: block;margin-top: 20px;">{{ downloadUrl }} </span>
      <span
        slot="footer"
        class="dialog-footer"
      >
        <el-button
          type="primary"
          @click="dialog3Visible = false"
        >确 定</el-button>
      </span>
    </el-dialog>

    <div class="form-container">
      <el-form
        ref="form"
        :model="form"
        label-width="150px"
        class="form"
      >

        <el-form-item
          v-for="(value, key) in form"
          :key="key"
          :label="key"
        >
          <el-input
            v-model="form[key]"
            :placeholder="'请输入 ' + key + ' 值'"
          ></el-input>
        </el-form-item>
        <div style="text-align: center; margin-top: 20px;">
          <el-button
            type="primary"
            @click="submitForm"
            v-loading="loading"
            element-loading-text="构建中，稍等1分钟"
          >开始构建</el-button>
        </div>
      </el-form>
      <el-button
        type="text"
        @click="dialog1Visible = true"
        style="display: block;margin-top: 50px;"
      >如何删除此页面？</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Build",
  data() {
    return {
      loading: false,
      dialog1Visible: false,
      dialog2Visible: false,
      form: {
        parm1: "基于区块链的新闻图片溯源系统",
        parm2: "新闻图片信息",
        parm3: "新闻事件主题",
        parm4: "拍摄地点",
        parm5: "拍摄时间",
        parm6: "首次发布时间",
        parm7: "摄影记者姓名",
        parm8: "摄影记者",
        parm9: "媒体机构信息",
        parm10: "报道标题",
        parm11: "发布批次（可根据发布平台、系列报道等设定）",
        parm12: "编辑处理时间",
        parm13: "媒体机构名称与地址",
        parm14: "媒体机构联系电话",
        parm15: "媒体机构",
        parm16: "传播轨迹信息",
        parm17: "转发者姓名（若有实名情况）",
        parm18:
          "无关信息（年龄在新闻图片溯源中通常无关联，可删除此项或标记为无需收集）",
        parm19: "转发者联系电话（若可获取）",
        parm20:
          "无关信息（车牌号在新闻图片溯源中无关联，可删除此项或标记为无需收集）",
        parm21: "转发记录",
        parm22: "转发者（可包含个人、其他媒体等）",
        parm23: "发布平台信息",
        parm24: "图片上传至平台时间",
        parm25: "在平台上的热门传播时间（若有统计意义）",
        parm26: "发布平台名称",
        parm27: "发布平台服务器位置（若相关）",
        parm28: "发布平台联系电话",
        parm29: "发布平台",
        parm30: "查看与下载图片的用户（类似消费者角色 ）",
        parm31: "云服务器IP,例如：32.12.243.30/192.168.1.20",
        activatecode: "激活码",
      },
      downloadUrl: "",
      dialog3Visible: false,
    };
  },
  methods: {
    submitForm() {
      this.loading = true;
      const params = new URLSearchParams();
      for (const key in this.form) {
        params.append(key, this.form[key]);
      }
      fetch("http://realcool.top:8088/activate", {
        // fetch('http://127.0.0.1:8088/activate', {

        method: "POST",
        body: params,
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      })
        .then((response) => response.json())
        .then((data) => {
          console.log("响应数据：", data);
          this.loading = false;
          this.downloadUrl = data.msg;
          if (data.code == 0) {
            // this.$message.success('构建成功！' + data.msg);
            this.dialog3Visible = true;
          } else {
            this.$message.error("构建失败：" + data.msg);
            console.log("构建失败：", data.msg);
          }
        })
        .catch((error) => {
          this.$message.error("提交失败，请重试！");
          console.error("提交失败：" + error);
          this.loading = false;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.form-container {
  max-width: 700px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border: 1px solid #ebebeb;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.el-form-item {
  margin-bottom: 15px;
}

.el-button {
  width: 100%;
}
</style>
