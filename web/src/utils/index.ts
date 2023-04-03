import { useAppStore } from '@/store';
import { INotice } from '@/types/global';
import { Notification } from '@arco-design/web-vue';
import { toInteger } from 'lodash';

export const hasPermission = (r: string) => {
  return useAppStore().permissions.includes(r);
};

type TargetContext = '_self' | '_parent' | '_blank' | '_top';

export const openWindow = (
  url: string,
  opts?: { target?: TargetContext; [key: string]: any }
) => {
  const { target = '_blank', ...others } = opts || {};
  window.open(
    url,
    target,
    Object.entries(others)
      .reduce((preValue: string[], curValue) => {
        const [key, value] = curValue;
        return [...preValue, `${key}=${value}`];
      }, [])
      .join(',')
  );
};

export const regexUrl = new RegExp(
  '^(?!mailto:)(?:(?:http|https|ftp)://)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(/|\\?|#)[^\\s]*)?$',
  'i'
);

export function parseTime(time: any, cFormat?: any) {
  if (arguments.length === 0 || !time) {
    return null;
  }
  const format = cFormat || '{y}-{m}-{d} {h}:{i}:{s}';
  let date;
  if (typeof time === 'object') {
    date = time;
  } else {
    if (typeof time === 'string') {
      if (/^[0-9]+$/.test(time)) {
        // support "1548221490638"
        time = parseInt(time as string, 10);
      } else {
        // support safari
        // https://stackoverflow.com/questions/4310953/invalid-date-in-safari
        time = time.replace(new RegExp(/-/gm), '/');
      }
    }

    if (typeof time === 'number' && time.toString().length === 10) {
      time *= 1000;
    }
    date = new Date(time);
  }
  const formatObj: any = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay(),
  };
  return format.replace(/{([ymdhisa])+}/g, (result: any, key: string) => {
    const value = formatObj[key];
    if (key === 'a') {
      return ['日', '一', '二', '三', '四2', '五', '六'][value];
    }
    return value.toString().padStart(2, '0');
  });
}

/**
 * @param {number} time
 * @param {string} option
 * @returns {string}
 */
export function formatTime(time: any, option?: any) {
  if (`${time}`.length === 10) {
    time = parseInt(time, 10) * 1000;
  } else {
    time = +time;
  }
  const d = new Date(time);
  const now = Date.now();

  const diff = (now - toInteger(d)) / 1000;

  if (diff < 30) {
    return '刚刚';
  }
  if (diff < 3600) {
    // less 1 hour
    return `${Math.ceil(diff / 60)} 分钟前`;
  }
  if (diff < 3600 * 24) {
    return `${Math.ceil(diff / 3600)} 小时前`;
  }
  if (diff < 3600 * 24 * 2) {
    return `1天前`;
  }

  if (option) {
    return parseTime(time, option);
  }
  return `${d.getFullYear()}年${
    d.getMonth() + 1
  }月${d.getDate()}日${d.getHours()}时${d.getMinutes()}分`;
}

export function readFile(file: any) {
  return new Promise((reslove, reject) => {
    try {
      if (window.FileReader) {
        const reader = new FileReader();
        reader.readAsText(file);
        reader.onload = (e) => {
          reslove(e.target?.result);
        };
        reader.onerror = (e: any) => {
          reject(new Error(e.message));
        };
      } else {
        reject(new Error('系统不支持读取文件'));
      }
    } catch (e: any) {
      reject(new Error(e.message));
    }
  });
}

export function deepClone(data: any) {
  return JSON.parse(JSON.stringify(data));
}

// 通知封装，方便简单调用
export const notice = ({
  type = 'success',
  content,
  duration = 3000,
  ...props
}: INotice): void => {
  Notification[type]({ content, duration, ...props });
};

export default null;
