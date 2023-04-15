import React, {useEffect, useState, useRef} from 'react'
import ArticlePreview from '../components/ArticlePreview'
import {getArticles} from '../api/api'
import '@styles/page.scss'
import {ShortArticle} from '../models/dto'

export const ArticlesPage: React.FC = () => {
  const [articles, setArticles] = useState<ShortArticle[]>([])

  const calledOnce = useRef(false)

  useEffect(() => {
    if (calledOnce.current) {
      return;
    }

    getArticles({
      count: 10,
      offset: 0,
      isDrafts: false
    }).then(data => {
      setArticles(data.data.articles)
    })

    calledOnce.current = true;
  }, [])

  return (
    <>
      <h1>📰 Последние статьи</h1>
      <div className="articles-list">
        {articles?.map((article, idx) =>
          <ArticlePreview
            title={article.title}
            shortText={article.shortContent}
            readingTime={4}
            thumbnailUrl={article.thumbnail}
            key={idx}
          />
        )}
      </div>
    </>
  )
}

export default ArticlesPage
