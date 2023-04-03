<template>
  <icon-arrow-down v-if="type === 'arrow-down'" :size="size" />
  <icon-arrow-fall v-if="type === 'arrow-fall'" :size="size" />
  <icon-arrow-left v-if="type === 'arrow-left'" :size="size" />
  <icon-arrow-right v-if="type === 'arrow-right'" :size="size" />
  <icon-arrow-rise v-if="type === 'arrow-rise'" :size="size" />
  <icon-arrow-up v-if="type === 'arrow-up'" :size="size" />
  <icon-caret-down v-if="type === 'caret-down'" :size="size" />
  <icon-caret-left v-if="type === 'caret-left'" :size="size" />
  <icon-caret-right v-if="type === 'caret-right'" :size="size" />
  <icon-caret-up v-if="type === 'caret-up'" :size="size" />
  <icon-double-down v-if="type === 'double-down'" :size="size" />
  <icon-double-left v-if="type === 'double-left'" :size="size" />
  <icon-double-right v-if="type === 'double-right'" :size="size" />
  <icon-double-up v-if="type === 'double-up'" :size="size" />
  <icon-down-circle v-if="type === 'down-circle'" :size="size" />
  <icon-down v-if="type === 'down'" :size="size" />
  <icon-drag-arrow v-if="type === 'drag-arrow'" :size="size" />
  <icon-expand v-if="type === 'expand'" :size="size" />
  <icon-left-circle v-if="type === 'left-circle'" :size="size" />
  <icon-left v-if="type === 'left'" :size="size" />
  <icon-menu-fold v-if="type === 'menu-fold'" :size="size" />
  <icon-menu-unfold v-if="type === 'menu-unfold'" :size="size" />
  <icon-right-circle v-if="type === 'right-circle'" :size="size" />
  <icon-right v-if="type === 'right'" :size="size" />
  <icon-rotate-left v-if="type === 'rotate-left'" :size="size" />
  <icon-rotate-right v-if="type === 'rotate-right'" :size="size" />
  <icon-shrink v-if="type === 'shrink'" :size="size" />
  <icon-swap v-if="type === 'swap'" :size="size" />
  <icon-to-bottom v-if="type === 'to-bottom'" :size="size" />
  <icon-to-left v-if="type === 'to-left'" :size="size" />
  <icon-to-right v-if="type === 'to-right'" :size="size" />
  <icon-to-top v-if="type === 'to-top'" :size="size" />
  <icon-up-circle v-if="type === 'up-circle'" :size="size" />
  <icon-up v-if="type === 'up'" :size="size" />
  <icon-check-circle-fill v-if="type === 'check-circle-fill'" :size="size" />
  <icon-close-circle-fill v-if="type === 'close-circle-fill'" :size="size" />
  <icon-exclamation-circle-fill
    v-if="type === 'exclamation-circle-fill'"
    :size="size"
  />
  <icon-exclamation-polygon-fill
    v-if="type === 'exclamation-polygon-fill'"
    :size="size"
  />
  <icon-info-circle-fill v-if="type === 'info-circle-fill'" :size="size" />
  <icon-minus-circle-fill v-if="type === 'minus-circle-fill'" :size="size" />
  <icon-plus-circle-fill v-if="type === 'plus-circle-fill'" :size="size" />
  <icon-question-circle-fill
    v-if="type === 'question-circle-fill'"
    :size="size"
  />
  <icon-check-circle v-if="type === 'check-circle'" :size="size" />
  <icon-check-square v-if="type === 'check-square'" :size="size" />
  <icon-check v-if="type === 'check'" :size="size" />
  <icon-clock-circle v-if="type === 'clock-circle'" :size="size" />
  <icon-close-circle v-if="type === 'close-circle'" :size="size" />
  <icon-close v-if="type === 'close'" :size="size" />
  <icon-exclamation-circle v-if="type === 'exclamation-circle'" :size="size" />
  <icon-exclamation v-if="type === 'exclamation'" :size="size" />
  <icon-info-circle v-if="type === 'info-circle'" :size="size" />
  <icon-info v-if="type === 'info'" :size="size" />
  <icon-minus-circle v-if="type === 'minus-circle'" :size="size" />
  <icon-minus v-if="type === 'minus'" :size="size" />
  <icon-plus-circle v-if="type === 'plus-circle'" :size="size" />
  <icon-plus v-if="type === 'plus'" :size="size" />
  <icon-question-circle v-if="type === 'question-circle'" :size="size" />
  <icon-question v-if="type === 'question'" :size="size" />
  <icon-stop v-if="type === 'stop'" :size="size" />
  <icon-heart-fill v-if="type === 'heart-fill'" :size="size" />
  <icon-star-fill v-if="type === 'star-fill'" :size="size" />
  <icon-thumb-down-fill v-if="type === 'thumb-down-fill'" :size="size" />
  <icon-thumb-up-fill v-if="type === 'thumb-up-fill'" :size="size" />
  <icon-at v-if="type === 'at'" :size="size" />
  <icon-cloud-download v-if="type === 'cloud-download'" :size="size" />
  <icon-code-block v-if="type === 'code-block'" :size="size" />
  <icon-code-square v-if="type === 'code-square'" :size="size" />
  <icon-code v-if="type === 'code'" :size="size" />
  <icon-customer-service v-if="type === 'customer-service'" :size="size" />
  <icon-download v-if="type === 'download'" :size="size" />
  <icon-export v-if="type === 'export'" :size="size" />
  <icon-eye-invisible v-if="type === 'eye-invisible'" :size="size" />
  <icon-eye v-if="type === 'eye'" :size="size" />
  <icon-heart v-if="type === 'heart'" :size="size" />
  <icon-history v-if="type === 'history'" :size="size" />
  <icon-home v-if="type === 'home'" :size="size" />
  <icon-import v-if="type === 'import'" :size="size" />
  <icon-launch v-if="type === 'launch'" :size="size" />
  <icon-list v-if="type === 'list'" :size="size" />
  <icon-message-banned v-if="type === 'message-banned'" :size="size" />
  <icon-message v-if="type === 'message'" :size="size" />
  <icon-more-vertical v-if="type === 'more-vertical'" :size="size" />
  <icon-more v-if="type === 'more'" :size="size" />
  <icon-poweroff v-if="type === 'poweroff'" :size="size" />
  <icon-refresh v-if="type === 'refresh'" :size="size" />
  <icon-reply v-if="type === 'reply'" :size="size" />
  <icon-save v-if="type === 'save'" :size="size" />
  <icon-scan v-if="type === 'scan'" :size="size" />
  <icon-search v-if="type === 'search'" :size="size" />
  <icon-select-all v-if="type === 'select-all'" :size="size" />
  <icon-send v-if="type === 'send'" :size="size" />
  <icon-settings v-if="type === 'settings'" :size="size" />
  <icon-share-alt v-if="type === 'share-alt'" :size="size" />
  <icon-share-external v-if="type === 'share-external'" :size="size" />
  <icon-share-internal v-if="type === 'share-internal'" :size="size" />
  <icon-star v-if="type === 'star'" :size="size" />
  <icon-sync v-if="type === 'sync'" :size="size" />
  <icon-thumb-down v-if="type === 'thumb-down'" :size="size" />
  <icon-thumb-up v-if="type === 'thumb-up'" :size="size" />
  <icon-translate v-if="type === 'translate'" :size="size" />
  <icon-upload v-if="type === 'upload'" :size="size" />
  <icon-voice v-if="type === 'voice'" :size="size" />
  <icon-align-center v-if="type === 'align-center'" :size="size" />
  <icon-align-left v-if="type === 'align-left'" :size="size" />
  <icon-align-right v-if="type === 'align-right'" :size="size" />
  <icon-attachment v-if="type === 'attachment'" :size="size" />
  <icon-bg-colors v-if="type === 'bg-colors'" :size="size" />
  <icon-bold v-if="type === 'bold'" :size="size" />
  <icon-brush v-if="type === 'brush'" :size="size" />
  <icon-copy v-if="type === 'copy'" :size="size" />
  <icon-delete v-if="type === 'delete'" :size="size" />
  <icon-edit v-if="type === 'edit'" :size="size" />
  <icon-eraser v-if="type === 'eraser'" :size="size" />
  <icon-filter v-if="type === 'filter'" :size="size" />
  <icon-find-replace v-if="type === 'find-replace'" :size="size" />
  <icon-font-colors v-if="type === 'font-colors'" :size="size" />
  <icon-formula v-if="type === 'formula'" :size="size" />
  <icon-h1 v-if="type === 'h1'" :size="size" />
  <icon-h2 v-if="type === 'h2'" :size="size" />
  <icon-h3 v-if="type === 'h3'" :size="size" />
  <icon-h4 v-if="type === 'h4'" :size="size" />
  <icon-h5 v-if="type === 'h5'" :size="size" />
  <icon-h6 v-if="type === 'h6'" :size="size" />
  <icon-h7 v-if="type === 'h7'" :size="size" />
  <icon-highlight v-if="type === 'highlight'" :size="size" />
  <icon-italic v-if="type === 'italic'" :size="size" />
  <icon-line-height v-if="type === 'line-height'" :size="size" />
  <icon-link v-if="type === 'link'" :size="size" />
  <icon-oblique-line v-if="type === 'oblique-line'" :size="size" />
  <icon-ordered-list v-if="type === 'ordered-list'" :size="size" />
  <icon-original-size v-if="type === 'original-size'" :size="size" />
  <icon-paste v-if="type === 'paste'" :size="size" />
  <icon-quote v-if="type === 'quote'" :size="size" />
  <icon-redo v-if="type === 'redo'" :size="size" />
  <icon-scissor v-if="type === 'scissor'" :size="size" />
  <icon-sort-ascending v-if="type === 'sort-ascending'" :size="size" />
  <icon-sort-descending v-if="type === 'sort-descending'" :size="size" />
  <icon-sort v-if="type === 'sort'" :size="size" />
  <icon-strikethrough v-if="type === 'strikethrough'" :size="size" />
  <icon-underline v-if="type === 'underline'" :size="size" />
  <icon-undo v-if="type === 'undo'" :size="size" />
  <icon-unordered-list v-if="type === 'unordered-list'" :size="size" />
  <icon-zoom-in v-if="type === 'zoom-in'" :size="size" />
  <icon-zoom-out v-if="type === 'zoom-out'" :size="size" />
  <icon-mute-fill v-if="type === 'mute-fill'" :size="size" />
  <icon-pause-circle-fill v-if="type === 'pause-circle-fill'" :size="size" />
  <icon-play-arrow-fill v-if="type === 'play-arrow-fill'" :size="size" />
  <icon-play-circle-fill v-if="type === 'play-circle-fill'" :size="size" />
  <icon-skip-next-fill v-if="type === 'skip-next-fill'" :size="size" />
  <icon-skip-previous-fill v-if="type === 'skip-previous-fill'" :size="size" />
  <icon-sound-fill v-if="type === 'sound-fill'" :size="size" />
  <icon-backward v-if="type === 'backward'" :size="size" />
  <icon-forward v-if="type === 'forward'" :size="size" />
  <icon-fullscreen-exit v-if="type === 'fullscreen-exit'" :size="size" />
  <icon-fullscreen v-if="type === 'fullscreen'" :size="size" />
  <icon-live-broadcast v-if="type === 'live-broadcast'" :size="size" />
  <icon-music v-if="type === 'music'" :size="size" />
  <icon-mute v-if="type === 'mute'" :size="size" />
  <icon-pause-circle v-if="type === 'pause-circle'" :size="size" />
  <icon-pause v-if="type === 'pause'" :size="size" />
  <icon-play-arrow v-if="type === 'play-arrow'" :size="size" />
  <icon-play-circle v-if="type === 'play-circle'" :size="size" />
  <icon-record-stop v-if="type === 'record-stop'" :size="size" />
  <icon-record v-if="type === 'record'" :size="size" />
  <icon-skip-next v-if="type === 'skip-next'" :size="size" />
  <icon-skip-previous v-if="type === 'skip-previous'" :size="size" />
  <icon-sound v-if="type === 'sound'" :size="size" />
  <icon-bytedance-color v-if="type === 'bytedance-color'" :size="size" />
  <icon-lark-color v-if="type === 'lark-color'" :size="size" />
  <icon-tiktok-color v-if="type === 'tiktok-color'" :size="size" />
  <icon-xigua-color v-if="type === 'xigua-color'" :size="size" />
  <icon-faceBook-circle-fill
    v-if="type === 'faceBook-circle-fill'"
    :size="size"
  />
  <icon-facebook-square-fill
    v-if="type === 'facebook-square-fill'"
    :size="size"
  />
  <icon-google-circle-fill v-if="type === 'google-circle-fill'" :size="size" />
  <icon-qq-circle-fill v-if="type === 'qq-circle-fill'" :size="size" />
  <icon-twitter-circle-fill
    v-if="type === 'twitter-circle-fill'"
    :size="size"
  />
  <icon-weibo-circle-fill v-if="type === 'weibo-circle-fill'" :size="size" />
  <icon-alipay-circle v-if="type === 'alipay-circle'" :size="size" />
  <icon-code-sandbox v-if="type === 'code-sandbox'" :size="size" />
  <icon-codepen v-if="type === 'codepen'" :size="size" />
  <icon-facebook v-if="type === 'facebook'" :size="size" />
  <icon-github v-if="type === 'github'" :size="size" />
  <icon-gitlab v-if="type === 'gitlab'" :size="size" />
  <icon-google v-if="type === 'google'" :size="size" />
  <icon-qq-zone v-if="type === 'qq-zone'" :size="size" />
  <icon-qq v-if="type === 'qq'" :size="size" />
  <icon-twitter v-if="type === 'twitter'" :size="size" />
  <icon-wechat v-if="type === 'wechat'" :size="size" />
  <icon-wechatpay v-if="type === 'wechatpay'" :size="size" />
  <icon-weibo v-if="type === 'weibo'" :size="size" />
  <icon-chinese-fill v-if="type === 'chinese-fill'" :size="size" />
  <icon-english-fill v-if="type === 'english-fill'" :size="size" />
  <icon-face-frown-fill v-if="type === 'face-frown-fill'" :size="size" />
  <icon-face-meh-fill v-if="type === 'face-meh-fill'" :size="size" />
  <icon-face-smile-fill v-if="type === 'face-smile-fill'" :size="size" />
  <icon-moon-fill v-if="type === 'moon-fill'" :size="size" />
  <icon-pen-fill v-if="type === 'pen-fill'" :size="size" />
  <icon-sun-fill v-if="type === 'sun-fill'" :size="size" />
  <icon-apps v-if="type === 'apps'" :size="size" />
  <icon-archive v-if="type === 'archive'" :size="size" />
  <icon-bar-chart v-if="type === 'bar-chart'" :size="size" />
  <icon-book v-if="type === 'book'" :size="size" />
  <icon-bookmark v-if="type === 'bookmark'" :size="size" />
  <icon-branch v-if="type === 'branch'" :size="size" />
  <icon-bug v-if="type === 'bug'" :size="size" />
  <icon-bulb v-if="type === 'bulb'" :size="size" />
  <icon-calendar-clock v-if="type === 'calendar-clock'" :size="size" />
  <icon-calendar v-if="type === 'calendar'" :size="size" />
  <icon-camera v-if="type === 'camera'" :size="size" />
  <icon-cloud v-if="type === 'cloud'" :size="size" />
  <icon-command v-if="type === 'command'" :size="size" />
  <icon-common v-if="type === 'common'" :size="size" />
  <icon-compass v-if="type === 'compass'" :size="size" />
  <icon-computer v-if="type === 'computer'" :size="size" />
  <icon-copyright v-if="type === 'copyright'" :size="size" />
  <icon-dashboard v-if="type === 'dashboard'" :size="size" />
  <icon-desktop v-if="type === 'desktop'" :size="size" />
  <icon-dice v-if="type === 'dice'" :size="size" />
  <icon-drag-dot-vertical v-if="type === 'drag-dot-vertical'" :size="size" />
  <icon-drag-dot v-if="type === 'drag-dot'" :size="size" />
  <icon-drive-file v-if="type === 'drive-file'" :size="size" />
  <icon-ear v-if="type === 'ear'" :size="size" />
  <icon-email v-if="type === 'email'" :size="size" />
  <icon-empty v-if="type === 'empty'" :size="size" />
  <icon-experiment v-if="type === 'experiment'" :size="size" />
  <icon-file-audio v-if="type === 'file-audio'" :size="size" />
  <icon-file-image v-if="type === 'file-image'" :size="size" />
  <icon-file-pdf v-if="type === 'file-pdf'" :size="size" />
  <icon-file-video v-if="type === 'file-video'" :size="size" />
  <icon-file v-if="type === 'file'" :size="size" />
  <icon-fire v-if="type === 'fire'" :size="size" />
  <icon-folder-add v-if="type === 'folder-add'" :size="size" />
  <icon-folder-delete v-if="type === 'folder-delete'" :size="size" />
  <icon-folder v-if="type === 'folder'" :size="size" />
  <icon-gift v-if="type === 'gift'" :size="size" />
  <icon-idcard v-if="type === 'idcard'" :size="size" />
  <icon-image-close v-if="type === 'image-close'" :size="size" />
  <icon-image v-if="type === 'image'" :size="size" />
  <icon-interaction v-if="type === 'interaction'" :size="size" />
  <icon-language v-if="type === 'language'" :size="size" />
  <icon-layers v-if="type === 'layers'" :size="size" />
  <icon-layout v-if="type === 'layout'" :size="size" />
  <icon-loading v-if="type === 'loading'" :size="size" />
  <icon-location v-if="type === 'location'" :size="size" />
  <icon-lock v-if="type === 'lock'" :size="size" />
  <icon-loop v-if="type === 'loop'" :size="size" />
  <icon-man v-if="type === 'man'" :size="size" />
  <icon-menu v-if="type === 'menu'" :size="size" />
  <icon-mind-mapping v-if="type === 'mind-mapping'" :size="size" />
  <icon-mobile v-if="type === 'mobile'" :size="size" />
  <icon-moon v-if="type === 'moon'" :size="size" />
  <icon-mosaic v-if="type === 'mosaic'" :size="size" />
  <icon-nav v-if="type === 'nav'" :size="size" />
  <icon-notification-close v-if="type === 'notification-close'" :size="size" />
  <icon-notification v-if="type === 'notification'" :size="size" />
  <icon-palette v-if="type === 'palette'" :size="size" />
  <icon-pen v-if="type === 'pen'" :size="size" />
  <icon-phone v-if="type === 'phone'" :size="size" />
  <icon-printer v-if="type === 'printer'" :size="size" />
  <icon-public v-if="type === 'public'" :size="size" />
  <icon-pushpin v-if="type === 'pushpin'" :size="size" />
  <icon-qrcode v-if="type === 'qrcode'" :size="size" />
  <icon-relation v-if="type === 'relation'" :size="size" />
  <icon-robot-add v-if="type === 'robot-add'" :size="size" />
  <icon-robot v-if="type === 'robot'" :size="size" />
  <icon-safe v-if="type === 'safe'" :size="size" />
  <icon-schedule v-if="type === 'schedule'" :size="size" />
  <icon-shake v-if="type === 'shake'" :size="size" />
  <icon-skin v-if="type === 'skin'" :size="size" />
  <icon-stamp v-if="type === 'stamp'" :size="size" />
  <icon-storage v-if="type === 'storage'" :size="size" />
  <icon-subscribe-add v-if="type === 'subscribe-add'" :size="size" />
  <icon-subscribe v-if="type === 'subscribe'" :size="size" />
  <icon-subscribed v-if="type === 'subscribed'" :size="size" />
  <icon-sun v-if="type === 'sun'" :size="size" />
  <icon-tag v-if="type === 'tag'" :size="size" />
  <icon-tags v-if="type === 'tags'" :size="size" />
  <icon-thunderbolt v-if="type === 'thunderbolt'" :size="size" />
  <icon-tool v-if="type === 'tool'" :size="size" />
  <icon-trophy v-if="type === 'trophy'" :size="size" />
  <icon-unlock v-if="type === 'unlock'" :size="size" />
  <icon-user-add v-if="type === 'user-add'" :size="size" />
  <icon-user-group v-if="type === 'user-group'" :size="size" />
  <icon-user v-if="type === 'user'" :size="size" />
  <icon-video-camera v-if="type === 'video-camera'" :size="size" />
  <icon-wifi v-if="type === 'wifi'" :size="size" />
  <icon-woman v-if="type === 'woman'" :size="size" />
</template>

<script lang="ts" setup>
  defineProps({
    type: String,
    size: {
      type: Number,
      default: 32,
    },
  });
</script>
