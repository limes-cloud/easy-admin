<template>
  <a-modal
    title="图像剪裁"
    :visible="visible"
    :mask-closable="false"
    :width="width"
    :hide-cancel="hideCancel"
    :footer="false"
    :body-style="{
      padding: 0,
    }"
    unmount-on-close
    @cancel="cancel"
  >
    <ImgCutter
      ref="imgCutterRef"
      label="选择本地图片"
      file-type="jpeg"
      :cross-origin="true"
      cross-origin-header="*"
      rate=""
      tool-bgc="none"
      :is-modal="false"
      :show-choose-btn="false"
      :lock-scroll="true"
      :box-width="width"
      :box-height="height"
      :cut-width="cutWith"
      :cut-height="cutHeight"
      :size-change="false"
      :move-able="true"
      :img-move="false"
      :original-graph="true"
      watermark-text=""
      watermark-text-font="20px Sans-serif"
      watermark-text-color="#00ff00"
      :watermark-text-x="1"
      :watermark-text-y="1"
      :small-to-upload="true"
      :save-cut-position="true"
      :scale-able="true"
      :preview-mode="false"
      :quality="1"
      :tool-box-overflow="false"
      :index="1"
      @cut-down="ok"
    >
      <template #cancel>
        <a-button class="btn" @click="cancel">取消</a-button>
      </template>
      <template #confirm>
        <a-button class="btn" type="primary" @click="ok">确认</a-button>
      </template>
    </ImgCutter>
  </a-modal>
</template>

<script lang="ts">
  // @ts-ignore
  import ImgCutter from 'vue-img-cutter';

  //   function dataURLtoBlob(dataurl: any) {
  //     const arr = dataurl.split(',');
  //     const mime = arr[0].match(/:(.*?);/)[1];
  //     const bstr = atob(arr[1]);
  //     let n = bstr.length;
  //     const u8arr = new Uint8Array(n);
  //     // eslint-disable-next-line no-plusplus
  //     while (n--) {
  //       u8arr[n] = bstr.charCodeAt(n);
  //     }
  //     return new Blob([u8arr], { type: mime });
  //   }

  export default {
    components: { ImgCutter },
    props: {
      image: {
        type: String,
        default: '',
      },
      hideCancel: {
        type: Boolean,
        default: false,
      },
      width: {
        type: Number,
        default: 350,
      },
      height: {
        type: Number,
        default: 350,
      },
      cutWith: {
        type: Number,
        default: 250,
      },
      cutHeight: {
        type: Number,
        default: 250,
      },
    },
    emits: ['confirm', 'cancel'],
    data: () => {
      return {
        visible: true,
      };
    },
    mounted() {
      this.open();
    },
    methods: {
      ok(data: any) {
        if (!data.file) return;
        this.$emit('confirm', {
          file: data.file,
          url: data.dataURL,
        });
        this.visible = false;
      },
      cancel() {
        this.visible = false;
        this.$emit('cancel', false);
      },
      open() {
        this.$nextTick(() => {
          this.$refs.imgCutterRef.handleOpen({
            name: 'file.png',
            src: this.image,
          });
        });
      },
    },
  };
</script>

<style scoped lang="less">
  .btn {
    margin-right: 15px;
  }
</style>
