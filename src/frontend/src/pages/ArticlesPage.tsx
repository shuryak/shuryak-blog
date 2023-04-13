import React, {useState} from 'react'
import ArticlePreview from '../components/ArticlePreview'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPenToSquare, faExpand, faCompress } from '@fortawesome/free-solid-svg-icons'
import '@styles/page.scss'

export const ArticlesPage: React.FC = () => {
  const [isExpand, setIsExpand] = useState<boolean>(false)

  return (
    <>
      <h1>📰 Последние статьи</h1>
      <div className="articles-list">
        <ArticlePreview
          title={'🔥 Какое-то клёвое название статьи'}
          shortText={'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusantium adipisci aliquid cupiditate dolores dolorum eligendi eveniet, ex illum incidunt maxime non odio officia possimus quam rem repudiandae sit, unde velit.'}
          readingTime={6}
          thumbnailUrl={'https://sun9-5.userapi.com/impg/nSonDo4MatowCsx2dnTxzcPwnJdCyB9LLBN7Sg/l8yxqCivrK4.jpg?size=1080x1080&quality=96&sign=34fa7e9a1417f982c0135f6f5cd77801&type=album'}
        />
        <ArticlePreview
          title={'Какое-то клёвое название статьи 2'}
          shortText={'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusantium adipisci aliquid cupiditate dolores dolorum eligendi eveniet, ex illum incidunt maxime non odio officia possimus quam rem repudiandae sit, unde velit.'}
          readingTime={4}
          thumbnailUrl={'https://sun9-5.userapi.com/impg/nSonDo4MatowCsx2dnTxzcPwnJdCyB9LLBN7Sg/l8yxqCivrK4.jpg?size=1080x1080&quality=96&sign=34fa7e9a1417f982c0135f6f5cd77801&type=album'}
        />
      </div>
    </>
  )
}

export default ArticlesPage
