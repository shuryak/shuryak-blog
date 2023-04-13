import React, {useState, useEffect, useRef} from 'react'
import {translate, GOST_52290} from 'iuliia'
import {kebabCase} from 'lodash'
import {OutputData} from '@editorjs/editorjs'

export const PublicationBlock: React.FC<{data: OutputData | undefined, goBack: () => void}> = ({data, goBack}) => {
  const [showError, setShowError] = useState<boolean>(false)
  const [title, setTitle] = useState<string>('')
  const [url, setUrl] = useState<string>('')

  const calledOnce = useRef(false);

  useEffect(() => {
    if (calledOnce.current) {
      return;
    }

    if (data) {
      data.blocks.forEach(block => {
        if (block.type === 'header') {
          setTitle(block.data.text)
          setUrl(kebabCase(translate(block.data.text, GOST_52290)))
        }
      })
    }

    calledOnce.current = true;
  }, [])

  return (
    <>
      <p className="error-message" style={showError ? {opacity: 100} : {opacity: 0}}>
        🚫 Неверное имя пользователя или пароль.
      </p>

      <form className="article-form">
        <div className="article-form-wrapper">
          <div className="thumbnail-input-wrapper">
            <p>Нажмите<br/>для загрузки тамбнейла...</p>
          </div>
          <div className="article-form-right-block">
            <div className="input-block">
              <label htmlFor="article-title">Название статьи:</label>
              <input
                className="account-form-input user-input"
                type="text"
                id="article-title"
                value={title}
                onChange={e => setTitle(e.target.value)}
              />
            </div>

            <div className="input-block">
              <label htmlFor="url">Короткий адрес статьи:</label>
              <input
                className="account-form-input user-input"
                type="text"
                id="url"
                value={url}
                onChange={e => setUrl(e.target.value)}
              />
            </div>
          </div>
        </div>

        <input
          className="default-btn user-input"
          type="submit"
          value="💾 Сохранить как черновик"
        />
        <input
          className="default-btn user-input"
          type="submit"
          value="🙌 Опубликовать"
        />

        <button className="default-btn user-input" onClick={goBack}>
          Вернуться к редактированию
        </button>
      </form>
    </>
  )
}

export default PublicationBlock
