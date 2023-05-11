<template>
  <ImageCropper
    v-if="showCropper"
    :image="cropImage.url"
    @confirm="handleCropImage"
    @cancel="showCropper = false"
  ></ImageCropper>
  <!--  -->

  <a-upload
    ref="uploadRef"
    :custom-request="customRequest"
    :list-type="(listType as any)"
    :file-list="uploadedFileList"
    :show-upload-button="true"
    :show-file-list="multiple"
    :auto-upload="autoUpload"
    image-preview
    :limit="limit"
    :accept="accept"
    @change="uploadChange"
  >
    <template #upload-button>
      <a-avatar
        v-if="uploadedFileList.length && !multiple"
        :size="size"
        :fade="true"
        :shape="shape as AvatarShape"
      >
        <template #trigger-icon>
          <icon-camera />
        </template>
        <img
          :src="
            uploadedFileList[0].url?.indexOf('://') == -1
              ? $staticUrl + uploadedFileList[0].url
              : uploadedFileList[0].url
          "
        />
      </a-avatar>
      <div v-else class="upload-card" :class="shape">
        <icon-camera v-if="accept == 'image'" class="icon" />
        <icon-plus v-else class="icon"></icon-plus>
        <span v-if="showText" class="text">{{ text }}</span>
      </div>
    </template>
  </a-upload>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import type {
    FileItem,
    RequestOption,
  } from '@arco-design/web-vue/es/upload/interfaces';
  import { AvatarShape } from '@arco-design/web-vue/es/avatar/interface';
  import ImageCropper from '@/components/image-cropper/index.vue';
  import { uploadFile } from '@/api/system/upload';
  import { useAppStore } from '@/store';

  const appStore = useAppStore();

  const emit = defineEmits(['confirm']);
  const props = defineProps({
    url: String,
    type: String,
    multiple: {
      type: Boolean,
      default: false,
    },
    shape: {
      type: String,
      default: 'square',
    },
    size: {
      type: Number,
      default: 80,
    },
    showText: {
      type: Boolean,
      default: true,
    },
    text: {
      type: String,
      default: '点击上传',
    },
    files: {
      type: Array,
      default: () => {
        return [];
      },
    },
    limit: {
      type: Number,
      default: 9,
    },
    cut: {
      type: Boolean,
      default: true,
    },
    listType: {
      type: String,
      default: 'picture-card',
    },
    accept: {
      type: String,
      default: 'image/*',
    },
    dir: {
      type: String,
      require: true,
    },
    autoUpload: Boolean,
    file: String,
  });
  const uploadRef = ref();

  const showCropper = ref(false);

  const uploadedFileList = ref<FileItem[]>([]);
  const domSize = ref(`${props.size}px`);
  const cropImage = ref({} as FileItem);

  const chooesFileList = ref<FileItem[]>([]);

  const uuid = () => {
    let d = new Date().getTime();
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
      // eslint-disable-next-line no-bitwise
      const r = (d + Math.random() * 16) % 16 | 0;
      d = Math.floor(d / 16);
      // eslint-disable-next-line no-bitwise
      return (c === 'x' ? r : (r & 0x3) | 0x8).toString(16);
    });
  };

  const fileName = (url: string) => {
    const index = url.lastIndexOf('/');
    if (index === -1) {
      return url;
    }
    return url.substring(index + 1);
  };

  // init 组件初始化
  const initComponent = () => {
    if (props.multiple) {
      props.files.forEach((fileUrl) => {
        uploadedFileList.value.push({
          uid: uuid(),
          url: fileUrl,
          name: fileName(fileUrl as string),
        } as FileItem);
      });
    } else if (props.file) {
      uploadedFileList.value.push({
        uid: uuid(),
        url: props.file,
        name: fileName(props.file as string),
      } as FileItem);
    }
  };

  initComponent();

  // isCutImage 是否为可裁剪的图片
  const isImage = (filename: string) => {
    const names = filename.split('.');
    const type = names[names.length - 1];
    const arr = ['png', 'jpg', 'jpeg'];
    return arr.includes(type);
  };

  //   const setUploadListUrl = (id: string, url: string) => {
  //     uploadedFileList.value.forEach((item, index) => {
  //       if (item.uid === id) {
  //         uploadedFileList.value[index].url = url;
  //       }
  //     });
  //   };

  // 进行文件对比并提交
  const compareEmit = () => {
    const uploadList: any[] = [];
    const uploadListItem: FileItem[] = [];
    uploadedFileList.value.forEach((item) => {
      let has = false;
      chooesFileList.value.forEach((citem) => {
        if (citem.uid === item.uid) {
          has = true;
        }
      });

      if (has) {
        uploadList.push(item);
        uploadListItem.push(item);
      }
    });
    uploadedFileList.value = uploadListItem;
    emit('confirm', uploadList);
  };

  // 上传修改触发
  const uploadChange = (fileItemList: FileItem[], fileItem: FileItem) => {
    if (!props.multiple) {
      if (
        props.cut &&
        isImage(fileItem.name as string) &&
        fileItem.status === 'init'
      ) {
        showCropper.value = true;
        cropImage.value = fileItem;
      }
      chooesFileList.value = [fileItem];
    } else {
      chooesFileList.value = fileItemList;
    }
  };

  // handleCropImage 裁剪图片
  const handleCropImage = (data: any) => {
    cropImage.value.file = data.file;
    cropImage.value.url = data.url;

    uploadRef.value.submit(cropImage.value);
  };

  const customRequest = (options: RequestOption) => {
    showCropper.value = false;
    appStore.startLoading('文件上传中');
    // docs: https://axios-http.com/docs/cancellation
    const controller = new AbortController();

    (async function requestWrap() {
      const { onError, onSuccess, fileItem } = options;
      const req = new FormData();
      req.append('dir', props.dir as any);
      req.append('file', fileItem.file as Blob);
      try {
        const { data } = await uploadFile(req);
        fileItem.url = data.file;
        uploadedFileList.value.push(fileItem);
        compareEmit();
        onSuccess('success');
      } catch (error) {
        onError(error);
      } finally {
        appStore.stopLoading();
      }
    })();
    return {
      abort() {
        controller.abort();
      },
    };
  };
</script>

<style scoped lang="less">
  .upload-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: v-bind(domSize);
    height: v-bind(domSize);
    background-color: #f4f5f7;
    border-radius: 2px;
    box-sizing: border-box;

    .text {
      font-size: 12px;
      margin-top: 10px;
      color: #999;
    }
    .icon {
      font-size: 22px;
      line-height: 22px;
      font-weight: normal;
      top: 0px;
      color: rgb(211, 212, 214);
    }

    &:hover {
      background-color: #f1f1f1;
    }
  }

  .circle {
    border-radius: 50%;
  }
  .square {
    border-radius: 2px;
  }
</style>
