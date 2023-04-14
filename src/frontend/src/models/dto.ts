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
