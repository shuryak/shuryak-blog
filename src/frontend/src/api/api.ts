import axios, {AxiosResponse, InternalAxiosRequestConfig} from 'axios'
import {mapKeys, snakeCase} from 'lodash'
import {camelCaseKeys} from './helpers'
import {CreateArticleRequest, GetManyArticlesRequest, Article, ShortArticle, ShortArticlesResponse, UpdateArticleRequest} from '../models/dto'

const api = axios.create({
  baseURL: 'http://localhost:8080/',
  timeout: 1000,
  withCredentials: true
})

// https://morioh.com/p/8e8b33c25ea1
api.interceptors.response.use((response: AxiosResponse) => {
  if (response.data && response.headers['content-type'].indexOf('application/json') !== 1) {
    response.data = camelCaseKeys(response.data)
  }

  return response
})
api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  const newConfig: InternalAxiosRequestConfig = {...config}
  // if (newConfig.headers['content-type'].indexOf('multipart/json') !== 1) {
  //   return newConfig
  // }
  if (config.params) {
    newConfig.params = mapKeys(config.params, (value, key) => snakeCase(key))
  }
  if (config.data) {
    newConfig.data = mapKeys(config.data, (value, key) => snakeCase(key))
  }
  return newConfig
})

export const getArticles = (payload: GetManyArticlesRequest) => api.get<ShortArticlesResponse>(
  'articles/get_many',
  {
    params: payload
  }
)

export const createArticle = (payload: CreateArticleRequest, accessToken: string) => api.post<ShortArticle>(
  'articles/create',
  payload,
  {
    headers: {
      Authorization: `Bearer ${accessToken}`
    }
  }
)

export const updateArticle = (payload: UpdateArticleRequest, accessToken: string) => api.patch<ShortArticle>(
  'articles/update',
  payload,
  {
    headers: {
      Authorization: `Bearer ${accessToken}`
    }
  }
)

export const getArticleByCustomId = (customId: string) => api.get<Article>(
  'articles/get_by_custom_id',
  {
    params: {
      customId: customId
    }
  }
)
