import Vue from 'vue';

declare module '@vue/runtime-core' {
  export interface ComponentCustomProperties {
    $staticUrl: any;
    $parseTime: any;
    $formatTime: any;
    $localeOptions: any;
    $densityList: any;
    $hasPermission: any;
  }
}

declare module 'json2yaml';
