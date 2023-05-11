<template>
  <div class="file-container">
    <div class="header">
      <div class="icon">
        <a-tooltip content="前往上层文件夹" mini :content-style="tipStyle">
          <span><icon-arrow-left size="22" :stroke-width="2" /></span>
        </a-tooltip>
      </div>
      <div class="path-input">
        <a-input placeholder="请输入文件路径">
          <template #prefix>
            <icon-home />
          </template>
        </a-input>
      </div>
      <a-badge
        class="badge"
        :offset="[-10, 2]"
        :count="isCut ? selectd.length : 0"
        :dot-style="{
          width: '15px',
          height: '15px',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }"
      >
        <div class="icon" :class="!selectd.length ? 'disable' : ''">
          <a-tooltip content="剪贴" mini :content-style="tipStyle">
            <span :class="isCut ? 'active' : ''" @click="cutFile"
              ><icon-scissor size="22" :stroke-width="2"
            /></span>
          </a-tooltip>
        </div>
      </a-badge>

      <div class="icon" :class="!(selectd.length && isCut) ? 'disable' : ''">
        <a-tooltip content="粘贴" mini :content-style="tipStyle">
          <span @click="pasteFile"
            ><icon-paste size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="icon" :class="!selectd.length ? 'disable' : ''">
        <a-tooltip content="删除" mini :content-style="tipStyle">
          <span @click="deleteFile"
            ><icon-delete size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="icon">
        <a-tooltip content="新建文件夹" mini :content-style="tipStyle">
          <span @click="newFileVisible = true"
            ><icon-folder-add size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="icon">
        <a-tooltip content="列表展示" mini :content-style="tipStyle">
          <span :class="!showCard ? 'active' : ''" @click="showCard = false"
            ><icon-mind-mapping size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="icon">
        <a-tooltip content="图标展示" mini :content-style="tipStyle">
          <span :class="showCard ? 'active' : ''" @click="showCard = true"
            ><icon-apps size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="icon">
        <a-tooltip content="上传文件" mini :content-style="tipStyle">
          <span @click="uploadFileVisible = true"
            ><icon-upload size="22" :stroke-width="2"
          /></span>
        </a-tooltip>
      </div>
      <div class="search-input">
        <a-input
          :style="{ width: '180px' }"
          placeholder="请输入文件名"
          allow-clear
        >
          <template #prefix>
            <icon-search />
          </template>
        </a-input>
      </div>
    </div>
    <div class="body">
      <div v-if="showCard" class="card">
        <div v-for="file in files" :key="file.id" class="item">
          <div
            class="file"
            :class="isSelectd(file.id) ? 'active' : ''"
            @click.exact="fileSelect(file)"
            @click.alt.exact="multipleFileSelect(file)"
          >
            <div class="item-icon">
              <img
                v-if="fileType(file) === 'image'"
                :src="$staticUrl + file.src"
              />
              <icon
                v-if="fileType(file) === 'file' || fileType(file) === 'dir'"
                :type="file.suffix"
              ></icon>

              <div v-if="fileType(file) === 'video'" class="video">
                <video :id="'video-' + file.id" preload="metadata">
                  <source :src="file.src + '#t=1'" :type="file.mime" />
                </video>
                <div class="cover">
                  <icon-play-circle
                    class="icon"
                    size="28"
                    @click="openFile(file)"
                  />
                </div>
              </div>

              <div v-if="fileType(file) === 'music'" class="music">
                <div class="cover">
                  <svgIcon
                    class="bg"
                    name="file-music-bg"
                    :size="100"
                  ></svgIcon>
                  <icon-play-circle
                    class="icon"
                    size="28"
                    @click="openFile(file)"
                  />
                </div>
              </div>
            </div>
            <a-tooltip :content="file.name" mini :content-style="tipStyle">
              <div class="item-desc">
                {{ file.name }}
              </div>
            </a-tooltip>
          </div>
        </div>
      </div>

      <div v-else class="list">
        <a-table
          v-model:selectedKeys="selectd"
          row-key="id"
          :columns="listColumns"
          :data="files"
          :sticky-header="400"
          :scroll="{
            x: '100%',
            y: '100%',
          }"
          :bordered="false"
          :pagination="false"
          :row-selection="{
            type: 'checkbox',
            showCheckedAll: true,
            onlyCurrent: false,
          }"
          row-class="table-row"
        >
          <template #name="{ record }">
            <div class="name">
              <icon class="icon" :type="record.suffix" :size="20" />
              <div class="text" @click="openFile(record)">{{
                record.name
              }}</div>
            </div>
          </template>

          <template #type="{ record }">
            <span v-if="record.is_dir">文件夹</span>
            <span v-else>文件</span>
          </template>

          <template #created_at="{ record }">
            {{ $formatTime(record.created_at) }}
          </template>

          <template #size="{ record }">
            {{ getFileSize(record.size) }}
          </template>
        </a-table>
      </div>

      <div v-if="!files.length" class="empty">
        <a-empty>
          <template #image>
            <svgIcon name="empty-data" size="220" />
          </template>
          文件夹为空
        </a-empty>
      </div>
    </div>

    <a-modal
      v-model:visible="playVideoVisible"
      width="480px"
      :hide-cancel="true"
      :footer="false"
      :body-style="{ padding: 0 }"
      unmount-on-close
    >
      <template #title>
        <div>{{ curFile.name }}</div>
      </template>
      <video class="video-model" autoplay controls preload="metadata">
        <source :src="curFile.src + '#t=1'" :type="curFile.mime" />
        当前浏览器不支持 video直接播放
      </video>
    </a-modal>

    <a-modal
      v-model:visible="playMusicVisible"
      width="480px"
      :hide-cancel="true"
      :footer="false"
      unmount-on-close
    >
      <template #title>
        <div>{{ curFile.name }}</div>
      </template>
      <audio
        class="music-model"
        autoplay
        controls
        preload="metadata"
        :src="curFile.src + '#t=1'"
      >
      </audio>
    </a-modal>

    <a-modal
      v-model:visible="showImageVisible"
      width="480px"
      :hide-cancel="true"
      :footer="false"
      :body-style="{ padding: 0 }"
    >
      <template #title>
        <div>{{ curFile.name }}</div>
      </template>
      <img style="width: 100%" :src="$staticUrl + curFile.src" />
    </a-modal>

    <a-modal
      v-model:visible="newFileVisible"
      width="380px"
      title="新建文件夹"
      :on-before-ok="newFileDir"
      :body-style="{ padding: '15px 15px' }"
      @ok="playMusicVisible = false"
    >
      <a-form :model="newFileForm" layout="vertical">
        <a-form-item label-width="0">
          <a-input v-model="newFileForm.name" placeholder="请输入文件夹名称" />
        </a-form-item>
      </a-form>
    </a-modal>

    <a-modal
      v-model:visible="uploadFileVisible"
      width="480px"
      title="文件上传"
      :on-before-ok="newFile"
      :body-style="{ padding: '15px 15px' }"
      @ok="playMusicVisible = false"
    >
      <!-- <Upload
        :file="userInfo.avatar"
        :multiple="false"
        :size="100"
        shape="circle"
        @confirm="uploadCallback"
      ></Upload> -->
      <Upload
        :size="100"
        :multiple="true"
        :cut="false"
        :auto-upload="true"
        list-type="text"
        accept="file"
        dir="file"
        @confirm="uploadFile"
      ></Upload>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import Message from '@arco-design/web-vue/es/message';
  import icon from './icon.vue';
  import Upload from '../upload/index.vue';

  interface File {
    id: number;
    user_id: number;
    src: string;
    name: string;
    is_dir: boolean;
    parent_id: number;
    size: number;
    suffix: string;
    mime: string;
    operator: string;
    operator_id: number;
    created_at: number;
  }

  const files = ref<File[]>([
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 100,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 1000,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 10000,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 100000,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 1,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件1.png',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 2,
      user_id: 1,
      src: 'static/avatar/ff431714d595fed8503dbb862e66e64e.jpg',
      name: '这是一个测试文件2.png',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'jpg',
      mime: 'image/jpg',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 3,
      user_id: 1,
      src: 'http://96.ierge.cn/15/238/477410.mp3?v=0524',
      name: '这是一个测试文件3.mp3',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'mp3',
      mime: 'music/mp3',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 3,
      user_id: 1,
      src: '',
      name: '这是一个文件夹',
      is_dir: true,
      parent_id: 1,
      size: 0,
      suffix: 'dir',
      mime: 'dir',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
    {
      id: 4,
      user_id: 1,
      src: 'https://video.pearvideo.com/mp4/short/20230510/cont-1782401-71090409-hd.mp4',
      name: '这是一个视频',
      is_dir: false,
      parent_id: 1,
      size: 0,
      suffix: 'mp4',
      mime: 'video/mp4',
      operator: '',
      operator_id: 0,
      created_at: 123456789,
    },
  ]);
  // const files = ref<File[]>([]);
  const props = defineProps({
    multipleSelect: {
      type: Boolean,
      default: true,
    },
  });

  const listColumns = [
    {
      title: '名称',
      dataIndex: 'name',
      slotName: 'name',
    },
    {
      title: '上传时间',
      dataIndex: 'created_at',
      slotName: 'created_at',
      width: 220,
    },
    {
      title: '文件类型',
      dataIndex: 'type',
      slotName: 'type',
      width: 120,
    },
    {
      title: '大小',
      dataIndex: 'size',
      slotName: 'size',
      width: 120,
    },
  ];

  const newFileForm = ref({ name: '' });
  const showCard = ref(false);
  const isCut = ref(false);
  const curFile = ref<File>({} as File);
  const tipStyle = { fontSize: '12px' };
  const selectd = ref<number[]>([]);
  const playVideoVisible = ref(false);
  const playMusicVisible = ref(false);
  const showImageVisible = ref(false);
  const newFileVisible = ref(false);
  const uploadFileVisible = ref(false);
  const newFileList = ref<any[]>([]);
  const fileType = (file: File): string => {
    if (file.mime.toLowerCase().indexOf('image') !== -1) {
      return 'image';
    }
    if (file.mime.toLowerCase().indexOf('video') !== -1) {
      return 'video';
    }
    if (file.mime.toLowerCase().indexOf('music') !== -1) {
      return 'music';
    }
    if (file.is_dir) {
      return 'dir';
    }
    return 'file';
  };

  const isSelectd = (id: number) => {
    return selectd.value.indexOf(id) !== -1;
  };

  const restSelectd = () => {
    selectd.value = [];
  };

  const openFile = (file: File) => {
    curFile.value = file;
    switch (fileType(file)) {
      case 'music':
        playMusicVisible.value = true;
        break;
      case 'video':
        playVideoVisible.value = true;
        break;
      case 'image':
        showImageVisible.value = true;
        break;
      case 'dir':
        console.log('open-dir');
        break;
      default:
        Message.error('暂不支持查看此类型的文件');
        break;
    }
  };

  const fileSelect = (file: File) => {
    restSelectd();
    selectd.value.push(file.id);
    selectd.value = [file.id];
  };

  const cutFile = () => {
    if (selectd.value.length <= 0) {
      return;
    }
    if (!isCut.value) {
      Message.success(`成功选择${selectd.value.length}个剪贴文件`);
    } else {
      Message.success(`取消剪贴成功`);
    }
    isCut.value = !isCut.value;
  };

  const pasteFile = () => {
    // 请求变更父目录的借口
  };
  const deleteFile = () => {
    // 请求删除文件接口
  };

  const newFileDir = () => {
    // 请求文件创建接口
  };

  const newFile = () => {
    // 请求文件上传接口
  };

  const uploadFile = (fs: any[]) => {
    newFileList.value = [];
    fs.forEach((file) => {
      newFileList.value.push({
        src: file.url,
        name: file.name,
        is_dir: false,
        mime: file.file.type,
        size: file.file.size,
      });
    });
  };

  const multipleFileSelect = (file: File) => {
    // 查找当前的文件是否存在列表
    const index = selectd.value.indexOf(file.id);

    if (props.multipleSelect) {
      if (index === -1) {
        selectd.value.push(file.id);
      } else {
        selectd.value.splice(index, 1);
      }
    } else {
      selectd.value = [file.id];
    }
  };

  function getFileSize(size: number) {
    if (!size) return '';
    const num = 1024.0; // byte
    if (size < num) return `${size}B`;
    if (size < num ** 2) return `${(size / num).toFixed(2)}K`; // kb
    if (size < num ** 3) return `${(size / num ** 2).toFixed(2)}M`; // M
    if (size < num ** 4) return `${(size / num ** 3).toFixed(2)}G`; // G
    return `${(size / num ** 4).toFixed(2)}T`; // T
  }
