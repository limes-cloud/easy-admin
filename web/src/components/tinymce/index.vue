<script lang="ts" setup>
  import { computed, onMounted, onUnmounted } from 'vue';

  import tinymce from 'tinymce/tinymce'; // tinymce核心文件
  import Editor from '@tinymce/tinymce-vue';

  //   import 'tinymce/models/dom'; // 引入dom模块。从 tinymce6，开始必须有此模块导入
  import 'tinymce/themes/silver'; // 默认主题
  import 'tinymce/icons/default'; // 引入编辑器图标icon，不引入则不显示对应图标
  import './zh-Hans'; // 引入编辑器语言包

  /* 引入编辑器插件
   * 位于 ./node_modules/tinymce/plugins 目录下，版本不同，插件会有所差异。根据自己版本来导入，若不存在的，不能导入，会报错。
   */
  import 'tinymce/plugins/advlist'; // 高级列表
  import 'tinymce/plugins/anchor'; // 锚点
  import 'tinymce/plugins/autolink'; // 自动链接
  import 'tinymce/plugins/autoresize'; // 编辑器高度自适应,注：plugins里引入此插件时，Init里设置的height将失效
  import 'tinymce/plugins/autosave'; // 自动存稿
  import 'tinymce/plugins/charmap'; // 特殊字符
  import 'tinymce/plugins/code'; // 编辑源码
  import 'tinymce/plugins/codesample'; // 代码示例
  import 'tinymce/plugins/directionality'; // 文字方向
  import 'tinymce/plugins/emoticons'; // 表情
  import 'tinymce/plugins/fullscreen'; // 全屏
  import 'tinymce/plugins/help'; // 帮助
  import 'tinymce/plugins/image'; // 插入编辑图片
  import 'tinymce/plugins/importcss'; // 引入css
  import 'tinymce/plugins/insertdatetime'; // 插入日期时间
  import 'tinymce/plugins/link'; // 超链接
  import 'tinymce/plugins/lists'; // 列表插件
  import 'tinymce/plugins/media'; // 插入编辑媒体
  import 'tinymce/plugins/nonbreaking'; // 插入不间断空格
  import 'tinymce/plugins/pagebreak'; // 插入分页符
  import 'tinymce/plugins/preview'; // 预览
  import 'tinymce/plugins/quickbars'; // 快速工具栏
  import 'tinymce/plugins/save'; // 保存
  import 'tinymce/plugins/searchreplace'; // 查找替换
  import 'tinymce/plugins/table'; // 表格
  import 'tinymce/plugins/template'; // 内容模板
  import 'tinymce/plugins/visualblocks'; // 显示元素范围
  import 'tinymce/plugins/visualchars'; // 显示不可见字符
  import 'tinymce/plugins/wordcount'; // 字数统计

  const props = defineProps({
    modelValue: {
      type: String,
      required: true,
      default: '',
    },
    menubar: {
      type: [Boolean, String],
      default: 'file edit insert view format table tools help',
    },

    height: {
      type: Number,
      default: 500,
    },
    id: {
      type: [String, Number],
      default: 'mytinymce',
    },
  });

  const emit = defineEmits(['update:modelValue']);

  const contentValue = computed({
    get() {
      return props.modelValue;
    },
    set(value) {
      emit('update:modelValue', value);
    },
  });

  const initOptions = {
    language: 'zh-Hans', // 汉化
    skin_url: '/tinymce/skins/ui/oxide', // 皮肤
    content_style:
      'body{font-size:14px;font-family:Microsoft YaHei,微软雅黑,宋体,Arial,Helvetica,sans-serif;line-height:1.5}img {max-width:100%;}',
    height: props.height,
    menubar: 'file edit view insert format tools table help',
    menu: {
      file: {
        title: 'File',
        items: 'newdocument | preview | export | deleteallconversations',
      },
      edit: {
        title: 'Edit',
        items: 'undo redo restoredraft | cut copy | selectall | searchreplace',
      },
      view: {
        title: 'View',
        items:
          'code | visualaid visualchars visualblocks | preview fullscreen | showcomments',
      },
      insert: {
        title: 'Insert',
        items:
          'image link media addcomment pageembed template codesample inserttable | charmap emoticons | pagebreak nonbreaking anchor tableofcontents | insertdatetime',
      },
      format: {
        title: 'Format',
        items:
          'bold italic underline strikethrough superscript subscript codeformat | styles blocks fontfamily fontsize align lineheight | forecolor backcolor | language | removeformat',
      },
      tools: { title: 'Tools', items: 'a11ycheck code wordcount' },
      table: {
        title: 'Table',
        items:
          'inserttable | cell row column | advtablesort | tableprops deletetable',
      },
      help: { title: 'Help', items: 'help' },
    },
    toolbar:
      'fullscreen | code forecolor backcolor bold italic underline strikethrough link anchor | alignleft aligncenter alignright alignjustify outdent indent  lineheight | styleselect formatselect fontselect fontsizeselect | bullist numlist | blockquote subscript superscript removeformat | table image media | indent2em formatpainter axupimgs',
    plugins:
      'code codesample preview searchreplace autolink directionality visualblocks visualchars fullscreen image link media template table charmap pagebreak nonbreaking anchor insertdatetime advlist lists wordcount autosave ',
    line_height_formats: '1 1.2 1.4 1.6 2', // 行高
    font_size_formats:
      '12px 14px 16px 18px 20px 22px 24px 28px 32px 36px 48px 56px 72px', // 字体大小
    font_family_formats:
      '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;',
    // 如需ajax上传可参考https://www.luweipai.cn/html/1670332196/
    images_file_types: 'jpeg,jpg,png,gif,bmp',
    images_upload_handler: (blobInfo: any, success: any, failure: any) => {
      const img = `data:image/jpeg;base64,${blobInfo.base64()}`;
      success(img);
    },
    placeholder: '在这里输入文字',
    branding: false, // tiny技术支持信息是否显示
    statusbar: false, // 最下方的元素路径和字数统计那一栏是否显示
    elementpath: false, // 元素路径是否显示
    custom_undo_redo_levels: 10, // 撤销和重做的次数
    draggable_modal: true, // 对话框允许拖拽
  };

  onMounted(async () => {
    tinymce.init({}); // 初始化
  });
  onUnmounted(() => {
    tinymce.remove(); // 销毁
  });
</script>

<template>
  <div class="tinymce-box">
    <Editor :id="props.id" v-model="contentValue" :init="initOptions" />
  </div>
</template>

<style scoped>
  .tinymce-box {
    width: 100%;
  }
</style>
