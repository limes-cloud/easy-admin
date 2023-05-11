<template>
  <svgIcon :name="name(type as string)" :size="size"></svgIcon>
</template>

<script lang="ts" setup>
  defineProps({
    type: String,
    size: {
      type: Number,
      default: 100,
    },
  });

  const name = (tp: string) => {
    let typeName = '';
    switch (tp) {
      case 'pdf':
      case 'pdfx':
        typeName = 'pdf';
        break;
      case 'xls':
      case 'xlsx':
        typeName = 'xls';
        break;
      case 'ppt':
      case 'pptx':
        typeName = 'ppt';
        break;
      case 'doc':
      case 'docx':
        typeName = 'doc';
        break;
      case 'wav':
      case 'mp3':
        typeName = 'music';
        break;
      case 'yaml':
      case 'conf':
      case 'yam':
        typeName = 'conf';
        break;
      case 'dir':
        typeName = 'dir';
        break;
      case 'mov':
      case 'avi':
      case 'mp4':
        typeName = 'video';
        break;
      case 'jpg':
      case 'jpeg':
      case 'png':
        typeName = 'image';
        break;
      case 'txt':
        typeName = 'txt';
        break;
      case 'html':
      case 'js':
        typeName = 'html';
        break;
      case 'sql':
        typeName = 'sql';
        break;
      case 'zip':
      case 'rar':
      case '7z':
        typeName = 'zip';
        break;
      default:
        typeName = 'other';
    }
    return `file-${typeName}`;
  };
</script>

<style lang="less" scoped>
  .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 40px;
    color: var(--color-text-2);
    text-align: center;
  }
</style>