</script>

<style lang="less">
  .table-row {
    .arco-table-td {
      border: none !important;
      font-size: 13px !important;
      color: #333 !important;
      .arco-table-cell {
        padding: 12px 16px;
      }
    }
  }
</style>

<style lang="less" scoped>
  .video-model {
    width: 100%;
    height: 320px;
  }
  .music-model {
    width: 100%;
  }

  .file-container {
    height: calc(100vh - 220px);
    overflow: hidden;
  }
  .header {
    display: flex;
    align-items: center;
    justify-content: space-around;
    height: 40px;
    .icon {
      padding: 0 10px;
      span {
        margin: auto;
        display: block;
        width: 32px;
        height: 32px;
        display: flex;
        justify-content: center;
        align-items: center;
        &:hover {
          background-color: var(--color-neutral-2);
        }
        cursor: pointer;
      }
      text-align: center;
    }
    .active {
      background-color: var(--color-neutral-2);
    }
    .disable {
      color: var(--color-text-4);
    }
    .path-input {
      padding: 0px 15px;
      min-width: 220px;
      flex: 1;
    }

    .search-input {
      padding: 0px 15px;
    }
  }

  .body {
    margin-top: 20px;
    height: calc(100% - 80px);
    position: relative;
    .empty {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100%;
    }

    .list {
      position: absolute;
      height: 100%;
      width: 100%;
      top: 0px;
      left: 0px;
      overflow-y: scroll;
      overflow-x: hidden;
      .name {
        display: flex;
        align-items: center;
        .icon {
          flex: 1;
        }
        .text {
          cursor: pointer;
          margin-left: 5px;
          white-space: nowrap;
          max-width: 180px;
          overflow: hidden;
          text-overflow: ellipsis;
        }
      }
    }
    .card {
      position: absolute;
      width: 100%;
      top: 0px;
      left: 0px;
      overflow-y: scroll;
      overflow-x: hidden;
      display: flex;
      flex-wrap: wrap;
      &::-webkit-scrollbar {
        display: none;
      }
      .item {
        padding: 5px;

        .file {
          padding: 10px 15px;
          border-radius: 4px;
          width: 100px;
          overflow: hidden;
          box-sizing: content-box;
          &:hover {
            background-color: var(--color-neutral-2);
          }
          .item-icon {
            width: 100px;
            height: 100px;
            img {
              width: 100%;
              height: 100%;
              border-radius: 4px;
            }
            .video {
              width: 100%;
              height: 100%;
              position: relative;

              video {
                width: 100%;
                height: 100%;
                object-fit: fill;
                border-radius: 4px;
              }
            }

            .music {
              width: 100%;
              height: 100%;
              position: relative;
              audio {
                width: 100%;
                height: 100%;
                object-fit: fill;
                border-radius: 4px;
              }
            }
          }

          .cover {
            position: absolute;
            top: 0px;
            left: 0px;
            width: 100%;
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            .bg {
              position: absolute;
              top: 0px;
              left: 0px;
              width: 100%;
              height: 100%;
            }
            .icon {
              z-index: 10;
            }
          }
          .item-desc {
            margin-top: 5px;
            font-size: 12px;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 2;
            overflow: hidden;
            text-align: center;
          }
        }
        .active {
          background-color: var(--color-neutral-2);
        }
      }
    }
  }
</style>
