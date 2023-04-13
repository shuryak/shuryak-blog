import React from 'react'
import '@styles/article-preview.scss'

export const ArticlePreview: React.FunctionComponent<{title: string, shortText: string, readingTime: number, thumbnailUrl: string}> =
  ({title, shortText, readingTime, thumbnailUrl}) => {
  function getMinutesString(minutes: number):string {
    const lastDigit = minutes % 10;
    if (lastDigit === 1 && minutes !== 11) {
      return `${minutes} минута`;
    } else if ([2, 3, 4].includes(lastDigit) && ![12, 13, 14].includes(minutes)) {
      return `${minutes} минуты`;
    } else {
      return `${minutes} минут`;
    }
  }

  return (
    <div className="article-preview">
      <div className="article-preview-text">
        <p className="article-preview-title">{title}</p>
        <p className="article-preview-short-text">{shortText}</p>
        <div className="article-preview-reading-time">{`Время чтения: ${getMinutesString(readingTime)}`}</div>
      </div>
      <img className="article-preview-thumbnail" src={thumbnailUrl} alt={title}/>
    </div>
  )
}

export default ArticlePreview
