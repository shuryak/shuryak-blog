import {camelCase} from 'lodash';

export function camelCaseKeys(obj: Record<string, any>): Record<string, any> {
  if (Array.isArray(obj)) {
    return obj.map(camelCaseKeys)
  } else if (obj !== null && typeof obj === 'object') {
    return Object.keys(obj).reduce((result: Record<string, any>, key: string) => {
      const value = obj[key]
      const camelCaseKey: string = camelCase(key)
      result[camelCaseKey] = camelCaseKeys(value)
      return result
    }, {})
  } else {
    return obj
  }
}
