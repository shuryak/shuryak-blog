import {OutputData} from '@editorjs/editorjs'

export interface LoginRequest {
  username: string
  password: string
}

export interface GetManyArticlesRequest {
  offset: number
  count: number
  isDrafts: boolean
}

export interface ShortArticlesResponse {
  articles: ShortArticle[]
}

export interface ShortArticle {
  id: number
  customId: string
  authorId: number
  title: string
  thumbnail: string
  shortContent: string
  isDraft: boolean
  createdAt: string
  updatedAt: string
}

export interface Article {
  id: number
  customId: string
  authorId: number
  title: string
  thumbnail: string
  content: OutputData
  isDraft: boolean
  createdAt: string
  updatedAt: string
}

export interface CreateArticleRequest {
  customId: string
  title: string
  thumbnail: string
  content: OutputData
  isDraft: boolean
}

export interface UpdateArticleRequest {
  customId: string
  title: string
  thumbnail: string
  content: OutputData
  isDraft: boolean
}
